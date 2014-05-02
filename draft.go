package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bradfitz/camlistore/pkg/misc/amazon/s3"
	"github.com/davecgh/go-spew/spew"
	"github.com/kr/pretty"
)

var ip int
var cmd string
var conf Config
var manifest []Post
var client *s3.Client
var DRAFT_DIR = os.Getenv("DRAFT_DIR") // Directory containing blog posts

func main() {
	initialize()

	parseArgs()

	runCmd()
}

func parseArgs() {
	// Parse flags
	flag.IntVar(&ip, "ip", 1234, "some dummy flag")
	flag.Parse()

	// Parse arguments
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("You need an argument!")
		return
	}

	// Set command
	cmd = args[0]
}

func runCmd() {
	switch cmd {
	case "ls":
		listPosts()
	case "add":
		addPost()
	default:
		fmt.Printf("The command \"%s\" doesn't exist.\n", cmd)
	}
}

func p(key string, obj interface{}) {
	fmt.Printf("%s: %#v", key, obj)
}

func pp(key string, obj interface{}) {
	pretty.Printf("%s: %# v", key, pretty.Formatter(obj))
}

func ps(key string, obj interface{}) {
	// spew.Printf("%s: %# v", key, pretty.Formatter(obj))
	spew.Dump(obj)
}

func initialize() {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)

	conf = Config{}
	decoder.Decode(&conf)

	client = &s3.Client{
		Auth: &s3.Auth{
			AccessKey:       conf.S3_access_key,
			SecretAccessKey: conf.S3_secret_access_key,
			Hostname:        conf.S3_hostname,
		},
		HTTPClient: http.DefaultClient,
	}

	readCloser, _, err := client.Get("cjohnsonstore", "draft/manifest.json")

	if err != nil {
		log.Fatal(err)
	} else {
		manifest, err = BuildManifest(readCloser)

		// pp("manifegst", manifest)
	}
}

func BuildManifest(buf io.ReadCloser) ([]Post, error) {
	var posts = make([]Post, 10)

	manifestBytes, err := ioutil.ReadAll(buf)
	if err != nil {
		log.Fatal(err)
	} else {
		err = json.Unmarshal(manifestBytes, &posts)
	}

	return posts, err
}

func flags() {

}

func arguments() {

}

func command() {

}

func listPosts() {
	fmt.Println("Listing your drafts...")

	if len(manifest) > 0 {
		ps("manifest", manifest)
	} else {
		fmt.Println("No drafts yet!")
	}
}

func addPost() {
	fmt.Println("Adding a post!")
}

// func sync() {
// 	fmt.Printf("Syncing...\n")

// 	buckets, err := client.Buckets()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Buckets are %s\n", &buckets)

// 	client.PutObject("test", "cjohnsonstore", nil, size, body)
// }
