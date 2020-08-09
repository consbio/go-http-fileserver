# go-http-fileserver #


## Purpose ##

This is a simple, lightweight HTTP fileserver that can be run easily from the command line.  We developed it so that we could easily interact with demo websites and their assets.  This even includes CORS support, so you could access resources from different domain names or ports.

*Warning: this does not have any security settings whatsoever.  Anything within any directory below that specified when this is launched will be available from your machine on whatever network can access it!*


## Installation ##

Install **go 1.1+** and setup [GOPATH](http://golang.org/doc/code.html#GOPATH) and GOBIN.  Then:

```
//Dependencies
> go get github.com/codegangsta/cli
> go get github.com/codegangsta/negroni
> go get github.com/phyber/negroni-gzip/gzip

> go get github.com/consbio/go-http-fileserver
```


## Usage ##

```> go-http-fileserver --help``` will print out help information.

Typically, just change to the directory you want served, and:

```> go-http-fileserver``` and it will serve that up on **port 80**.

```> go-http-fileserver -p 8080``` to serve on a different port, such as port 8080.

Then just

```CTRL-C``` to kill the server when you are done.
