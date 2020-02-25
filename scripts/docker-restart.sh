containerId=$(docker ps -q -a -f "name=ny-data-api")
echo "restarting docker container"
if [ -z "$containerId" ]
then
      echo container not running
else
    echo stopping existing containerId $containerId
    docker stop $containerId
    docker rm $containerId
fi
echo pulling latest $NY_APP_DOCKER_USER/ny-data-api:0.0.1 image
docker pull $NY_APP_DOCKER_USER/ny-data-api:0.0.1
docker run \
  -e NY_APP_SERVE_HOST\
  -e NY_APP_SERVE_SCHEME\
  -e NY_APP_POSTGRES_PASS\
  -e NY_APP_POSTGRES_HOST\
  -d\
  -p 80:8000/tcp\
  --expose 8000\
  --name ny-data-api\
  $NY_APP_DOCKER_USER/ny-data-api:0.0.1