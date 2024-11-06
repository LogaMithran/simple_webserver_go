package people

import "net/http"

func Call(url string) *http.Response {
	if res, err := http.Get(url); err != nil {
		panic(err)
	} else {
		return res
	}

	return nil
}
