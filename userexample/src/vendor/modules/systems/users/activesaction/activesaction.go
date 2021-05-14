package activesaction

import (
	"modules/members"
	"net/http"
	"strconv"

	"github.com/herb-go/usersystem/modules/activesessions"

	"github.com/herb-go/usersystem-drivers/commonpayload"

	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/herb/middleware/action"
	"github.com/herb-go/herb/ui/render"
)

type Result struct {
	LastActive int64
	LoginTime  int64
	RevokeCode string
	ID         string
	SessionID  string
	Ip         string
}

var ActionActives = action.New(func(w http.ResponseWriter, r *http.Request) {
	session := members.WebSession.MustGetRequestSession(r)

	uid := session.UID()
	actives := members.ActiveSessions.MustGetActiveSessions("web", uid)

	result := make([]*Result, 0, len(actives))
	for k := range actives {
		session := usersession.MustExecGetSession(members.User, "web", actives[k].SessionID)
		if session == nil {
			continue
		}
		lt, err := strconv.ParseInt(session.Payloads.LoadString(commonpayload.PayloadNameLogintime), 10, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, &Result{
			LoginTime:  lt,
			LastActive: actives[k].LastActive.Unix(),
			SessionID:  actives[k].SessionID,
			Ip:         session.Payloads.LoadString(commonpayload.PayloadNameHTTPIp),
			ID:         session.Payloads.LoadString(activesessions.PayloadSerialNumber),
			RevokeCode: session.RevokeCode(),
		})
	}
	render.MustJSON(w, map[string]interface{}{"Items": result}, 200)
})
