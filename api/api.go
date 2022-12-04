package api

import (
	"errors"
	"io"
	"net/http"
	"os"
)

var baseUrl = "https://papermc.io/api/v2/projects/"
var project = "paper"

func Get(buildVersion string, buildNumber string, buildName string) (err error) {
	if buildVersion == "" || buildNumber == "" || buildName == "" {
		return errors.New("Empty fields")
	}

	out, err := os.Create("paper.jar")
	defer out.Close()
	if err != nil {
		return err
	}

	res, err := http.Get(baseUrl + project + "/versions/" + buildVersion + "/builds/" + buildNumber + "/downloads/" + buildName)
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}
