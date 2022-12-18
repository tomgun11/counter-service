# Build Stage
FROM golang:1.17-alpine as build-env
 
ENV APP_NAME counter-app
ENV CMD_PATH main.go
 
# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
 


# Run Stage
FROM alpine:3.14
 
ENV APP_NAME counter-app
COPY --from=build-env /$APP_NAME .
EXPOSE 80

CMD ./$APP_NAME