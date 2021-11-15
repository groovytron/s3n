COVERAGE_FILE=coverage.out

.PHONY: fix
fix:
	go fmt

.PHONY: test
test:
	go test

.PHONY: coverage
coverage:
	go test -coverprofile=$(COVERAGE_FILE)

.PHONY: report
report: coverage
	go tool cover -html=$(COVERAGE_FILE)

.PHONY: clean
clean:
	rm -rf $(COVERAGE_FILE)
