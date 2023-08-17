# Todo API

### Routes

The following routes are available.

GET /todos
GET /todos/{id}
POST /todos
PUT /todos/{id}
DELETE /todos/{id}

Generally, the todo response will be
```
id
title
description
completed
created_at
updated_at
```

Example POST request
```
curl --location --request POST 'https://vorto-test-todo-api-afb06cc2f659.herokuapp.com/todos' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "title": "Buy Tokens",
    "description": "Bitcoins are good",
    "completed": true
}'
```

### Production
This app is deployed via a container to Heroku. The link to the app is at `https://vorto-test-todo-api-afb06cc2f659.herokuapp.com/todos`


### Running locally
You can run the app locally by running `docker compose up` which will start up a postgres container and the todo-api.
Then you can hit the api via `localhost:3000/todos`
