package app_user

import (
	"errors"
	"github.com/watermint/toolbox/experimental/app_root"
	"github.com/watermint/toolbox/experimental/app_workspace"
	"go.uber.org/zap"
)

type Repository interface {
	Resolve(userHash string) (User, error)
}

type RootUserRepository interface {
	RootUser() User
}

var (
	localRepository Repository
)

func newSingleUserRepository(app app_workspace.Application) (r Repository, err error) {
	user, err := NewLocalUser(app)
	if err != nil {
		return nil, err
	}

	r = &singleUserRepository{
		User: user,
	}
	return r, nil
}

func SingleUserRepository(app app_workspace.Application) (r Repository, err error) {
	if localRepository == nil {
		localRepository, err = newSingleUserRepository(app)
		if err != nil {
			return nil, err
		}
	}
	return localRepository, nil
}

type singleUserRepository struct {
	User User
}

func (z *singleUserRepository) RootUser() User {
	return z.User
}

func (z *singleUserRepository) Resolve(userHash string) (User, error) {
	if z.User.UserHash() == userHash {
		return z.User, nil
	} else {
		app_root.Log().Debug("User not found for userHash", zap.String("hash", userHash))
		return nil, errors.New("user not found")
	}
}
