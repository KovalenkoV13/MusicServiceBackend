FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod download && go build MusicServiceBackend/cmd/app/

EXPOSE 8080

CMD ./app