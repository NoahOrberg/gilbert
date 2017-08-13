# gilbert
[![Build Status](https://travis-ci.org/NoahOrberg/gilbert.svg?branch=master)](https://travis-ci.org/NoahOrberg/gilbert)

## Installation
```
$ go get github.com/NoahOrberg/gilbert
```
And you should set ENVIRONMENT VARIABLE
```
$ export GIST_TOKEN=XXXXXXXX
```
Or Basic Auth
```
$ gilbert -f README.md
Please login
Username: <USER NAME>
Password: <PASSWORD>

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

