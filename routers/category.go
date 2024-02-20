package routers

import (
	"encoding/json"

	"github.com/luispalacio22/gambit/bd"
	"github.com/luispalacio22/gambit/models"
)

func InsertCategory(body string, User string) (int, string) {
	var t models.Category
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}
	if len(t.CategName) == 0 {
		return 400, "Debe especificar el nombre (tittle) de la categoria"
	}

	if len(t.CategPath) == 0 {
		return 400, "Debe especificar el Path (ruta) de la categoria"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

}
