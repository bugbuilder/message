FROM golang:1.11-alpine3.8 AS build
RUN apk add --no-cache \
	git \
	upx \
	make && \
	go get -v -u github.com/golang/dep/cmd/dep

FROM build AS message
WORKDIR /go/src/github.com/bugbuilder/message
COPY . .
RUN dep ensure
RUN make build && mv message /usr/bin/message

FROM scratch
COPY --from=message /usr/bin/message /message
ENTRYPOINT [ "/message"  ]
