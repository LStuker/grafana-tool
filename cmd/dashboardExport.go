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

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/lstuker/grafana-tool/grafana"
	"github.com/spf13/cobra"
)

var path string
var folderName string

// dashboardExportCmd represents the dashboardExport command
var dashboardExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports all dashboard of a Grafana folder",
	Run: func(cmd *cobra.Command, args []string) {
		exportDashboard()
	},
}

func init() {
	dashboardCmd.AddCommand(dashboardExportCmd)
	dashboardExportCmd.Flags().StringVarP(&path, "path", "p", "", "Path to save dashboards (required)")
	dashboardExportCmd.MarkFlagRequired("path")
	dashboardExportCmd.Flags().StringVarP(&folderName, "folder", "f", "", "Grafana folder name. Dashboards of this folder will be exported")

}

func exportDashboard() {
	c := grafana.NewClient(grafanaURL, apiToken, username, password, http.DefaultClient)
	folderID := ""
	if folderName != "" {
		folders, err := c.GetFolders()
		if err != nil {
			log.Fatal(err)
		}
		folder, err := folders.FolderFindByName(folderName)
		folderID = strconv.Itoa(folder.ID)
	}

	searchResults, err := c.SearchDashboard("", folderID, "dash-db")
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range searchResults {
		dashboardFull, err := c.GetDashboardByUID(result.UID)
		if err != nil {
			log.Fatal(err)
		}
		dash, err := json.MarshalIndent(dashboardFull.Dashboard, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		filePath := fmt.Sprintf("%s/%s/%s_dashboard.json", path, dashboardFull.Dashboard.TitelFirstWord(), dashboardFull.Dashboard.TitelForFile())

		dashboardPath := strings.TrimRight(path, "/") + "/" + dashboardFull.Dashboard.TitelFirstWord()
		err = os.MkdirAll(dashboardPath, 0755)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Writing dashboard to: %s\n", filePath)
		err = ioutil.WriteFile(filePath, dash, 0644)
		if err != nil {
			log.Fatal(err)
		}

	}
}
