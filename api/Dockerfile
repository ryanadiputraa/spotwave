FROM go:1.21.3-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN gp build -o api cmd/api.go

CMD [ "api" ]