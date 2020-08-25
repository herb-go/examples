package users

import (
	"regexp"

	"github.com/herb-go/util"
)

//ModuleName module name
const ModuleName = "900systems.users"

var AccountRegexp = regexp.MustCompile("^(\\w){6,18}$")

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		util.InitOrderByName(ModuleName)
	})
}
