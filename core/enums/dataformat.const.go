package enums

type DataFormat string
type DeletionPolicy string

const (
	JSON DataFormat = "json"
	YAML DataFormat = "yaml"
	TOML DataFormat = "toml"
	PROP DataFormat = "properties"
	ENV  DataFormat = "env"
)

const (
	Delete   DeletionPolicy = "delete"
	Retain   DeletionPolicy = "retain"
	Snapshot DeletionPolicy = "snapshot"
)
