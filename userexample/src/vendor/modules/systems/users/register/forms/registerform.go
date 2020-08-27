package forms

import (
	"modules/members"
	"modules/systems/users"
	"net/http"
	"strings"

	"github.com/herb-go/user"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
	"github.com/herb-go/util/form/commonform"
)

//RegisterFormFieldLabels form field labels map.
var RegisterFormFieldLabels = map[string]string{
	"Account":        "Account",
	"Name":           "Name",
	"Company":        "Company",
	"Password":       "Password",
	"RepeatPassword": "RepeatPassword",
}

//RegisterForm form struct forregister
type RegisterForm struct {
	formdata.Form
	Account        string
	Password       string
	RepeatPassword string
	Name           string
	Company        string
}

//RegisterFormID form id of formregister
const RegisterFormID = "formsystems.users.register"

//NewRegisterForm create newregister form
func NewRegisterForm() *RegisterForm {
	form := &RegisterForm{}
	return form
}

//ComponentID return form component id.
func (f *RegisterForm) ComponentID() string {
	return RegisterFormID
}

//Validate Validate form and return any error if raised.
func (f *RegisterForm) Validate() error {
	commonform.ValidateRequiredString(f, f.Account, "Account")
	commonform.ValidateRequiredString(f, f.Name, "Name")
	commonform.ValidateRequiredString(f, f.Company, "Company")
	commonform.ValidateRequiredString(f, f.Password, "Password")
	commonform.ValidateRequiredString(f, f.RepeatPassword, "RepeatPasword")
	if !f.HasError() {
		commonform.ValidateStringLength(f, f.Password, "Password", 6, 32)
	}
	if !f.HasError() {
		f.ValidateFieldf(f.Password == f.RepeatPassword, "RepeatPassword", "重复密码不匹配")
	}
	if !f.HasError() {
		f.ValidateFieldf(users.AccountRegexp.MatchString(f.Account), "Account", "帐号格式错误")
	}
	if !f.HasError() {
		a := user.NewAccount()
		a.Keyword = "account"
		a.Account = f.Account
		uid, err := members.Account.AccountToUID(a)
		if err != nil {
			return err
		}
		f.ValidateFieldf(uid == "", "Account", "帐号已存在")
	}
	return nil
}

//Exec execwhen form validated.
func (f *RegisterForm) Exec() error {
	return users.RegisterUser(f.Account, f.Password, f.Name, f.Company)
}

//InitWithRequest init register form  with http request.
func (f *RegisterForm) InitWithRequest(r *http.Request) error {

	//Put your request form code here.
	//such as get current user id or ip address.

	//Set form labels with translated messages
	f.SetComponentLabels(ui.GetMessages(f.Lang(), "app").Collection(RegisterFormFieldLabels))
	f.Account = strings.ToLower(f.Account)
	return nil
}
