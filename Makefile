transfer: build
	scp centralog personal@192.168.0.23:/home/personal
build:
	go build -o centralog
