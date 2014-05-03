package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bradfitz/camlistore/pkg/misc/amazon/s3"
)

var ip int
var cmd string
var args []string
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
	args = flag.Args()
	if len(args) == 0 {
		exit("You need an argument!")
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
		exit("The command you wrote doesn't exist.")
	}
}

func initialize() {
	file, err := os.Open("conf.json")

	if err != nil {
		exit("There is no conf.json file in your local directory.")
	}

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
		logxit(err)
	} else {
		manifest, err = BuildManifest(readCloser)
	}
}

func BuildManifest(buf io.ReadCloser) ([]Post, error) {
	var posts = make([]Post, 10)

	manifestBytes, err := ioutil.ReadAll(buf)
	if err != nil {
		logxit(err)
	} else {
		err = json.Unmarshal(manifestBytes, &posts)
	}

	return posts, err
}
