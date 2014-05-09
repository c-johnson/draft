package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bradfitz/camlistore/pkg/misc/amazon/s3"
)

type draftS3 struct {
	conf Config
}

func initS3Client(conf Config) {
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

func GetManifest(conf Config) (io.ReadCloser, error) {
	initS3Client(conf)
	return ReadS3(conf.S3_bucket, conf.S3_manifest)
}

func WriteManifest(conf Config, manifest Manifest) error {
	initS3Client(conf)

	manifestStr, err := json.Marshal(manifest)

	pln(string(manifestStr))

	if err != nil {
		return err
	}

	err = WriteS3(conf.S3_bucket, conf.S3_manifest, string(manifestStr))
	if err != nil {
		return err
	}

	return nil
}

func WriteDraft(fullpath string, bucket string) error {
	contents, err := ioutil.ReadFile(fullpath)
	if err != nil {
		return err
	}

	tokens := strings.Split(fullpath, "/")
	shortname := tokens[len(tokens)-1]

	return WriteS3(bucket, "draft/drafts/"+shortname, string(contents))
}

func GetDraft(shortname string, bucket string) (string, error) {
	rc, err := ReadS3(bucket, "draft/drafts/"+shortname)
	return readerToString(rc), err
}

func WriteS3(bucket string, key string, content string) error {
	var buf bytes.Buffer
	reader := strings.NewReader(content)
	md5h := md5.New()

	size, err := io.Copy(io.MultiWriter(&buf, md5h), reader)
	if err != nil {
		return err
	}

	err = s3_client.PutObject(key, bucket, md5h, size, &buf)
	if err != nil {
		return err
	}

	return nil
}

func ReadS3(bucket string, key string) (io.ReadCloser, error) {
	readCloser, _, err := s3_client.Get(bucket, key)
	return readCloser, err
}
