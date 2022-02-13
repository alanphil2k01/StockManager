# Build the go binary
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/alanphil2k01

COPY server/go.mod .
COPY server/go.sum .
RUN go mod download

COPY server .

COPY frontend pkg/server/static

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /go/bin/SSMC cmd/ssmcserver/main.go

#  Build image for running the program
FROM scratch

COPY --from=builder /go/bin/SSMC /bin/SSMC

ENTRYPOINT ["/bin/SSMC"]
