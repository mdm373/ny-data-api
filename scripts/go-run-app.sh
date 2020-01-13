sh ./scripts/swagger-generate.sh --local
NY_APP_PASS=$(cat ./.secrets.json | jq -r '.CONFIG.LOCAL.PASS')
NY_APP_HOST=$(cat ./.secrets.json | jq -r '.CONFIG.LOCAL.HOST')
go run ./app --host=$NY_APP_HOST --pass=$NY_APP_PASS --serveHost=http://localhost:8000