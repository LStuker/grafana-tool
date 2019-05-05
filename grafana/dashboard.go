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
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// DashboardFullJSON Dashboard export with meta data
// more info: https://grafana.com/docs/reference/dashboard/
type DashboardFullJSON struct {
	Meta struct {
		Type        string    `json:"type"`
		CanSave     bool      `json:"canSave"`
		CanEdit     bool      `json:"canEdit"`
		CanAdmin    bool      `json:"canAdmin"`
		CanStar     bool      `json:"canStar"`
		Slug        string    `json:"slug"`
		URL         string    `json:"url"`
		Expires     time.Time `json:"expires"`
		Created     time.Time `json:"created"`
		Updated     time.Time `json:"updated"`
		UpdatedBy   string    `json:"updatedBy"`
		CreatedBy   string    `json:"createdBy"`
		Version     int       `json:"version"`
		HasACL      bool      `json:"hasAcl"`
		IsFolder    bool      `json:"isFolder"`
		FolderID    int       `json:"folderId"`
		FolderTitle string    `json:"folderTitle"`
		FolderURL   string    `json:"folderUrl"`
		Provisioned bool      `json:"provisioned"`
	} `json:"meta"`
	Dashboard DashboardJSON `json:"dashboard"`
}

// DashboardJSON more info:
// https://grafana.com/docs/http_api/dashboard/#dashboard-api
type DashboardJSON struct {
	Annotations  Annotations `json:"annotations"`
	Description  string      `json:"description"`
	Editable     bool        `json:"editable"`
	GnetID       interface{} `json:"gnetId"`
	GraphTooltip int         `json:"graphTooltip"`
	ID           int         `json:"id"`
	Iteration    int64       `json:"iteration"`
	Links        []Links     `json:"links"`
	Panels       []struct {
		AliasColors struct {
		} `json:"aliasColors,omitempty"`
		Bars       bool        `json:"bars,omitempty"`
		DashLength int         `json:"dashLength,omitempty"`
		Dashes     bool        `json:"dashes,omitempty"`
		Datasource interface{} `json:"datasource,omitempty"`
		Fill       int         `json:"fill,omitempty"`
		GridPos    struct {
			H int `json:"h"`
			W int `json:"w"`
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"gridPos"`
		ID     int `json:"id"`
		Legend struct {
			Avg     bool `json:"avg"`
			Current bool `json:"current"`
			Max     bool `json:"max"`
			Min     bool `json:"min"`
			Show    bool `json:"show"`
			Total   bool `json:"total"`
			Values  bool `json:"values"`
		} `json:"legend,omitempty"`
		Lines           bool          `json:"lines,omitempty"`
		Linewidth       int           `json:"linewidth,omitempty"`
		Links           []interface{} `json:"links,omitempty"`
		NullPointMode   string        `json:"nullPointMode,omitempty"`
		Percentage      bool          `json:"percentage,omitempty"`
		Pointradius     float32       `json:"pointradius,omitempty"`
		Points          bool          `json:"points,omitempty"`
		Renderer        string        `json:"renderer,omitempty"`
		SeriesOverrides []interface{} `json:"seriesOverrides,omitempty"`
		SpaceLength     int           `json:"spaceLength,omitempty"`
		Stack           bool          `json:"stack,omitempty"`
		SteppedLine     bool          `json:"steppedLine,omitempty"`
		Targets         []struct {
			Alias   string `json:"alias"`
			GroupBy []struct {
				Params []string `json:"params"`
				Type   string   `json:"type"`
			} `json:"groupBy"`
			Measurement  string `json:"measurement"`
			OrderByTime  string `json:"orderByTime"`
			Policy       string `json:"policy"`
			Query        string `json:"query"`
			RawQuery     bool   `json:"rawQuery"`
			RefID        string `json:"refId"`
			ResultFormat string `json:"resultFormat"`
			Select       [][]struct {
				Params []string `json:"params"`
				Type   string   `json:"type"`
			} `json:"select"`
			Tags []interface{} `json:"tags"`
		} `json:"targets,omitempty"`
		Thresholds json.RawMessage `json:"thresholds,omitempty"`
		TimeFrom   interface{}     `json:"timeFrom,omitempty"`
		TimeShift  interface{}     `json:"timeShift,omitempty"`
		Title      string          `json:"title"`
		Tooltip    struct {
			Shared    bool   `json:"shared"`
			Sort      int    `json:"sort"`
			ValueType string `json:"value_type"`
		} `json:"tooltip,omitempty"`
		Type  string `json:"type"`
		Xaxis struct {
			Buckets interface{}   `json:"buckets"`
			Mode    string        `json:"mode"`
			Name    interface{}   `json:"name"`
			Show    bool          `json:"show"`
			Values  []interface{} `json:"values"`
		} `json:"xaxis,omitempty"`
		Yaxes []struct {
			Format  string      `json:"format"`
			Label   interface{} `json:"label"`
			LogBase int         `json:"logBase"`
			Max     interface{} `json:"max"`
			Min     interface{} `json:"min"`
			Show    bool        `json:"show"`
		} `json:"yaxes,omitempty"`
		Yaxis struct {
			Align      bool        `json:"align"`
			AlignLevel interface{} `json:"alignLevel"`
		} `json:"yaxis,omitempty"`
		Description string        `json:"description,omitempty"`
		Collapsed   bool          `json:"collapsed,omitempty"`
		Panels      []interface{} `json:"panels,omitempty"`
		Repeat      string        `json:"repeat,omitempty"`
		ScopedVars  struct {
			Esxi struct {
				Selected bool   `json:"selected"`
				Text     string `json:"text"`
				Value    string `json:"value"`
			} `json:"esxi"`
		} `json:"scopedVars,omitempty"`
		CacheTimeout    interface{} `json:"cacheTimeout,omitempty"`
		ColorBackground bool        `json:"colorBackground,omitempty"`
		ColorValue      bool        `json:"colorValue,omitempty"`
		Colors          []string    `json:"colors,omitempty"`
		Decimals        int         `json:"decimals,omitempty"`
		Format          string      `json:"format,omitempty"`
		Gauge           struct {
			MaxValue         int  `json:"maxValue"`
			MinValue         int  `json:"minValue"`
			Show             bool `json:"show"`
			ThresholdLabels  bool `json:"thresholdLabels"`
			ThresholdMarkers bool `json:"thresholdMarkers"`
		} `json:"gauge,omitempty"`
		Interval     interface{} `json:"interval,omitempty"`
		MappingType  int         `json:"mappingType,omitempty"`
		MappingTypes []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"mappingTypes,omitempty"`
		MaxDataPoints   int         `json:"maxDataPoints,omitempty"`
		NullText        interface{} `json:"nullText,omitempty"`
		Postfix         string      `json:"postfix,omitempty"`
		PostfixFontSize string      `json:"postfixFontSize,omitempty"`
		Prefix          string      `json:"prefix,omitempty"`
		PrefixFontSize  string      `json:"prefixFontSize,omitempty"`
		RangeMaps       []struct {
			From string `json:"from"`
			Text string `json:"text"`
			To   string `json:"to"`
		} `json:"rangeMaps,omitempty"`
		Sparkline struct {
			FillColor string `json:"fillColor"`
			Full      bool   `json:"full"`
			LineColor string `json:"lineColor"`
			Show      bool   `json:"show"`
		} `json:"sparkline,omitempty"`
		TableColumn   string `json:"tableColumn,omitempty"`
		ValueFontSize string `json:"valueFontSize,omitempty"`
		ValueMaps     []struct {
			Op    string `json:"op"`
			Text  string `json:"text"`
			Value string `json:"value"`
		} `json:"valueMaps,omitempty"`
		ValueName string `json:"valueName,omitempty"`
	} `json:"panels"`
	SchemaVersion int      `json:"schemaVersion"`
	Style         string   `json:"style"`
	Tags          []string `json:"tags"`
	Templating    struct {
		List []struct {
			Current struct {
				Text  string          `json:"text"`
				Value json.RawMessage `json:"value"`
			} `json:"current"`
			Hide           int           `json:"hide"`
			Label          string        `json:"label"`
			Name           string        `json:"name"`
			Options        []interface{} `json:"options"`
			Query          string        `json:"query"`
			Refresh        int           `json:"refresh"`
			Regex          string        `json:"regex,omitempty"`
			SkipURLSync    bool          `json:"skipUrlSync"`
			Type           string        `json:"type"`
			Auto           bool          `json:"auto,omitempty"`
			AutoCount      int           `json:"auto_count,omitempty"`
			AutoMin        string        `json:"auto_min,omitempty"`
			AllValue       interface{}   `json:"allValue,omitempty"`
			Datasource     string        `json:"datasource,omitempty"`
			IncludeAll     bool          `json:"includeAll,omitempty"`
			Multi          bool          `json:"multi,omitempty"`
			Sort           int           `json:"sort,omitempty"`
			TagValuesQuery string        `json:"tagValuesQuery,omitempty"`
			Tags           []interface{} `json:"tags,omitempty"`
			TagsQuery      string        `json:"tagsQuery,omitempty"`
			UseTags        bool          `json:"useTags,omitempty"`
		} `json:"list"`
	} `json:"templating"`
	Time       Time       `json:"time"`
	Timepicker Timepicker `json:"timepicker"`
	Timezone   string     `json:"timezone"`
	Title      string     `json:"title"`
	UID        string     `json:"uid"`
	Version    int        `json:"version"`
}

// Links is part of Grafana dashboard json
// more info: https://grafana.com/docs/reference/dashboard/
type Links struct {
	Icon        string   `json:"icon"`
	IncludeVars bool     `json:"includeVars"`
	KeepTime    bool     `json:"keepTime"`
	Tags        []string `json:"tags"`
	TargetBlank bool     `json:"targetBlank"`
	Type        string   `json:"type"`
	AsDropdown  bool     `json:"asDropdown,omitempty"`
	Title       string   `json:"title,omitempty"`
}

// Time is part of Grafana dashboard json
// more info: https://grafana.com/docs/reference/dashboard/
type Time struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// Timepicker is part of Grafana dashboard json
// more info: https://grafana.com/docs/reference/dashboard/
type Timepicker struct {
	RefreshIntervals []string `json:"refresh_intervals"`
	TimeOptions      []string `json:"time_options"`
}

// Annotations is part of Grafana dashboard json
// more info: https://grafana.com/docs/reference/dashboard/
type Annotations struct {
	List []struct {
		BuiltIn    int    `json:"builtIn"`
		Datasource string `json:"datasource"`
		Enable     bool   `json:"enable"`
		Hide       bool   `json:"hide"`
		IconColor  string `json:"iconColor"`
		Name       string `json:"name"`
		Type       string `json:"type"`
	} `json:"list"`
}

// SearchResult is part of Grafana dashboard json
// more info: https://grafana.com/docs/reference/dashboard/
type SearchResult []struct {
	ID          int           `json:"id"`
	UID         string        `json:"uid"`
	Title       string        `json:"title"`
	URI         string        `json:"uri"`
	URL         string        `json:"url"`
	Type        string        `json:"type"`
	Tags        []interface{} `json:"tags"`
	IsStarred   bool          `json:"isStarred"`
	FolderID    int           `json:"folderId"`
	FolderUID   string        `json:"folderUid"`
	FolderTitle string        `json:"folderTitle"`
	FolderURL   string        `json:"folderUrl"`
}

// GetDashboardByUID returns all folders users has permissions to view.
// It reflects GET /api/dashboards/uid/:uid API call.
// More info: http://docs.grafana.org/http_api/dashboard/
func (r *Client) GetDashboardByUID(UID string) (DashboardFullJSON, error) {
	var (
		raw  []byte
		code int
		err  error
	)

	path := fmt.Sprintf("/api/dashboards/uid/%s", UID)

	raw, code, err = r.getRequest(path, nil)
	records := DashboardFullJSON{}

	if err != nil && code != 200 {
		return records, err
	}
	if code != 200 {
		return records, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}

	err = json.Unmarshal(raw, &records)
	return records, err
}

// SearchDashboard returns all folders users has permissions to view.
// It reflects GET /api/dashboards/uid/:uid API call.
// More info: http://docs.grafana.org/http_api/dashboard/
func (r *Client) SearchDashboard(query string, folderIDs string, queryType string) (SearchResult, error) {
	var (
		raw  []byte
		code int
		err  error
	)

	u := url.URL{}
	q := u.Query()

	if query != "" {
		q.Set("query", query)
	}
	if folderIDs != "" {
		q.Set("folderIds", folderIDs)
	}
	if queryType != "" {
		q.Set("type", queryType)
	}

	path := "/api/search"

	raw, code, err = r.getRequest(path, q)

	records := SearchResult{}

	if err != nil && code != 200 {
		return records, err
	}
	if code != 200 {
		return records, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}

	err = json.Unmarshal(raw, &records)
	return records, err
}

// TitelForFile return the dashboard titel in a file friendly style
// ex: "Telegraf: Workshop System Dashboard (Windows)" will return
// telegraf_workshop_system_dashboard_windows
func (d DashboardJSON) TitelForFile() string {
	reg1, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	spaces := regexp.MustCompile(`\s+`)
	titel := reg1.ReplaceAllString(d.Title, " ")
	titel = spaces.ReplaceAllString(titel, " ")
	titel = strings.TrimSpace(titel)
	return strings.ToLower(strings.Replace(titel, " ", "_", -1))
}

// TitelFirstWord reurns from titel the first word
// ex: "Linux CPUs" will return linux
func (d DashboardJSON) TitelFirstWord() string {
	name := d.TitelForFile()
	return strings.Split(name, "_")[0]
}
