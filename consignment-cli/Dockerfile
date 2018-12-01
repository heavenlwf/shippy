FROM golang:1.10 as builder

WORKDIR /Users/heaven/go/src/shippy/consignment-cli
ENV GOPATH /Users/heaven/go
COPY . .

RUN go get

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY consignment.json /app/consignment.json
COPY --from=builder /Users/heaven/go/src/shippy/consignment-cli .

ENTRYPOINT ["./consignment-cli", "consignment.json"]
CMD ["./consignment-cli"]
