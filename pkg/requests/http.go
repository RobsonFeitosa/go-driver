package requests

import (
	"io"
	"io/ioutil"
)

func Post(path string, body io.Reader) ([]byte, error) {
	resp, err := doRequest("POST", path, body, nil, true)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func AuthenticatedPostWithHeaders(path string, body io.Reader, headers map[string]string) ([]byte, error) {
	resp, err := doRequest("POST", path, body, nil, true)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func AuthenticatedPost(path string, body io.Reader) ([]byte, error) {
	resp, err := doRequest("POST", path, body, nil, true)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func AuthenticatedPut(path string, body io.Reader) ([]byte, error) {
	resp, err := doRequest("PUT", path, body, nil, true)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func AuthenticatedGet(path string) ([]byte, error) {
	resp, err := doRequest("GET", path, nil, nil, true)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func AuthenticatedDelete(path string) error {
	_, err := doRequest("GET", path, nil, nil, true)

	return err
}
