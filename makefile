run:
	go build -o dynapgen
	./dynapgen

run-release:
	go build -o dynapgen
	ENV=production ./dynapgen

compose-build:
	docker compose up -d --build

compose-down:
	docker compose down --remove-orphans