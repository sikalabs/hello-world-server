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
	@make build-push-color-trz
	@make build-push-color-thermofisher
	@make build-push-color-aardwark
	@make build-push-color-ttc
	@make build-push-color-tietoevry
	@make build-push-color-animato
	@make build-push-color-digiteq
	@make build-push-color-heidelbergmaterials
	@make build-push-color-eon
	@make build-push-color-ceps
	@make build-push-color-generali
	@make build-push-color-elm
	@make build-push-color-t-mobile
	@make build-push-color-jt
	@make build-push-color-rozhlas
	@make build-push-color-era
	@make build-push-color-ogilvy
	@make build-push-color-dhl
	@make build-push-color-kcd
	@make build-push-color-gpe
	@make build-push-color-exainfra

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

build-push-color-fingood:
	docker build \
		--build-arg BACKGROUND_COLOR="#F3F4F6" \
		--build-arg COLOR="#434750" \
		--build-arg TEXT="Hello Fingood!" \
		--tag sikalabs/hello-world-server:fingood \
		- < Dockerfile
	docker push sikalabs/hello-world-server:fingood

build-push-color-trz:
	docker build \
		--build-arg BACKGROUND_COLOR="#9C1012" \
		--build-arg COLOR="white" \
		--build-arg TEXT="Hello TRZ!" \
		--tag sikalabs/hello-world-server:trz \
		- < Dockerfile
	docker push sikalabs/hello-world-server:trz

build-push-color-thermofisher:
	docker build \
		--build-arg COLOR="#DB1215" \
		--build-arg TEXT="Hello Thermo Fisher!" \
		--tag sikalabs/hello-world-server:thermofisher \
		- < Dockerfile
	docker push sikalabs/hello-world-server:thermofisher

build-push-color-aardwark:
	docker build \
		--build-arg COLOR="#734FAA" \
		--build-arg TEXT="Hello Aardwark!" \
		--tag sikalabs/hello-world-server:aardwark \
		- < Dockerfile
	docker push sikalabs/hello-world-server:aardwark

build-push-color-ttc:
	docker build \
		--build-arg COLOR="#0AB7E6" \
		--build-arg TEXT="Hello TTC!" \
		--tag sikalabs/hello-world-server:ttc \
		- < Dockerfile
	docker push sikalabs/hello-world-server:ttc

build-push-color-tietoevry:
	docker build \
		--build-arg TEXT="Hello Tietoevry!" \
		--build-arg BACKGROUND_COLOR="#280071" \
		--build-arg COLOR="white" \
		--tag sikalabs/hello-world-server:tietoevry \
		- < Dockerfile
	docker push sikalabs/hello-world-server:tietoevry

build-push-color-animato:
	docker build \
		--build-arg COLOR="#F80200" \
		--build-arg TEXT="Hello Animato!" \
		--tag sikalabs/hello-world-server:animato \
		- < Dockerfile
	docker push sikalabs/hello-world-server:animato

build-push-color-spcss:
	docker build \
		--build-arg COLOR="#009ADB" \
		--build-arg TEXT="Hello SPCSS!" \
		--tag sikalabs/hello-world-server:spcss \
		- < Dockerfile
	docker push sikalabs/hello-world-server:spcss

build-push-color-frequentis:
	docker build \
		--build-arg COLOR="#F80200" \
		--build-arg TEXT="Hello Frequentis!" \
		--tag sikalabs/hello-world-server:frequentis \
		- < Dockerfile
	docker push sikalabs/hello-world-server:frequentis

build-push-color-digiteq:
	docker build \
		--build-arg COLOR="#185DB2" \
		--build-arg TEXT="Hello Digiteq!" \
		--tag sikalabs/hello-world-server:digiteq \
		- < Dockerfile
	docker push sikalabs/hello-world-server:digiteq

build-push-color-heidelbergmaterials:
	docker build \
		--build-arg COLOR="#004E2B" \
		--build-arg TEXT="Hello Heidelberg Materials!" \
		--tag sikalabs/hello-world-server:heidelbergmaterials \
		- < Dockerfile
	docker push sikalabs/hello-world-server:heidelbergmaterials

build-push-color-eon:
	docker build \
		--build-arg COLOR="#EA1B0A" \
		--build-arg TEXT="Hello EON!" \
		--tag sikalabs/hello-world-server:eon \
		- < Dockerfile
	docker push sikalabs/hello-world-server:eon

build-push-color-ceps:
	docker build \
		--build-arg BACKGROUND_COLOR="#C1251F" \
		--build-arg COLOR="white" \
		--build-arg TEXT="Hello CEPS!" \
		--tag sikalabs/hello-world-server:ceps \
		- < Dockerfile
	docker push sikalabs/hello-world-server:ceps

build-push-color-generali:
	docker build \
		--build-arg COLOR="#C21B18" \
		--build-arg TEXT="Hello Generali!" \
		--tag sikalabs/hello-world-server:generali \
		- < Dockerfile
	docker push sikalabs/hello-world-server:generali

build-push-color-elm:
	docker build \
		--build-arg COLOR="#818285" \
		--build-arg TEXT="Hello ELM!" \
		--tag sikalabs/hello-world-server:elm \
		- < Dockerfile
	docker push sikalabs/hello-world-server:elm

build-push-color-t-mobile:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#E20074" \
		--build-arg TEXT="Hello T-Mobile!" \
		--tag sikalabs/hello-world-server:t-mobile \
		- < Dockerfile
	docker push sikalabs/hello-world-server:t-mobile

build-push-color-jt:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="black" \
		--build-arg TEXT="Hello J&T!" \
		--tag sikalabs/hello-world-server:jt \
		- < Dockerfile
	docker push sikalabs/hello-world-server:jt


build-push-color-rozhlas:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#07183D" \
		--build-arg TEXT="Hello Rozhlas!" \
		--tag sikalabs/hello-world-server:rozhlas \
		- < Dockerfile
	docker push sikalabs/hello-world-server:rozhlas

build-push-color-era:
	docker build \
		--build-arg COLOR="#B8173E" \
		--build-arg TEXT="Hello ERA!" \
		--tag sikalabs/hello-world-server:era \
		- < Dockerfile
	docker push sikalabs/hello-world-server:era

build-push-color-ogilvy:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#EB3F42" \
		--build-arg TEXT="Hello Ogilvy!" \
		--tag sikalabs/hello-world-server:ogilvy \
		- < Dockerfile
	docker push sikalabs/hello-world-server:ogilvy

build-push-color-dhl:
	docker build \
		--build-arg COLOR="#D40410" \
		--build-arg BACKGROUND_COLOR="#FFCC01" \
		--build-arg TEXT="Hello DHL!" \
		--tag sikalabs/hello-world-server:dhl \
		- < Dockerfile
	docker push sikalabs/hello-world-server:dhl

build-push-color-kcd:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#0D304E" \
		--build-arg TEXT="Hello KCD!" \
		--tag sikalabs/hello-world-server:kcd \
		- < Dockerfile
	docker push sikalabs/hello-world-server:kcd

build-push-color-gpe:
	docker build \
		--build-arg COLOR="#013C73" \
		--build-arg BACKGROUND_COLOR="white" \
		--build-arg TEXT="Hello GPE!" \
		--tag sikalabs/hello-world-server:gpe \
		- < Dockerfile
	docker push sikalabs/hello-world-server:gpe

build-push-color-exainfra:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#426690" \
		--build-arg TEXT="Hello EXA Infra!" \
		--tag sikalabs/hello-world-server:exainfra \
		- < Dockerfile
	docker push sikalabs/hello-world-server:exainfra
