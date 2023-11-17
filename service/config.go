package service

import (
	"strings"
)

const alreadyExists = "already exists: "

func getConfigPath(svc Service) string {
	err := svc.Install()
	if err == nil {
		return ""
	}
	esg := err.Error()
	idx := strings.Index(esg, alreadyExists)
	if idx == -1 {
		return ""
	}
	return esg[idx+len(alreadyExists):]
}
