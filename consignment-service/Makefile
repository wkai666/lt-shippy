build:
	protoc -I. --go_out=. --go_micro=. proto/consignment/consignment.proto

run:
	docker run -p 50051:50051 \
	  -e MICRO_SERVER_ADDRESS=:50051 \
	  -e MICRO_REGISTRY=mdns \
	  consignment-service
