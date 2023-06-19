package websocket

import (
	"fmt"
	"net/http"
	"time"

	"gowebsocket/helper"
	"gowebsocket/models"

	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

const (
	defaultAppId = 101 // 默认平台Id
)

var (
	clientManager = NewClientManager()                    // 管理者
	appIds        = []uint32{defaultAppId, 102, 103, 104} // 全部的平台

	serverIp   string
	serverPort string
)

func GetAppIds() []uint32 {

	return appIds
}

// 获取本机Ip和port
func GetServer() (server *models.Server) {
	server = models.NewServer(serverIp, serverPort)

	return
}

// 对比参数的ip地址和初始化时设定的本机ip地址，鉴别是否是本机
func IsLocal(server *models.Server) (isLocal bool) {
	if server.Ip == serverIp && server.Port == serverPort {
		isLocal = true
	}

	return
}

func InAppIds(appId uint32) (inAppId bool) {

	for _, value := range appIds {
		if value == appId {
			inAppId = true

			return
		}
	}

	return
}

func GetDefaultAppId() (appId uint32) {
	appId = defaultAppId

	return
}

// 启动websocket服务器程序，设定serverPort和webSocketPort
func StartWebSocket() {

	serverIp = helper.GetServerIp()

	webSocketPort := viper.GetString("app.webSocketPort") //8089
	rpcPort := viper.GetString("app.rpcPort")

	serverPort = rpcPort //9001

	http.HandleFunc("/acc", wsPage)

	// 添加处理程序
	go clientManager.start()
	fmt.Println("WebSocket 启动程序成功", serverIp, serverPort)

	http.ListenAndServe(":"+webSocketPort, nil)
}

// 升级协议，在tpl文件中访问对应的链路实现函数调用
func wsPage(w http.ResponseWriter, req *http.Request) {

	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])

		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)

		return
	}

	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)

	go client.read()
	go client.write()

	// 用户连接事件
	clientManager.Register <- client
}
