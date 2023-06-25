/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 12:11
 */

package home

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gowebsocket/servers/websocket"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 加载聊天页面
func Index(c *gin.Context) {

	appIdStr := c.Query("appId")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)
	if !websocket.InAppIds(appId) {
		appId = websocket.GetDefaultAppId()
	}

	fmt.Println("http_request 聊天首页", appId)
	//设定模板index.tpl所需要使用的数据，通过c.HTML传输过去
	//TODO:在实质部署到服务器的时候，这里httpUrl应该传入server的外网ip地址以及对应的port

	// serverIp, err := GetOutBoundIP()
	// if err != nil {
	// 	fmt.Println("serverIp Get Err", err)
	// }
	//Docker serverIp需要在参数中添加
	serverIp := os.Getenv("HOST_IP")
	httpPort := viper.GetString("app.httpPort")
	httpUrl_out := serverIp + ":" + httpPort
	webSocketPort := viper.GetString("app.webSocketPort")
	webSocketUrl_out := serverIp + ":" + webSocketPort
	data := gin.H{
		"title":        "聊天首页",
		"appId":        appId,
		"httpUrl":      httpUrl_out,
		"webSocketUrl": webSocketUrl_out,
	}
	// "httpUrl":      viper.GetString("app.httpUrl"),
	// 	"webSocketUrl": viper.GetString("app.webSocketUrl"),
	c.HTML(http.StatusOK, "index.tpl", data)
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
