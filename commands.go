package main

import (
	"fmt"
	"os"
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

	draftDir, err := os.Open(DRAFT_DIR) // For read access.
	if err != nil {
		logxit(err)
	}

	files, err := draftDir.Readdir(0)
	if err != nil {
		logxit(err)
	}

	for _, file := range files {
		if file.IsDir() {
			// Do nothing
		} else {

			pln("name = " + file.Name())
		}
	}

	pln("names = !")
	fmt.Printf("how many files? = %v", len(files))
	// ps(names)
}
