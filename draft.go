package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"os"

	"github.com/bradfitz/camlistore/pkg/misc/amazon/s3"
)

var ip int
var cmd string
var args []string
var conf Config
var manifest Manifest
var s3_client *s3.Client
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
		exit("You need an argument!!")
	}

	// Set command
	cmd = args[0]
}

func runCmd() {
	switch cmd {
	case "ls":
		listPosts()
	case "add":
		if len(args) <= 1 {
			pln("You need an argument for the \"add\" command.")
		} else {
			addPost(args[1])
		}
	case "test":
		addPost("jimmy")
		addPost("loves")
		addPost("toast")
		listPosts()
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

	readCloser, err := GetManifest(conf)

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
