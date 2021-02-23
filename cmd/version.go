/**------------------------------------------------------------**
 * @filename cmd/version.go
 * @author   jiny - caojingyin@baimaohui.net
 * @version  1.0.0
 * @date     2017/07/02 12:05
 * @desc     cmd-version - summary
 **------------------------------------------------------------**/

package cmd

import (
	"bytes"
	"fmt"
	"runtime"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

const (
	Version     = "1.0.0"
	versionTemp = `
  Jiny version: {{.JinyVersion}}
  ├── Go version:     {{.GoVersion}}
  ├── Jinygo version: {{.MVersion}}
  ├── OS/Arch:        {{.Os}}/{{.Arch}}
  └── Date:           {{.Date}}
`
)

// VersionOptions include version
type VersionOptions struct {
	JinyVersion string
	MVersion    string
	GoVersion   string
	Os          string
	Arch        string
	Date        string
}

// addVersion augments our CLI surface with version.
func addVersion(cmd *cobra.Command) {
	cmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: `Print jiny version.`,
		Run: func(cmd *cobra.Command, args []string) {
			version()
		},
	})
}

func version() {
	var doc bytes.Buffer
	today := time.Now()
	vo := VersionOptions{
		JinyVersion: Version,
		MVersion:    "1.0.0",
		GoVersion:   runtime.Version(),
		Os:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		Date:        today.Format("2006-01-02"),
	}
	tmpl, _ := template.New("version").Parse(versionTemp)
	_ = tmpl.Execute(&doc, vo)
	fmt.Println(doc.String())
}
