FROM golang:alpine AS go
RUN apk add nodejs npm
WORKDIR /backend

COPY ./swagger.sh .
RUN chmod +x swagger.sh
RUN ./swagger.sh install

COPY ./go.mod .
RUN go mod download

COPY ./package.json .
COPY ./package-lock.json .
RUN npm install

COPY . .
RUN npm run build
RUN ./swagger.sh init
RUN go build -o app

FROM alpine AS logo
RUN apk add figlet
WORKDIR /logo

RUN figlet GoDash > logo.txt

FROM alpine AS final
RUN apk add tzdata
WORKDIR /app

COPY --from=logo /logo/logo.txt .

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

COPY --from=go /backend/config/ ./config/
COPY --from=go /backend/templates ./templates/
COPY --from=go /backend/static ./static/
COPY --from=go /backend/app .

ENTRYPOINT ["/app/entrypoint.sh"]
