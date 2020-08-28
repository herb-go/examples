package revokeaction

import (
	"modules/members"
	"net/http"

	"github.com/herb-go/util/action/commonaction"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/herb/middleware/router"
)

var ActionRevoke = action.New(func(w http.ResponseWriter, r *http.Request) {
	code := router.GetParams(r).Get("id")
	if code == "" {
		http.NotFound(w, r)
		return
	}
	session, err := members.WebSession.GetSession
	_, err := members.WebSession.RevokeSession(code)
	if err != nil {
		panic(err)
	}
	members.ActiveSessions.PurgeActiveSession(members.WebSession)
	commonaction.SuccessAction(w, r)
})
