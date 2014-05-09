package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/russross/blackfriday"
)

func listPosts() {
	pln("Listing your drafts...")

	if len(manifest) > 0 {
		ps(manifest)
	} else {
		pln("No drafts yet!")
	}

	sync()
}

func addPost(shortname string) {
	pln("Adding a post!")

	err, _ := manifest.Add(shortname)
	if err != nil {
		logxit(err)
	} else {
		pln("Successfully added the item!")
		pln("Now saving the manifest to S3...")

		err = WriteManifest(conf, manifest)

		if err != nil {
			exit(err.Error())
		}

		exit("Successfully saved!")
	}
}

func sync() {
	if DRAFT_DIR == "" {
		exit("You need to set your $DRAFT_DIR variable to point to your drafts folder")
	}

	files, err := filesFromDirString(DRAFT_DIR)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			// Do nothing
		} else {
			fragments := strings.Split(file.Name(), ".")
			shortname := fragments[0]

			inManifest, _ := manifest.Find(shortname)

			if inManifest {
				fullpath := path.Join(DRAFT_DIR, file.Name())
				err = WriteDraft(fullpath, conf.S3_bucket)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func generate() {
	for _, post := range manifest {
		draft, err := GetDraft(post.Shortname, conf.S3_bucket)
		if err != nil {
			log.Fatal(err)
		}

		postBytes := blackfriday.MarkdownBasic([]byte(draft))
		ioutil.WriteFile(post.Shortname+".html", postBytes, 0644)
	}
}

func filesFromDirString(str string) ([]os.FileInfo, error) {
	draftDir, err := os.Open(str) // For read access.
	var files []os.FileInfo
	if err == nil {
		files, err = draftDir.Readdir(0)
	}

	return files, err
}

func test() {
	sync()
}
