

test/sample/sample.go : spack.go 
	go run spack.go -r static-test/sample > test/sample/sample.go

test/staticmap/staticmap.go: spack.go
	go run spack.go -r static-test/ > test/staticmap/staticmap.go

test : test/sample/sample.go test/staticmap/staticmap.go
	go test -v ./

clean: 
	rm test/sample/sample.go test/staticmap/staticmap.go

