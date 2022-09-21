test:
	docker-compose down
	rm -rf ./db/db_data
	docker-compose up test-sadko-archive-db
#	go test github.com/Vityalimbaev/Example-Backend/tests