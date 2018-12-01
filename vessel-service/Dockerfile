FROM golang:1.10 as builder

WORKDIR /Users/heaven/go/src/shippy/vessel-service
ENV GOPATH /Users/heaven/go

RUN go get

RUN CGO_ENABLED=0 GOOS=linux go build  -o vessel-service -a -installsuffix cgo main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /User/heaven/go/src/shippy/vessel-service .

CMD ["./vessel-service"]