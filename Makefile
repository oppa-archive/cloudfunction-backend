export AWS_DEFAULT_REGION ?= us-east-1
export AWS_PROFILE := oremi

build:
	sam build
build-cached:
	sam build --cached

clean:
	rm -f ./aws-sam/build

dev:
	make -j dev-watch dev-sam
dev-sam:
	sam local start-api
dev-watch:
	watchexec -f '*.go' 'make -j build'

deploy: build
	sam deploy --no-confirm-changeset --force-upload --no-fail-on-empty-changeset

test:
	go test -v ./...
