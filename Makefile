#DEV

build-dev:
	docker build -t Video-Chat-App -f containers/images/Dockerfile . && docker build -t turn -f containers/images/Dockerfile.turn .

run-dev:
	docker-compose -f containers/composes/dc.dev.yml up