


- Build and run:

```bash
$ export GO111MODULE=on
$ export GOFLAGS=-mod=vendor
$ go mod download
$ go build -o wiki-short.bin
$ ./wiki-short.bin
```

Server is listening on [localhost:8010/search/Yoshua%20Bengio](http://localhost:8010/search/Yoshua%20Bengio)

to enter new search/input change the last part of url e.g

[localhost:8010/search/Bill Gates](http://localhost:8010/search/Bill%20Gates)
