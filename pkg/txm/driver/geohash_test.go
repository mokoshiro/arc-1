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

func TestLookupGeoHashes(t *testing.T) {
	lat := 36.103774791666666
	lng := 140.08785504166664
	radius := 100.0 // meter
	indexes := lookupGeoHashIndexes(lat, lng, 7, radius)
	if len(indexes) != 4 {
		t.Errorf("expected lookup index numbers=%d, but got=%d", 4, len(indexes))
	}
}
