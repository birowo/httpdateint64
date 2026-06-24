package httpdateint64

import (
	"bytes"
	"net/http"
	"testing"
	"time"
)

func TestConv(t *testing.T) {
	now := time.Now()
	httpdate := Conv(now.Unix())
	if !bytes.Equal(
		httpdate[:],
		[]byte(
			now.UTC().Format(http.TimeFormat),
		),
	) {
		t.Error("Hasil salah")
	}
}
