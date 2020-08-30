package listaction

import (
	"modules/members"
	"net/http"

	"github.com/herb-go/usersystem/userdataset"

	"github.com/herb-go/user/status"

	"github.com/herb-go/herb/ui/render"

	"github.com/herb-go/herb/middleware/action"
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
	users, err := members.Status.Service.ListUsersByStatus("", 0, status.StatusBanned, status.StatusNormal)
	if err != nil {
		panic(err)
	}
	results := []*Result{}
	ds, err := userdataset.ExecNewDataset(members.User)
	if err != nil {
		panic(err)
	}
	pfs, err := members.Profile.LoadProfiles(ds, false, users...)
	for _, v := range users {
		acc, err := members.Account.Account(v)
		if err != nil {
			panic(err)
		}
		st, err := members.Status.LoadStatus(v)
		if err != nil {
			panic(err)
		}
		label, err := members.Status.Service.Label(st)
		if err != nil {
			panic(err)
		}
		var name string
		var company string
		pf := pfs[v]
		if pf != nil {
			name = pf.Load("name")
			company = pf.Load("company")
		}
		result := &Result{
			ID:          v,
			Account:     (*acc)[0].Account,
			Name:        name,
			Company:     company,
			Status:      st,
			StatusLabel: label,
		}
		results = append(results, result)
	}
	render.MustJSON(w, map[string]interface{}{"Items": results}, 200)
})
