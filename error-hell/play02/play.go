package play02

import (
	"encoding/binary"
	"error-hell/model"
	"io"
)

func Parse(r io.Reader) (*model.Point, error) {
	var p model.Point
	var err error
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return nil, err
	}

	return &p, nil
}
