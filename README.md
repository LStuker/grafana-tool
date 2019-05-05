# Grafana Tool

A small helper for Grafana

## Usage

### Global Flags

Use API token to authenticated:
```
grafana-tool COMMAND --grafana-url http://foo.bar:3000 --api-token eyJrIjoieVBIMnIzTVl0YlFWbFlBckN== 
```

Use user and password to authenticated:
```
grafana-tool COMMAND --grafana-url http://foo.bar:3000 --user john --password mylittlesecret
```

If you don’t want to pollute your command line, or if don't want that your sensitive data are show up in the history or log, it’s a good idea to use environment variables to authenticated:
```
GRAFANA_URL="http://foo.bar:3000"
GRAFANA_API_TOKEN="eyJrIjoieVBIMnIzTVl0YlFWbFlBckN=="
grafana-tool COMMAND
```

### Export dashboards


Export all dashboards:
```
grafana-tool dashboard export --grafana-url http://foo.bar:3000 --api-token eyJrIjoieVBIMnIzTVl0YlFWbFlBckN== --path ~/backup
```


Export all dashboards of a folder:
```
grafana-tool dashboard export --grafana-url http://foo.bar:3000 --api-token eyJrIjoieVBIMnIzTVl0YlFWbFlBckN== --path ~/backup --folder devBot
```

## Installation

### From Source:

Grafana Tool requires golang version go1.11 or newer, the Makefile requires GNU make.

1. [Install Go](https://golang.org/doc/install) >=go1.11
2. [Install dep](https://golang.github.io/dep/docs/installation.html) ==v0.5.0
3. Download Grafana Tool source:
```
go get -d github.com/LStuker/grafana-tool
```
4. Run build from the source directory
```
go build main.go
```

### From Docker:

```
docker run -it lstuker/grafana-tool:0.0.1
```

or build your own:

```
git clone github.com/lstuker/grafana-tool
docker build .
```

### Changelog

View the [changelog](/CHANGELOG.md) for the latest updates and changes by
version.