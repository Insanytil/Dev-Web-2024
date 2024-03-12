FROM node:16.20-alpine3.18 AS JS_BUILD
COPY web /web
WORKDIR /web
RUN npm install && npm run build

FROM golang:1.22.0-alpine3.18 AS GO_BUILD
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.18.3
COPY --from=JS_BUILD /web/build* ./web/
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server
