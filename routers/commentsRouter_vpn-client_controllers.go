package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["vpn-client/controllers:IpController"] = append(beego.GlobalControllerRouter["vpn-client/controllers:IpController"],
		beego.ControllerComments{
			"GiveIpAdd",
			`/give`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["vpn-client/controllers:IpController"] = append(beego.GlobalControllerRouter["vpn-client/controllers:IpController"],
		beego.ControllerComments{
			"StartupAddComputer",
			`/startup`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["vpn-client/controllers:IpController"] = append(beego.GlobalControllerRouter["vpn-client/controllers:IpController"],
		beego.ControllerComments{
			"CheckComputer",
			`/check`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["vpn-client/controllers:IpController"] = append(beego.GlobalControllerRouter["vpn-client/controllers:IpController"],
		beego.ControllerComments{
			"Change",
			`/change`,
			[]string{"get"},
			nil})

}
