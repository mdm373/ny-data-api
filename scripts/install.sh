brew install jq
brew install gettext
brew link --force gettext
go get github.com/go-swagger/go-swagger
cp -n .env.shadow .env
cp -n .env.shadow .env.docker
sh ./scripts/go-get-deps.sh