// api project api.go
package api

import (
	"net/http"
)

type ApiClient struct {
	Source string
	hc     http.Client
}
