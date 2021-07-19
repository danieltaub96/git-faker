
build-all:
	./go-build-all.sh

.PHONY: test
test-dev:
	richgo test ./... -v
test:
	go test ./... -v
test-ci:
	mkdir out || echo
	go install github.com/vakenbolt/go-test-report/
	go test -json ./... | go-test-report --output out/faker_test_report.html --title "Faker-Tests"

dockertest:
	docker rmi -f git-faker-test-runner:latest || echo
	docker build . -t git-faker-test-runner:latest -f Tests.Dockerfile