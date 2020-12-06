FROM golang:1.15.5-alpine3.12

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Dmitry Roitman"

# Set the Current Working Directory inside the container
WORKDIR /app


# Copy the source from the current directory to the Working Directory inside the container
COPY . .
#COPY ./src/models /usr/local/go/src/models
# Build the Go app
RUN go get -d github.com/gorilla/mux && \
    go get -d github.com/jackc/pgx && \
    go get database/sql && \
    go get html/template && \
    go get github.com/go-swagger/go-swagger && \
    go get -u gorm.io/gorm && \
    go get github.com/jinzhu/gorm/dialects/postgres && \
    go get github.com/gorilla/websocket
RUN export GOPATH=$GOPATH:./src
# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["go", "run", "main.go"]

