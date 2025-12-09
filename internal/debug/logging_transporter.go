package debug

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

var Enabled = false

type LoggingTransport struct {
	Next http.RoundTripper
}

func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Next == nil {
		t.Next = http.DefaultTransport
	}

	if Enabled {
		dumpReq, _ := httputil.DumpRequestOut(req, true)
		fmt.Fprintf(os.Stderr, "\n----- HTTP Request -----\n%s\n------------------------\n", string(dumpReq))
	}

	// Perform request
	resp, err := t.Next.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if Enabled {
		// Read & clone body so caller can still read it
		var buf bytes.Buffer
		if resp.Body != nil {
			_, _ = io.Copy(&buf, resp.Body)
			_ = resp.Body.Close()
			resp.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
		}

		dumpResp, _ := httputil.DumpResponse(resp, true)
		fmt.Fprintf(os.Stderr, "\n----- HTTP Response -----\n%s\n-------------------------\n", string(dumpResp))
	}

	return resp, nil
}
