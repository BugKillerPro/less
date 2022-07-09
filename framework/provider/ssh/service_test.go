package ssh

import (
	"github.com/BugKillerPro/less/framework/provider/config"
	"github.com/BugKillerPro/less/framework/provider/log"
	tests "github.com/BugKillerPro/less/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestlessSSHService_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.LessConfigProvider{})
	container.Bind(&log.LessLogServiceProvider{})

	Convey("test get client", t, func() {
		LessRedis, err := NewlessSSH(container)
		So(err, ShouldBeNil)
		service, ok := LessRedis.(*lessSSH)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("ssh.web-01"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		session, err := client.NewSession()
		So(err, ShouldBeNil)
		out, err := session.Output("pwd")
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		session.Close()
	})
}
