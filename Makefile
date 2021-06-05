
build-all:
	./go-build-all.sh

.PHONY: test
test:
	go test
test-ci:
	go install github.com/vakenbolt/go-test-report/
	go test -json | go-test-report --output out/faker_test_report.html --title "Faker-Tests"

dockertest: genbuild
	docker build . -t go-test-report-test-runner:$(VERSION) -f Tests.Dockerfile