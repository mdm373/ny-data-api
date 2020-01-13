# NY Data API
> An api layer for ny data

See [ingest](https://github.com/mdm373/ny-data-ingest) for data source details and [web](https://github.com/mdm373/ny-data-web) for the pretty front end.

## Tooling
* `mux` for general routing
* `squirrel` / `structable` for SQL querying / structuring
*  postgres driver via `pq`
* `npm` for easy script running / versioning
* `swagger` for api documentation

## Install
```
npm run install
```
create configuration `./.secrets.json` file in the following format
```
{
    "DOCKER_HOST" : "hostname to ssh into for remote hosted docker updates",
    "DOCKER_USER" : "name of user for docker image pull / push",
    "CONFIG" : {
        "LOCAL" : {
            "HOST" : "host of posstgres db when running locally",
            "PASS" : "password for the same"
            "SERVE_HOST": "externally visible host (xyz.abc.com) when served",
            "SCHEME": "externally visible scheme (https or http) when served"
        },
        "HOSTED" : "same format as local but for values to use when running in hosted env"
    }
}
```

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