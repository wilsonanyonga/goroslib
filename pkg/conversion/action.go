package conversion

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var tplAction = template.Must(template.New("").Parse(
	`//autogenerated:yes
//nolint:revive,lll
package {{ .GoPkgName }}

import (
{{- range $k, $v := .Imports }}
    "{{ $k }}"
{{- end }}
)
{{ .Goal }}
{{ .Result }}
{{ .Feedback }}
type {{ .Name }}Action struct {
{{- if .RosPkgName }}
    msg.Package ` + "`" + `ros:"{{ .RosPkgName }}"` + "`" + `
{{- end }}
    {{ .Name }}ActionGoal
    {{ .Name }}ActionResult
    {{ .Name }}ActionFeedback
}
`))

// ImportAction generates Go file from an .action file and writes to the io.Writer.
func ImportAction(path string, goPkgName string, rosPkgName string, w io.Writer) error {
	name := strings.TrimSuffix(filepath.Base(path), ".action")

	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(buf)

	parts := strings.Split(content, "---")
	if len(parts) != 3 {
		return fmt.Errorf("definition must contain a goal a result and a feedback")
	}

	goalDef, err := parseMessageDefinition(goPkgName, name+"ActionGoal", parts[0])
	if err != nil {
		return err
	}

	resultDef, err := parseMessageDefinition(goPkgName, name+"ActionResult", parts[1])
	if err != nil {
		return err
	}

	feedbackDef, err := parseMessageDefinition(goPkgName, name+"ActionFeedback", parts[2])
	if err != nil {
		return err
	}

	imports := make(map[string]struct{})
	for i := range goalDef.Imports {
		imports[i] = struct{}{}
	}
	for i := range resultDef.Imports {
		imports[i] = struct{}{}
	}
	for i := range feedbackDef.Imports {
		imports[i] = struct{}{}
	}

	goal, err := goalDef.write()
	if err != nil {
		return err
	}

	result, err := resultDef.write()
	if err != nil {
		return err
	}

	feedback, err := feedbackDef.write()
	if err != nil {
		return err
	}

	return tplAction.Execute(w, map[string]interface{}{
		"GoPkgName":  goPkgName,
		"RosPkgName": rosPkgName,
		"Imports":    imports,
		"Goal":       goal,
		"Result":     result,
		"Feedback":   feedback,
		"Name":       name,
	})
}