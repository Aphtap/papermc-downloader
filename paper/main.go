package paper

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

var baseUrl = "https://papermc.io/api/v2/projects/"
var project = "paper"

func getJson(url string, target interface{}) error {
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func Get(buildVersion string, buildNumber uint16, buildName string) (err error) {
	if buildVersion == "" || buildNumber != 0 || buildName == "" {
		return errors.New("Empty fields")
	}

	out, err := os.Create("paper.jar")
	defer out.Close()
	if err != nil {
		return err
	}

	res, err := httpClient.Get(baseUrl + project + "/versions/" + buildVersion + "/builds/" + strconv.Itoa(int(buildNumber)) + "/downloads/" + buildName)
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}
