package sv_project

import (
	"errors"
	"github.com/watermint/toolbox/domain/figma/api/fg_client"
	"github.com/watermint/toolbox/domain/figma/model/mo_file"
	"github.com/watermint/toolbox/domain/figma/model/mo_project"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"regexp"
)

type MsgProject struct {
	VerifyOk                         app_msg.Message
	VerifyTeamIdTooLong              app_msg.Message
	VerifyTeamIdInvalidCharacters    app_msg.Message
	VerifyProjectIdTooLong           app_msg.Message
	VerifyProjectIdInvalidCharacters app_msg.Message
}

var (
	MProject = app_msg.Apply(&MsgProject{}).(*MsgProject)
)

type Project interface {
	List(teamId string) ([]*mo_project.Project, error)
	Files(projectId string) ([]*mo_file.File, error)
}

const (
	VerifyTeamIdLooksOkay VerifyResult = iota
	VerifyTeamIdTooLong
	VerifyTeamInvalidChar
)

const (
	// TeamIdMaxLength maximum length of the team id. There is no clear definition in the
	// API doc as of implementation. But keep it smaller for safe.
	// Current team id length is ~20 numeric characters. This threshold includes buffer for
	// future changes in Figma API side.
	TeamIdMaxLength = 36
)

var (
	// TeamIdRegex is the allowed Team Id pattern.
	// There is no clear definition in the API doc as of implementation.
	// Current team id is numeric. But this pattern definition allows alphabetic characters
	// for future changes in Figma API side.
	TeamIdRegex = regexp.MustCompile(`^[a-zA-Z0-9]*$`)
)

type VerifyResult int

func VerifyTeamId(teamId string) (VerifyResult, app_msg.Message) {
	if TeamIdMaxLength < len(teamId) {
		return VerifyTeamIdTooLong, MProject.VerifyTeamIdTooLong
	}
	if !TeamIdRegex.MatchString(teamId) {
		return VerifyTeamInvalidChar, MProject.VerifyTeamIdInvalidCharacters
	}
	return VerifyTeamIdLooksOkay, MProject.VerifyOk
}

const (
	VerifyProjectIdLooksOkay VerifyResult = iota
	VerifyProjectIdTooLong
	VerifyProjectIdInvalidCharacter
)

const (
	// ProjectIdMaxLength maximum length of the project id.
	// There is no clear definition in the API doc as of implementation.
	// But keep it smaller for safe. Current team id length is ~8 numeric characters.
	// This threshold includes buffer for future changes in Figma API side.
	ProjectIdMaxLength = 20
)

var (
	// ProjectIdRegex is the allowed Project Id pattern.
	// There is no clear definition in the API doc as of implementation.
	// Current project id is numeric. But this pattern definition allows alphabetic characters
	// for future changes in Figma API side.
	ProjectIdRegex = regexp.MustCompile(`^[a-zA-Z0-9]*$`)
)

func VerifyProjectId(projectId string) (VerifyResult, app_msg.Message) {
	if ProjectIdMaxLength < len(projectId) {
		return VerifyProjectIdTooLong, MProject.VerifyProjectIdTooLong
	}
	if !TeamIdRegex.MatchString(projectId) {
		return VerifyProjectIdInvalidCharacter, MProject.VerifyProjectIdInvalidCharacters
	}
	return VerifyProjectIdLooksOkay, MProject.VerifyOk
}

func New(client fg_client.Client) Project {
	return &projectImpl{
		client: client,
	}
}

type projectImpl struct {
	client fg_client.Client
}

func (z projectImpl) List(teamId string) (projects []*mo_project.Project, err error) {
	if r, m := VerifyTeamId(teamId); r != VerifyTeamIdLooksOkay {
		return nil, errors.New(m.Key())
	}
	res := z.client.Get("teams/" + teamId + "/projects")
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	projects = make([]*mo_project.Project, 0)
	err = res.Success().Json().FindArrayEach("projects", func(e es_json.Json) error {
		prj := &mo_project.Project{}
		if err := e.Model(prj); err != nil {
			return err
		}
		projects = append(projects, prj)
		return nil
	})
	return projects, err
}

func (z projectImpl) Files(projectId string) (files []*mo_file.File, err error) {
	res := z.client.Get("projects/" + projectId + "/files")
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	files = make([]*mo_file.File, 0)
	err = res.Success().Json().FindArrayEach("files", func(e es_json.Json) error {
		f := &mo_file.File{}
		if err := e.Model(f); err != nil {
			return err
		}
		files = append(files, f)
		return nil
	})
	return files, err
}
