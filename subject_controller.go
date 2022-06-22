package main

import (
	
	"fmt"
	"github.com/BugKillerPro/less/fk"
)

func SubjectAddController(c *fk.Context) error {
	c.SetOkStatus().Json("ok, SubjectAddController")
	return nil
}

func SubjectListController(c *fk.Context) error {
	c.SetOkStatus().Json("ok, SubjectListController")
	return nil
}

func SubjectDelController(c *fk.Context) error {
	c.SetOkStatus().Json("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *fk.Context) error {
	c.SetOkStatus().Json("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *fk.Context) error {
	subjectId, _ := c.ParamInt("id", 0)
	c.SetOkStatus().Json("ok, SubjectGetController:" + fmt.Sprint(subjectId))

	return nil
}

func SubjectNameController(c *fk.Context) error {
	c.SetOkStatus().Json("ok, SubjectNameController")
	return nil
}
