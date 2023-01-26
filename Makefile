all: clean build

build:
	go build -o dest/server

clean:
	rm -rf dest