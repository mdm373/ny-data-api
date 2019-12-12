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
sh ./scripts/docker-run.sh $1