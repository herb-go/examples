package actions

//Actions for form updateprofile .
//You can  bind actions to  router by using below code :
//import   updateprofileactions "modules/systems/users/updateprofile/actions"
//
//	Router.POST("/updateprofile").
//		Handle(updateprofileactions.ActionUpdateprofile)

import (
	"net/http"

	"modules/systems/users/updateprofile/forms"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/herb/ui/render"
	"github.com/herb-go/herb/ui/validator/formdata"
)

//ActionUpdateprofile action that verifyupdateprofile form in json format.
var ActionUpdateprofile = action.New(func(w http.ResponseWriter, r *http.Request) {
	form := forms.NewUpdateprofileForm()
	if formdata.MustValidateJSONRequest(r, form) {
		err := form.Exec()
		if err != nil {
			panic(err)
		}
		render.MustJSON(w, "ok", 200)
	} else {
		formdata.MustRenderErrorsJSON(w, form)
	}
})
