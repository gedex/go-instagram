package instagram

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func ExternalIP() string {

RETRY:

	resp, err := http.Get("http://myexternalip.com/raw")

	if err != nil {
		time.Sleep(time.Second * 5)
		fmt.Println("error accessing myexternalip gonna retry")
		goto RETRY
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		time.Sleep(time.Second * 5)
		fmt.Println("error accessing myexternalip gonna retry")
		goto RETRY
	}

	return strings.TrimSpace(string(b))

}
