package tools

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func CheckHttpErr(r *http.Response, err error) error {
	/*
		Check for HTTP error and error codes
	*/

	// Host Blocked or Connect errors
	if err != nil {
		if strings.Contains(err.Error(), "no Host in request URL") {
			return errors.New("no domain was provied. please init client with domain.")
		} else if strings.Contains(err.Error(), "connection was forcibly closed by the remote host") {
			return errors.New(fmt.Sprintf("error: %s.\nsidenote: are you sure the domain isn't blocked by your isp?", err.Error()))
		}
		return err
	}

	// Response error
	if r.StatusCode != 200 {
		return errors.New(fmt.Sprintf("http status code not 200. %s", r.StatusCode))
	}

	return nil
}
