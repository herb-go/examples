package members

import (
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/httpusersystem/services/websession"
	"github.com/herb-go/usersystem/services/activesessions"
	"github.com/herb-go/usersystem/services/useraccount"
	"github.com/herb-go/usersystem/services/userpassword"
	"github.com/herb-go/usersystem/services/userprofile"
	"github.com/herb-go/usersystem/services/userrole"
	"github.com/herb-go/usersystem/services/userstatus"
	"github.com/herb-go/usersystem/services/userterm"
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

func init() {
	util.Must(User.Ready())
	util.RegisterModule(ModuleName, func() {
		util.Must(User.Configuring())
		util.Must(app.Members.ApplyTo(User))
		util.Must(User.Start())
		util.OnQuitAndLogError(User.Stop)
	})
}
