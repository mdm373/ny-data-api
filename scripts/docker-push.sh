sh ./scripts/docker-build.sh
dockerUser=$NY_APP_DOCKER_USER
echo "pushing changes to docker for user $dockerUser..."
docker image tag ny-data-api:0.0.1 $dockerUser/ny-data-api:0.0.1
docker image push $dockerUser/ny-data-api:0.0.1