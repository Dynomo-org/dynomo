run:
	go build -o dynapgen
	./dynapgen

run-release:
	go build -o dynapgen
	ENV=production ./dynapgen