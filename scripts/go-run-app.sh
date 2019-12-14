NY_APP_PASS=$(cat ./.secrets.json | jq -r '.CONFIG.LOCAL.PASS')
NY_APP_HOST=$(cat ./.secrets.json | jq -r '.CONFIG.LOCAL.HOST')
go run ./app/main.go --host=$NY_APP_HOST --pass=$NY_APP_PASS