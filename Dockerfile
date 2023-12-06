# syntax=docker/dockerfile:1

# First, retrieve the most recent go version from the collection of go Docker images
FROM golang:1.20

# Afterwards, create the working directory into the image
WORKDIR /app

# Copy all of the files over into the image
COPY . /app

# Create the binary needed to run the go program.
RUN go build -o /deploy

# EXPOSE 8080

CMD [ "/deploy" ]