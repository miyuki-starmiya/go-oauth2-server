# go-oauth2-server
> OAuth 2.0 Authorization server and Resource server in Golang.

## Quick Start
create `.env` in your project root directory.
```sh:.env
CLIENT_ID=
CLIENT_SECRET=
REDIRECT_URI=

MONGO_INITDB_ROOT_USERNAME=
MONGO_INITDB_ROOT_PASSWORD=
MONGO_INITDB_DATABASE=
MONGO_USER=
MONGO_PASSWORD=
MONGO_DB=
MONGO_HOST=
MONGO_PORT=
```

run containers via docker-compose and create collections.
```sh
$ docker-compose up --build
$ docker-compose exec db bash

# in db container
$ mongosh -u <MONGO_USER> -p <MONGO_PASSWORD>
$ use go-oauth2-server
$ db.createCollection('codes')
$ db.createCollection('tokens')
```

### Open in your web browser
Authorization Request: localhost:9001/authorize?response_type=code&redirect_uri=http://example.com/&state=xyz&client_id=123456

Grant Token Request: localhost:9001/token?grant_type=authorization_code&code=YmY3YTc3ODItZjRhZi0zNTA4LWIyMzMtOGU2YmJhMWY0M2Vh&redirect_uri=http://example.com/

Sample Response
```sh
{
    "access_token": "OGZLYMY4MWITMGVJOS0ZODAZLTG4OGUTYTJIOGM0ZGU2YTQY",
    "expires_in": 3600,
    "refresh_token": "",
    "token_type": "Bearer"
}
```

## Satisfied Requirements
- Authorization Request parameters
  - response_type
  - client_id
  - redirect_uri
  - state
- Authorization Response parameters
  - code
  - state
- Grant Types
  - Authorization Code Grant


## MIT License
Copyright (c) 2024 @miyuki-starmiya
