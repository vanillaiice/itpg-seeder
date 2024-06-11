run-sqlite:
	go run . --db-backend sqlite --db-url "file:itpg.db?journal_mode=memory&sync_mode=off&mode=rwc"
run-postgres:
	go run . --db-backend postgres --db-url "postgres://ark@localhost:5432/ark"
