package config

import (
	"path/filepath"
	"testing"

	"github.com/BugKillerPro/less/framework/contract"
	tests "github.com/BugKillerPro/less/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestlessConfig_GetInt(t *testing.T) {
	container := tests.InitBaseContainer()

	Convey("test less env normal case", t, func() {
		appService := container.MustMake(contract.AppKey).(contract.App)
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		folder := filepath.Join(appService.ConfigFolder(), envService.AppEnv())

		serv, err := NewlessConfig(container, folder, map[string]string{})
		So(err, ShouldBeNil)
		conf := serv.(*lessConfig)
		timeout := conf.GetString("database.default.timeout")
		So(timeout, ShouldEqual, "10s")
	})
}
