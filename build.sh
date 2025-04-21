rm -r ./build || mkdir ./build && CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -x -o ./build/Sandboxiest\ Sandbox.exe -ldflags "-s -w" && go build -x -o ./build/Sandboxiest\ Sandbox .




