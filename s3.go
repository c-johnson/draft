package main

import (
	"io"
	"net/http"

	"github.com/bradfitz/camlistore/pkg/misc/amazon/s3"
)

func GetManifest(conf Config) (io.ReadCloser, error) {
	client = &s3.Client{
		Auth: &s3.Auth{
			AccessKey:       conf.S3_access_key,
			SecretAccessKey: conf.S3_secret_access_key,
			Hostname:        conf.S3_hostname,
		},
		HTTPClient: http.DefaultClient,
	}

	readCloser, _, err := client.Get("cjohnsonstore", "draft/manifest.json")
	return readCloser, err
}
