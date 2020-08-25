package actions

//Actions for form register .
//You can  bind actions to  router by using below code :
//import   registeractions "modules/systems/users/register/actions"
//
//	Router.POST("/register").
//		Handle(registeractions.ActionRegister)

import (
	"net/http"

	"modules/systems/users/register/forms"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/herb/ui/render"
	"github.com/herb-go/herb/ui/validator/formdata"
)

//ActionRegister action that verifyregister form in json format.
var ActionRegister = action.New(func(w http.ResponseWriter, r *http.Request) {
	form := forms.NewRegisterForm()
	if formdata.MustValidateJSONRequest(r, form) {
		err := form.Exec()
		if err != nil {
			panic(err)
		}
		render.MustJSON(w, form, 200)
	} else {
		formdata.MustRenderErrorsJSON(w, form)
	}
})
