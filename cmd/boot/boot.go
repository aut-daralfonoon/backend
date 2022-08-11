package boot

import (
	"log"

	"github.com/AUT-Daralfonoon/back-end/gateway/api"
	"github.com/AUT-Daralfonoon/back-end/gateway/api/holding/event"
	em "github.com/AUT-Daralfonoon/back-end/modules/holding/event"
)

var managers = struct{ event *em.EventManager }{
	event: &em.EventManager{},
}

func Boot() {
	log.Println("Loading Providers...")
    NewProvider()

    log.Println("Loading Server")
    ServerRouter()
}

func NewProvider() {
	NewEventProvider()
}

func bootManagers() {
	managers.event = em.InitManager()
}

func NewEventProvider() {
	sc := api.NewGateway()
	ec := event.NewEventController(managers.event)
	sc.EventManager = ec
}

func ServerRouter() {
	r := api.InitServer()
	if err := r.Run("localhost:8080"); err != nil {
		log.Fatalf("server could not start")
	}
}
