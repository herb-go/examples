package activesaction

//Auto generated code for hiring workers.
//DO NOT EDIT THIS FILE.
import worker "github.com/herb-go/worker"

func init() {
	//Worker "systems/users/activesaction.ActionActives"
	//You can add Introduction by add comment in form WORKER(ActionActives):Introduction
	worker.Hire("systems/users/activesaction.ActionActives", &ActionActives)

}