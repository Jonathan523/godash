FROM golang:alpine AS go
RUN apk add nodejs npm
WORKDIR /backend

COPY ./go.mod .
RUN go mod download

COPY ./package.json .
COPY ./package-lock.json .
RUN npm install

COPY . .
RUN npm run build
RUN go build -o app

FROM alpine AS logo
RUN apk add figlet
WORKDIR /logo

RUN figlet Launchpad > logo.txt

FROM alpine AS final
WORKDIR /app

COPY --from=logo /logo/logo.txt .

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

COPY --from=go /backend/config/ ./config/
COPY --from=go /backend/templates .
COPY --from=go /backend/static .
COPY --from=go /backend/app .

ENTRYPOINT ["/app/entrypoint.sh"]
