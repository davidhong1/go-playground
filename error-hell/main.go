package main

import (
	"bytes"
	"encoding/json"
	"error-hell/main01"
	"error-hell/main02"
	"error-hell/main03"
	"error-hell/model"
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

	p, err := main01.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", p)

	r = bytes.NewReader(bs)

	p, err = main02.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", p)

	r = bytes.NewReader(bs)

	p, err = main03.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", p)
}
