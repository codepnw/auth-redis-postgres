include .env

goose-up:
	cd sql/schema && goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up

goose-down:
	cd sql/schema && goose $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down

sqlc-gen:
	sqlc generate