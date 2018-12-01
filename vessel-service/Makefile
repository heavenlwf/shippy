build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shippy/vessel-service proto/vessel/vessel.proto\
	docker build -t vessel-service .

run:
	docker run -it --name vessel-service vessel-service