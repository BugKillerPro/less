package command

import (
	"fmt"
	"github.com/BugKillerPro/less/framework/cobra"
	"github.com/BugKillerPro/less/framework/contract"
)

// helpCommand show current envionment
var DemoCommand = &cobra.Command{
	Use:   "demo",
	Short: "demo for framework",
	Run: func(c *cobra.Command, args []string) {
		container := c.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		fmt.Println("app base folder:", appService.BaseFolder())
	},
}
