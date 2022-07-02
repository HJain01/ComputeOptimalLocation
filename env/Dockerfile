FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY ../go.mod ./
COPY ../go.sum ./
RUN go mod download

COPY ../api ./api
COPY ../cmd ./cmd

RUN echo $(ls)

RUN go build -o /compute-optimal-location ./api

CMD [ "/compute-optimal-location" ]
