package redis

import (
	"context"
	"github.com/BugKillerPro/less/framework/provider/config"
	"github.com/BugKillerPro/less/framework/provider/log"
	tests "github.com/BugKillerPro/less/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestlessService_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.lessConfigProvider{})
	container.Bind(&log.lessLogServiceProvider{})

	Convey("test get client", t, func() {
		lessRedis, err := NewlessRedis(container)
		So(err, ShouldBeNil)
		service, ok := lessRedis.(*lessRedis)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("redis.write"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		ctx := context.Background()
		err = client.Set(ctx, "foo", "bar", 1*time.Hour).Err()
		So(err, ShouldBeNil)
		val, err := client.Get(ctx, "foo").Result()
		So(err, ShouldBeNil)
		So(val, ShouldEqual, "bar")
		err = client.Del(ctx, "foo").Err()
		So(err, ShouldBeNil)
	})
}
