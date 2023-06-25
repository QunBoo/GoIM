# GoIM
Go+Gin+websocket+Redis实现的IM即时通讯系统

#Docker build, docker run
docker build -t go-im
docker run --link redis-test:redis -e HOST_IP=10.104.9.173 -p 8080:8080 -p 8089:8089 -p 9001:9001 go-im    
