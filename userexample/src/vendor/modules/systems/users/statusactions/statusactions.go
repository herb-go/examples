package statusactions

import (
	"modules/members"
	"net/http"

	"github.com/herb-go/user/status"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/herb/middleware/router"
	"github.com/herb-go/util/action/commonaction"
)

var ActionEnable = action.New(func(w http.ResponseWriter, r *http.Request) {
	id := router.GetParams(r).Get("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	members.Status.MustUpdateStatus(id, status.StatusNormal)
	commonaction.SuccessAction(w, r)
})

var ActionDisable = action.New(func(w http.ResponseWriter, r *http.Request) {
	id := router.GetParams(r).Get("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	members.Status.MustUpdateStatus(id, status.StatusBanned)
	commonaction.SuccessAction(w, r)
})
