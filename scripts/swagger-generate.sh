cp -R ./node_modules/swagger-ui-dist/ ./static/swagger-ui
swagger generate spec -m -o ./.temp/swagger_template.yaml
mv ./static/swagger-ui/index.html ./static/swagger-ui/_back.index.html
cat ./static/swagger-ui/_back.index.html | sed "s,https://petstore.swagger.io/v2/swagger.json,./swagger.yaml,g" > ./static/swagger-ui/index.html
SWAGGER_VERSION=$(cat ./package.json | jq -r '.version')
if [ $1 == "--hosted" ]
then
  SWAGGER_HOST=$(cat ./.secrets.json | jq -r '.CONFIG.HOSTED.SERVE_HOST')
  SWAGGER_SCHEME=$(cat ./.secrets.json | jq -r '.CONFIG.HOSTED.SCHEME')
else
  SWAGGER_HOST=$(cat ./.secrets.json | jq -r '.CONFIG.LOCAL.SERVE_HOST')
  SWAGGER_SCHEME=$(cat ./.secrets.json | jq -r '.CONFIG.LOCAL.SCHEME')
fi
SWAGGER_HOST=$SWAGGER_HOST SWAGGER_SCHEME=$SWAGGER_SCHEME SWAGGER_VERSION=$SWAGGER_VERSION envsubst \
  '${SWAGGER_HOST} ${SWAGGER_SCHEME} ${SWAGGER_VERSION}' \
  < ./.temp/swagger_template.yaml > ./static/swagger-ui/swagger.yaml
swagger validate ./static/swagger-ui/swagger.yaml
