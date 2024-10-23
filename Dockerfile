# official golang image as base image
FROM golang:1.23-alpine

# set working directory
WORKDIR /app

# copy app code to working directory
COPY . .

# build app
RUN go build -o main .

# expose the app's port
EXPOSE 8080

# run the app
CMD ["./main"]