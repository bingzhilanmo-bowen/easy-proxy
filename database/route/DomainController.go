package route

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"strconv"
)

type RouteJson struct {
	Domain string `json:"domain" binding:"required"`
	RouteType int `json:"routeType" `
	Expression string `json:"expression" `
}

type Route struct {
	gorm.Model
	Domain string
	RouteType int
	Expression string
}

func AddRoute(json RouteJson, db *gorm.DB, render render.Render) {
	route := Route{Domain:json.Domain,
		           RouteType:json.RouteType,
	               Expression:json.Expression}
	db.Create(&route)
	render.JSON(200, route)
}

func GetById(params martini.Params, db *gorm.DB, render render.Render)  {
	idStr := params["id"]
	id,err := strconv.Atoi(idStr)

	if err != nil {
		panic("id must is a number !")
	}

	var rs = Route{}
	db.First(&rs,id)

	render.JSON(200, rs)
}


