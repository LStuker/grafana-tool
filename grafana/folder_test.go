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
package grafana_test

import (
	"encoding/json"
	"testing"

	"github.com/lstuker/grafana-tool/grafana"
)

func TestFindFolderByName(t *testing.T) {
	var folders grafana.FolderListJSON
	folders_str := []byte(`[{"id":98,"uid":"HeZIp-Qmk","title":"bank-now"},{"id":83,"uid":"5_oQ-G8mk","title":"Fun"},{"id":22,"uid":"p0S7PPezk","title":"newbit"},{"id":50,"uid":"DOqaFO2mk","title":"sql-dev-dashboard"},{"id":76,"uid":"1CaRYRymz","title":"vmware"},{"id":37,"uid":"J09uHF4iz","title":"Workshop"}]`)
	json.Unmarshal(folders_str, &folders)
	folder, _ := folders.FolderFindByName("Fun")
	if folder.ID != 83 {
		t.Errorf("Is was  incorrect, got: %d, want: %d.", folder.ID, 83)
	}
}
