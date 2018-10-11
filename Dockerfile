FROM golang:1.11-alpine3.8 AS build
WORKDIR /go/src/app
ADD *.go Gopkg.lock Gopkg.toml /go/src/app/

RUN apk add --no-cache git upx && \
	go get -v -u github.com/golang/dep/cmd/dep && \
	dep ensure && \
    CGO_ENABLED=0 GOOS=linux go build -o bin/message -a -ldflags="-s -w" -installsuffix cgo && \
    upx --best bin/message

FROM scratch
COPY --from=build /go/src/app/bin/message /message
ENTRYPOINT [ "/message"  ]
