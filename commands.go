package main

func listPosts() {
	pln("Listing your drafts...")

	if len(manifest) > 0 {
		ps("manifest", manifest)
	} else {
		pln("No drafts yet!")
	}
}

func addPost() {
	pln("Adding a post!")

	if len(args) > 1 {
		p("Shortname is", args[1])
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
