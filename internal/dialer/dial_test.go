package dialer

import "testing"

func TestValidateDestinationAddr(t *testing.T) {
	in := []struct {
		addr     string // remote server's address
		expected bool   // whether occurred error or not
	}{
		{
			addr:     "",
			expected: true,
		},
		{
			addr:     "localhost:",
			expected: true,
		},
		{
			addr:     "localhost",
			expected: true,
		},
		{
			addr:     "localhost:50051",
			expected: false,
		},
	}

	for _, c := range in {
		err := validateDestinationAddr(c.addr)
		if (err != nil) != c.expected {
			t.Errorf("Expected (err != nil) => %v, but got=%v, addr=%s", c.expected, err != nil, c.addr)
		}
	}
}
