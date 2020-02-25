sh ./scripts/swagger-generate.sh
go run ./app --host=$NY_APP_POSTGRES_HOST --pass=$NY_APP_POSTGRES_PASS --serveHost=$NY_APP_SERVE_SCHEME://$NY_APP_SERVE_HOST