# Inspiration: https://github.com/kjcodeacct/go-tflite-cross-platform/blob/cross_compilation_support/_example/cross_complie/Makefile
# In this project `make` is only used to generate native libraries to be bridged in a non go language.
# If you simply want to use an application, you're looking at the wrong file :-).

.PHONY: all, linux_amd64, linter

all: linux_amd64, build_library_android_arm64, build_library_darwin

linux_amd64:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o $(CURDIR)/build/libraries/linux_amd64/libthyra.so -buildmode=c-shared internal/c/main.go

linter:
	chmod +x runlinters.sh; ./runlinters.sh
	
build_library_android_arm64:
# CGO_LDFLAGS+="-fuse-ld=gold"
	CGO_ENABLED=1 GOOS=android GOARCH=arm64 CC=/Users/adrien/Library/Android/sdk/ndk/24.0.8215888/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang CXX=/Users/adrien/Library/Android/sdk/ndk/24.0.8215888/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang++ \
	go build -o $(CURDIR)/build/libraries/arm64-v8a/libthyra.so -buildmode=c-shared internal/c/main.go

build_library_android_x86_64:
	CGO_ENABLED=1 GOOS=android GOARCH=amd64 CC=/Users/adrien/Library/Android/sdk/ndk/24.0.8215888/toolchains/llvm/prebuilt/darwin-x86_64/bin/x86_64-linux-android21-clang CXX=/Users/adrien/Library/Android/sdk/ndk/24.0.8215888/toolchains/llvm/prebuilt/darwin-x86_64/bin/x86_64-linux-android21-clang++ \
	go build -o $(CURDIR)/build/libraries/android_x86_64/libthyra.so -buildmode=c-shared internal/c/main.go

build_library_darwin:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o $(CURDIR)/libraries/darwin_amd64/libthyra.so -buildmode=c-shared internal/c/main.go
