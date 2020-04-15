# make sure we use Go modules for dependency management
export GO111MODULE := auto

# set up working directories
BUILD_TOOL := cmd
TOOLS := cmd build

CMD := ./cmd
BINARY_SEARCH := github.com/AnNiran/karate-chop/binarysearch

test:
	go test ${CMD}
	go test ${BINARY_SEARCH}

mod:
	go mod tidy

# cover:

	cd cmd && go build && cd
	