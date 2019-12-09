
if [ -z "$1" ]
then
    echo "skipping pull for local image"
else
    echo pulling latest image
    docker pull $1ny-data-api:0.0.1
fi
docker run -d -p 80:443 --name ny-data-api $1ny-data-api:0.0.1