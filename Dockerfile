FROM golang:1.6

WORKDIR /app
ADD lib/little-bastard.go /app/little-bastard.go

CMD go run little-bastard.go
