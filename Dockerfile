FROM node:alpine AS build
WORKDIR /build

COPY package.json .
COPY yarn.lock .
RUN yarn install

COPY tailwind.config.js .
COPY templates/ ./templates/
COPY static/ ./static/
RUN yarn run tailwind:build

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

COPY bookmarks/config.yaml ./bookmarks/config.yaml
COPY templates/ ./templates/
COPY --from=build /build/static/favicon/ ./static/favicon/
COPY --from=build /build/static/css/style.css ./static/css/style.css
COPY godash .

ARG VERSION
ENV VERSION=$VERSION
ARG DATE
ENV DATE=$DATE

ENTRYPOINT ["/app/entrypoint.sh"]
