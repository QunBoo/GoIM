package main

//TODO：代码重构，实现逻辑上端到端传输
//TODO: 性能测试，
//TODO：增加单对单聊天
//TODO: 实现文件传输
//TODO：增加用户登录、token鉴权，MySQL

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"gowebsocket/lib/redislib"
	"gowebsocket/routers"
	"gowebsocket/servers/grpcserver"
	"gowebsocket/servers/task"
	"gowebsocket/servers/websocket"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//读取配置文件
	initConfig()
	//初始化日志
	initFile()
	//初始化Redis
	initRedis()

	router := gin.Default()
	// 初始化路由
	routers.Init(router)
	routers.WebsocketInit()

	// 定时清理连接任务
	task.Init()
	// 服务器注册
	task.ServerInit()

	go websocket.StartWebSocket()
	// grpc
	go grpcserver.Init()
	// fmt.Println("Current number of goroutines AFTER grpcserver.Init(): ", runtime.NumGoroutine())   //5个goroutine
	//打开网页
	// go open()

	httpPort := viper.GetString("app.httpPort")
	http.ListenAndServe(":"+httpPort, router)
	// fmt.Println("Current number of goroutines AFTER ListenAndServe: ", runtime.NumGoroutine())

}

// 初始化日志
func initFile() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	logFile := viper.GetString("app.logFile")
	f, _ := os.Create(logFile)
	gin.DefaultWriter = io.MultiWriter(f)
}

func initConfig() {
	viper.SetConfigName("config/app")
	viper.AddConfigPath(".") // 添加搜索路径

	err := viper.ReadInConfig()
	if err != nil {
		errors := fmt.Errorf("fatal error config file: %s ", err)
		panic(errors)
	}

	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config redis:", viper.Get("redis"))

}

func initRedis() {
	redislib.ExampleNewClient()
}

/*
// 打开网页
func open() {

	time.Sleep(1000 * time.Millisecond)

	httpUrl := viper.GetString("app.httpUrl")
	httpUrl = "http://" + httpUrl + "/home/index"

	fmt.Println("访问页面体验:", httpUrl)

	cmd := exec.Command("open", httpUrl)
	cmd.Output()
}
*/
