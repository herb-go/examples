package app

import (
	"sync/atomic"

	"github.com/herb-go/herbconfig/source"
	"github.com/herb-go/usersystem-drivers/userconfig"
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
	"github.com/herb-go/util/config/tomlconfig"
)

//Members usersystem config
var Members = &userconfig.Config{}

var syncMembers atomic.Value

//StoreMembers atomically store usersystem config
func (a *appSync) StoreMembers(m *userconfig.Config) {
	syncMembers.Store(m)
}

//LoadMembers atomically load usersystem config
func (a *appSync) LoadMembers() *userconfig.Config {
	v := syncMembers.Load()
	if v == nil {
		return nil
	}
	return v.(*userconfig.Config)
}

func init() {
	config.RegisterLoader(util.ConfigFile("/members.toml"), func(configpath source.Source) {
		util.Must(tomlconfig.Load(configpath, Members))
		Sync.StoreMembers(Members)
	})
}
