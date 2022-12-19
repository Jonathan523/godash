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

RUN figlet GoDash > logo.txt

FROM alpine AS final
RUN apk add tzdata
WORKDIR /app

COPY --from=logo /logo/logo.txt .

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

# default bookmarks.json
COPY --from=go /backend/bookmarks/bookmarks.json ./bookmarks/bookmarks.json
# go templates
COPY --from=go /backend/templates/ ./templates/
# build static files and favicons
COPY --from=go /backend/static/favicon/ ./static/favicon/
COPY --from=go /backend/static/css/style.css ./static/css/style.css
COPY --from=go /backend/static/js/app.min.js ./static/js/app.min.js
# go executable
COPY --from=go /backend/app .

ENTRYPOINT ["/app/entrypoint.sh"]
