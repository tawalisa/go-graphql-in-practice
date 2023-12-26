package myecho

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-grapohql-in-practice/graphql"
)

const delimiter = "__!!__"

var myserver *echo.Echo

func init() {
	e := initServer()
	initGraphQL(e)

	myserver = e
}

func initGraphQL(e *echo.Echo) {
	graphql.MyGraphql{}.RunwithMyEcho(e)
}

func initServer() *echo.Echo {
	e := echo.New()

	return e
}

func Run() {

	err := myserver.Start(":8080")
	if err != nil {
		log.Error(err)
		return
	}

}
