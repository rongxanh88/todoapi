FROM golang:1.21.0-alpine3.17

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the application code to the container
COPY . .

RUN go build -o /todo-api

# API port
EXPOSE 3000

# POSTGRES port (only necessary for running locally)
EXPOSE 5432

CMD [ "/todo-api" ]
