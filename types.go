package main

import "time"

type Config struct {
	S3_access_key        string
	S3_secret_access_key string
	S3_hostname          string
}

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

type Manifest []Post

func (m Manifest) Find(shortname string) (bool, Post) {
	for _, post := range m {
		if post.Shortname == shortname {
			return true, post
		}
	}
	return false, Post{}
}

func (m Manifest) Add(shortname string) (bool, Post) {
	found, post := m.Find(shortname)
	if found {
		return false, post
	}

	newpost := Post{Shortname: shortname}
	_ = append(m, newpost)

	return true, newpost
}
