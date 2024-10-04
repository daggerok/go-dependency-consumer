# go-dependency-consumer [![ci](https://github.com/daggerok/go-dependency-consumer/actions/workflows/ci.yaml/badge.svg)](https://github.com/daggerok/go-dependency-consumer/actions/workflows/ci.yaml)
This repository demonstrates how to develop go application, use unit testing and use a custom library as a dependency

```bash
brwe reinstall go
#echo 'export GOPATH=$HOME/go' >> ~/.zsh
#echo 'export PATH=$PATH:/path/to/go/bin:$GOPATH/bin' >> ~/.zsh
go env GOPATH
go version

mkdir go-dependency-consumer && cd $_
go mod init github.com/daggerok/go-dependency-consumer

go get github.com/daggerok/go-dependency-producer/lib
#go: downloading github.com/daggerok/go-dependency-producer v0.0.0-20241004031955-70e75ca163aa
#go: added github.com/daggerok/go-dependency-producer v0.0.0-20241004031955-70e75ca163aa

touch main.go # ...and implement

go run main.go
go run main.go Maksimko

go build        main.go && ./main
go build -o app main.go && ./app

touch main_test.go # ...and implement

go test
go test main_test.go

GOPATH="$(pwd):/home/runner/go/bin" go install .
./bin/go-dependency-consumer Max
./bin/go-dependency-consumer

mkdir .github/workflows
touch .github/workflows/ci.yaml # ...and implement

# git commit...
# git push...
```
