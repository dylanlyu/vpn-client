package controllers

import (
	"encoding/json"
	"os/exec"
	"time"
	"vpn-client/helpers"
	"vpn-client/structures"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

var (
	serurl = "http://" + beego.AppConfig.String("ServerAdd")
)

type IpController struct {
	beego.Controller
}

// @Title 顯示IP
// @Description 顯示IP
// @Param	body		body 	Json	flase
// @Failure 403 body is Json
// @router /give [get]
func (ip *IpController) GiveIpAdd() {
	var responses structures.ResponsesIp
	responses.Ip = helpers.GetConnectIp()
	responses.Computer = helpers.GetLocalName()
	//responses.Ip = "192.168.0.18"
	//responses.Name = "Taipei-01-03"
	responses.Status = 1
	json.Marshal(&responses)

	req := httplib.Post(serurl + "/v1/server/update?token=")
	req.Header("Content-Type", "application/json")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36"+
		"(KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
	req.JSONBody(&responses)

	str, err := req.Bytes()

	if err != nil {
		beego.Debug(err)
	}
	beego.Debug(str)

	var msg structures.Message

	json.Unmarshal(str, &msg)

	ip.Data["json"] = msg

	ip.ServeJSON()

}

// @Title 顯示IP
// @Description 顯示IP
// @Param	body		body 	Json	flase
// @Failure 403 body is Json
// @router /startup [get]
func (ip *IpController) StartupAddComputer() {
	var responses structures.ResponsesIp
	responses.Ip = helpers.GetConnectIp()
	responses.Computer = helpers.GetLocalName()
	//responses.Ip = "192.168.0.12"
	//responses.Name = "Taipei-01-03"
	responses.Status = 1
	json.Marshal(&responses)

	req := httplib.Post(serurl + "/v1/server/add?token=")
	req.Header("Content-Type", "application/json")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36"+
		"(KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
	req.JSONBody(&responses)

	str, err := req.Bytes()

	if err != nil {
		beego.Debug(err)
	}
	beego.Debug(str)

	var msg structures.Message

	json.Unmarshal(str, &msg)

	ip.Data["json"] = msg

	ip.ServeJSON()
}

// @Title 顯示IP
// @Description 顯示IP
// @Param	body		body 	Json	flase
// @Failure 403 body is Json
// @router /check [get]
func (ip *IpController) CheckComputer() {
	var responses structures.ResponsesIp
	responses.Computer = helpers.GetLocalName()
	req := httplib.Post(serurl + "/v1/server/check?token=")
	req.Header("Content-Type", "application/json")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36"+
		"(KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
	req.JSONBody(&responses)

	str, err := req.Bytes()

	if err != nil {
		beego.Debug(err)
	}
	beego.Debug(str)

	json.Unmarshal(str, &responses)

	ipadd := helpers.GetConnectIp()

	if responses.Ip != ipadd {

		req := httplib.Get("http://localhost:8080/v1/client/give")
		req.Header("Content-Type", "application/json")
		req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36"+
			"(KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
		str, err := req.Bytes()
		if err != nil {
			beego.Debug(err)
		}
		beego.Debug(str)

		var message structures.Message

		json.Unmarshal(str, &message)
		if message.Error == true {
			//beego.Error(message.Message)
		}

	}

	ip.Data["json"] = map[string]interface{}{"error": false, "message": "check success"}

	ip.ServeJSON()
}

// @Title 置換IP
// @Description 置換IP
// @Param	body		body 	Json	flase
// @Failure 403 body is Json
// @router /change [get]
func (ip *IpController) Change() {

	var change structures.ChangeIp
	json.Unmarshal(ip.Ctx.Input.RequestBody, &change)

	beego.Debug(string(ip.Ctx.Input.RequestBody))

	go func() {
		time.Sleep(5 * time.Second)
		cmd := exec.Command("poff", "dsl-provider")
		out, _ := cmd.CombinedOutput()
		beego.Debug("%s", string(out))

		time.Sleep(10 * time.Second)
		cmd = exec.Command("pon", "dsl-provider")
		out, _ = cmd.CombinedOutput()
		beego.Debug("%s", string(out))
		time.Sleep(15 * time.Second)

		req := httplib.Get("http://localhost:8080/v1/client/give")
		req.Header("Content-Type", "application/json")
		req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36"+
			"(KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
		str, err := req.Bytes()
		if err != nil {
			beego.Debug(err)
		}
		beego.Debug(str)

		var message structures.Message

		json.Unmarshal(str, &message)
		if message.Error == true {
			//beego.Error(message.Message)
		}

	}()

	ip.Data["json"] = map[string]interface{}{"error": false, "message": "wait"}

	ip.ServeJSON()

}
