package driver

import "testing"

func TestGeoHashEncode(t *testing.T) {
	lat := 36.103774791666666
	lng := 140.08785504166664
	y, x := encodeLatLngToMeter(lat, lng, phi0, lambda)
	t.Log(y, x)

	latitude, longitude := decodeMeterToLatLng(x, y, phi0, lambda)
	t.Log(latitude, longitude)
}
