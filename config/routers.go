package config

import "uus/controller"

var Routers map[string]interface{} = map[string]interface{}{
	"index":  &controller.IndexController{},
	"login":  &controller.LoginController{},
	"update": &controller.UpdateController{},
}
