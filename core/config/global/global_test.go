package global

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/codec/yaml"
	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/utils/strutil"
	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

func TestConfig(t *testing.T) {
	t.Run("json", func(t *testing.T) { testutil.Log(t, "\n"+json.MustMarshalIndentString(C())) })
	t.Run("yaml", func(t *testing.T) { testutil.Log(t, "\n"+yaml.MustMarshalString(C())) })

	t.Run("all", func(t *testing.T) { testutil.Log(t, "\n"+yaml.MustMarshalString(config.Raw())) })
}

func TestGenStruct(t *testing.T) {
	lines := []string{
		"log.console.disabled		bool",
		"log.console.format			string",
		"log.console.level			level.Level",
		"log.file.enabled			bool",
		"log.file.format			string",
		"log.file.level				level.Level",
		"log.file.filename			string",
		"log.file.max_size			int",
		"log.file.max_age			int",
		"log.file.max_backups		int",
		"log.file.daily_rotate		bool",
	}

	split := func(s string) []string {
		ss := strings.Split(s, "\t")
		ns := make([]string, 0, len(ss))
		for i := 0; i < len(ss); i++ {
			if ss[i] != "" {
				ns = append(ns, ss[i])
			}
		}
		return ns
	}

	get := func(ss []string, idx int) string {
		if len(ss) > idx {
			return ss[idx]
		}
		return ""
	}

	mas := func() int {
		i := 0
		for _, line := range lines {
			if l := len(get(split(line), 0)); l > i {
				i = l
			}
		}
		return i
	}()

	jsonTpl := `json:"%s,omitempty"`
	yamlTpl := `yaml:"%s,omitempty"`

	buf := &bytes.Buffer{}
	buf.WriteString("\ntype Config struct {\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			buf.WriteString("\n")
			continue
		}

		fs := split(line)

		rn := get(fs, 0)
		sn := strings.ReplaceAll(rn, ".", "_")
		pn := strutil.ToPascalCase(sn)
		tp := get(fs, 1)

		buf.WriteString("\t")

		// filed name
		buf.WriteString(pn)
		buf.WriteString(" ")

		// field type
		buf.WriteString(tp)
		buf.WriteString(" ")

		// tag start
		buf.WriteString("`")

		// json tag
		jsonTag := fmt.Sprintf(jsonTpl, rn)
		buf.WriteString(jsonTag)
		buf.WriteString(strings.Repeat(" ", mas+len(jsonTpl)-2-len(jsonTag)))

		// sep
		buf.WriteString(" ")

		// yaml tag
		yamlTag := fmt.Sprintf(yamlTpl, rn)
		buf.WriteString(yamlTag)
		buf.WriteString(strings.Repeat(" ", mas+len(yamlTpl)-2-len(yamlTag)))

		// tag end
		buf.WriteString("`")

		buf.WriteString("\n")
	}

	buf.WriteString("}\n")

	testutil.LogNoErr(t, os.WriteFile("global", buf.Bytes(), 0644), buf.String())
}
