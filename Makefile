
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

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
	 go list -m github.com/danieltaub96/git-faker@v0.0.1