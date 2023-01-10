FROM --platform=linux/arm64 golang:1.18 AS build

ENV GO111MODULE=on

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git

WORKDIR /go/src/github.com/ambiez-go/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .

FROM alpine
COPY --from=build /go/src/github.com/ambiez-go/main go/bin/ambiez-go/

COPY --from=build /go/src/github.com/ambiez-go/.env go/bin/ambiez-go/

EXPOSE 9000
WORKDIR go/bin/ambiez-go/



CMD ["./main"]