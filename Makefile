MONGO_URI=mongodb://web:web@localhost:27017/local

up:
	docker-compose up --build

migrate.up:
	docker-compose exec server migrate -database "$(MONGO_URI)" -path migrations up

migrate.down:
	docker-compose exec server migrate -database "$(MONGO_URI)" -path migrations down

migrate.create:
	docker-compose exec server migrate create -tz Europe/Moscow -ext bson -dir ./migrations ${name}

mock:
	cd app && go generate ./...

## cat ~/gophkeeper/config.json
## ./cli change_addr -a server:1234
## ./cli login_user -l dkmelnik -p testtest
## ./cli text -c test
