qa: analyze test

analyze:
	@go vet ./...
	@go run honnef.co/go/tools/cmd/staticcheck@latest --checks=all ./...

build: qa
	@GOOS=windows GOARCH=amd64 go build -o ./build/pwmanager.exe .

build-docker:
	@docker build -t programmierigel/pwmanager .

build-docker-arm :
	@docker build -t programmierigel/pwmanager -f ./Dockerfile_ARM

coverage: test
	@mkdir -p ./coverage
	@go test -coverprofile=./coverage/coverage.out ./...
	@go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
	@open ./coverage/coverage.html

docker-push: build-docker
	@docker push programmierigel/pwmanager
	@docker system prune --all --volumes --force

docker-run: docker-push
	@docker pull programmierigel/pwmanager
	docker run -d -p 3000:3000 -e PASSWORD=123 programmierigel/pwmanager

test:
	@go test -cover ./...

.PHONY: analyze \
	build \
	build-docker \
	coverage \
  docker-push \
  docker-run \
	qa \
	test
