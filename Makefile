.PHONY: \
    all \
	build_frontend \
    build_backend \
    build_server_side \
    deploy_server_side \
    deploy_device

all:
	echo "Usage 'make build_server_side' or 'make deploy_server_side' or 'make deploy_device'"

build_frontend:
	cd server/frontend &&\
	yarn install &&\
	yarn build

build_backend:
	sudo rm -Rf server_side_deploy &&\
	glide install &&\
	go build -v cmd/server/ws-service.go &&\
	go build -v cmd/server/mqtt-service.go

build_server_side: build_frontend build_backend
	rm -Rf server_side_deploy &&\
	mkdir -p server_side_deploy &&\
	mv server/frontend/dist server_side_deploy/public &&\
	mv ws-service server_side_deploy/ &&\
	mv mqtt-service server_side_deploy/ &&\
	cp server/Dockerfile.ws server_side_deploy/ &&\
	cp server/Dockerfile.mqtt server_side_deploy/ &&\
	cp server/Dockerfile.nginx server_side_deploy/ &&\
	cp server/services/nginx/nginx.conf server_side_deploy &&\
	cp server/docker-compose.yaml server_side_deploy/ &&\
	cp server/.env server_side_deploy/

deploy_server_side: build_server_side
	ansible-playbook -i ansible/hosts ansible/deploy.yaml

deploy_device:
	rm -Rf device_deploy &&\
	mkdir -p device_deploy &&\
	glide install &&\
	export GOARCH=arm &&\
	export GOARM=7 &&\
    go build -v cmd/device/cloud-service.go &&\
    go build -v cmd/device/telemetry-service.go
	mv cloud-service device_deploy/ &&\
	mv telemetry-service device_deploy/ &&\
	cp cmd/device/nats device_deploy/ &&\
    cp cmd/device/cloud.sh device_deploy/ &&\
    cp cmd/device/telemetry.sh device_deploy/ &&\
	cp cmd/device/iot-cloud.service device_deploy/ &&\
	cp cmd/device/iot-telemetry.service device_deploy/ &&\
	cp cmd/device/nats.service device_deploy/
