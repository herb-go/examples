package forms

import (
	"modules/members"
	"net/http"

	"github.com/herb-go/user/profile"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/herb/ui/validator/formdata"
	"github.com/herb-go/util/form/commonform"
)

//UpdateprofileFormFieldLabels form field labels map.
var UpdateprofileFormFieldLabels = map[string]string{
	"Name":    "姓名",
	"Company": "公司",
}

//UpdateprofileForm form struct forupdateprofile
type UpdateprofileForm struct {
	formdata.Form
	uid     string
	Name    string
	Company string
}

//UpdateprofileFormID form id of formupdateprofile
const UpdateprofileFormID = "formsystems.users.updateprofile"

//NewUpdateprofileForm create newupdateprofile form
func NewUpdateprofileForm() *UpdateprofileForm {
	form := &UpdateprofileForm{}
	return form
}

//ComponentID return form component id.
func (f *UpdateprofileForm) ComponentID() string {
	return UpdateprofileFormID
}

//Validate Validate form and return any error if raised.
func (f *UpdateprofileForm) Validate() error {
	commonform.ValidateRequiredString(f, f.Name, "Name")
	commonform.ValidateRequiredString(f, f.Company, "Company")
	if !f.HasError() {
	}
	return nil
}

//Exec execwhen form validated.
func (f *UpdateprofileForm) Exec() error {
	p := profile.NewProfile()
	p.With("name", f.Name)
	p.With("company", f.Company)
	return members.Profile.UpdateProfile(nil, f.uid, p)
}

//InitWithRequest init updateprofile form  with http request.
func (f *UpdateprofileForm) InitWithRequest(r *http.Request) error {

	//Put your request form code here.
	//such as get current user id or ip address.

	//Set form labels with translated messages
	var err error
	f.uid, err = members.WebSession.IdentifyRequest(r)
	if err != nil {
		return err
	}
	f.SetComponentLabels(ui.GetMessages(f.Lang(), "app").Collection(UpdateprofileFormFieldLabels))
	return nil
}
