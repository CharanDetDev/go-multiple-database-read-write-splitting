package route

import (
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/domain"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/http/handler"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/repository"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/service"
	"github.com/CharanDetDev/go-multiple-database-read-write-splitting/util/database"
	"github.com/gofiber/fiber/v2"
)

type (
	route struct {
		PersonHandler domain.PersonHandler
	}

	Route interface {
		InitRoute(app *fiber.App)
	}
)

func NewRoute() Route {

	newAddressRepo := repository.NewAddressRepo(database.Conn)

	newPersonRepo := repository.NewPersonRepo(database.Conn)
	newPersonService := service.NewPersonService(newPersonRepo, newAddressRepo)
	newPersonHandle := handler.NewPersonHandler(newPersonService)

	return &route{
		newPersonHandle,
	}
}

func (route *route) InitRoute(app *fiber.App) {

	personGroup := app.Group("/person")
	route.personGroup(personGroup)

}
