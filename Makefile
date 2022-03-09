docker.start.components:
	docker-compose -f docker-compose.yml up -d --remove-orphans postgres;

run.test:
	go test ./... -v -coverprofile=coverage.out