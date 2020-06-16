FROM golang:1.12-alpine
WORKDIR /app 
COPY ./main .
CMD main
EXPOSE 20080