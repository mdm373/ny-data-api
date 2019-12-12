sh ./scripts/docker-build.sh
dockerUser=$(cat ./.secrets.json | jq -r '.DOCKER_USER')
echo "pushing changes to docker for user $dockerUser..."
docker image tag ny-data-api:0.0.1 $dockerUser/ny-data-api:0.0.1
docker image push $dockerUser/ny-data-api:0.0.1