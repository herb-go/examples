package activesaction

import (
	"modules/members"
	"net/http"

	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/herbsecurity/authority"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/herb/ui/render"
)

type Result struct {
	LastActive int64
	Payloads   *authority.Payloads
}

var ActionActives = action.New(func(w http.ResponseWriter, r *http.Request) {
	session, err := members.WebSession.GetRequestSession(r)
	if err != nil {
		panic(err)
	}
	uid := session.UID()
	actives, err := members.ActiveSessions.GetActiveSessions("web", uid)
	if err != nil {
		panic(err)
	}
	result := make([]*Result, 0, len(actives))
	for k := range actives {
		session, err := usersession.ExecGetSession(members.User, "web", actives[k].SessionID)
		if err != nil {
			panic(err)
		}
		if session == nil {
			continue
		}
		result = append(result, &Result{Payloads: session.Payloads, LastActive: actives[k].LastActive.Unix()})
	}
	render.MustJSON(w, result, 200)
})
