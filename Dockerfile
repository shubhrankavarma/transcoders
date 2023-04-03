FROM golang:alpine

# golang specific variables
# ENV GO111MODULE=on \
#   CGO_ENABLED=0 \
#   GOOS=linux \

# current working directory is /build in the container
WORKDIR /build

# copy over go.mod and go.sum (module dependencies and checksum)
# over to working directory
COPY go.mod .
COPY go.sum .

# download the dependencies
RUN go mod download

# copy our application code into the container
COPY ./config .
COPY ./dbiface .
COPY ./docs .
COPY ./handlers .
COPY ./go.mod .
COPY ./go.sum .
COPY ./main.go .
COPY ./.gitignore .
COPY ./Dockerfile .

# Expose port 51000 to the outside world
EXPOSE 51000

# Run app
CMD ["go", "run", "main.go"]