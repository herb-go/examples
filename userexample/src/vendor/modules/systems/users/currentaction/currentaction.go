package currentaction

import (
	"modules/members"
	"net/http"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/herb/ui/render"
	"github.com/herb-go/user/profile"
)

var ActionCurrent = action.New(func(w http.ResponseWriter, r *http.Request) {
	s, err := members.WebSession.GetRequestSession(r)
	if err != nil {
		panic(err)
	}
	var uid string
	if s != nil {
		uid = s.UID()
	}
	var p *profile.Profile
	if s != nil && uid != "" {
		p, err = members.Profile.LoadProfile(uid)
		if err != nil {
			panic(err)
		}
	}
	var acc string
	if s != nil && uid != "" {
		a, err := members.Account.Accounts(uid)
		if err != nil {
			panic(err)
		}
		acc = (*a)[0].Account
	}
	render.MustJSON(w, map[string]interface{}{"UID": uid, "Name": p.Load("name"), "Company": p.Load("company"), "Account": acc}, 200)
})
