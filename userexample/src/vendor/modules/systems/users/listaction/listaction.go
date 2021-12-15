package listaction

import (
	"modules/members"
	"net/http"
	"sort"

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
	var q = r.URL.Query()
	var last = q.Get("last")
	var rev = q.Get("rev") != ""
	var limit = 10
	users := members.Status.Service.MustListUsersByStatus(last, limit+1, rev, status.StatusBanned, status.StatusNormal)
	sort.Strings(users)

	results := []*Result{}

	pfs := members.Profile.MustLoadProfiles(users...)
	for _, v := range users {
		acc := members.Account.MustAccounts(v)
		st, _ := members.Status.MustLoadStatus(v)

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
	var iter string
	if len(results) > limit {
		iter = results[limit-1].ID
		results = results[:limit]
	}
	var data = map[string]interface{}{
		"Items": results,
		"Iter":  iter,
	}
	render.MustJSON(w, data, 200)
})
