build:
	GOOS=linux GOARCH=amd64 go build 
	docker build -t vessel-service .
run:
	docker run -p 50051:50051 \
	  -e MICRO_SERVER_ADDRESS=:50051 \
	  -e MICRO_REGISTRY=mdns \
	  vessel-service
