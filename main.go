package main

import (
	_ "vpn-client/docs"
	_ "vpn-client/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	var FilterToken = func(ctx *context.Context) {

		if ctx.Request.URL.Query()["token"] == nil {
			beego.Info("URL=error")
			beego.Error("URL:" + ctx.Input.URL() + "  " + ctx.Input.IP())
			ctx.Redirect(302, "/")
			return
		}

		key := ""

		if ctx.Request.URL.Query()["token"][0] != key {
			beego.Info("Token=error")
			beego.Error("URL:" + ctx.Input.URL() + "  " + ctx.Input.IP())
			ctx.Redirect(302, "/")
			return
		}

	}

	beego.InsertFilter("/v1/client/change", beego.BeforeRouter, FilterToken)
	beego.Run()
}
