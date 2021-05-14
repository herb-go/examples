package members

import (
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/httpusersystem/services/websession"
	"github.com/herb-go/usersystem/modules/activesessions"
	"github.com/herb-go/usersystem/modules/sessionpayload"
	"github.com/herb-go/usersystem/modules/useraccount"
	"github.com/herb-go/usersystem/modules/userpassword"
	"github.com/herb-go/usersystem/modules/userprofile"
	"github.com/herb-go/usersystem/modules/userrole"
	"github.com/herb-go/usersystem/modules/userstatus"
	"github.com/herb-go/usersystem/modules/userterm"
	"github.com/herb-go/util"

	"modules/app"
)

//ModuleName module name
const ModuleName = "800members"

//User members user system  module.
var User = usersystem.New().WithKeyword("Members")

//Status user status module.Comment if not use.
var Status = userstatus.MustNewAndInstallTo(User)

//Account user account module.Comment if not use.
var Account = useraccount.MustNewAndInstallTo(User)

//Password  user password module.Comment if not use.
var Password = userpassword.MustNewAndInstallTo(User)

//Role user role module.Comment if not use.
var Role = userrole.MustNewAndInstallTo(User)

//Term user session term module.Comment if not use.
var Term = userterm.MustNewAndInstallTo(User)

//Profile user profile module.Comment if not use.
var Profile = userprofile.MustNewAndInstallTo(User)

//WebSession user web session module.Comment if not use.
var WebSession = websession.MustNewAndInstallTo(User)

//ActiveSessions user active sessions module.Comment if not used.
var ActiveSessions = activesessions.MustNewAndInstallTo(User)

//Payload user session payload module.Coment if not used
var Payload = sessionpayload.MustNewAndInstallTo(User)

func init() {
	herbsystem.MustReady(User)
	util.RegisterModule(ModuleName, func() {
		herbsystem.MustConfigure(User)
		util.Must(app.Members.ApplyTo(User))
		herbsystem.MustStart(User)
		util.OnQuit(func() {
			herbsystem.MustStop(User)
		})
	})
}
