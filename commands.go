package main

func listPosts() {
	pln("Listing your drafts...")

	if len(manifest) > 0 {
		ps(manifest)
	} else {
		pln("No drafts yet!")
	}
}

func addPost() {
	pln("Adding a post!")

	if len(args) > 1 {
		shortname := args[1]
		found, post := manifest.Find(shortname)
		if found {
			exit("Shit already exists!")
			ps(post)
		} else {
			success, post := manifest.Add(shortname)
			if success {
				pln("Successfully added the item!")
				ps(post)
				exit("")
			}
		}

	} else {
		pln("You need an argument for the \"add\" command.")
	}
}

// func sync() {
//  fmt.Printf("Syncing...\n")

//  buckets, err := client.Buckets()
//  if err != nil {
//    panic(err)
//  }

//  fmt.Printf("Buckets are %s\n", &buckets)

//  client.PutObject("test", "cjohnsonstore", nil, size, body)
// }
