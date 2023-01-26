all: clean dependencies build

dependencies:
	go mod download

build:
	go build -o dest/server

clean:
	rm -rf dest