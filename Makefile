include .env

SEED := $(shell perl -e "print int(rand(1000000))")

.PHONY: build
build:
	#${GO} test ./...
	${GO} env -w GOOS=linux GOARCH=amd64
	${GO} build -o ./build/server ./

.PHONY: docker-run
docker-run:
	docker build . --tag video_web:${VERSION}-${SEED}
	docker rm -f video_web
	docker run -d --restart=always -p 9010:9010 --name video_web -v /mnt/e/code/video_web/configs/conf.yaml:/app/configs/conf.yaml video_web:${VERSION}-${SEED}

.PHONY: docker-push
docker-push:
	docker build . --tag video_web:${VERSION}-${SEED}
	docker tag video_web:${VERSION}-${SEED} registry.cn-hangzhou.aliyuncs.com/kkako/video_web:${VERSION}-${SEED}
	docker login --username=${USERNAME} registry.cn-hangzhou.aliyuncs.com -p ${PASSWORD}
	docker push ${ADDR}:${VERSION}-${SEED}

.PHONY: test
test:
	${GO} test ./...

.PHONY: echo
echo:
	echo GOBASE = ${GOBASE}
	echo GOPATH = ${GO}
	echo GOBIN = ${GOBIN}
	echo GOFILES = ${GOFILES}

.PHONY: protoc
protoc:
		protoc -I. -I ./third_party -I ./api \
			--go_out . --go_opt paths=source_relative \
			--go-grpc_out . --go-grpc_opt paths=source_relative \
			--grpc-gateway_out . \
			--grpc-gateway_opt logtostderr=true \
			--grpc-gateway_opt paths=source_relative \
			--grpc-gateway_opt generate_unbound_methods=true \
			 ./api/user/v1/*.proto ./api/video/v1/*.proto ./api/comment/v1/*.proto