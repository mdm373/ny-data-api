export NY_APP_PASS=$(cat ./.secrets.json | jq -r '.CONFIG.LOCAL.PASS')
export NY_APP_HOST=$(cat ./.secrets.json | jq -r '.CONFIG.LOCAL.HOST')
export SERVE_HOST="http://localhost:8000"
sh ./scripts/bin-run.sh