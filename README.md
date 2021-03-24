![](poster.jpg)

---
# IoT Stand

This solution is boilerplate structure that provides base *microservices* for:

- **ARM based Board (Raspberry Pi 3, Toradex, Boundery Devices, etc.)**
- **MQTT based IoT Hub**
- **Self-hosted cloud server for client who communicated with devices in IoT Hub**
- **Example web client**

---
# ARM based Board

---
## Requires

- [Golang](https://golang.org)
- Linux (x86 or amd64)
- Any ARMv7 microcomputer board (Raspberry Pi 3, Toradex, Boundery Devices, etc.)
- [Yocto Linux with QT](http://code.qt.io/cgit/yocto/meta-boot2qt.git/)
- [Qt Open Source for Desktop](https://www.qt.io/)
- [Nats for Desktop&ARMv7](https://nats.io/)

---
## Microservices

- ```Telemetry Service [Golang]``` send any telemetry data (sensors, circels, temperature, etc.) to ```Cloud Service```
- ```Cloud Service [Golang]``` - send any data to IoT hub based from MQTT protocol
- ```HMI Service [QT Quick]``` - user interface for sensor displays

All microservices communicate with NATS pub/sub and provides hight level abstract RPC API.

---
## Build Yocto Linux

Build you board image with [Yocto meta-boot2qt layer](http://code.qt.io/cgit/yocto/meta-boot2qt.git/) and [Instructions](http://doc.qt.io/QtForDeviceCreation/qtee-custom-embedded-linux-image.html).

---
## Build Microservices

Run in terminal ```make deploy_device``` and see ```deploy_device``` folder.

Deploy all files to ```/deploy_device``` on your device with ```ssh```.

Start ```systemd``` services (```.service```) and ```NATS ARMv7``` service.

---
## Development on local host machine

Use ```cmd/device/telemetry-service.go``` and ```cmd/device/telemetry-service.go``` for analyze of golang based microservice.

Use ```device/service/hmi``` for analyze of QT based microservice with ```Qt Creator```.

For start all of microservices required ```NATS x86 or amd64 targets```.

---
# MQTT based IoT Hub, Websocket self-hosted Cloud server and Web client

---
## Requires

- [Golang](https://golang.org)
- [Nats for Desktop](https://nats.io/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---
## Build

```
make build_server_side
```

See folder ```server_side_deploy```.

---
## Run

```
cd server_side_deploy

docker-compose up -d --force-recreate --build
```

or use ansible for remote server:

```
ansible-playbook -i YOUR_INVENTORY_FILE ansible/deploy.yaml
```

You can see result front-end based react app on port ```:80```.

If you want analyze server-side code, see folder ```server/service/**```.

If you want to analyze front-end code, see folder ```server/frontend/**```.

```Websocket service``` and ```MQTT service``` communicate with NATS pub/sub, provides hight level abstract RPC API and can be replicated with ```Kubernetes```|```Mesos```|```Swarm```.
