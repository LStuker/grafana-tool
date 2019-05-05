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
	"encoding/json"
	"errors"
	"fmt"
)

// FolderListJSON is a list of folders from the Gragana API
// More info: https://grafana.com/docs/http_api/folder/
type FolderListJSON []FolderJSON

// FolderJSON is a folders from the Gragana API
// More info: https://grafana.com/docs/http_api/folder/
type FolderJSON struct {
	ID    int    `json:"id"`
	UID   string `json:"uid"`
	Title string `json:"title"`
	// URL       string    `json:"url"`
	// HasAcl    bool      `json:"hasAcl"`
	// CanSave   bool      `json:"canSave"`
	// CanEdit   bool      `json:"canEdit"`
	// CanAdmin  bool      `json:"canAdmin"`
	// CreatedBy string    `json:"createdBy"`
	// Created   time.Time `json:"created"`
	// UpdatedBy string    `json:"updatedBy"`
	// Updated   time.Time `json:"updated"`
	// Version   int       `json:"version"`
}

// GetFolders returns all folders users has permissions to view.
// It reflects GET /api/folders API call.
// More info: http://docs.grafana.org/http_api/folder/
func (r *Client) GetFolders() (FolderListJSON, error) {
	var (
		records FolderListJSON
		raw     []byte
		code    int
		err     error
	)

	raw, code, err = r.getRequest("/api/folders", nil)

	if err != nil && code != 200 {
		return records, err
	}
	if code != 200 {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}

	err = json.Unmarshal(raw, &records)
	return records, err
}

// FolderFindByName search in a FolderListJSON the folder by Titel name and
// returns the FolderJSON object
func (f FolderListJSON) FolderFindByName(title string) (FolderJSON, error) {
	var empty FolderJSON

	for _, folder := range f {
		if folder.Title == title {
			return folder, nil
		}
	}
	return empty, errors.New("Folder not found")
}
