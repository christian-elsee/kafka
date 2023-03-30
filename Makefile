export NAME := $(shell basename "$$PWD" )
export ORG := christianelsee
export SHA := $(shell git rev-parse --short HEAD)
export TS  := $(shell date +%s)

.DEFAULT_GOAL := @goal
.ONESHELL:
.POSIX:

## workflow
@goal: distclean dist build

dist:
	mkdir $@
	go mod init github.com/christian-elsee/kafka ||:
	go mod tidy

	cat manifest/cli.yaml \
		| envsubst \
		| tee $@/manifest.yaml
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo update
	helm template -f helm/values.yaml \
		$(NAME) \
			--create-namespace \
			--namespace=$(NAME) \
		bitnami/$(NAME) \
	| tee -a $@/manifest.yaml

build: dist
	docker build \
		-t local/$(NAME) \
		-t docker.io/$(ORG)/$(NAME) \
		-t docker.io/$(ORG)/$(NAME):$(SHA) \
		.

namespace:
	kubectl create namespace $(NAME) \
		--dry-run=client \
		-oyaml \
	| kubectl apply -f-
	kubectl config set-context \
		--current \
		--namespace $(NAME)

install: build namespace
	<secrets/docker.io.token.gpg gpg -d \
		| xargs -- \
			docker login \
				-u $(ORG) \
				-p

	# if push fails, it may be due to docker image
	# with current sha, not having been built yet
	docker push docker.io/$(ORG)/$(NAME):$(SHA)
	docker push docker.io/$(ORG)/$(NAME):latest
	kubectl apply \
		-f dist/manifest.yaml \
		-n $(NAME)

distclean:
	rm -rvf dist

clean:
	kubectl delete -f dist/manifest.yaml

## ad hoc
push: branch := $(shell git branch --show-current)
push: ;: ## push
	test "$(branch)"

	# ensure working tree is clean for push
	git status --porcelain \
		| xargs \
		| grep -qv .

	ssh-agent bash -c \
		"<secrets/key.gpg gpg -d | ssh-add - \
			&& git push origin $(branch) -f"

lint:
	goimports -lv .
	golint ./...
	go vet ./... ||:

local-build:
	go build -o dist/build main.go
