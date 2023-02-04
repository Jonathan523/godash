FROM node:alpine AS build
WORKDIR /build

COPY ./package.json .
COPY ./package-lock.json .
RUN npm install

COPY . .
RUN npm run build

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

# default config.yaml
COPY --from=build /build/bookmarks/config.yaml ./bookmarks/config.yaml
# go templates
COPY --from=build /build/templates/ ./templates/
# build static files and favicons
COPY --from=build /build/static/favicon/ ./static/favicon/
COPY --from=build /build/static/css/style.css ./static/css/style.css
COPY --from=build /build/static/js/app.min.js ./static/js/app.min.js
# go executable
COPY godash .

ENTRYPOINT ["/app/entrypoint.sh"]
