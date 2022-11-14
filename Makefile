build:
	docker-compose build
up:
	docker-compose up -d
down:
	docker-compose down
ssh:
	docker-compose exec -it goapp  sh