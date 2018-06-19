FROM golang as builder
ARG version
WORKDIR /go/src/todo
ADD . /go/src/todo
RUN go get -u github.com/go-redis/redis
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todo main.go
RUN mkdir /opt/bin && mv todo /opt/bin/

FROM alpine:latest
WORKDIR /opt/bin
COPY --from=builder /opt/bin/todo .
EXPOSE 8080
CMD [ "./todo" ]
