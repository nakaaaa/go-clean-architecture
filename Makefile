DSN="admin:admin@tcp(127.0.0.1:3306)/example?charset=utf8mb4"

.PHONY: db/migrate db/status db/down

db/migrate:
	(cd ./database/migrations; goose mysql $(DSN) up)
		
db/status:
	(cd ./database/migrations; goose mysql $(DSN) status)

db/down:
	(cd ./database/migrations; goose mysql $(DSN) down)

db/reset:
	MYSQL_PWD=admin mysql -h 127.0.0.1 -u admin -e 'DROP DATABASE IF EXISTS example;'

.PHONY: mysql/connect

mysql/connect:
	mysql -u admin -padmin -h 127.0.0.1 example