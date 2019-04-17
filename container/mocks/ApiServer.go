package mocks

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
)

func NewMockApiServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logrus.Debug("Mock server has received a HTTP call on ", r.URL)
			var response = ""

			if  isRequestFor("containers/json?limit=0", r) {
				response = getMockJsonFromDisk("./mocks/data/containers.json")
			} else if isRequestFor("ae8964ba86c7cd7522cf84e09781343d88e0e3543281c747d88b27e246578b65", r) {
				response = getMockJsonFromDisk("./mocks/data/container_stopped.json")
			} else if isRequestFor("b978af0b858aa8855cce46b628817d4ed58e58f2c4f66c9b9c5449134ed4c008", r) {
				response = getMockJsonFromDisk("./mocks/data/container_running.json")
			} else if isRequestFor("sha256:19d07168491a3f9e2798a9bed96544e34d57ddc4757a4ac5bb199dea896c87fd", r) {
				response = getMockJsonFromDisk("./mocks/data/image01.json")
			} else if isRequestFor("sha256:4dbc5f9c07028a985e14d1393e849ea07f68804c4293050d5a641b138db72daa", r) {
				response = getMockJsonFromDisk("./mocks/data/image02.json")
			}
			fmt.Fprintln(w, response)
		},
	))
}

func isRequestFor(urlPart string, r *http.Request) bool {
	return strings.Contains(r.URL.String(), urlPart)
}

func getMockJsonFromDisk(relPath string) string {
	absPath, _ := filepath.Abs(relPath)
	logrus.Error(absPath)
	buf, err := ioutil.ReadFile(absPath)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	return string(buf)
}
