package controller

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/jackc/pgx/v4"
)

type tvSpecs struct {
	ID           int64  `json:"ID"`
	Brand        string `json:"Brand"`
	Manufacturer string `json:"Manufacturer"`
	Model        string `json:"Model"`
	Year         int64  `json:"Year"`
}

type tvs struct {
	TVs []tvSpecs
}

func wrapToJson(rows pgx.Rows) ([]byte, error) {
	t := tvs{}
	var err error

	for rows.Next() {
		tv := tvSpecs{}
		if err := rows.Scan(&tv.ID, &tv.Brand, &tv.Manufacturer, &tv.Model, &tv.Year); err != nil {
			log.Fatal(err)
		}
		t.TVs = append(t.TVs, tv)
	}

	if t.TVs == nil {
		return nil, errors.New("slice is empty")
	}
	out, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return out, nil
}
