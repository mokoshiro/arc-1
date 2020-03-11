package driver

import (
	"fmt"
	"math"

	"github.com/mmcloughlin/geohash"
)

var accRadius = map[int]float32{ // meter
	1: 4989600.00,
	2: 1012500.00,
	3: 155925.00,
	4: 31640.62,
	5: 4872660.0,
	6: 98.877,
	7: 15.227,
}

const (
	phi0       = 36.0
	lambda     = 139 + 50/60
	lambda0Rad = 2.4405520707054045
)

func encodeGeoHash(lat float64, lng float64, chars uint) string {
	return geohash.EncodeWithPrecision(lat, lng, chars)
}

func lookupGeoHashIndexes(lat, lng float64, chars uint, radius float64) []string {
	roundIndexes := map[string]struct{}{}
	indexes := make([]string, 0, 360)
	// Step 1: 円の中心座標を算出
	centerY, centerX := encodeLatLngToMeter(lat, lng, phi0, lambda)
	// Step 2: 360度分のGeoHashを重複を除いて算出する
	for i := 1; i <= 360; i++ {
		x := radius*math.Cos(float64(i)) + centerX
		y := radius*math.Sin(float64(i)) + centerY
		decodedLat, decodedLng := decodeMeterToLatLng(x, y, lat, lng)
		index := encodeGeoHash(decodedLat, decodedLng, chars)
		if _, ok := roundIndexes[index]; !ok {
			roundIndexes[index] = struct{}{}
			indexes = append(indexes, index)
		}
	}
	return indexes[:len(roundIndexes)]
}

func decodeMeterToLatLng(x, y, phi0Deg, lambda0Deg float64) (float64, float64) {
	phi0Rad := phi0 * math.Pi / 180
	// 定数 (a, F: 世界測地系-測地基準系1980（GRS80）楕円体)
	m0 := 0.9999
	a := 6378137.0
	F := 298.257222101

	n := 1.0 / (2*F - 1)
	aarray := aArray(n)
	barray := betaArray(n)
	darray := deltaArray(n)

	A_ := ((m0 * a) / (1.0 + n)) * aarray[0]
	S_ := ((m0 * a) / (1.0 + n)) * (aarray[0]*phi0Rad + sDot(aarray, phi0Rad))

	xi := (x + S_) / A_
	eta := y / A_

	xisin := func(xi float64) []float64 {
		a := []float64{1, 2, 3, 4, 5}
		for i := range a {
			a[i] = math.Sin(2 * xi * a[i])
		}
		return a
	}(xi)
	xicosh := func(eta float64) []float64 {
		a := []float64{1, 2, 3, 4, 5}
		for i := range a {
			a[i] = math.Cosh(2 * eta * a[i])
		}
		return a
	}(eta)
	mul := multiply(xisin, xicosh)
	mul = multiply(barray[1:], mul)
	xi2 := xi - sum(mul)

	xicos := func(xi float64) []float64 {
		a := []float64{1, 2, 3, 4, 5}
		for i := range a {
			a[i] = math.Cos(2 * xi * a[i])
		}
		return a
	}(xi)
	etasinh := func(eta float64) []float64 {
		a := []float64{1, 2, 3, 4, 5}
		for i := range a {
			a[i] = math.Sinh(2 * eta * a[i])
		}
		return a
	}(eta)

	mul = multiply(xicos, etasinh)
	mul = multiply(barray[1:], mul)
	eta2 := eta - sum(mul)
	chi := math.Asin(math.Sin(xi2) / math.Cosh(eta2))
	latitude := chi + deltaDot(darray, chi)
	longitude := lambda0Rad + math.Atan(math.Sinh(eta2)/math.Cos(xi2))
	return latitude * (180 / math.Pi), longitude * (180 / math.Pi)
}

func deltaDot(darray []float64, chi float64) float64 {
	a := []float64{1, 2, 3, 4, 5, 6}
	res := 0.0
	for i := range a {
		res += darray[i] * math.Sin((2 * chi * a[i]))
	}
	return res
}

func betaArray(n float64) []float64 {
	b0 := 0.0
	b1 := (1./2.0)*n - (2./3)*(math.Pow(n, 2)) + (37./96)*(math.Pow(n, 3)) - (1./360)*(math.Pow(n, 4)) - (81./512)*(math.Pow(n, 5))
	b2 := (1./48.)*(math.Pow(n, 2)) + (1./15.)*(math.Pow(n, 3)) - (437./1440)*(math.Pow(n, 4)) + (46./105)*(math.Pow(n, 5))
	b3 := (17./480.)*(math.Pow(n, 3)) - (37./840.)*(math.Pow(n, 4)) - (209./4480)*(math.Pow(n, 5))
	b4 := (4397./161280.)*(math.Pow(n, 4)) - (11./504)*(math.Pow(n, 5))
	b5 := (4583. / 161280.) * (math.Pow(n, 5))
	return []float64{b0, b1, b2, b3, b4, b5}
}

func deltaArray(n float64) []float64 {
	d0 := 0.0
	d1 := 2.*n - (2./3)*(math.Pow(n, 2)) - 2.*(math.Pow(n, 3)) + (116./45)*(math.Pow(n, 4)) + (26./45)*(math.Pow(n, 5)) - (2854./675)*(math.Pow(n, 6))
	d2 := (7./3)*(math.Pow(n, 2)) - (8./5)*(math.Pow(n, 3)) - (227./45)*(math.Pow(n, 4)) + (2704./315)*(math.Pow(n, 5)) + (2323./945)*(math.Pow(n, 6))
	d3 := (56./15)*(math.Pow(n, 3)) - (136./35)*(math.Pow(n, 4)) - (1262./105)*(math.Pow(n, 5)) + (73814./2835)*(math.Pow(n, 6))
	d4 := (4279./630)*(math.Pow(n, 4)) - (332./35)*(math.Pow(n, 5)) - (399572./14175)*(math.Pow(n, 6))
	d5 := (4174./315)*(math.Pow(n, 5)) - (144838./6237)*(math.Pow(n, 6))
	d6 := (601676. / 22275) * (math.Pow(n, 6))
	return []float64{d0, d1, d2, d3, d4, d5, d6}
}

func encodeLatLngToMeter(lat, lng float64, phi0, lambda0 float64) (float64, float64) {
	phiRad := lat * math.Pi / 180
	lambdaRad := lng * math.Pi / 180
	phi0Rad := phi0 * math.Pi / 180

	// 定数 (a, F: 世界測地系-測地基準系1980（GRS80）楕円体)
	m0 := 0.9999
	a := 6378137.0
	F := 298.257222101

	// Step 1: n, A_i, alpha_iの算出
	n := 1.0 / (2*F - 1)
	aarray := aArray(n)
	alArray := alphaArray(n)

	// Step 2: S, Aの計算
	A_ := ((m0 * a) / (1.0 + n)) * aarray[0]
	S_ := ((m0 * a) / (1.0 + n)) * (aarray[0]*phi0Rad + sDot(aarray, phi0Rad))

	// Step 3: lambda-C, lamda-S
	lambdaC := math.Cos(lambdaRad - lambda0Rad)
	lambdaS := math.Sin(lambdaRad - lambda0Rad)

	fmt.Println(lambdaC, lambdaS)

	// Step 4: t, t_
	t := math.Sinh(math.Atanh(math.Sin(phiRad)) - ((2*math.Sqrt(n))/(1+n))*math.Atanh(((2*math.Sqrt(n))/(1+n))*math.Sin(phiRad)))
	t_ := math.Sqrt(1 + t*t)

	// Step 5: xi, eta
	xi2 := math.Atan(t / lambdaC)
	eta2 := math.Atanh(lambdaS / t_)
	xsin := func(xi2 float64) []float64 {
		a := []float64{1, 2, 3, 4, 5}
		for i := range a {
			a[i] = math.Sin(2 * xi2 * a[i])
		}
		return a
	}(xi2)
	xcosh := func(eta2 float64) []float64 {
		a := []float64{1, 2, 3, 4, 5}
		for i := range a {
			a[i] = math.Cosh(2 * eta2 * a[i])
		}
		return a
	}(eta2)
	xsincos := multiply(xsin, xcosh)
	alphaMul := multiply(alArray[1:], xsincos)
	x := A_*(xi2+sum(alphaMul)) - S_

	ycos := func(xi2 float64) []float64 {
		a := []float64{1, 2, 3, 4, 5}
		for i := range a {
			a[i] = math.Cos(2 * xi2 * a[i])
		}
		return a
	}(xi2)
	ysinh := func(eta2 float64) []float64 {
		a := []float64{1, 2, 3, 4, 5}
		for i := range a {
			a[i] = math.Sinh(2 * eta2 * a[i])
		}
		return a
	}(eta2)
	ysincos := multiply(ycos, ysinh)
	alphaMul = multiply(alArray[1:], ysincos)
	y := A_ * (eta2 + sum(alphaMul))

	return x, y
}

func multiply(a, b []float64) []float64 {
	res := make([]float64, len(a))
	for i := range res {
		res[i] = a[i] * b[i]
	}
	return res
}

func sum(a []float64) float64 {
	s := 0.0
	for i := range a {
		s += a[i]
	}
	return s
}

func sDot(aarray []float64, phi0Rad float64) float64 {
	arange := []float64{1, 2, 3, 4, 5}
	dot := 0.0
	for i := range arange {
		s := math.Sin(2 * phi0Rad * arange[i])
		dot += aarray[i+1] * s
	}
	return dot
}

func aArray(n float64) []float64 {
	A0 := 1 + (math.Pow(n, 2) / 4.0) + math.Pow(n, 4)/64.0
	A1 := -(3.0 / 2.0) * (n - (math.Pow(n, 3)/8.0 - math.Pow(n, 5)/64.0))
	A2 := 15.0 / 16.0 * (math.Pow(n, 2) - math.Pow(n, 4)/4)
	A3 := -(35.0 / 48.0) * (math.Pow(n, 3) - (5.0/16.0)*(math.Pow(n, 5)))
	A4 := (315.0 / 512.0) * math.Pow(n, 4)
	A5 := -(693.0 / 1280.0) * math.Pow(n, 5)
	return []float64{A0, A1, A2, A3, A4, A5}
}

func alphaArray(n float64) []float64 {
	a0 := 0.0
	a1 := (1./2)*n - (2./3)*(math.Pow(n, 2)) + (5.0/16.0)*(math.Pow(n, 3)) + (41./180)*(math.Pow(n, 4)) - (127.0/288.0)*(math.Pow(n, 5))
	a2 := (13.0/48.0)*(math.Pow(n, 2)) - (3.0/5.0)*math.Pow(n, 3) + (557./1440)*(math.Pow(n, 4)) + (281.0/630.0)*(math.Pow(n, 5))
	a3 := (61.0/240.0)*math.Pow(n, 3) - (103.0/140.0)*math.Pow(n, 4) + (15061.0/26880.0)*math.Pow(n, 5)
	a4 := (49561.0/161280.0)*math.Pow(n, 4) - (179./168)*math.Pow(n, 5)
	a5 := (34729.0 / 80640.0) * math.Pow(n, 5)
	return []float64{a0, a1, a2, a3, a4, a5}
}

func decodeKmToLatLng(x, y, lat, lng float64) (float64, float64) {
	perLatKm := 91.142519
	r := 6378.137
	cf := math.Cos(lat/180*math.Pi) * 2 * math.Pi * r
	perLngKm := cf / 360
	return y / perLatKm, x / perLngKm
}
