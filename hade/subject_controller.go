package main

import "github.com/goPractise/hade/framework"

func SubjectAddController(c *framework.Context) error {
	c.Json("ok, SubjectAddController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.Json("ok, SubjectListController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	c.Json("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.Json("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	c.Json("ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.Json("ok, SubjectNameController")
	return nil
}
