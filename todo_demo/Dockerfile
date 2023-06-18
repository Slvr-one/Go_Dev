FROM golang:1.14-alpine as build

WORKDIR /app

RUN apk add --no-cache git curl

COPY . .

RUN go build -o todo

# # #### #
# FROM alpine:3.9 as runtime

# COPY --from=build /app/todo /app/todo

ENTRYPOINT [ "/app/todo" ]