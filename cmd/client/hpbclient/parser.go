package hpbclient

import "regexp"

func ParseRPCVersion(version string) (string, string) {
	trimreg := regexp.MustCompile("/0")
	trimed := trimreg.ReplaceAllString(version, " ")
	reg := regexp.MustCompile("[0-9.]+")
	versions := reg.FindAllString(trimed, -1)

	if len(versions) > 1 {
		return versions[0], versions[1]
	} else if len(versions) == 1 {
		return versions[0], ""
	} else {
		return "", ""
	}
}

func ParseRPCTime(timestr string) string {
	length := len("2020-09-18 11:37:32.342")
	return timestr[:length]
}
