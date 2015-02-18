package instagram

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func ExternalIP() string {
	resp, err := http.Get("http://myexternalip.com/raw")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(b))

}
