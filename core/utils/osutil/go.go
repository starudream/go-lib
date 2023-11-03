package osutil

import (
	"encoding/json"
	"os/exec"
)

func GoListRoot() string {
	output, err := exec.Command("go", "list", "-json").Output()
	if err == nil {
		var res = struct {
			Root string `json:"Root"`
		}{}
		err = json.Unmarshal(output, &res)
		if err == nil {
			return res.Root
		}
	}
	return ""
}
