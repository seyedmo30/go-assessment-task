./goose -allow-missing -dir ./migrations  postgres "host=${POSTGRES_HOST} user=${POSTGRES_USER} dbname=${POSTGRES_DB} password=${POSTGRES_PASSWORD} sslmode=${POSTGRES_SSL_MODE}" up
./goose -allow-missing -table goose_seeds -dir ./seeds  postgres "host=${POSTGRES_HOST} user=${POSTGRES_USER} dbname=${POSTGRES_DB} password=${POSTGRES_PASSWORD} sslmode=${POSTGRES_SSL_MODE}" up
./main