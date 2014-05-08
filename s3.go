package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"io"
	"net/http"
	"strings"

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

func WriteManifest() error {
	InitS3Client()

	b, err := json.Marshal(manifest)

	if err != nil {
		return err
	}

	err = WriteS3(string(b))
	if err != nil {
		return err
	}

	return nil
}

func WriteS3(str string) error {
	var buf bytes.Buffer
	reader := strings.NewReader(str)
	md5h := md5.New()

	size, err := io.Copy(io.MultiWriter(&buf, md5h), reader)
	if err != nil {
		return err
	}

	err = s3_client.PutObject("draft/manifest2.json", "cjohnsonstore", md5h, size, &buf)
	if err != nil {
		return err
	}

	return nil
}
