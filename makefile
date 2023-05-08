run:
	go build -o dynapgen
	./dynapgen

run-release:
	go build -o dynapgen
	ENV=production ./dynapgen

compose-build:
	docker compose up -d

compose-down:
	docker compose down --remove-orphans