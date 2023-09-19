run:
	go build -o dynapgen
	./dynapgen

compose-build:
	docker compose up -d

compose-down:
	docker compose down --remove-orphans