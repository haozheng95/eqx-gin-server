FROM golang:1.19 AS build
WORKDIR /go/src
RUN git clone https://github.com/haozheng95/eqx-gin-server.git
WORKDIR /go/src/eqx-gin-server/gin-server
ENV CGO_ENABLED=0
RUN go env -w GO111MODULE=on
RUN go mod tidy

RUN go build -a -installsuffix cgo -o sw .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/eqx-gin-server/gin-server/sw ./
EXPOSE 8080/tcp
ENTRYPOINT ["./startapi"]
