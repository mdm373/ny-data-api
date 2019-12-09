sh ./scripts/docker-build.sh
echo pushing changes to docker...
docker image tag ny-data-api:0.0.1 mdm373/ny-data-api:0.0.1
docker image push mdm373/ny-data-api:0.0.1