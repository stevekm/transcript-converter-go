# go test -v ./...
test:
	set -euo pipefail
	go clean -testcache && \
	go test -v . | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''
.PHONY:test

format:
	go fmt .

run:
	go run main.go "tests/input1.txt" "tests/input2.txt"
