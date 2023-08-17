FROM golang:1.21.0-alpine3.17

ENV DATABASE_URL="postgres://siwqbqwwndrxxe:41215a90b5dff4ef242529a67f965fad9d0c17149e7e366bb7d75e4a69ff1a80@ec2-44-214-132-149.compute-1.amazonaws.com:5432/d270rc33e5e2cn"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the application code to the container
COPY . .

RUN go build -o /todo-api

# API port
EXPOSE 3000

CMD [ "/todo-api" ]
