package main

import (
    "fmt"
    "github.com/hirokidaichi/spack/test/sample"
    "testing"
)

var P = fmt.Println

func TestReader(t *testing.T) {
    element := staticmap.Get("hello.js")
    if element == nil {
        t.Errorf("cannot create gzip")
    }
    bytes := element.UnzipBytes()
    if content := string(bytes); content != "console.log(\"hello.js\")\n" {
        t.Errorf("does not match file content:\n%s\n%s", content)
    }

}
