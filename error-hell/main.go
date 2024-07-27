package main

import (
	"bytes"
	"encoding/json"
	"error-hell/model"
	"error-hell/play01"
	"error-hell/play02"
	"error-hell/play03"
	"fmt"
)

func main() {
	point := model.Point{
		Longitude:     111,
		Latitude:      2,
		Distance:      4,
		ElevationGain: 5,
		ElevationLoss: 6,
	}
	bs, err := json.Marshal(point)
	if err != nil {
		panic(err)
	}
	fmt.Println(bs)

	r := bytes.NewReader(bs)

	p, err := play01.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", p)

	r = bytes.NewReader(bs)

	p, err = play02.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", p)

	r = bytes.NewReader(bs)

	p, err = play03.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", p)
}
