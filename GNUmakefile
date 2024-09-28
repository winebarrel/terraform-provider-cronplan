.PHONY: build
build: vet
	go build

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	rm -f terraform-provider-cronplan

dev.tfrc: dev.tfrc.tpl
	sed "s|{{PATH_TO_PROVIDER}}|$(shell pwd)|" dev.tfrc.tpl > dev.tfrc

.PHONY: tf-plan
tf-plan: build dev.tfrc
	TF_CLI_CONFIG_FILE=dev.tfrc terraform plan

.PHONY: tf-apply
tf-apply: build dev.tfrc
	TF_CLI_CONFIG_FILE=dev.tfrc terraform apply -auto-approve

.PHONY: tf-clean
tf-clean: clean
	rm -f dev.tfrc terraform.tfstate*

.PHONY: docs
docs:
	cd tools; go generate ./...
