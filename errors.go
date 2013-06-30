package octogopher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ClientError(resp *http.Response) error {
	return httpError(resp)
}

func ServerError(resp *http.Response) error {
	return httpError(resp)
}

func httpError(resp *http.Response) error {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	e := fmt.Sprintf("%v: %v\n", resp.Status, string(body))
	return errors.New(e)
}
