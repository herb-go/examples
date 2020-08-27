package actions

//Actions for form updatepassword .
//You can  bind actions to  router by using below code :
//import   updatepasswordactions "modules/systems/users/updatepassword/actions"
//
//	Router.POST("/updatepassword").
//		Handle(updatepasswordactions.ActionUpdatepassword)

import (
	"net/http"

	"modules/systems/users/updatepassword/forms"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/herb/ui/render"
	"github.com/herb-go/herb/ui/validator/formdata"
)

//ActionUpdatepassword action that verifyupdatepassword form in json format.
var ActionUpdatepassword = action.New(func(w http.ResponseWriter, r *http.Request) {
	form := forms.NewUpdatepasswordForm()
	if formdata.MustValidateJSONRequest(r, form) {
		err := form.Exec()
		if err != nil {
			panic(err)
		}
		render.MustJSON(w, "success", 200)
	} else {
		formdata.MustRenderErrorsJSON(w, form)
	}
})
