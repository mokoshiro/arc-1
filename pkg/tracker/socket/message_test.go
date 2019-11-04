package socket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	b := []byte{0x00, 0x01, 0x00, 0x00}

	body := []byte(`{
        "peer_id": "aaaa",
        "mode": 1,
        "addr": "127.0.0.1:8080",
        "longitude": 127.000001,
        "latitude": 34.0001010
    }`)
	b = append(b, body...)
	r, err := handleMessage(b)
	if err != nil {
		t.Fatal(err)
	}
	gr := r.(*greetResponse)
	assert.Equal(t, 1, gr.Status)
}
