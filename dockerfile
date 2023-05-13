FROM golang:1.20

RUN apt-get update && apt install -y openjdk-17-jre && apt install -y openjdk-17-jdk

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app app.go

CMD ["app"]