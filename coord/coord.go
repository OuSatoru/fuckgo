package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

const (
	pi  = math.Pi
	xPi = pi * 3000.0 / 180.0
	a   = 6378245.0
	ee  = 0.00669342162296594323
)

func main() {
	var typ string
	flag.StringVar(&typ, "type", "bd09", "type is bd09, wgs84, gcj02")
	flag.Parse()
	coords := flag.Args()
	if typ != "bd09" && typ != "wgs84" && typ != "gcj02" {
		flag.Usage()
		return
	}
	if len(coords) != 2 {
		flag.Usage()
		panic("请输入两个坐标，lng, lat")
	}
	lng, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		panic(err)
	}
	lat, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		panic(err)
	}
	// fmt.Println("trans: ", transLng(lng, lat), transLat(lng, lat))
	switch typ {
	case "bd09":
		fmt.Print("火星坐标 ")
		fmt.Println(bdToGc(lng, lat))
		fmt.Print("GPS坐标 ")
		fmt.Println(bdToWg(lng, lat))
	case "wgs84":
		fmt.Print("火星坐标 ")
		fmt.Println(wgToGc(lng, lat))
		fmt.Print("百度坐标 ")
		fmt.Println(wgToBd(lng, lat))
	case "gcj02":
		fmt.Print("百度坐标 ")
		fmt.Println(gcToBd(lng, lat))
		fmt.Print("GPS坐标 ")
		fmt.Println(gcToWg(lng, lat))
	}

}

func bdToGc(lng, lat float64) (float64, float64) {
	x := lng - 0.0065
	y := lat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*xPi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*xPi)
	gLng := z * math.Cos(theta)
	gLat := z * math.Sin(theta)
	return gLng, gLat
}

func gcToBd(lng, lat float64) (float64, float64) {
	z := math.Sqrt(lng*lng+lat*lat) + 0.00002*math.Sin(lat*xPi)
	theta := math.Atan2(lat, lng) + 0.000003*math.Cos(lng*xPi)
	bLng := z*math.Cos(theta) + 0.0065
	bLat := z*math.Sin(theta) + 0.006
	return bLng, bLat
}

func wgToGc(lng, lat float64) (float64, float64) {
	if outOfChina(lng, lat) {
		return lng, lat
	}
	dlng := transLng(lng-105.0, lat-35.0)
	dlat := transLat(lng-105.0, lat-35.0)
	radlat := lat / 180.0 * pi
	magic := math.Sin(radlat)
	magic = 1 - ee*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dlng = (dlng * 180.0) / (a / sqrtMagic * math.Cos(radlat) * pi)
	dlat = (dlat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * pi)
	return lng + dlng, lat + dlat
}

func gcToWg(lng, lat float64) (float64, float64) {
	if outOfChina(lng, lat) {
		return lng, lat
	}
	dlng := transLng(lng-105.0, lat-35.0)
	dlat := transLat(lng-105.0, lat-35.0)
	radlat := lat / 180.0 * pi
	magic := math.Sin(radlat)
	magic = 1 - ee*magic*magic
	sqrtMagic := math.Sqrt(magic)
	dlng = (dlng * 180.0) / (a / sqrtMagic * math.Cos(radlat) * pi)
	dlat = (dlat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * pi)
	mglng := lng + dlng
	mglat := lat + dlat
	return lng*2 - mglng, lat*2 - mglat
}

func bdToWg(lng, lat float64) (float64, float64) {
	ln, la := bdToGc(lng, lat)
	return gcToWg(ln, la)
}

func wgToBd(lng, lat float64) (float64, float64) {
	ln, la := wgToGc(lng, lat)
	return gcToBd(ln, la)
}

func outOfChina(lng, lat float64) bool {
	return !(lng > 73.66 && lng < 135.05 && lat > 3.86 && lat < 53.55)
}

func transLng(lng, lat float64) float64 {
	ret := 300.0 + lng + 2.0*lat + 0.1*lng*lng + 0.1*lng*lat + 0.1*math.Sqrt(math.Abs(lng))
	ret += (20.0*math.Sin(6.0*lng*pi) + 20.0*math.Sin(2.0*lng*pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lng*pi) + 40.0*math.Sin(lng/3.0*pi)) * 2.0 / 3.0
	ret += (150.0*math.Sin(lng/12.0*pi) + 300.0*math.Sin(lng/30.0*pi)) * 2.0 / 3.0
	return ret
}

func transLat(lng, lat float64) float64 {
	ret := -100.0 + 2.0*lng + 3.0*lat + 0.2*lat*lat + 0.1*lng*lat + 0.2*math.Sqrt(math.Abs(lng))
	ret += (20.0*math.Sin(6.0*lng*pi) + 20.0*math.Sin(2.0*lng*pi)) * 2.0 / 3.0
	ret += (20.0*math.Sin(lat*pi) + 40.0*math.Sin(lat/3.0*pi)) * 2.0 / 3.0
	ret += (160.0*math.Sin(lat/12.0*pi) + 320.0*math.Sin(lat*pi/30.0)) * 2.0 / 3.0
	return ret
}
