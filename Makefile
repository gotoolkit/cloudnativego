APP?=cloudnativego
APP_PORT?=8585
APP_DATASOURCE?=root:root@tcp(docker.for.mac.localhost:3306)/rest?charset=utf8&parseTime=true
APP_DATASTOREPATH=/data
PROJECT?=github.com/gotoolkit/cloudnativego

RELEASE?=0.1.5
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

GOOS?=linux
GOARCH?=amd64

CONTAINER_IMAGE?=containerize/${APP}

clean:
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/pkg/version.Release=${RELEASE} \
		-X ${PROJECT}/pkg/version.Commit=${COMMIT} \
		-X ${PROJECT}/pkg/version.BuildTime=${BUILD_TIME} " \
		-o ${APP}

test:
	go test -v -race ./...

container: build
	docker build -t ${CONTAINER_IMAGE}:${RELEASE} .
	docker build -t ${CONTAINER_IMAGE}:latest .

run: container
	docker stop ${APP} || true && docker rm ${APP} || true
	docker run --name ${APP} -p ${APP_PORT}:${APP_PORT} --rm \
		-v cloudnativego-data:/data \
		-e "APP_PORT=${APP_PORT}" \
		-e "APP_DATASTOREPATH=${APP_DATASTOREPATH}" \
		${CONTAINER_IMAGE}:${RELEASE}

push: container
	docker push ${CONTAINER_IMAGE}:${RELEASE}
	docker push ${CONTAINER_IMAGE}:latest
