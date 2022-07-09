package id

import (
	tests "github.com/BugKillerPro/less/test"
	"testing"

	"github.com/BugKillerPro/less/framework/contract"
	"github.com/BugKillerPro/less/framework/provider/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConsoleLog_Normal(t *testing.T) {
	Convey("test less console log normal case", t, func() {
		c := tests.InitBaseContainer()
		c.Bind(&config.lessConfigProvider{})

		err := c.Bind(&lessIDProvider{})
		So(err, ShouldBeNil)

		idService := c.MustMake(contract.IDKey).(contract.IDService)
		xid := idService.NewID()
		So(xid, ShouldNotBeEmpty)
	})
}
