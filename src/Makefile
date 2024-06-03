all:
	docker exec -it psqldb psql -U admin -d db

.PHONY: bash
bash:
	docker exec -it psqldb bash

.PHONY: db
db:
	cat ./sql/create_tables.sql | docker exec -i psqldb psql -U postgres -d db
	cat ./sql/alter_tables.sql  | docker exec -i psqldb psql -U postgres -d db
	cat ./sql/create_funcs.sql  | docker exec -i psqldb psql -U postgres -d db
	cat ./sql/create_roles.sql  | docker exec -i psqldb psql -U postgres -d db
	cat ./sql/copy_data_from_csv.sql  | docker exec -i psqldb psql -U postgres -d db

.PHONY: restart
restart:
	docker-compose -f docker-compose.yml down
	docker-compose -f docker-compose.yml down
	# sudo rm -rf postgresql_data
	docker-compose -f docker-compose.yml build
	docker-compose -f docker-compose.yml up -d

go:
	go run cmd/app/main.go

# .PHONY: clean
# clean:
# 	docker-compose -f docker-compose.yml down
# 	sudo rm -rf postgresql_data
