package httpdateint64

import (
	"bytes"
	"net/http"
	"testing"
	"time"
)

func TestConv(t *testing.T) {
	now := time.Now()
	httpdate, err := Conv(now.Unix())
	if err != nil || !bytes.Equal(
		httpdate[:],
		[]byte(
			now.UTC().Format(http.TimeFormat),
		),
	) {
		t.Errorf("e r r o r: %s != %s\n",
			string(httpdate[:]),
			now.UTC().Format(http.TimeFormat),
		)
	}
}
