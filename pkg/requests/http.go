package requests

import (
	"io"
	"io/ioutil"
)

func AuthenticatedPost(path string, body io.Reader) ([]byte, error) {
	resp, err := doRequest("POST", path, body, true)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func AuthenticatedGet(path string) ([]byte, error) {
	resp, err := doRequest("GET", path, nil, true)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}
