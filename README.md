


- Build and run:

```bash
$ export GO111MODULE=on
$ export GOFLAGS=-mod=vendor
$ go mod download
$ go build -o wiki-short.bin
$ ./wiki-short.bin
```

Server is listening on localhost:8010

## Test

```bash
$ go test -v
=== RUN   TestEmptyTable
--- PASS: TestEmptyTable (0.00s)
=== RUN   TestGetNonExistentProduct
--- PASS: TestGetNonExistentProduct (0.00s)
=== RUN   TestCreateProduct
--- PASS: TestCreateProduct (0.00s)
=== RUN   TestGetProduct
--- PASS: TestGetProduct (0.00s)
=== RUN   TestUpdateProduct
--- PASS: TestUpdateProduct (0.01s)
=== RUN   TestDeleteProduct
--- PASS: TestDeleteProduct (0.01s)
PASS
ok      _/home/tom/r/wiki-short 0.034s
```
