package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/kr/pretty"
)

/* Print debugging statements */

func p(key string, obj interface{}) {
	fmt.Printf("%s: %#v\n", key, obj)
}

func pf(str string, vars ...interface{}) {
	fmt.Printf("str", vars)
}

func pln(str string) {
	fmt.Println(str)
}

func pp(key string, obj interface{}) {
	pretty.Printf("%s: %# v", key, pretty.Formatter(obj))
}

func ps(obj interface{}) {
	spew.Dump(obj)
}

/* Type converstion */

func readerToString(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf.String()
}

/* Program exit and logging */

func exit(msg string) {
	pln(msg)
	os.Exit(0)
}

func logxit(err error) {
	log.Fatal(err)
	exit("There was a fatal error.  Sorry about that.")
}
