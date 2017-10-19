package tree

import (
	"flag"
	"os"
	"github.com/watermint/toolbox/infra"
	"fmt"
	"strings"
	"errors"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra/util"
	"database/sql"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
)

type ReportOpts struct {
	db *sql.DB

	Infra      *infra.InfraOpts
	Target     string
	Anonymise  bool
	ReportPath string
}

const (
	TARGET_FULL_DROPBOX = "account"
	TARGET_TEAM_FOLDER  = "team-folder"

	REPORT_VERSION = 1
)

var (
	SUPPORTED_TARGETS = []string{
		TARGET_FULL_DROPBOX,
		TARGET_TEAM_FOLDER,
	}
)

func ParseReportOptions(args []string) (*ReportOpts, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	opts := &ReportOpts{}
	opts.Infra = infra.PrepareInfraFlags(f)

	descTarget := fmt.Sprintf("Target (%s)", strings.Join(SUPPORTED_TARGETS, ","))
	f.StringVar(&opts.Target, "target", "", descTarget)

	descAnonymise := "Anonymise report (remove actual file/folder names, email address etc.)"
	f.BoolVar(&opts.Anonymise, "anon", true, descAnonymise)

	descReportPath := "Path and filename of the report"
	f.StringVar(&opts.ReportPath, "out", "", descReportPath)

	f.SetOutput(os.Stderr)
	f.Parse(args)

	return opts, nil
}

func ExecReport(args []string) error {
	opts, err := ParseReportOptions(args)
	if err != nil {
		return err
	}
	if opts.ReportPath == "" {
		return errors.New("please specify report file path")
	}

	defer opts.Infra.Shutdown()
	err = opts.Infra.Startup()
	if err != nil {
		seelog.Errorf("Unable to start operation: %s", err)
		return err
	}
	seelog.Tracef("options: %s", util.MarshalObjectToString(opts))

	switch opts.Target {
	case TARGET_FULL_DROPBOX:
		token, err := opts.Infra.LoadOrAuthDropboxFull()
		if err != nil || token == "" {
			seelog.Errorf("Unable to acquire token (error: %s)", err)
			return err
		}
		return opts.ReportFullDropbox(token)

	case TARGET_TEAM_FOLDER:
		token, err := opts.Infra.LoadOrAuthBusinessFile()
		if err != nil || token == "" {
			seelog.Errorf("Unable to acquire token (error: %s)", err)
			return err
		}
		return opts.ReportTeamFolder(token)

	default:
		return errors.New(fmt.Sprintf("unsupported target (%s)", opts.Target))
	}
}

func (r *ReportOpts) Prepare() error {
	var err error
	r.db, err = sql.Open("sqlite3", r.ReportPath)
	if err != nil {
		seelog.Errorf("Unable to open file: path[%s] error[%s]", r.ReportPath, err)
		return err
	}

	q := `
	DROP TABLE IF EXISTS config
	`
	_, err = r.db.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to drop table : error[%s]", err)
		return err
	}

	q = `
	CREATE TABLE config (
	  ver	   INT8,
	  target   VARCHAR
	)
	`
	_, err = r.db.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to create table : error[%s]", err)
		return err
	}

	q = `
	INSERT OR REPLACE INTO config (
	  ver,
	  target
	) VALUES (?, ?)
	`
	_, err = r.db.Exec(q, REPORT_VERSION, r.Target)
	if err != nil {
		seelog.Warnf("Unable to write config : error[%s]", err)
		return err
	}
	return nil
}

func (r *ReportOpts) ReportTeamFolder(teamFileToken string) error {
	err := r.Prepare()
	if err != nil {
		seelog.Warnf("Unable to prepare report database : error[%s]", err)
		return err
	}

	q := `
	DROP TABLE IF EXISTS team_folder
	`
	_, err = r.db.Exec(q)
	if err == nil {
		seelog.Errorf("Unable to drop table: %s", err)
		return err
	}

	q = `
	CREATE TABLE team_folder (
	  team_folder_id         VARCHAR PRIMARY KEY,
	  name                   VARCHAR,
      status                 VARCHAR,
	  is_team_shared_dropbox BOOLEAN
	)
	`
	_, err = r.db.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to create table : error[%s]", err)
		return err
	}

	q = `
	DROP TABLE IF EXISTS shared_folder
	`
	_, err = r.db.Exec(q)
	if err == nil {
		seelog.Errorf("Unable to drop table: %s", err)
		return err
	}
	
	q = `
	CREATE TABLE shared_folder (
	  shared_folder_id      VARCHAR PRIMARY KEY,
	  is_inside_team_folder BOOL,
	  is_team_folder        BOOL,
	  name                  VARCHAR,
	  path_lower            VARCHAR,

	  policy_acl_update     VARCHAR,
	  policy_shared_link    VARCHAR,
	  policy_member         VARCHAR,
	  policy_resolved_member_policy VARCHAR,
	  viewer_team_member_id VARCHAR
	)
	`
	_, err = r.db.Exec(q)
	if err != nil {
		seelog.Warnf("Unable to create table : error[%s]", err)
		return err
	}

	err = r.reportTeamFolders(teamFileToken)
	if err != nil {
		seelog.Warnf("Unable to create report [team-folder] : error[%s]", err)
		return err
	}

	return nil
}

func (r *ReportOpts) reportTeamFolders(teamFileToken string) error {
	seelog.Info("Loading team folders")

	client := team.New(dropbox.Config{Token: teamFileToken})

	listArg := team.NewTeamFolderListArg()
	list, err := client.TeamFolderList(listArg)
	if err != nil {
		seelog.Warnf("Unable to load team folder list : error[%s]", err)
		return err
	}

	insertQuery := `
	INSERT OR REPLACE INTO team_folder (
	  team_folder_id,
	  name,
	  status,
	  is_team_shared_dropbox
    )
	`

	hasMore := true

	for hasMore {
		for _, tf := range list.TeamFolders {
			name := ""
			if !r.Anonymise {
				name = tf.Name
			}
			_, err = r.db.Exec(
				insertQuery,
				tf.TeamFolderId,
				name,
				tf.Status.Tag,
				tf.IsTeamSharedDropbox,
			)

			seelog.Tracef(
				"Loading team folder: TeamFolderId[%s] Status[%s] IsTeamSharedDropbox[%t]",
				tf.TeamFolderId,
				tf.Status.Tag,
				tf.IsTeamSharedDropbox,
			)

			if err != nil {
				seelog.Warnf("Unable to insert data : error[%s]", err)
				return err
			}
		}

		hasMore = list.HasMore

		if hasMore {
			contArg := team.NewTeamFolderListContinueArg(list.Cursor)
			list, err = client.TeamFolderListContinue(contArg)
			if err != nil {
				seelog.Warnf("Unable to load team folder list continue : error[%s]", err)
				return err
			}
		}
	}
	return nil
}

func (r *ReportOpts) reportSharedFolders() error {
	//DROP
	q := ``


}

func (r *ReportOpts) reportSharedFolderMembers() error {


}

func (r *ReportOpts) ReportFullDropbox(fullDropboxToken string) error {
	return errors.New("not implemented")
}
