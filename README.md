# go-oauth2-server
> OAuth 2.0 Authorization server and Resource server in Golang.

## Quick Start

```sh
$ docker-compose up --build
```

### Open in your web browser
Authorization Request: http://localhost:9091/authorize?client_id=000000&response_type=code

Grant Token Request: http://localhost:9091/token?grant_type=client_credentials&client_id=000000&client_secret=999999

```sh
# res
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
- Access Token
  - JWT


## MIT License
Copyright (c) 2024 miyuki-starmiya
