package forms

import (
	"modules/members"
	"net/http"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
	"github.com/herb-go/util/form/commonform"
)

//UpdatepasswordFormFieldLabels form field labels map.
var UpdatepasswordFormFieldLabels = map[string]string{
	"Passsword":      "Passsword",
	"NewPassword":    "NewPassword",
	"RepeatPassword": "RepeatPassword",
}

//UpdatepasswordForm form struct forupdatepassword
type UpdatepasswordForm struct {
	formdata.Form
	uid            string
	Password       string
	NewPassword    string
	RepeatPassword string
}

//UpdatepasswordFormID form id of formupdatepassword
const UpdatepasswordFormID = "formsystems.users.updatepassword"

//NewUpdatepasswordForm create newupdatepassword form
func NewUpdatepasswordForm() *UpdatepasswordForm {
	form := &UpdatepasswordForm{}
	return form
}

//ComponentID return form component id.
func (f *UpdatepasswordForm) ComponentID() string {
	return UpdatepasswordFormID
}

//Validate Validate form and return any error if raised.
func (f *UpdatepasswordForm) Validate() error {
	commonform.ValidateRequiredString(f, f.Password, "Passsword")
	commonform.ValidateRequiredString(f, f.NewPassword, "NewPassword")
	commonform.ValidateRequiredString(f, f.RepeatPassword, "RepeatPassword")
	if !f.HasError() {
		commonform.ValidateStringLength(f, f.NewPassword, "NewPassword", 6, 32)
	}
	if !f.HasError() {
		f.ValidateFieldf(f.NewPassword == f.RepeatPassword, "Repeatpassword", "重复密码不匹配")
	}
	if !f.HasError() {
		ok := members.Password.MustVerifyPassword(f.uid, f.Password)
		f.ValidateFieldf(ok, "Password", "密码错误")
	}
	return nil
}

//Exec execwhen form validated.
func (f *UpdatepasswordForm) Exec() error {
	members.Password.MustUpdatePassword(f.uid, f.NewPassword)
	members.Term.MustStartNewTerm(f.uid)
	return nil
}

//InitWithRequest init updatepassword form  with http request.
func (f *UpdatepasswordForm) InitWithRequest(r *http.Request) error {

	//Put your request form code here.
	//such as get current user id or ip address.

	//Set form labels with translated messages
	var err error
	f.uid, err = members.WebSession.IdentifyRequest(r)
	if err != nil {
		return err
	}
	f.SetComponentLabels(ui.GetMessages(f.Lang(), "app").Collection(UpdatepasswordFormFieldLabels))
	return nil
}
