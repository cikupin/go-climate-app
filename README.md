# Simple Go Climate
Only a really simple web app created from Go. This web obtains data by calling API from openweathermap.org


## Quick Start
### 1. Go and Dep Installation
##### A. Go
1. Download Go & follow the installation instruction found here: https://golang.org/doc/install

##### B. Go Dep
1. Go Dep is the official *"experiment"* Go package manager. Please follow the instruction here https://github.com/golang/dep

### 2. Build
To build :

    $ [ -d $GOPATH/src/github.com/cikupin ] || mkdir -p $GOPATH/src/github.com/cikupin
    $ cd $GOPATH/src/github.com/cikupin
    $ git clone git@github.com:cikupin/go-climate-app.git
    $ cd go-climate-app
    $
    $ dep ensure
    $ go build -race

### 3. Run
To run :

    $ ./go-climate-app 