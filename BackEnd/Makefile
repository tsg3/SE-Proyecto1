GOFLAGS=env GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=1 CC=arm-linux-gnueabihf-gcc

build: cmake_build
	go build

cross_compile: cmake_build
	@echo "MAKING" 
	$(GOFLAGS) go build

cmake_build:
	mkdir -p build
	cd build && cmake .. && make