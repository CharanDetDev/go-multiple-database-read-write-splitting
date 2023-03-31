package main

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/http/route"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/config"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/database"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/logg"
	"github.com/gofiber/fiber/v2"
)

func init() {

	isConfig := config.ConfigInit()
	isDatabase := database.InitDatabase()
	if isConfig && isDatabase {
		logg.PrintloggerVariadicHasHeader("\t ***** Initail :: Configuration & Database :: SUCCESS **** | ", "Results", *database.Conn)
	} else {
		logg.PrintloggerVariadicHasHeader("\t ***** Initail :: Configuration & Database :: ERROR **** | ", "Results", *database.Conn, logg.GetCallerPathNameFileNameLineNumber())
	}

}

func main() {
	defer database.ConnectionClose()

	app := fiber.New()
	router := route.NewRoute()
	router.InitRoute(app)

	app.Listen(config.Env.API_PORT)
}
