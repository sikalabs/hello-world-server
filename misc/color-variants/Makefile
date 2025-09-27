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
	@make build-push-color-linuxdays
	@make build-push-color-eventlook
	@make build-push-color-adbros
	@make build-push-color-cloud-native-bratislava
	@make build-push-color-continental
	@make build-push-color-cez
	@make build-push-color-vitkovice
	@make build-push-color-ostrava

build-push-color-variant:
	docker build \
		--build-arg BACKGROUND_COLOR=${BACKGROUND_COLOR} \
		--tag sikalabs/hello-world-server:${BACKGROUND_COLOR} \
		--tag ghcr.io/sikalabs/hello-world-server:${BACKGROUND_COLOR} \
		- < Dockerfile
	docker push sikalabs/hello-world-server:${BACKGROUND_COLOR}
	docker push ghcr.io/sikalabs/hello-world-server:${BACKGROUND_COLOR}

build-push-color-ondrejsika:
	docker build \
		--build-arg COLOR="#C7BFB0" \
		--build-arg BACKGROUND_COLOR="#131480" \
		--tag sikalabs/hello-world-server:ondrejsika \
		--tag ghcr.io/sikalabs/hello-world-server:ondrejsika \
		- < Dockerfile
	docker push sikalabs/hello-world-server:ondrejsika
	docker push ghcr.io/sikalabs/hello-world-server:ondrejsika

build-push-color-devopslive:
	docker build \
		--build-arg COLOR="#C7BFB0" \
		--build-arg BACKGROUND_COLOR="#131480" \
		--build-arg TEXT="Hello DevOps Live!" \
		--tag sikalabs/hello-world-server:devopslive \
		--tag ghcr.io/sikalabs/hello-world-server:devopslive \
		- < Dockerfile
	docker push sikalabs/hello-world-server:devopslive
	docker push ghcr.io/sikalabs/hello-world-server:devopslive

build-push-color-cra:
	docker build \
		--build-arg COLOR="#F99D1C" \
		--build-arg TEXT="Hello CRA!" \
		--tag sikalabs/hello-world-server:cra \
		--tag ghcr.io/sikalabs/hello-world-server:cra \
		- < Dockerfile
	docker push sikalabs/hello-world-server:cra
	docker push ghcr.io/sikalabs/hello-world-server:cra

build-push-color-lekarna-cz:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="green " \
		--build-arg TEXT="Hello Lekarna.cz!" \
		--tag sikalabs/hello-world-server:lekarna-cz \
		--tag ghcr.io/sikalabs/hello-world-server:lekarna-cz \
		- < Dockerfile
	docker push sikalabs/hello-world-server:lekarna-cz
	docker push ghcr.io/sikalabs/hello-world-server:lekarna-cz

build-push-color-o2:
	docker build \
		--build-arg COLOR="#010066" \
		--build-arg TEXT="Hello O2!" \
		--tag sikalabs/hello-world-server:o2 \
		--tag ghcr.io/sikalabs/hello-world-server:o2 \
		- < Dockerfile
	docker push sikalabs/hello-world-server:o2
	docker push ghcr.io/sikalabs/hello-world-server:o2

build-push-color-trask:
	docker build \
		--build-arg BACKGROUND_COLOR="#37B0E3" \
		--build-arg COLOR="#204895" \
		--build-arg TEXT="Hello TRASK!" \
		--tag sikalabs/hello-world-server:trask \
		--tag ghcr.io/sikalabs/hello-world-server:trask \
		- < Dockerfile
	docker push sikalabs/hello-world-server:trask
	docker push ghcr.io/sikalabs/hello-world-server:trask

build-push-color-mstp:
	docker build \
		--build-arg BACKGROUND_COLOR="#ED8906" \
		--build-arg COLOR="white" \
		--build-arg TEXT="Hello MSTP!" \
		--tag sikalabs/hello-world-server:mstp \
		--tag ghcr.io/sikalabs/hello-world-server:mstp \
		- < Dockerfile
	docker push sikalabs/hello-world-server:mstp
	docker push ghcr.io/sikalabs/hello-world-server:mstp

build-push-color-atos:
	docker build \
		--build-arg BACKGROUND_COLOR="#0696FF" \
		--build-arg COLOR="white" \
		--build-arg TEXT="Hello AtoS!" \
		--tag sikalabs/hello-world-server:atos \
		--tag ghcr.io/sikalabs/hello-world-server:atos \
		- < Dockerfile
	docker push sikalabs/hello-world-server:atos
	docker push ghcr.io/sikalabs/hello-world-server:atos

build-push-color-mavenir:
	docker build \
		--build-arg BACKGROUND_COLOR="white" \
		--build-arg COLOR="#00558A" \
		--build-arg TEXT="Hello Mavenir!" \
		--tag sikalabs/hello-world-server:mavenir \
		--tag ghcr.io/sikalabs/hello-world-server:mavenir \
		- < Dockerfile
	docker push sikalabs/hello-world-server:mavenir
	docker push ghcr.io/sikalabs/hello-world-server:mavenir

build-push-color-ogsoft:
	docker build \
		--build-arg BACKGROUND_COLOR="black" \
		--build-arg COLOR="#F07D00" \
		--build-arg TEXT="Hello OGSOFT!" \
		--tag sikalabs/hello-world-server:ogsoft \
		--tag ghcr.io/sikalabs/hello-world-server:ogsoft \
		- < Dockerfile
	docker push sikalabs/hello-world-server:ogsoft
	docker push ghcr.io/sikalabs/hello-world-server:ogsoft

build-push-color-fingood:
	docker build \
		--build-arg BACKGROUND_COLOR="#F3F4F6" \
		--build-arg COLOR="#434750" \
		--build-arg TEXT="Hello Fingood!" \
		--tag sikalabs/hello-world-server:fingood \
		--tag ghcr.io/sikalabs/hello-world-server:fingood \
		- < Dockerfile
	docker push sikalabs/hello-world-server:fingood
	docker push ghcr.io/sikalabs/hello-world-server:fingood

build-push-color-trz:
	docker build \
		--build-arg BACKGROUND_COLOR="#9C1012" \
		--build-arg COLOR="white" \
		--build-arg TEXT="Hello TRZ!" \
		--tag sikalabs/hello-world-server:trz \
		--tag ghcr.io/sikalabs/hello-world-server:trz \
		- < Dockerfile
	docker push sikalabs/hello-world-server:trz
	docker push ghcr.io/sikalabs/hello-world-server:trz

build-push-color-thermofisher:
	docker build \
		--build-arg COLOR="#DB1215" \
		--build-arg TEXT="Hello Thermo Fisher!" \
		--tag sikalabs/hello-world-server:thermofisher \
		--tag ghcr.io/sikalabs/hello-world-server:thermofisher \
		- < Dockerfile
	docker push sikalabs/hello-world-server:thermofisher
	docker push ghcr.io/sikalabs/hello-world-server:thermofisher

build-push-color-aardwark:
	docker build \
		--build-arg COLOR="#734FAA" \
		--build-arg TEXT="Hello Aardwark!" \
		--tag sikalabs/hello-world-server:aardwark \
		--tag ghcr.io/sikalabs/hello-world-server:aardwark \
		- < Dockerfile
	docker push sikalabs/hello-world-server:aardwark
	docker push ghcr.io/sikalabs/hello-world-server:aardwark

build-push-color-ttc:
	docker build \
		--build-arg COLOR="#0AB7E6" \
		--build-arg TEXT="Hello TTC!" \
		--tag sikalabs/hello-world-server:ttc \
		--tag ghcr.io/sikalabs/hello-world-server:ttc \
		- < Dockerfile
	docker push sikalabs/hello-world-server:ttc
	docker push ghcr.io/sikalabs/hello-world-server:ttc

build-push-color-tietoevry:
	docker build \
		--build-arg TEXT="Hello Tietoevry!" \
		--build-arg BACKGROUND_COLOR="#280071" \
		--build-arg COLOR="white" \
		--tag sikalabs/hello-world-server:tietoevry \
		--tag ghcr.io/sikalabs/hello-world-server:tietoevry \
		- < Dockerfile
	docker push sikalabs/hello-world-server:tietoevry
	docker push ghcr.io/sikalabs/hello-world-server:tietoevry

build-push-color-animato:
	docker build \
		--build-arg COLOR="#F80200" \
		--build-arg TEXT="Hello Animato!" \
		--tag sikalabs/hello-world-server:animato \
		--tag ghcr.io/sikalabs/hello-world-server:animato \
		- < Dockerfile
	docker push sikalabs/hello-world-server:animato
	docker push ghcr.io/sikalabs/hello-world-server:animato

build-push-color-spcss:
	docker build \
		--build-arg COLOR="#009ADB" \
		--build-arg TEXT="Hello SPCSS!" \
		--tag sikalabs/hello-world-server:spcss \
		--tag ghcr.io/sikalabs/hello-world-server:spcss \
		- < Dockerfile
	docker push sikalabs/hello-world-server:spcss
	docker push ghcr.io/sikalabs/hello-world-server:spcss

build-push-color-frequentis:
	docker build \
		--build-arg COLOR="#F80200" \
		--build-arg TEXT="Hello Frequentis!" \
		--tag sikalabs/hello-world-server:frequentis \
		--tag ghcr.io/sikalabs/hello-world-server:frequentis \
		- < Dockerfile
	docker push sikalabs/hello-world-server:frequentis
	docker push ghcr.io/sikalabs/hello-world-server:frequentis

build-push-color-digiteq:
	docker build \
		--build-arg COLOR="#185DB2" \
		--build-arg TEXT="Hello Digiteq!" \
		--tag sikalabs/hello-world-server:digiteq \
		--tag ghcr.io/sikalabs/hello-world-server:digiteq \
		- < Dockerfile
	docker push sikalabs/hello-world-server:digiteq
	docker push ghcr.io/sikalabs/hello-world-server:digiteq

build-push-color-heidelbergmaterials:
	docker build \
		--build-arg COLOR="#004E2B" \
		--build-arg TEXT="Hello Heidelberg Materials!" \
		--tag sikalabs/hello-world-server:heidelbergmaterials \
		--tag ghcr.io/sikalabs/hello-world-server:heidelbergmaterials \
		- < Dockerfile
	docker push sikalabs/hello-world-server:heidelbergmaterials
	docker push ghcr.io/sikalabs/hello-world-server:heidelbergmaterials

build-push-color-eon:
	docker build \
		--build-arg COLOR="#EA1B0A" \
		--build-arg TEXT="Hello EON!" \
		--tag sikalabs/hello-world-server:eon \
		--tag ghcr.io/sikalabs/hello-world-server:eon \
		- < Dockerfile
	docker push sikalabs/hello-world-server:eon
	docker push ghcr.io/sikalabs/hello-world-server:eon

build-push-color-ceps:
	docker build \
		--build-arg BACKGROUND_COLOR="#C1251F" \
		--build-arg COLOR="white" \
		--build-arg TEXT="Hello CEPS!" \
		--tag sikalabs/hello-world-server:ceps \
		--tag ghcr.io/sikalabs/hello-world-server:ceps \
		- < Dockerfile
	docker push sikalabs/hello-world-server:ceps
	docker push ghcr.io/sikalabs/hello-world-server:ceps

build-push-color-generali:
	docker build \
		--build-arg COLOR="#C21B18" \
		--build-arg TEXT="Hello Generali!" \
		--tag sikalabs/hello-world-server:generali \
		--tag ghcr.io/sikalabs/hello-world-server:generali \
		- < Dockerfile
	docker push sikalabs/hello-world-server:generali
	docker push ghcr.io/sikalabs/hello-world-server:generali

build-push-color-elm:
	docker build \
		--build-arg COLOR="#818285" \
		--build-arg TEXT="Hello ELM!" \
		--tag sikalabs/hello-world-server:elm \
		--tag ghcr.io/sikalabs/hello-world-server:elm \
		- < Dockerfile
	docker push sikalabs/hello-world-server:elm
	docker push ghcr.io/sikalabs/hello-world-server:elm

build-push-color-t-mobile:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#E20074" \
		--build-arg TEXT="Hello T-Mobile!" \
		--tag sikalabs/hello-world-server:t-mobile \
		--tag ghcr.io/sikalabs/hello-world-server:t-mobile \
		- < Dockerfile
	docker push sikalabs/hello-world-server:t-mobile
	docker push ghcr.io/sikalabs/hello-world-server:t-mobile

build-push-color-jt:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="black" \
		--build-arg TEXT="Hello J&T!" \
		--tag sikalabs/hello-world-server:jt \
		--tag ghcr.io/sikalabs/hello-world-server:jt \
		- < Dockerfile
	docker push sikalabs/hello-world-server:jt
	docker push ghcr.io/sikalabs/hello-world-server:jt


build-push-color-rozhlas:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#07183D" \
		--build-arg TEXT="Hello Rozhlas!" \
		--tag sikalabs/hello-world-server:rozhlas \
		--tag ghcr.io/sikalabs/hello-world-server:rozhlas \
		- < Dockerfile
	docker push sikalabs/hello-world-server:rozhlas
	docker push ghcr.io/sikalabs/hello-world-server:rozhlas

build-push-color-era:
	docker build \
		--build-arg COLOR="#B8173E" \
		--build-arg TEXT="Hello ERA!" \
		--tag sikalabs/hello-world-server:era \
		--tag ghcr.io/sikalabs/hello-world-server:era \
		- < Dockerfile
	docker push sikalabs/hello-world-server:era
	docker push ghcr.io/sikalabs/hello-world-server:era

build-push-color-ogilvy:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#EB3F42" \
		--build-arg TEXT="Hello Ogilvy!" \
		--tag sikalabs/hello-world-server:ogilvy \
		--tag ghcr.io/sikalabs/hello-world-server:ogilvy \
		- < Dockerfile
	docker push sikalabs/hello-world-server:ogilvy
	docker push ghcr.io/sikalabs/hello-world-server:ogilvy

build-push-color-dhl:
	docker build \
		--build-arg COLOR="#D40410" \
		--build-arg BACKGROUND_COLOR="#FFCC01" \
		--build-arg TEXT="Hello DHL!" \
		--tag sikalabs/hello-world-server:dhl \
		--tag ghcr.io/sikalabs/hello-world-server:dhl \
		- < Dockerfile
	docker push sikalabs/hello-world-server:dhl
	docker push ghcr.io/sikalabs/hello-world-server:dhl

build-push-color-kcd:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#0D304E" \
		--build-arg TEXT="Hello KCD!" \
		--tag sikalabs/hello-world-server:kcd \
		--tag ghcr.io/sikalabs/hello-world-server:kcd \
		- < Dockerfile
	docker push sikalabs/hello-world-server:kcd
	docker push ghcr.io/sikalabs/hello-world-server:kcd

build-push-color-gpe:
	docker build \
		--build-arg COLOR="#013C73" \
		--build-arg BACKGROUND_COLOR="white" \
		--build-arg TEXT="Hello GPE!" \
		--tag sikalabs/hello-world-server:gpe \
		--tag ghcr.io/sikalabs/hello-world-server:gpe \
		- < Dockerfile
	docker push sikalabs/hello-world-server:gpe
	docker push ghcr.io/sikalabs/hello-world-server:gpe

build-push-color-exainfra:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#426690" \
		--build-arg TEXT="Hello EXA Infra!" \
		--tag sikalabs/hello-world-server:exainfra \
		--tag ghcr.io/sikalabs/hello-world-server:exainfra \
		- < Dockerfile
	docker push sikalabs/hello-world-server:exainfra
	docker push ghcr.io/sikalabs/hello-world-server:exainfra

build-push-color-linuxdays:
	docker build \
		--build-arg COLOR="#FAAF3A" \
		--build-arg BACKGROUND_COLOR="#EEEEEE" \
		--build-arg TEXT="Hello LinuxDays!" \
		--tag sikalabs/hello-world-server:linuxdays \
		--tag ghcr.io/sikalabs/hello-world-server:linuxdays \
		- < Dockerfile
	docker push sikalabs/hello-world-server:linuxdays
	docker push ghcr.io/sikalabs/hello-world-server:linuxdays

build-push-color-eventlook:
	docker build \
		--build-arg COLOR="#F9541C" \
		--build-arg BACKGROUND_COLOR="#160F1A" \
		--build-arg TEXT="Hello EventLook!" \
		--tag sikalabs/hello-world-server:eventlook \
		--tag ghcr.io/sikalabs/hello-world-server:eventlook \
		- < Dockerfile
	docker push sikalabs/hello-world-server:eventlook
	docker push ghcr.io/sikalabs/hello-world-server:eventlook

build-push-color-adbros:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#B51A82" \
		--build-arg TEXT="Hello Adbros!" \
		--tag sikalabs/hello-world-server:adbros \
		--tag ghcr.io/sikalabs/hello-world-server:adbros \
		- < Dockerfile
	docker push sikalabs/hello-world-server:adbros
	docker push ghcr.io/sikalabs/hello-world-server:adbros

build-push-color-cloud-native-bratislava:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#042285" \
		--build-arg TEXT="Hello Cloud Native Bratislava!" \
		--tag sikalabs/hello-world-server:cloud-native-bratislava \
		--tag ghcr.io/sikalabs/hello-world-server:cloud-native-bratislava \
		- < Dockerfile
	docker push sikalabs/hello-world-server:cloud-native-bratislava
	docker push ghcr.io/sikalabs/hello-world-server:cloud-native-bratislava

build-push-color-continental:
	docker build \
		--build-arg COLOR="black" \
		--build-arg BACKGROUND_COLOR="#FFA500" \
		--build-arg TEXT="Hello Continental!" \
		--tag sikalabs/hello-world-server:continental \
		--tag ghcr.io/sikalabs/hello-world-server:continental \
		- < Dockerfile
	docker push sikalabs/hello-world-server:continental
	docker push ghcr.io/sikalabs/hello-world-server:continental

build-push-color-cez:
	docker build \
		--build-arg COLOR="#F24F01" \
		--build-arg BACKGROUND_COLOR="white" \
		--build-arg TEXT="Hello CEZ!" \
		--tag sikalabs/hello-world-server:cez \
		--tag ghcr.io/sikalabs/hello-world-server:cez \
		- < Dockerfile
	docker push sikalabs/hello-world-server:cez
	docker push ghcr.io/sikalabs/hello-world-server:cez

build-push-color-vitkovice:
	docker build \
		--build-arg COLOR="white" \
		--build-arg BACKGROUND_COLOR="#21334C" \
		--build-arg TEXT="Hello VITKOVICE!" \
		--tag sikalabs/hello-world-server:vitkovice \
		--tag ghcr.io/sikalabs/hello-world-server:vitkovice \
		- < Dockerfile
	docker push sikalabs/hello-world-server:vitkovice
	docker push ghcr.io/sikalabs/hello-world-server:vitkovice
		--tag ghcr.io/sikalabs/hello-world-server:vitkovice \
		- < Dockerfile
	docker push sikalabs/hello-world-server:vitkovice
	docker push ghcr.io/sikalabs/hello-world-server:vitkovice

build-push-color-ostrava:
	docker build \
		--build-arg COLOR="#00AFD3" \
		--build-arg BACKGROUND_COLOR="white" \
		--build-arg TEXT="Hello OSTRAVA!!!" \
		--tag sikalabs/hello-world-server:ostrava \
		--tag ghcr.io/sikalabs/hello-world-server:ostrava \
		- < Dockerfile
	docker push sikalabs/hello-world-server:ostrava
	docker push ghcr.io/sikalabs/hello-world-server:ostrava
