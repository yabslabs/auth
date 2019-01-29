package storage

import (
	"github.com/yabslabs/auth/model"
)

type Auth interface {
	Login(username, password string) bool
	GetRoles(username string) *model.Roles
}

func InitAuth(authType string) Auth {
	switch authType {
	case "mock":
		return mock.Init()
	}
}
