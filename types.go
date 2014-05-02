package main

import "time"

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

type Config struct {
	S3_access_key        string
	S3_secret_access_key string
	S3_hostname          string
}
