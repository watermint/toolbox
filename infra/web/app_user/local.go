package app_user

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/json"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	localUserHashFile = "local_user.json"
)

type localUserContext struct {
	UserHash string `json:"user_hash"`
}

func NewLocalUser(app app_workspace.Application) (user User, err error) {
	lu := &localUser{
		app: app,
	}
	if err = lu.setup(); err != nil {
		return nil, err
	}
	return lu, nil
}

type localUser struct {
	app      app_workspace.Application
	user     app_workspace.MultiUser
	userHash string
}

func (z *localUser) toFile(hash string) (err error) {
	lu := &localUserContext{
		UserHash: hash,
	}
	b, err := json.Marshal(&lu)
	if err != nil {
		app_root.Log().Debug("Unable to marshal", zap.Error(err))
		return err
	}
	path := filepath.Join(z.app.Home(), localUserHashFile)
	return ioutil.WriteFile(path, b, 0600)
}

func (z *localUser) fromFile() (hash string, err error) {
	lu := &localUserContext{}

	path := filepath.Join(z.app.Home(), localUserHashFile)
	l := app_root.Log().With(zap.String("path", path))

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return "", err
	}

	f, err := os.Open(path)
	if err != nil {
		l.Debug("Unable to open file", zap.Error(err))
		return "", err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		l.Debug("Unable to read file", zap.Error(err))
		return "", err
	}
	if err = json.Unmarshal(b, lu); err != nil {
		l.Debug("Unable to unmarshal", zap.Error(err))
		return "", err
	}
	return lu.UserHash, nil
}

func (z *localUser) generate() (hash string, err error) {
	l := app_root.Log()
	b := make([]byte, 40)
	_, err = rand.Read(b)
	if err != nil {
		l.Debug("Unable to generate", zap.Error(err))
		return "", err
	}
	hash = base32.StdEncoding.EncodeToString(b)

	if err = z.toFile(hash); err != nil {
		l.Debug("Unable to store", zap.Error(err))
		return "", err
	}
	return
}

func (z *localUser) setup() error {
	hash, err := z.fromFile()
	if err != nil {
		hash, err = z.generate()
		if err != nil {
			return err
		}
	}
	z.userHash = hash
	z.user, err = app_workspace.NewMultiUser(z.app, z.userHash)
	if err != nil {
		return err
	}
	return nil
}

func (z *localUser) UserHash() string {
	return z.userHash
}

func (z *localUser) Workspace() app_workspace.MultiUser {
	return z.user
}
