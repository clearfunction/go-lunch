all: osx windows linux

osx: hello_osx
windows: hello_windows.exe
linux: hello_linux

clean:
	rm ./*.exe hello_linux hello_osx

hello_linux:
	GOOS=linux GOARCH=386 go build -o hello_linux hello.go

hello_windows.exe:
	GOOS=windows GOARCH=386 go build -o hello_windows.exe hello.go

hello_osx:
	GOOS=darwin GOARCH=386 go build -o hello_osx hello.go
