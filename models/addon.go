package models

type Addon struct {
	Name string
  Slug string
  Advanced bool
  Description string
  Repository string
  Version NullString
  LatestVersion string
  Update bool
  Installed string
  Available bool
  Icon bool
  Logo bool
  State string
}
