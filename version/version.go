package version

import (
	"embed"
	"fmt"
	"io/fs"
	"runtime/debug"
	"strings"
)

// https://stackoverflow.com/questions/66285635/how-do-you-use-go-1-16-embed-features-in-subfolders-packages
var VersionFS embed.FS

// String displays all the version values
func String() string {
	bi := ""
	info, ok := debug.ReadBuildInfo()
	if ok {
		bi = "\n-----------\n"
		bi = bi + fmt.Sprintf("vcs.revision (%s): %s", buildSetting(info, "vcs"), buildSetting(info, "vcs.revision"))
		if buildSetting(info, "vcs.modified") == "true" {
			bi = bi + " (dirty)"
		}
		bi = bi + " " + buildSetting(info, "vcs.time")
	}
	//spew.Dump(info)
	ver, err := fs.ReadFile(VersionFS, "version/version")
	if err != nil {
		return fmt.Sprintf("Unknown version (%+v)"+bi, err)
	}
	res := strings.TrimSpace(string(ver)) + bi
	ver, err = fs.ReadFile(VersionFS, "version/version.private")
	if err == nil {
		res = res + "\n-----------\n" + strings.TrimSpace(string(ver))
	}
	return res
}

func buildSetting(bi *debug.BuildInfo, key string) string {
	for _, bs := range bi.Settings {
		if bs.Key == key {
			return bs.Value
		}
	}
	return "<unknown>"
}
