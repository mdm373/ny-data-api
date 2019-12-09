sh ./scripts/docker-build.sh
sh ./scripts/docker-push.sh
export DOCKER_HOST=$(cat ./.secrets.json | jq -r '.DOCKER_HOST')
sh ./scripts/docker-restart.sh mdm373/
export DOCKER_HOST=
