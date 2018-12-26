# gopgweb

[![Go Report Card](https://goreportcard.com/badge/github.com/sayanarijit/gopgweb)](https://goreportcard.com/report/github.com/sayanarijit/gopgweb)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/sayanarijit/gopgweb/blob/master/LICENSE)

This tool uses [gopassgen](https://github.com/sayanarijit/gopassgen) library to generate password based on given policy.

### [Live demo](https://gopgweb.herokuapp.com) | [Download](https://github.com/sayanarijit/gopgweb/archive/master.zip)

***GoPGWeb uses relative path. Hence binary needs to be executed from where it is using `./` syntax via command-line or double-click via GUI***

### Compile and run

```bash
git clone https://github.com/sayanarijit/gopgweb.git
go build
./gopgweb
```

### Command-line help menu

```bash
Usage of gopgweb:
  -host string
        HTTP server url (default "0.0.0.0")
  -port int
        HTTP server port (default 8080)
```
