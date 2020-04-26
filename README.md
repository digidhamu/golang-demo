# Golang Demo

## Awesome Go
https://awesome-go.com

## HTTP Router & Server
http://www.gorillatoolkit.org/pkg/mux

## Setup Workspace
https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos

https://ahmadawais.com/install-go-lang-on-macos-with-homebrew/

nano ~/.zshrc
```zsh
export GOPATH=$HOME/Documents/digidhamu/go
export PATH=$PATH:$GOPATH/bin
```

## Tutorial
https://www.bogotobogo.com/GoLang/GoLang_Web.php

## Provide Go Path
brew --prefix golang

## Testing
https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http/

https://golang.org/pkg/testing/

https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318

https://github.com/stretchr/testify/

## Tooling
https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/


##  netstat for Ports (on macOS)
`netstat -an -ptcp | grep LISTEN`

`lsof -i -P | grep -i "listen"`

`sudo lsof -i -P | grep -i "listen"`
