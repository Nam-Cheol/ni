package profile

import (
	"fmt"
	"strings"
)

const Default = "prototype"

var valid = map[string]struct{}{
	"concept":    {},
	"prototype":  {},
	"mvp":        {},
	"beta":       {},
	"production": {},
}

func IsValid(name string) bool {
	_, ok := valid[name]
	return ok
}

func Validate(name string) error {
	if IsValid(name) {
		return nil
	}
	return fmt.Errorf("unsupported readiness profile %q", name)
}

func Names() []string {
	return []string{"concept", "prototype", "mvp", "beta", "production"}
}

func NamesText() string {
	return strings.Join(Names(), ", ")
}
