FROM golang:1.24.0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN ls

RUN go build -o /kryptonim-app .

EXPOSE 8081

CMD ["/kryptonim-app"]
