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

func TestTitelForFile(t *testing.T) {
	tables := []struct {
		dashboard string
		expect    string
	}{
		{`{"title":"Linux Memory"}`, "linux_memory"},
		{`{"title":"Linux$ Memory"}`, "linux_memory"},
		{`{"title":"Linux/Memory"}`, "linux_memory"},
		{`{"title":"Linux (Memory) system"}`, "linux_memory_system"},
	}

	for _, table := range tables {
		var dashboard grafana.DashboardJSON
		dashboard_json := []byte(table.dashboard)
		json.Unmarshal(dashboard_json, &dashboard)
		filename := dashboard.TitelForFile()
		if filename != table.expect {
			t.Errorf("Is was  incorrect, got: %s, want: %s.", filename, table.expect)
		}
	}
}

func TestTitelFirstWord(t *testing.T) {
	tables := []struct {
		dashboard string
		expect    string
	}{
		{`{"title":"Linux Memory"}`, "linux"},
		{`{"title":"Linux$ Memory"}`, "linux"},
		{`{"title":"Linux/Memory"}`, "linux"},
		{`{"title":"Linux (Memory) system"}`, "linux"},
	}

	for _, table := range tables {
		var dashboard grafana.DashboardJSON
		dashboard_json := []byte(table.dashboard)
		json.Unmarshal(dashboard_json, &dashboard)
		filename := dashboard.TitelFirstWord()
		if filename != table.expect {
			t.Errorf("Is was  incorrect, got: %s, want: %s.", filename, table.expect)
		}
	}
}
