package file

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

var idp *IDP

func (i *IDP) Validate(username, password string) bool {
	for _, login := range idp.logins.Logins {
		if username != login.Username {
			continue
		}
		if password == login.Password {
			return true
		}
	}
	return false
}

func init() {
	idp = &IDP{logins: model.Logins{}}
	readDefaultLogins()
}

func readDefaultLogins() {
	loginsFile, err := os.Open(loginsPath)
	if err != nil {
		log.Fatalf("unable to open logins file: %v", err)
	}
	defer loginsFile.Close()

	loginBytes, err := ioutil.ReadAll(loginsFile)
	if err != nil {
		log.Fatalf("unable to read logins file %v", err)
	}

	if err = json.Unmarshal(loginBytes, &idp.logins); err != nil {
		log.Fatalf("unable to unmarshal logins %v", err)
	}
}
