package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/luispalacio22/gambit/auth"
	"github.com/luispalacio22/gambit/routers"
)

func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Voy aprocesar " + path + " > " + method)
	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)
	isOK, statusCode, user := ValidoAuthorization(path, method, headers)
	if !isOK {
		return statusCode, user
	}

	switch path[0:4] {
	case "user":
		return ProcesoUsers(body, path, method, user, id, request)
	case "prod":
		return ProcesoProducts(body, path, method, user, idn, request)
	case "stoc":
		return ProcesoStock(body, path, method, user, idn, request)
	case "addr":
		return ProcesoAddres(body, path, method, user, idn, request)
	case "cate":
		return ProcesoCategory(body, path, method, user, idn, request)
	case "orde":
		return ProcesoOrder(body, path, method, user, idn, request)
	}

	return 400, "Method invalid"
}

func ValidoAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") || (path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "token requerido"
	}

	todoOk, err, msg := auth.ValidoToken(token)
	if !todoOk {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error en el token " + msg)
			return false, 401, msg
		}
	}
	fmt.Println("Token ok")
	return true, 200, msg

}

func ProcesoUsers(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}

func ProcesoProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}

func ProcesoCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	case "PUT":
		return routers.UpdateCategory(body, user, id)
	case "DELETE":
		return routers.DeleteCategory(user, id)
	case "GET":
		return routers.SelectCategories(body, request)
	}
	return 400, "Method invalid"
}

func ProcesoStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}

func ProcesoAddres(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}

func ProcesoOrder(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}
