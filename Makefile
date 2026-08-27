space := $(empty) $(empty)
goose_env := GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_URL) GOOSE_MIGRATION_DIR=./migrations

templ:
	templ generate --watch --cmd="go run ."

.PHONY: migration
migration:
	@$(goose_env) goose $(word 2,$(MAKECMDGOALS)) $(subst $(space),_,$(wordlist 3,100,$(MAKECMDGOALS))) sql

test:
	@echo $(subst $(empty) $(empty),_,$(wordlist 2,100,$(MAKECMDGOALS)))

%:
	@: