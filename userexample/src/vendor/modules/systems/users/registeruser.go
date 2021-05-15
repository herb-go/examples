package users

import (
	"modules/members"

	"github.com/herb-go/user/status"

	"github.com/herb-go/user"

	"github.com/herb-go/uniqueid"
	"github.com/herb-go/user/profile"
	"github.com/herb-go/usersystem/usercreate"
)

func RegisterUser(account string, password string, name string, company string) {
	var uid string
	uid = uniqueid.DefaultGenerator.MustGenerateID()

	defer func() {
		if r := recover(); r != nil {
			usercreate.MustExecRemove(members.User, uid)
			panic(r)
		}
	}()
	acc := user.NewAccount()
	acc.Account = account
	acc.Keyword = "account"

	usercreate.MustExecCreate(members.User, uid)

	members.Account.MustBindAccount(uid, acc)
	members.Password.MustUpdatePassword(uid, password)
	p := profile.NewProfile().With("name", name).With("company", company)
	members.Profile.MustUpdateProfile(nil, uid, p)
	members.Status.MustUpdateStatus(uid, status.StatusNormal)
}
