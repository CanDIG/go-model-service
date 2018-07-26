# Build stage
FROM golang AS builder

WORKDIR /go/src/github.com/CanDIG/go-model-service
ADD . /go/src/github.com/CanDIG/go-model-service

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN go get -u -v -tags sqlite github.com/gobuffalo/pop/... && \
    go install -tags sqlite github.com/gobuffalo/pop/soda

RUN cd /go/src/github.com/CanDIG/go-model-service && \
    dep ensure && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags sqlite -o main ./variant-service/api/cmd/variant-service-server/main.go

# Final stage

FROM scratch

WORKDIR /
COPY --from=builder /go/src/github.com/CanDIG/go-model-service/main /

EXPOSE 3000

ENTRYPOINT ["/main", "--port=3000"]
