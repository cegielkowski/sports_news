run-server:
	go run ./cmd/api/main.go

build-api:
	go build ./cmd/api/main.go

create-docker-image:
	docker build -t sport_news .

run-docker:
	docker run -p 3000:3000 sport_news

test:
	go test -v ./...

mock:
	mockery --all

