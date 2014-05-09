package main

import (
	"errors"
	"time"
)

type Config struct {
	S3_access_key        string
	S3_secret_access_key string
	S3_hostname          string
	S3_bucket            string
	S3_manifest          string
}

type Manifest []Post

type Post struct {
	Title       string
	Shortname   string
	DateDisplay time.Time
	DateCreated time.Time
	DateUpdated time.Time
	Public      bool
	Checksum    int
	Tags        []string
}

func (m *Manifest) Find(shortname string) (bool, Post) {
	for _, post := range *m {
		if post.Shortname == shortname {
			return true, post
		}
	}
	return false, Post{}
}

func (m *Manifest) Add(shortname string) (error, Post) {
	found, post := m.Find(shortname)
	if found {
		return errors.New("Post already exists"), post
	}

	newpost := Post{Shortname: shortname}
	*m = append(*m, newpost)
	// listPosts()
	ps(m)
	pln("===")
	ps(newpost)
	pln("This is here")

	return nil, newpost
}
