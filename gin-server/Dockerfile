FROM golang:1.19 AS build
WORKDIR /go/src
RUN git clone https://github.com/haozheng95/eqx-gin-server.git
WORKDIR /go/src/eqx-gin-server/gin-server
ENV CGO_ENABLED=0
RUN go mod tidy

RUN go build -a -installsuffix cgo -o startapi .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/eqx-gin-server/gin-server/startapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./startapi"]
