
if [ -z "$1" ]
then
    export NY_APP_PASS=$(cat ./.secrets.json | jq -r '.CONFIG.LOCAL.PASS')
    export NY_APP_HOST=$(cat ./.secrets.json | jq -r '.CONFIG.LOCAL.HOST')
    export SERVE_HOST=$(cat ./.secrets.json | jq -r '.CONFIG.LOCAL.SCHEME')://$(cat ./.secrets.json | jq -r '.CONFIG.LOCAL.SERVE_HOST')
    echo "skipping pull for local image"
else
    export NY_APP_PASS=$(cat ./.secrets.json | jq -r '.CONFIG.HOSTED.PASS')
    export NY_APP_HOST=$(cat ./.secrets.json | jq -r '.CONFIG.HOSTED.HOST')
    export SERVE_HOST=$(cat ./.secrets.json | jq -r '.CONFIG.HOSTED.SCHEME')://$(cat ./.secrets.json | jq -r '.CONFIG.HOSTED.SERVE_HOST')
    echo pulling latest image
    docker pull $1ny-data-api:0.0.1
fi
docker run -e SERVE_HOST -e NY_APP_PASS -e NY_APP_HOST -d -p 80:8000/tcp --expose 8000 --name ny-data-api $1ny-data-api:0.0.1