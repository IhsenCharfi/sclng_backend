.PHONY: modsync
modsync:
	go get -u 
	go mod tidy
	go mod vendor

.PHOCY: docker
docker: modsync
	docker build -t sclng_backend .
	docker run --rm -it -p 3000:3000 sclng_backend:latest