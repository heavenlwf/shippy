FROM golang:1.10 as builder

WORKDIR /Users/heaven/go/src/shippy/vessel-service
ENV GOPATH /Users/heaven/go
COPY . .

RUN go get
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build  -o vessel-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /Users/heaven/go/src/shippy/vessel-service .

CMD ["./vessel-service"]