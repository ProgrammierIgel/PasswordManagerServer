FROM golang:1.23-alpine AS builder

RUN apk update

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/voting

# ---

FROM alpine:3.21.2

COPY --from=builder /go/bin/voting /usr/local/bin/voting

ENTRYPOINT [ "voting" ]
