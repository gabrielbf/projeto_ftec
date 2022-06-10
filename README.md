# sample

project description
## Open-API

\*Require npm installed

```
npm i -g redoc-cli
go install github.com/swaggo/swag/cmd/swag@latest

npx redoc-cli serve docs/swagger.yaml -p <port> --watch

OR

npx @redocly/openapi-cli preview-docs docs/swagger.yaml --port <port>
```

Or starting the application with the make-file command:

```
make run
```

And accessing your application at `http://localhost:<port>/swagger/` in your web browser.

---

<br>

## Updating Open API documentation

You can update the Open-API documentation in every change in the controller/ handler methods or its DTOs (request or response) by running:

```
make generate-swagger

or

swag init -g cmd/main.go
```

The .yaml or .json files will be generated automatically.
Access https://github.com/swaggo/swag for more informations.

<br>

---

## Environment Variables

<br>
To run this project is necessary to add some environment variables.

<br>

> This is an Example Environments from Development environment

| KEY                                      | VALUE                                            | OPTIONAL | Description                                |
| ---------------------------------------- | ------------------------------------------------ | -------- | ------------------------------------------ |
| ENVIRONMENT                              | local                                            |          | Environment name. (used to log on datadog) |
| LOG_LEVEL                                | DEBUG                                            |          | Log Level                                  |
| SERVER_PORT                              | 3000                                             |          | Server Port. (used to run the http server) |
| DB_NAME                                  | sample                                           |          | Database name                              |
| DB_HOST                                  | localhost                                        |          | Database Hostname                          |
| DB_PORT                                  | 5432                                             |          | Database port service                      |
| DB_USER                                  | admin                                            |          | Database User                              |
| DB_PASSWORD                              | admin                                            |          | Database Password                          |

---

<br>

## Running the app Locally

You can run the app by setting the environment variables above in a `.env` file in the project root path. You can copy the `.env.example` and change its name to `.env`. Then, you can run the following command do set up the other services needed (Postgres):

```

docker compose up -d

```

Done that, run:

```
make run
```

---

<br>

## Project Design

We are working to apply the best practices and following the clean code and the clean architecture knowledge, even though it's a long road to go.
Try to remember The Boy Scout Rule

```
Always leave the campground cleaner than you found it
```

Actually, this project works with the following layers:

### app

Contains the `service` and `app` packages. The `service` package handles with the concrete implementation of the business logic domain. And the `app` is responsible for build and initializing the application.

### domain

`domain` have all the files about a specific domain, like <s>its entity</s>, enums and interfaces implemented by the `service` layer. In there you can find the packages `constants`, `enum`, and `sample`, . \*<s>its entity</s> is commented because we're actually also using the entity as a datastructure object too, so its located at the infra layer.

### infra

`infra` handles all that is external or common on the application packages. Like, database and messaging broker connections, environment and logs setup, repository and web client implementations etc. Actually, infra also holds the `entity` package for the domain models, something i particular don't like, and maybe we change it in the future. My opinion about it is that i'd rather it on the domain layer than in here, even being used as datastructure object and not as an entity. The folder also contains the migration files, witch is common to put in a external folder, at the same hierarchy of `/internal`.

### interface

`interface` is a layer used as an entry-point for the application, containing the amqp and http interfaces, like controllers, middlewares, handlers, consumers and its DTOs (request, response and incoming messages).

### mock

`Mock` is a folder with auto generate files for mocked tests.
You should generate this files always when changing your real implementations. For doing this, run:

```
make mocks
```

---

<br>
