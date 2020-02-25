# NY Data API
> An api layer for ny data

See [ingest](https://github.com/mdm373/ny-data-ingest) for data source details and [web](https://github.com/mdm373/ny-data-web) for the pretty front end.

## Tooling
* `mux` for general routing
* `squirrel` / `structable` for SQL querying / structuring
*  postgres driver via `pq`
* `npm` for easy script running, env loading / versioning
* `swagger` for api documentation

## Install
```
npm run install
```
* populate `.env` for local server instance config
* populate `.env.docker` for hosted server instance config

## Development Pre-Requisits
  * go (~ 1.13.x)
  * npm
  * docker
  * docker-cli
## Requirements (installed during npm install)
  * jq
  * go-swagger
  * gettext (envsubst)

## Scripts
* `npm run serve-dev` - run a dev server
* `npm run deploy` - update docker, restart/update remote docker host

## Swagger

visit `/swagger-ui/` or `/` once serving for swagger generated api documentation