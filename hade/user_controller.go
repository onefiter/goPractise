package main

import "github.com/goPractise/hade/framework"

func UserLoginController(c *framework.Context) error {
	c.Json("ok, UserLoginController")
	return nil
}
