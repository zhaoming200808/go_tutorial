package main

import "github.com/astaxie/beego"

func main() {
	beego.AdminHttpPort = 8901
	beego.Run()
}
