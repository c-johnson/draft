package main

func listPosts() {
	pln("Listing your drafts...")

	if len(manifest) > 0 {
		ps(manifest)
	} else {
		pln("No drafts yet!")
	}
}

func addPost(shortname string) error {
	pln("Adding a post!")

	err, _ := manifest.Add(shortname)
	if err != nil {
		return err
	} else {
		pln("Successfully added the item!")
		pln("Now saving the manifest to S3...")

		err = WriteManifest(conf, manifest)

		if err != nil {
			exit(err.Error())
		}

		exit("Successfully saved!")
	}

	return nil
}
