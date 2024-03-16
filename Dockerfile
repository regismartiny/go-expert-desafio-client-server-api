FROM golang:1.22-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD cmd /app/cmd
ADD configs /app/configs
ADD internal /app/internal

WORKDIR /app/cmd/server

RUN apk add --no-cache gcc g++
RUN GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server

EXPOSE 8080

CMD [ "./server" ]