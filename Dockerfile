FROM golang:1.19 AS build
WORKDIR /go/src

COPY gin-server/go ./go
COPY gin-server/main.go .

ENV CGO_ENABLED=0
RUN go mod init github.com/haozheng95/eqx-gin-server/gin-server
RUN go get -d -v ./...
RUN go mod tidy

RUN go build -a -installsuffix cgo -o startapi .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/startapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./startapi"]
