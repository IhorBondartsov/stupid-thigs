run:
	go build .
	./static-service -port=8000 -host=127.0.0.1

clean:
	rm ./static-service

docker:
	docker build -t stastic/service:latest -f ./Dockerfile .
	docker run  -p 8001:8001 stastic/service:latest