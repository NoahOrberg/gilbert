# gilbert
## Build Status
### master branch
[![Build Status](https://travis-ci.org/NoahOrberg/gilbert.svg?branch=master)](https://travis-ci.org/NoahOrberg/gilbert)
### develop branch
[![Build Status](https://travis-ci.org/NoahOrberg/gilbert.svg?branch=develop)](https://travis-ci.org/NoahOrberg/gilbert)

## Installation
```
$ go get github.com/NoahOrberg/gilbert
```
And you should set ENVIRONMENT VARIABLE
```
$ export GILBERT_GISTTOKEN=XXXXXXXX
$ export GILBERT_GISTURL=https://api.github.com/gists
```

## Usage
Example
- README.md will be published to Gist.
```
### upload only one file.
$ gilbert -f README.md
### upload file and description of this.
$ gilbert -f README.md -d "this file is README"
```

