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
	session, err := members.WebSession.GetSession(id)
	if session != nil {
		_, err = members.WebSession.RevokeSession(session.RevokeCode())
		if err != nil {
			panic(err)
		}
		err = members.ActiveSessions.PurgeActiveSession(session)
		if err != nil {
			panic(err)
		}
	}
	commonaction.SuccessAction(w, r)
})
