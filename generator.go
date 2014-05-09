package main

// func GenerateHTML() {
// 	for i, post := range manifest {

// 	}
// }

/*
	This is all just reference stuff
*/

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path"
// 	"strings"
// 	"time"

// 	"github.com/robfig/revel"
// 	"github.com/russross/blackfriday"
// )

// var (
// 	gopath      = os.Getenv("GOPATH")
// 	APP_ROOT    = path.Join(gopath, "src/github.com/c-johnson/chome")
// 	PUBLIC_ROOT = path.Join(APP_ROOT, "public")
// 	BLOG_OUT    = path.Join(PUBLIC_ROOT, "posts")
// 	BLOG_ROOT   = os.Getenv("BLOG_ROOT")
// )

// /* This is the root of the function that will be called on a CRON / Save file basis which translates .md files into static pages and generates metadata around files in your BLOG_ROOT directory */
// func Generate(overwrite bool, manifest bool) {
// 	if overwrite {
// 		revel.INFO.Printf("Overwriting all posts HTML")
// 		GenerateHtml()
// 	}
// 	if manifest {
// 		revel.INFO.Printf("Saving a new manifest.json file")
// 		SaveManifest()
// 	}
// }

// /*
//   This function returns all of the public posts in the web's manifest file
// */
// func PublicPosts() (posts []Post, err error) {
// 	manifest, err := Manifest()
// 	if err == nil {
// 		for _, post := range manifest {
// 			if post.Public {
// 				posts = append(posts, post)
// 			}
// 		}
// 	}
// 	return
// }

// /* This function scans the manifest file looking for public posts, creating a .html file in the BLOG_OUT directory for each one */
// func GenerateHtml() {
// 	posts_src := path.Join(BLOG_ROOT)
// 	posts_target := path.Join(BLOG_OUT, "out")
// 	os.MkdirAll(posts_target, os.ModePerm)
// 	posts, err := PublicPosts()
// 	if err == nil {
// 		for _, post := range posts {
// 			if post.Public {
// 				html := Compile(posts_src, post.Shortname)
// 				outPath := path.Join(posts_target, post.Shortname+".html")
// 				err := ioutil.WriteFile(outPath, []byte(html), os.ModePerm)
// 				if err != nil {
// 					fmt.Println("Fuck, an err ", err)
// 				}
// 			}
// 		}
// 	}
// }

// /* Retrieves the manifest file as a byte slice, used in some obscure function */
// func ManifestBytes() ([]byte, error) {
// 	fmt.Printf("blog_out = %s", BLOG_OUT)
// 	fullpath := path.Join(BLOG_OUT, "manifest.json")
// 	return ioutil.ReadFile(fullpath)
// }

// /* Retrieves the manifest file */
// func Manifest() ([]Post, error) {
// 	manifestBytes, err := ManifestBytes()
// 	var posts = make([]Post, 1)

// 	if err == nil {
// 		err = json.Unmarshal(manifestBytes, &posts)
// 	}

// 	return posts, err
// }

// /* Writes the manifest object into a .json file in the BLOG_OUT directory */
// func WriteManifest(posts_target string, manifest []Post) {
// 	os.MkdirAll(posts_target, os.ModePerm)
// 	manifestPath := path.Join(posts_target, "manifest.json")

// 	file, err := os.Create(manifestPath)

// 	if err == nil {
// 		manifestJson, err := json.Marshal(manifest)
// 		if err == nil {
// 			var b bytes.Buffer
// 			json.Indent(&b, manifestJson, "", "\t")
// 			b.WriteTo(file)
// 		}
// 	} else {
// 		fmt.Println(err)
// 	}
// }

// /* Gathers the metadata from files in BLOG_ROOT, then calls WriteManifest to persist it */
// func SaveManifest() {
// 	postsData, _ := ioutil.ReadDir(BLOG_ROOT)

// 	manifest := make([]Post, 1)

// 	for _, postData := range postsData {
// 		newPost := Post{Shortname: Shortname(postData.Name()), DateCreated: time.Now(), Public: false}
// 		manifest = append(manifest, newPost)
// 	}

// 	WriteManifest(BLOG_OUT, manifest)
// }

// func Shortname(str string) string {
// 	return strings.Split(str, ".")[0]
// }

// func Compile(root string, shortname string) string {
// 	path := path.Join(root, shortname+".md")
// 	file, err := ioutil.ReadFile(path)

// 	// checksum := crc64.Checksum(file, crc64.MakeTable(crc64.ISO))
// 	// fmt.Println(checksum, "awef")

// 	if err == nil {
// 		return string(blackfriday.MarkdownBasic(file))
// 	} else {
// 		return ""
// 	}
// }

// func FindPublicPost(filename string) (string, error) {
// 	manifest, err := Manifest()

// 	if err == nil {
// 		for _, post := range manifest {
// 			if post.Public && filename == post.Shortname {
// 				var postPath = path.Join(BLOG_OUT, "out", post.Shortname+".html")
// 				return postPath, nil
// 			}
// 		}
// 	}

// 	return "", err
// }
