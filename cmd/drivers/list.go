package drivers

import (
	"os/exec"
	"strings"
)

func ListDrivers() map[string]string {
	rawOut, err := exec.Command("ubuntu-drivers", "list").Output()
	if err != nil {
		panic(err)
	}

	out := strings.Split(strings.TrimSpace(string(rawOut)), "\n")

	drivers := make(map[string]string)

	for _, v := range out {
		parts := strings.Split(v, ",")
		drivers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	return drivers
}