# whiplash.go
whiplash.go is a command line game that a parody of the movie "Whiplash".

## Overview
When you type a sentence, whiplash.go checks it's tempo and print feedback. (During the testing, I recognized that I have no sense of rhythm.:))

Most of sentences are a parody of the movie '[Whiplash](http://www.imdb.com/title/tt2582802/)'. And the intro sentence is inspired by ["The Best Programming Advice I Ever Got" with Rob Pike](http://www.informit.com/articles/article.aspx?p=1941206).

## Installation
Make sure you have a working Go environment (go 1.3 or above is *required*). [See the install instructions](http://golang.org/doc/install.html).

To install `whiplash.go`, simply run:
```
$ go get github.com/dorajistyle/whiplash
```

Make sure your `PATH` includes to the `$GOPATH/bin` directory so your commands can be easily used:
```
export PATH=$PATH:$GOPATH/bin
```

## Getting Started
```
$ whiplash
```

## Help
```
$ whiplash help
```

## Options

* --bpm "400.00"		BPM. The game's tempo.
* --text "hello world."	Text that you want to type
* --skipIntro "false"		Use 'true' when you want to skip intro
* --skipChecking "false"	Use 'true' when you want to skip checking that you were rushing or dragging.
* --level "normal"		easy|normal|hard|guru
