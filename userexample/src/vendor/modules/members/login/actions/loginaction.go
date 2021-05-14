package actions

import (
	"net/http"

	"github.com/herb-go/herb/ui/render"
	"github.com/herb-go/herb/ui/validator/formdata"

	usermodule "modules/members"
	"modules/members/login/forms"

	"github.com/herb-go/herb/middleware/action"
)

//ActionLogin action that verify login form in json format.
var ActionLogin = action.New(func(w http.ResponseWriter, r *http.Request) {
	form := forms.NewLoginForm()
	if formdata.MustValidateJSONRequest(r, form) {
		usermodule.WebSession.MustLogin(r, form.UID())
		render.MustJSON(w, form.Username, 200)
		return
	}
	formdata.MustRenderErrorsJSON(w, form)
})
