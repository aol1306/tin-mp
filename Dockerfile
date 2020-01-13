FROM golang:alpine
WORKDIR /workdir
COPY . .
RUN apk add --no-cache git g++
EXPOSE 8080
CMD ["go", "run", "."]
