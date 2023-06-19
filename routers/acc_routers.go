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

// Websocket 服务初始化，login、心跳、ping,就是将三个函数映射到全局map中可以通过"cmd":"login"调用
func WebsocketInit() {
	websocket.Register("login", websocket.LoginController)
	websocket.Register("heartbeat", websocket.HeartbeatController)
	websocket.Register("ping", websocket.PingController)
}
