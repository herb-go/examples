package forms

import (
	usermodule "modules/members"
	"net/http"
	"strings"

	"github.com/herb-go/herbsystem"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
	"github.com/herb-go/user"
	"github.com/herb-go/usersystem-drivers/userform"
	"github.com/herb-go/util/form/commonform"
)

//AccountKeyword account keyword used to login.
const AccountKeyword = "account"

//LoginFormFieldLabels form field labels map.
var LoginFormFieldLabels = map[string]string{
	"Username": "Username",
	"Password": "Password",
}

//LoginForm form struct for login
type LoginForm struct {
	formdata.Form
	Username string
	Password string
	uid      string
}

//LoginFormID form id of  login form
const LoginFormID = "formmembers.login"

//NewLoginForm create new login form
func NewLoginForm() *LoginForm {
	form := &LoginForm{}
	return form
}

//ComponentID return form component id.
func (f *LoginForm) ComponentID() string {
	return LoginFormID
}

//UID return form uid
func (f *LoginForm) UID() string {
	return f.uid
}

//ValidateUserStatus validate user status
func (f *LoginForm) ValidateUserStatus() error {
	var a bool
	err := herbsystem.Catch(func() {
		a = usermodule.Status.MustIsUserAvaliable(f.uid)
	})
	if err == user.ErrUserNotExists {
		f.ValidateFieldMessagef(true, "Username", userform.MsgIncorrectUsernameOrPassword)
		return nil
	}
	if err != nil {
		return err
	}
	if f.HasError() {
		return nil
	}
	f.ValidateFieldMessagef(a, "Username", userform.MsgUserNotAvailable)
	return nil

}

//GetUserID get user id from username field
//Return user id and any error if raised
func (f *LoginForm) GetUserID() (string, error) {
	account := user.NewAccount()
	account.Keyword = AccountKeyword
	account.Account = strings.ToLower(f.Username)
	uid := usermodule.Account.MustAccountToUID(account)
	f.uid = uid
	return f.uid, nil
}

//Validate Validate form and return any error if raised.
func (f *LoginForm) Validate() error {
	var uid string
	var err error
	f.Username = strings.TrimSpace(f.Username)
	commonform.ValidateRequiredString(f, f.Username, "Username")
	commonform.ValidateRequiredString(f, f.Password, "Password")
	if !f.HasError() {
		uid, err = f.GetUserID()
		if err != nil {
			return err
		}
		f.ValidateFieldMessagef(uid != "", "Username", userform.MsgIncorrectUsernameOrPassword)
	}
	if !f.HasError() {
		result, err := usermodule.Password.VerifyPassword(uid, f.Password)
		if err != nil {
			return err
		}
		f.ValidateFieldMessagef(result == true, "Username", userform.MsgIncorrectUsernameOrPassword)
	}
	if !f.HasError() {
		err = f.ValidateUserStatus()
		if err != nil {
			return err
		}
	}
	return nil
}

//InitWithRequest init  login form  with http request.
func (f *LoginForm) InitWithRequest(r *http.Request) error {
	//Put your request form code here.
	//such as get current user id or ip address.

	//Set form labels with translated messages
	f.SetComponentLabels(ui.GetMessages(f.Lang(), "herbgo.usersystem").Collection(LoginFormFieldLabels))

	return nil
}
