# Build stage
FROM golang AS builder

WORKDIR /go/src/github.com/CanDIG/go-model-service
ADD . /go/src/github.com/CanDIG/go-model-service

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN cd /go/src/github.com/CanDIG/go-model-service && \
    dep ensure && \
    cd variant-service/api && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../../main ./cmd/variant-service-server/main.go

# Final stage

FROM scratch

WORKDIR /
COPY --from=builder /go/src/github.com/CanDIG/go-model-service/main /

EXPOSE 3000

ENTRYPOINT ["/main", "--port=3000"]
