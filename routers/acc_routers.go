/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 16:02
 */

package routers

import (
	"gowebsocket/servers/websocket"
)

// Websocket 路由初始化，login、心跳、ping
func WebsocketInit() {
	websocket.Register("login", websocket.LoginController)
	websocket.Register("heartbeat", websocket.HeartbeatController)
	websocket.Register("ping", websocket.PingController)
}
