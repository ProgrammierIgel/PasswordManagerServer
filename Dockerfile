FROM golang:1.24-alpine AS builder

RUN apk update

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/voting

# ---

FROM alpine:latest

COPY --from=builder /go/bin/pwmanager /usr/local/bin/pwmanager

ENTRYPOINT [ "voting" ]
