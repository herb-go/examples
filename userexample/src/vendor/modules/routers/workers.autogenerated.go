package routers

//Auto generated code for hiring workers.
//DO NOT EDIT THIS FILE.
import worker "github.com/herb-go/worker"

func init() {
	//Worker "routers.LoginRequired"
	//You can add Introduction by add comment in form WORKER(LoginRequired):Introduction
	worker.Hire("routers.LoginRequired", &LoginRequired)

	//Worker "routers.RouterAPIFactory"
	//You can add Introduction by add comment in form WORKER(RouterAPIFactory):Introduction
	worker.Hire("routers.RouterAPIFactory", &RouterAPIFactory)

	//Worker "routers.RouterFactory"
	//You can add Introduction by add comment in form WORKER(RouterFactory):Introduction
	worker.Hire("routers.RouterFactory", &RouterFactory)

}