FROM golang:1.20

WORKDIR /usr/src/gowebserver

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build 
ENV GIN_MODE=release

CMD ["./gowebserver"]
