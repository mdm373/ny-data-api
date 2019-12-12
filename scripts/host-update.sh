sh ./scripts/docker-build.sh
sh ./scripts/docker-push.sh
export DOCKER_HOST=$(cat ./.secrets.json | jq -r '.DOCKER_HOST')
dockerUser=$(cat ./.secrets.json | jq -r '.DOCKER_USER')
sh ./scripts/docker-restart.sh $dockerUser/
export DOCKER_HOST=
