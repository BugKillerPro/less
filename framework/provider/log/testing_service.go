package log

import (
	"github.com/BugKillerPro/less/framework/contract"
	"github.com/BugKillerPro/less/framework/provider/log/formatter"
	"github.com/BugKillerPro/less/framework/provider/log/services"
	"os"
)

// lessTestingLog 是 Env 的具体实现
type lessTestingLog struct {
}

// NewlessTestingLog 测试日志，直接打印到控制台
func NewlessTestingLog(params ...interface{}) (interface{}, error) {
	log := &services.lessConsoleLog{}

	log.SetLevel(contract.DebugLevel)
	log.SetCtxFielder(nil)
	log.SetFormatter(formatter.TextFormatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	return log, nil
}
