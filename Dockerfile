FROM golang:1.12-alpine
WORKDIR /app 
COPY drone-demo /app
ENTRYPOINT "./drone-demo"
EXPOSE 20080