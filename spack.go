package main

import (
    "bytes"
    "compress/gzip"
    "encoding/base64"
    "github.com/hirokidaichi/go-static-gen/template"
    "github.com/jessevdk/go-flags"
    "io/ioutil"
    "mime"
    "os"
    "path/filepath"
    "strings"
)

type SpackOption struct {
    Root     string `short:"r" long:"root" default:"./static-test" description:"target dir"`
    AllowExt string `short:"e" long:"ext" default:"css,js,html,json,ico,png,jpeg,jpg"`
    allowExt map[string]bool
}

var spackOption *SpackOption

func GetSpackOption() *SpackOption {
    if spackOption != nil {
        return spackOption
    }

    option := new(SpackOption)
    _, err := flags.Parse(option)
    if err != nil {
        os.Exit(1)
    }
    splittedExt := strings.Split(option.AllowExt, ",")
    allowExt := make(map[string]bool)
    for _, e := range splittedExt {
        allowExt[strings.Join([]string{".", e}, "")] = true
    }
    option.allowExt = allowExt
    spackOption = option
    return option
}

func (self *SpackOption) IsAllowExt(item *SpackItem) bool {
    _, ok := self.allowExt[item.Ext]
    return ok
}

type SpackItem struct {
    FilePath string
    Rel      string
    Ext      string
    MimeType string
    length   int
    data     string
}

func (self *SpackItem) Data() string {
    if len(self.data) != 0 {
        return self.data
    }
    data, err := ioutil.ReadFile(self.FilePath)
    if err != nil {
        self.data = ""
        return ""
    }

    self.data = encodeBase64(compress(data))
    return self.data
}

func (self *SpackOption) FindAllFiles() chan *SpackItem {
    emitter := make(chan *SpackItem, 0)
    go func() {
        defer close(emitter)
        filepath.Walk(self.Root,
            func(path string, info os.FileInfo, err error) error {
                if err != nil {
                    return err
                }
                if info.IsDir() {
                    return nil
                }
                ext := filepath.Ext(path)
                rel, _ := filepath.Rel(self.Root, path)
                r := &SpackItem{
                    FilePath: path,
                    Ext:      ext,
                    Rel:      rel,
                    MimeType: mime.TypeByExtension(ext),
                }
                if !self.IsAllowExt(r) {
                    return nil
                }
                emitter <- r
                return nil
            })
    }()
    return emitter
}

type AssignValue struct {
    Name   string
    Option *SpackOption
    Items  chan *SpackItem
}

func encodeBase64(data []byte) string {
    return base64.StdEncoding.EncodeToString(data)
}

func compress(data []byte) []byte {
    b := new(bytes.Buffer)
    w := gzip.NewWriter(b)
    w.Write(data)
    w.Close()
    return b.Bytes()
}

func main() {
    option := GetSpackOption()
    items := option.FindAllFiles()
    err := template.GetTemplate().Execute(os.Stdout, &AssignValue{
        Name:   "staticmap",
        Option: option,
        Items:  items,
    })
    if err != nil {
        panic(err)
    }

}
