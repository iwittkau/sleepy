.PHONY: docker-build
docker-build:
	docker build --build-arg GOPROXY=${GOPROXY} --rm -t  iwittkau/sleepy .

.PHONY: docker-run
docker-run:
	docker run --rm iwittkau/sleepy 10s