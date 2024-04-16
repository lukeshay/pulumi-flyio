//go:build exclude
// +build exclude

package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/iancoleman/strcase"
)

var alwaysRequire = []string{"flyId"}

var requireMap = map[string][]string{
	"App":                  {"name", "organization", "status"},
	"CreateAppRequest":     {"app_name", "org_slug"},
	"CreateMachineRequest": {"config"},
	"Organization":         {"name", "slug"},
	"Machine":              {"name", "instanceId", "config", "state"},
	"FlyMachineConfig":     {"image"},
	"FlyMachineGuest":      {"cpus", "cpu_kind", "memory_mb"},
}

func main() {
	gen, _ := os.ReadFile("../../pkg/flyio/flyio.gen.go")

	re := regexp.MustCompile("`json:\"[^,]*.*\"`")
	re2 := regexp.MustCompile("`.*:\"([^,]*)(.*)\"`")

	result := ""

	currentStruct := ""

	lines := strings.Split(string(gen), "\n")

	for idx, line := range lines {
		lineParts := strings.Split(strings.TrimSpace(line), " ")

		if len(lineParts) == 4 && lineParts[0] == "type" && lineParts[2] == "struct" {
			currentStruct = lineParts[1]
		} else {
			line = strings.ReplaceAll(line, "FlyDuration", "string")
		}

		if strings.HasPrefix(line, "}") {
			currentStruct = ""
		}

		if idx > 0 {
			partsBefore := strings.Split(strings.TrimSpace(lines[idx-1]), " ")
			if len(lineParts) > 0 && len(partsBefore) > 2 && partsBefore[1] == lineParts[0] && partsBefore[2] == "Deprecated:" {
				continue
			}
		}

		result += re.ReplaceAllStringFunc(line, func(s string) string {
			parts := re2.FindStringSubmatch(line)

			name := parts[1]

			if name == "id" {
				name = "flyId"
			}

			pulumi := strcase.ToLowerCamel(name)

			required := alwaysRequire
			if value, ok := requireMap[currentStruct]; ok {
				required = append(required, value...)
			}

			if strings.Contains(parts[2], "omitempty") && !slices.Contains(required, name) {
				pulumi += ",optional"
			}

			return fmt.Sprintf("`json:\"%s%s\" pulumi:\"%s\"`", parts[1], parts[2], pulumi)
		})

		result += "\n"
	}

	// result = strings.Replace(result, "import (", "import (\n  \"time\"", 1)

	os.WriteFile("../../pkg/flyio/flyio.gen.go", []byte(result), 0644)
}
