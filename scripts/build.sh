region=${1:---test}
if [ $region == "--prod" ]
then
    echo "building for prod"
    go build -a -ldflags '-extldflags "-static"' -o ./.temp/ny-data-api
else
    echo "building for test"
    go build -o ./.temp/ny-data-api
fi