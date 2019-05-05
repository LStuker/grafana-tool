// Copyright © 2019 Lucien Stuker <lucien.stuker@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grafana

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

// Client uses Grafana REST API for interacting with Grafana server.
type Client struct {
	baseURL    string
	basicAuth  bool
	apiToken   string
	username   string
	password   string
	httpClient *http.Client
}

// NewClient initializes client for interacting with Grafana API;
func NewClient(apiURL, apiToken string, username string, password string, client *http.Client) *Client {
	baseURL, _ := url.Parse(apiURL)
	basicAuth := username != ""
	return &Client{baseURL: baseURL.String(),
		basicAuth:  basicAuth,
		apiToken:   apiToken,
		username:   username,
		password:   password,
		httpClient: client}
}

func (r *Client) getRequest(query string, params url.Values) ([]byte, int, error) {
	return r.request("GET", query, params, nil)
}

func (r *Client) postRequest(query string, params url.Values, body []byte) ([]byte, int, error) {
	return r.request("POST", query, params, bytes.NewBuffer(body))
}

func (r *Client) deleteRequest(query string) ([]byte, int, error) {
	return r.request("DELETE", query, nil, nil)
}

func (r *Client) request(method, query string, params url.Values, buf io.Reader) ([]byte, int, error) {
	u, _ := url.Parse(r.baseURL)
	u.Path = path.Join(u.Path, query)
	if params != nil {
		u.RawQuery = params.Encode()
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		// dial tcp: lookup failed with no such host
		return nil, 504, err
	}

	if r.basicAuth {
		req.SetBasicAuth(r.username, r.password)
	} else {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.apiToken))
	}

	log.Printf("Request: %s\n", u.String())

	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "autograf")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := r.httpClient.Do(req)
	if err != nil {
		// dial tcp: lookup failed with no such host
		if resp == nil {
			return nil, 503, err
		}
		return nil, resp.StatusCode, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return data, resp.StatusCode, err
}
