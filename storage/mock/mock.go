package mock

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/yabslabs/simple-auth/model"
)

const loginsPath = "../../resources/logins.json"

type IDP struct {
	logins model.Logins
}

func Init() *IDP {
	idp := &IDP{logins: model.Logins{}}
	idp.readDefaultLogins()
	return idp
}

func (i *IDP) Login(username, password string) bool {
	for _, login := range i.logins.Logins {
		if username != login.Username {
			continue
		}
		if password == login.Password {
			return true
		}
	}
	return false
}

func (i *IDP) GetUserRoles(username string) []model.Role {
	for _, login := range i.logins.Logins {
		if login.Username == username {
			return login.Roles
		}
	}
	return nil
}

func (i *IDP) readDefaultLogins() {
	loginsFile, err := os.Open(loginsPath)
	if err != nil {
		log.Fatalf("unable to open logins file: %v", err)
	}
	defer loginsFile.Close()

	loginBytes, err := ioutil.ReadAll(loginsFile)
	if err != nil {
		log.Fatalf("unable to read logins file %v", err)
	}

	if err = json.Unmarshal(loginBytes, &i.logins); err != nil {
		log.Fatalf("unable to unmarshal logins %v", err)
	}
}
