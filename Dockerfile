FROM golang:1.20-alpine

WORKDIR /app

COPY . .
RUN ls -la
#RUN go mod download && go build -a -installsuffix cgo -o main /MusicServiceBackend/cmd/app/main.go

EXPOSE 8080
#CMD ["/MusicServiceBackend"]