FROM golang:1.19 AS build
WORKDIR /go/src
COPY /Users/yhaozheng/go/src/github.com/equinix-labs/metal-go ../../metal-go
COPY go ./go
COPY go.mod ./go.mod
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o startapi .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/startapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./startapi"]
