package bd

import (
	"database/sql"
	"fmt"

	"github.com/luispalacio22/gambit/models"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("comienza registro de InsertCategory")
	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := "INSERT INTO category (Categ_Name, Cated:Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"
	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInserId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	fmt.Println("Insert category ejecucion exitosa")
	return LastInserId, nil
}
