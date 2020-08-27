package listaction

import (
	"modules/members"
	"net/http"

	"github.com/herb-go/user/status"

	"github.com/herb-go/herb/ui/render"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/usersystem-drivers/tomluser"
)

type Result struct {
	ID          string
	Account     string
	Name        string
	Company     string
	Status      status.Status
	StatusLabel string
}

var ActionList = action.New(func(w http.ResponseWriter, r *http.Request) {
	users := members.Status.Service.(*tomluser.Users)
	data := users.GetAllUsers()
	results := []*Result{}
	for _, v := range data.Users {
		label, err := members.Status.Service.Label(v.Status())
		if err != nil {
			panic(err)
		}
		result := &Result{
			ID:          v.UID,
			Account:     v.Accounts[0].Account,
			Name:        v.Profiles.Load("name"),
			Company:     v.Profiles.Load("company"),
			Status:      v.Status(),
			StatusLabel: label,
		}
		results = append(results, result)
	}
	render.MustJSON(w, map[string]interface{}{"Items": results}, 200)
})
