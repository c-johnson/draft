package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/bradfitz/camlistore/pkg/misc/amazon/s3"
)

var conf Config

var cmd string
var args []string

var manifest Manifest
var s3_client *s3.Client
var DRAFT_DIR = os.Getenv("DRAFT_DIR") // Directory containing blog posts

func main() {
	loadConfig()

	initialize()

	parseArgs()

	runCmd()
}

func loadConfig() {
	file, err := os.Open("conf.json")

	if err != nil {
		exit("There is no conf.json file in your local directory.")
	}

	decoder := json.NewDecoder(file)

	conf = Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}

	if DRAFT_DIR == "" {
		exit("You need a $DRAFT_DIR variable!  See the readme.md for details")
	}
}

func initialize() {
	readCloser, err := GetManifest(conf)

	if err != nil {
		logxit(err)
	} else {
		manifest, err = BuildManifest(readCloser)
	}
}

func parseArgs() {
	// Parse flags
	// flag.IntVar(&ip, "ip", 1234, "some dummy flag")
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
	case "sync":
		sync()
	case "test":
		test()
	case "add":
		if len(args) <= 1 {
			pln("You need an argument for the \"add\" command.")
		} else {
			addPost(args[1])
		}
	case "generate":
		generate()
	default:
		exit("The command you wrote doesn't exist.")
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
