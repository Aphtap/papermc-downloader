package paper

import (
	"strconv"
)

type PaperVersions struct {
	Project_id     string
	Project_name   string
	Version_groups []string
	Versions       []string
}

type PaperBuilds struct {
	Project_id   string
	Project_name string
	Version      string
	Builds       []uint16
}

type PaperDownloads struct {
	Project_id   string
	Project_name string
	Version      string
	Build        uint16
	Time         string
	Channel      string
	Promoted     bool
	Downloads    struct {
		Application struct {
			Name string
		}
	}
}

func GetLatestVersion() string {
	version := PaperVersions{}
	getJson(baseUrl+project, &version)
	return version.Version_groups[len(version.Version_groups)-1]
}

func GetLatestBuild(version string) uint16 {
	build := PaperBuilds{}
	getJson(baseUrl+project+"/versions/"+version, &build)
	return build.Builds[len(build.Builds)-1]
}

func GetLatestDownload(version string, build uint16) string {
	dl := PaperDownloads{}
	getJson(baseUrl+project+"/versions/"+version+"/builds/"+strconv.Itoa(int(build)), &dl)
	return dl.Downloads.Application.Name
}
