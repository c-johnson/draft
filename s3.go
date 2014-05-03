package main

import (
	"io"
	"net/http"

	"github.com/bradfitz/camlistore/pkg/misc/amazon/s3"
)

func InitS3Client() {
	if s3_client == nil && &conf != nil {
		s3_client = &s3.Client{
			Auth: &s3.Auth{
				AccessKey:       conf.S3_access_key,
				SecretAccessKey: conf.S3_secret_access_key,
				Hostname:        conf.S3_hostname,
			},
			HTTPClient: http.DefaultClient,
		}
	}
}

func GetManifest() (io.ReadCloser, error) {
	InitS3Client()
	readCloser, _, err := s3_client.Get("cjohnsonstore", "draft/manifest.json")
	return readCloser, err
}

func WriteManifest() {
	InitS3Client()
}
