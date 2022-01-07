build-push-color-variants:
	@make build-push-color-ondrejsika
	@make build-push-color-variant BACKGROUND_COLOR=green
	@make build-push-color-variant BACKGROUND_COLOR=blue
	@make build-push-color-variant BACKGROUND_COLOR=yellow
	@make build-push-color-variant BACKGROUND_COLOR=red
	@make build-push-color-cra

build-push-color-variant:
	docker build \
		--build-arg BACKGROUND_COLOR=${BACKGROUND_COLOR} \
		--tag sikalabs/hello-world-server:${BACKGROUND_COLOR} \
		- < Dockerfile
	docker push sikalabs/hello-world-server:${BACKGROUND_COLOR}

build-push-color-ondrejsika:
	docker build \
		--build-arg COLOR="#C7BFB0" \
		--build-arg BACKGROUND_COLOR="#131480" \
		--tag sikalabs/hello-world-server:ondrejsika \
		- < Dockerfile
	docker push sikalabs/hello-world-server:ondrejsika

build-push-color-cra:
	docker build \
		--build-arg COLOR="#F99D1C" \
		--build-arg TEXT="Hello CRA!" \
		--tag sikalabs/hello-world-server:cra \
		- < Dockerfile
	docker push sikalabs/hello-world-server:cra
