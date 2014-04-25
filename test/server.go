package main

import (
    "github.com/hirokidaichi/spack/test/staticmap"
    "log"
    "net/http"
)

func main() {
    http.Handle(staticmap.NewStaticServer("/static/"))
    log.Fatal(http.ListenAndServe(":8080", nil))
}
