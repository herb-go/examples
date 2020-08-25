package users

import (
	"modules/members"

	"github.com/herb-go/user/status"

	"github.com/herb-go/user"

	"github.com/herb-go/user/profile"
	"github.com/herb-go/uniqueid"
	"github.com/herb-go/usersystem/usercreate"
)

func RegisterUser(account string, password string, name string, company string) error {
	var err error
	var uid string
	uid, err = uniqueid.DefaultGenerator.GenerateID()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			usercreate.ExecRemove(members.User, uid)
		}
	}()
	acc := user.NewAccount()
	acc.Account = account
	acc.Keyword = "account"

	err = usercreate.ExecCreate(members.User, uid)
	if err != nil {
		return err
	}
	err = members.Account.BindAccount(uid, acc)
	if err != nil {
		return err
	}
	err = members.Password.UpdatePassword(uid, password)
	if err != nil {
		return err
	}
	p := profile.NewProfile().With("name", name).With("company", company)
	err = members.Profile.UpdateProfile(nil, uid, p)
	if err != nil {
		return err
	}
	err = members.Status.UpdateStatus(uid, status.StatusNormal)
	if err != nil {
		return err
	}
	return nil
}
