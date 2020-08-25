package listaction

import (
	"modules/members"
	"net/http"

	"github.com/herb-go/herb/ui/render"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/usersystem-drivers/tomluser"
)

var ActionList = action.New(func(w http.ResponseWriter, r *http.Request) {
	users := members.Status.Service.(*tomluser.Users)
	data := users.GetAllUsers()
	render.MustJSON(w, data, 200)
})
