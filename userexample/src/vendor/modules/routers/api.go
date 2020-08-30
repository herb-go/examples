package routers

import (
	"modules/members"
	loginaction "modules/members/login/actions"
	"modules/middlewares"
	"modules/systems/users/activesaction"
	"modules/systems/users/currentaction"
	"modules/systems/users/revokeaction"
	"modules/systems/users/statusactions"

	"modules/systems/users/listaction"
	registeraction "modules/systems/users/register/actions"
	updatepasswordaction "modules/systems/users/updatepassword/actions"
	updateprofileaction "modules/systems/users/updateprofile/actions"

	"github.com/herb-go/herb/identifier"

	"github.com/herb-go/util/action/commonaction"

	"github.com/herb-go/herb/middleware"
	"github.com/herb-go/herb/middleware/errorpage"
	"github.com/herb-go/herb/middleware/router"
	"github.com/herb-go/herb/middleware/router/httprouter"
)

var LoginRequired = identifier.NewLoggedInFilter(members.WebSession, nil).ServeMiddleware

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
	Router.GET("/current").Use(LoginRequired).Handle(currentaction.ActionCurrent)
	Router.GET("/actives").Use(LoginRequired).Handle(activesaction.ActionActives)
	Router.POST("/updatepassword").Use(LoginRequired).Handle(updatepasswordaction.ActionUpdatepassword)
	Router.POST("/updateprofile").Use(LoginRequired).Handle(updateprofileaction.ActionUpdateprofile)

	Router.POST("/revoke/:id").Use(LoginRequired).Handle(revokeaction.ActionRevoke)
	Router.POST("/enable/:id").Use(LoginRequired).Handle(statusactions.ActionEnable)
	Router.POST("/disable/:id").Use(LoginRequired).Handle(statusactions.ActionDisable)
	Router.POST("/logout").Use(LoginRequired, members.WebSession.LogoutMiddleware).HandleFunc(commonaction.SuccessAction)
	Router.POST("/login").Handle(loginaction.ActionLogin)
	Router.POST("/register").Handle(registeraction.ActionRegister)
	return Router
})
