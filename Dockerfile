# build
FROM golang:1.19-alpine AS build
COPY . /src
WORKDIR /src
RUN go build -o /build/app

# iamge
FROM alpine:latest
COPY --from=build /build/app /bin/app
ENTRYPOINT [ "/bin/app" ]