# NY Data API
> An api layer for ny data

See [ingest](https://github.com/mdm373/ny-data-ingest) for data source details and [web](https://github.com/mdm373/ny-data-web) for the pretty front end.

## Tooling
* `mux` for general routing
* `squirrel` / `structable` for SQL querying / structuring
*  postgres driver via `pq`
* `npm` for easy script running

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
        },
        "HOSTED" : "same format as local but for values to use when running in hosted env"
    }
}
```
## Requirements
* development
  * go (~ 1.13.x)
  * npm
  * jq
* deployment
  * docker
  * docker-cli

## Scripts
* `npm run serve-dev` - run a dev server
* `npm run deploy` - update docker, restart/update remote docker host
