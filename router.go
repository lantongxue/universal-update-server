package main

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine            *gin.Engine
	defaultController string
	defaultAction     string
}

func NewRouter(r *gin.Engine) *Router {
	return &Router{
		engine:            r,
		defaultController: "Index",
		defaultAction:     "Index",
	}
}

func (r *Router) AutoRouter(v interface{}) {
	reflectVal := reflect.TypeOf(v)
	obj := reflect.ValueOf(v).Interface()
	ctlType := reflectVal.Elem()
	controllerFullName := ctlType.Name()
	controllerName := strings.TrimSuffix(controllerFullName, "Controller")
	controllerLayer := strings.TrimPrefix(ctlType.PkgPath(), "uus/controller")
	if controllerLayer == "" {
		controllerLayer = "/"
	}

	for i := 0; i < reflectVal.NumMethod(); i++ {
		method := reflectVal.Method(i)
		// get method parameters count
		// the http handler has two parameters
		// number 0 is Controller object
		if method.Type.NumIn() != 2 {
			continue
		}
		// number 1 is gin.Context
		if method.Type.In(1).Elem().String() != "gin.Context" {
			continue
		}
		ar := strings.Split(method.Name, "__")
		methodName := ar[0]
		requestMethod := "Any"
		if len(ar) == 2 {
			requestMethod = ar[1]
		}

		requestPath := "/"
		if controllerName == r.defaultController && methodName == r.defaultAction {
			requestPath = controllerLayer
		}

		if controllerName == r.defaultController && methodName != r.defaultAction {
			requestPath = controllerLayer + "/" + strings.ToLower(methodName)
		}

		if controllerName != r.defaultController && methodName == r.defaultAction {
			requestPath = controllerLayer + "/" + strings.ToLower(controllerName)
		}

		if controllerName != r.defaultController && methodName != r.defaultAction {
			requestPath = controllerLayer + "/" + strings.ToLower(controllerName) + "/" + strings.ToLower(methodName)
		}

		switch strings.ToUpper(requestMethod) {
		case "GET":
			r.engine.GET(requestPath, func(ctx *gin.Context) {
				params := make([]reflect.Value, 2)
				params[0] = reflect.ValueOf(obj)
				params[1] = reflect.ValueOf(ctx)
				method.Func.Call(params)
			})
		case "POST":
			r.engine.POST(requestPath, func(ctx *gin.Context) {
				params := make([]reflect.Value, 2)
				params[0] = reflect.ValueOf(obj)
				params[1] = reflect.ValueOf(ctx)
				method.Func.Call(params)
			})
		case "PUT":
			r.engine.PUT(requestPath, func(ctx *gin.Context) {
				params := make([]reflect.Value, 2)
				params[0] = reflect.ValueOf(obj)
				params[1] = reflect.ValueOf(ctx)
				method.Func.Call(params)
			})
		case "PATCH":
			r.engine.PATCH(requestPath, func(ctx *gin.Context) {
				params := make([]reflect.Value, 2)
				params[0] = reflect.ValueOf(obj)
				params[1] = reflect.ValueOf(ctx)
				method.Func.Call(params)
			})
		case "HEAD":
			r.engine.HEAD(requestPath, func(ctx *gin.Context) {
				params := make([]reflect.Value, 2)
				params[0] = reflect.ValueOf(obj)
				params[1] = reflect.ValueOf(ctx)
				method.Func.Call(params)
			})
		case "OPTIONS":
			r.engine.OPTIONS(requestPath, func(ctx *gin.Context) {
				params := make([]reflect.Value, 2)
				params[0] = reflect.ValueOf(obj)
				params[1] = reflect.ValueOf(ctx)
				method.Func.Call(params)
			})
		case "DELETE":
			r.engine.DELETE(requestPath, func(ctx *gin.Context) {
				params := make([]reflect.Value, 2)
				params[0] = reflect.ValueOf(obj)
				params[1] = reflect.ValueOf(ctx)
				method.Func.Call(params)
			})
		default:
			r.engine.Any(requestPath, func(ctx *gin.Context) {
				params := make([]reflect.Value, 2)
				params[0] = reflect.ValueOf(obj)
				params[1] = reflect.ValueOf(ctx)
				method.Func.Call(params)
			})
		}
	}
}
