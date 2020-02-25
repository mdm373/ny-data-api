#!/bin/sh
./.temp/ny-data-api\
  -pass $NY_APP_POSTGRES_PASS\
  -host $NY_APP_POSTGRES_HOST\
  -serveHost $NY_APP_SERVE_SCHEME://$NY_APP_SERVE_HOST