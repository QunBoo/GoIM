# GoIM
Go+Gin+websocket+Redis实现的IM即时通讯系统

## Docker build, docker run
docker build -t go-im  
docker run -d --link redis-test:redis -e HOST_IP=[`YOUR IP ADDRESS`] -p 8080:8080 -p 8089:8089 -p 9001:9001 go-im    

## TODO
//TODO：代码重构，实现逻辑上端到端传输  
//TODO: 性能测试，  
//TODO：增加单对单聊天  
//TODO: 实现文件传输  
//TODO：增加用户登录、token鉴权，MySQL  