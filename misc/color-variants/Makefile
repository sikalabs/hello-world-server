build-push-color-variants:
	@make build-push-color-ondrejsika
	@make build-push-color-devopslive
	@make build-push-color-variant BACKGROUND_COLOR=green
	@make build-push-color-variant BACKGROUND_COLOR=blue
	@make build-push-color-variant BACKGROUND_COLOR=yellow
	@make build-push-color-variant BACKGROUND_COLOR=red
	@make build-push-color-cra
	@make build-push-color-lekarna-cz
	@make build-push-color-o2
	@make build-push-color-trask
	@make build-push-color-mstp
	@make build-push-color-atos
	@make build-push-color-mavenir
	@make build-push-color-ogsoft

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

build-push-color-devopslive:
	docker build \
		--build-arg COLOR="#C7BFB0" \
		--build-arg BACKGROUND_COLOR="#131480" \
		--build-arg TEXT="Hello DevOps Live!" \
		--tag sikalabs/hello-world-server:devopslive \
		- < Dockerfile
	docker push sikalabs/hello-world-server:devopslive

build-push-color-cra:
	docker build \
		--build-arg COLOR="#F99D1C" \
		--build-arg TEXT="Hello CRA!" \
		--tag sikalabs/hello-world-server:cra \
		- < Dockerfile
	docker push sikalabs/hello-world-server:cra

build-push-color-lekarna-cz:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="green " \
		--build-arg TEXT="Hello Lekarna.cz!" \
		--tag sikalabs/hello-world-server:lekarna-cz \
		- < Dockerfile
	docker push sikalabs/hello-world-server:lekarna-cz

build-push-color-o2:
	docker build \
		--build-arg COLOR="#010066" \
		--build-arg TEXT="Hello O2!" \
		--tag sikalabs/hello-world-server:o2 \
		- < Dockerfile
	docker push sikalabs/hello-world-server:o2

build-push-color-trask:
	docker build \
		--build-arg BACKGROUND_COLOR="#37B0E3" \
		--build-arg COLOR="#204895" \
		--build-arg TEXT="Hello TRASK!" \
		--tag sikalabs/hello-world-server:trask \
		- < Dockerfile
	docker push sikalabs/hello-world-server:trask

build-push-color-mstp:
	docker build \
		--build-arg BACKGROUND_COLOR="#ED8906" \
		--build-arg COLOR="white" \
		--build-arg TEXT="Hello MSTP!" \
		--tag sikalabs/hello-world-server:mstp \
		- < Dockerfile
	docker push sikalabs/hello-world-server:mstp

build-push-color-atos:
	docker build \
		--build-arg BACKGROUND_COLOR="#0696FF" \
		--build-arg COLOR="white" \
		--build-arg TEXT="Hello AtoS!" \
		--tag sikalabs/hello-world-server:atos \
		- < Dockerfile
	docker push sikalabs/hello-world-server:atos

build-push-color-mavenir:
	docker build \
		--build-arg BACKGROUND_COLOR="white" \
		--build-arg COLOR="#00558A" \
		--build-arg TEXT="Hello Mavenir!" \
		--tag sikalabs/hello-world-server:mavenir \
		- < Dockerfile
	docker push sikalabs/hello-world-server:mavenir

build-push-color-ogsoft:
	docker build \
		--build-arg BACKGROUND_COLOR="black" \
		--build-arg COLOR="#F07D00" \
		--build-arg TEXT="Hello OGSOFT!" \
		--tag sikalabs/hello-world-server:ogsoft \
		- < Dockerfile
	docker push sikalabs/hello-world-server:ogsoft
