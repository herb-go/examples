package revokeaction

import (
	"modules/members"
	"net/http"

	"github.com/herb-go/util/action/commonaction"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/herb/middleware/router"
)

var ActionRevoke = action.New(func(w http.ResponseWriter, r *http.Request) {
	id := router.GetParams(r).Get("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	session := members.WebSession.MustGetSession(id)
	if session != nil {
		members.WebSession.MustRevokeSession(session.RevokeCode())
		members.ActiveSessions.MustPurgeActiveSession(session)
	}
	commonaction.SuccessAction(w, r)
})
