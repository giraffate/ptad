package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	URLFormat = "https://papertrailapp.com/api/v1/archives/%s/download"
)

// PaperTrailClient is the client for papertrail.
type PaperTrailClient struct {
	token string
	*http.Client
}

// NewPaperTrailClient is the constructor of PaperTrailClient.
func NewPaperTrailClient(token string) *PaperTrailClient {
	return &PaperTrailClient{
		token:  token,
		Client: &http.Client{},
	}
}

// NewRequest returns a new http.Request.
func (c *PaperTrailClient) NewRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Papertrail-Token", c.token)
	return req, nil
}

// DownloadArchive retrieve an archive hourly.
//
// Papertrail docs: https://help.papertrailapp.com/kb/how-it-works/permanent-log-archives#downloading-logs
func (c *PaperTrailClient) DownloadArchive(date string) error {
	req, err := c.NewRequest(fmt.Sprintf(URLFormat, date))
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.OpenFile(fmt.Sprintf("./%s.tsv.gz", date), os.O_RDWR|os.O_CREATE, 0444)
	if err != nil {
		return err
	}

	if _, err := io.Copy(file, resp.Body); err != nil {
		return err
	}
	return nil
}
