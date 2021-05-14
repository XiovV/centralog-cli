personal: build
	scp centralog personal@192.168.0.23:/home/personal

lab1: build
		scp centralog testlab1@192.168.0.25:/home/testlab1

build:
	go build -o centralog
