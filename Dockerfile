FROM golang:alpine AS go
WORKDIR /backend

COPY ./go.mod .
RUN go mod download

COPY . .
RUN npm build
RUN go build -o app

FROM alpine AS logo
RUN apk add figlet
WORKDIR /logo

RUN figlet Launchpad > logo.txt

FROM alpine AS final
WORKDIR /app

COPY --from=logo /logo/logo.txt .

# copy all the configuration files and default bookmark json
COPY --from=go /backend/bookmark/bookmarks.json ./bookmark/bookmarks.json
COPY --from=go /backend/logging/logging.json ./logging/logging.json

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

COPY --from=go /backend/static .
COPY --from=go /backend/app .

ENTRYPOINT ["/app/entrypoint.sh"]
