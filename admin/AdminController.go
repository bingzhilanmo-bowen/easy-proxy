package admin

import (
	"easy-proxy/database"
	"easy-proxy/database/route"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"log"
	"net/http"
)

func AdminServer(addr string) {
	m := martini.Classic()
	db := database.GormDb()
	m.Map(db)
	m.Use(render.Renderer())

	m.Group("/route", func(router martini.Router) {
		router.Post("/create",binding.Bind(route.RouteJson{}) ,route.AddRoute)
		router.Get("/get/:id", route.GetById)
	})

	m.Get("/", func() string {
		return "hello !!!"
	})

	err := http.ListenAndServe(addr, m)

	if err == nil  {
       log.Panicf("start admin %s", addr)
	}
}
