package cors

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("cors"))
	})
	//cors := New(mux) //默认参数为 *
	cors := New(
		mux,
		WithOrigin("*"),
		WithMethods("*"),
		WithHeaders("*"),
	)
	ts := httptest.NewServer(cors)

	//check 1
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}
	origin := resp.Header.Get("Access-Control-Allow-Origin")
	if origin != "*" {
		t.Errorf("origin: %s\n", origin)
	}
	methods := resp.Header.Get("Access-Control-Allow-Methods")
	if methods != "*" {
		t.Errorf("methods: %s\n", methods)
	}
	headers := resp.Header.Get("Access-Control-Allow-Headers")
	if headers != "*" {
		t.Errorf("headers: %s\n", headers)
	}

	//check 2
	req, err := http.NewRequest(http.MethodOptions, ts.URL, nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Access-Control-Request-Method", "OPTIONS")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("options status code %d\n", resp.StatusCode)
	}

	//
	fmt.Println("success")
}
