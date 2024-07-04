FROM golang:1.21.3-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN chmod +x ./config/config.sh

ARG PORT
ARG WEB_URL
ARG GOOGLE_API_KEY
ARG SPOTIFY_CLIENT_ID
ARG SPOTIFY_CLIENT_SECRET
ARG SPOTIFY_OAUTH_STATE
ARG SPOTIFY_REDIRECT_URI
ARG RAPID_API_KEY

# Write config file
RUN sh config/config.sh ${PORT} ${WEB_URL} ${GOOGLE_API_KEY} ${SPOTIFY_CLIENT_ID} ${SPOTIFY_CLIENT_SECRET} ${SPOTIFY_OAUTH_STATE} ${SPOTIFY_REDIRECT_URI} ${RAPID_API_KEY}

RUN go build -o spotwave cmd/api.go

FROM alpine:3.20

WORKDIR /app

# Copy the config file from the build stage
COPY --from=build /app/config/config.yml /app/config/config.yml

# Copy app from build stage
COPY --from=build /app/spotwave /app/spotwave

EXPOSE 80

CMD [ "./spotwave" ]