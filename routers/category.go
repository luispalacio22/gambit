package routers

import (
	"encoding/json"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
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

	result, err2 := bd.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el registo de la categoria" + t.CategName + " > " + err2.Error()
	}
	return 200, "{CategID: " + strconv.Itoa(int(result)) + "}"
}

func UpdateCategory(body string, User string, id int) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}
	if len(t.CategName) == 0 || len(t.CategPath) == 0 {
		return 400, "Debe especificar el nombre (tittle) o el path de la categoria"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.CategID = id
	err2 := bd.UpdateCategory(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el update de la categoria" + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Update Ok"
}

func DeleteCategory(User string, id int) (int, string) {
	if id == 0 {
		return 400, "Debe especificar el id a borrar"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err := bd.DeleteCategory(id)
	if err != nil {
		return 400, "Ocurrio un error al realizar el delete de la categoria" + strconv.Itoa(id) + ">" + err.Error()
	}

	return 200, "Delete ok"
}
