package routers

import (
	"modules/members"
	loginaction "modules/members/login/actions"
	"modules/middlewares"
	"modules/systems/users/activesaction"
	"modules/systems/users/currentaction"
	"modules/systems/users/listaction"
	registeraction "modules/systems/users/register/actions"

	"github.com/herb-go/util/action/commonaction"

	"github.com/herb-go/herb/middleware"
	"github.com/herb-go/herb/middleware/errorpage"
	"github.com/herb-go/herb/middleware/router"
	"github.com/herb-go/herb/middleware/router/httprouter"
)

//APIMiddlewares middlewares that should used in api requests
var APIMiddlewares = func() middleware.Middlewares {
	return middleware.Middlewares{
		middlewares.MiddlewareCsrfVerifyHeader,
		errorpage.MiddlewareDisable,
	}
}

//RouterAPIFactory api router factory.
var RouterAPIFactory = router.NewFactory(func() router.Router {
	var Router = httprouter.New()
	//Put your router configure code here
	Router.GET("/list").Handle(listaction.ActionList)
	Router.GET("/current").Handle(currentaction.ActionCurrent)
	Router.GET("/actives").Handle(activesaction.ActionActives)
	Router.POST("/logout").Use(members.WebSession.LogoutMiddleware).HandleFunc(commonaction.SuccessAction)
	Router.POST("/login").Handle(loginaction.ActionLogin)
	Router.POST("/register").Handle(registeraction.ActionRegister)
	return Router
})