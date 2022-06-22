package middleware

import (
	"github.com/BugKillerPro/less/fk"
	"fmt"
)

func Test1() fk.Handler {
	// 使用函数回调
	return func(c *fk.Context) error {
		fmt.Println("middleware pre test1")
		c.Next()
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() fk.Handler {
	// 使用函数回调
	return func(c *fk.Context) error {
		fmt.Println("middleware pre test2")
		c.Next()
		fmt.Println("middleware post test2")
		return nil
	}
}

func Test3() fk.Handler {
	// 使用函数回调
	return func(c *fk.Context) error {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware post test3")
		return nil
	}
}
