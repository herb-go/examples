package actions

//Auto generated code for hiring workers.
//DO NOT EDIT THIS FILE.
import worker "github.com/herb-go/worker"

func init() {
	//Worker "members/login/actions.ActionLogin"
	//You can add Introduction by add comment in form WORKER(ActionLogin):Introduction
	worker.Hire("members/login/actions.ActionLogin", &ActionLogin)

}