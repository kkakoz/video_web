include .env

SEED := $(shell perl -e "print int(rand(1000000))")

.PHONY: build
build:
	#${GO} test ./...
	${GO} env -w GOOS=linux GOARCH=amd64
	${GO} build -o ./build/server ./cmd
	docker build ./build --tag video_web:${VERSION}-${SEED}
	docker tag video_web:${VERSION}-${SEED} registry.cn-hangzhou.aliyuncs.com/kkako/video_web:${VERSION}-${SEED}
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
