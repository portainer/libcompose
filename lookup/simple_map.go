package lookup

import (
	"fmt"

	"github.com/portainer/libcompose/config"
)

// MapLookup is a structure that implements the project.EnvironmentLookup interface.
// It holds a map of variables where to lookup environment values.
type MapLookup struct {
	Vars map[string]string
}

// Lookup creates a string slice of string containing a "docker-friendly" environment string
// in the form of 'key=value'. It gets environment values from the vars map.
func (l *MapLookup) Lookup(key string, config *config.ServiceConfig) []string {
	ret := l.Vars[key]
	if ret == "" {
		return []string{}
	}
	return []string{fmt.Sprintf("%s=%s", key, ret)}
}
