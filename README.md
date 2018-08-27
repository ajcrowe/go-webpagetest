go-webpagetest
==============

go-webpagetest is a simple library to interface with the [WebPageTest](http://www.webpagetest.org) RestAPI

[![Build Status](https://travis-ci.org/ajcrowe/go-webpagetest.svg?branch=master)](https://travis-ci.org/ajcrowe/go-webpagetest)
[![GoDoc](https://godoc.org/github.com/ajcrowe/go-webpagetest?status.svg)](https://godoc.org/github.com/ajcrowe/go-webpagetest)

## Features

*	
	### Tests
	* create tests and submit them to a webpagetest instance
	* poll tests as they run and have the results automatically unmarshalled to the Test struct
	* retrieve historic test results from the API by Request ID

* 	
	### Locations
	* list available test locations
	* get the default location tests are run


## Install

You can install this library with `go get`

`go get github.com/ajcrowe/go-webpagetest`


Then import 

```go
import (
	"github.com/ajcrowe/go-webpagetest"
)
```

## Examples

You can find some examples in the `examples` folder for usage

## Variable Name Changes

A number of variables were renamed to correct casing such as `Id -> ID` and `Url -> URL` if you want to continuing using these please see the [`old-variable-name`](https://github.com/ajcrowe/go-webpagetest/tree/old-variable-name) branch