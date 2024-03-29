package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidoToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		fmt.Println("El token no es valido")
		return false, nil, "Eltoken no es valido"
	}

	UserInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Print("No se puede decodificar la parte del token: ", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(UserInfo, &tkj)
	if err != nil {
		fmt.Println("No se puede decodificar la estrutura JSON ", err.Error())
	}

	ahora := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(ahora) {
		fmt.Println("Fecha expiracion token = " + tm.String())
		fmt.Println("Token expirado!")
		return false, err, "Token expirado !!"
	}

	return true, nil, string(tkj.Username)
}
