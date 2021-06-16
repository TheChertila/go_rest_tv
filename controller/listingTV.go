package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"gitlab.com/TheChertila/REST_API_TV_Market/datastore"
)

func GetTVsList(key string, w http.ResponseWriter) {
	db := datastore.NewDBConnection()
	var rows pgx.Rows
	var err error
	defer db.Close(context.Background())
	query := `select t.id, b.brand_name as brand, m.manufacturer_name as manufacturer, model, "year"
	from tv t
	join manufacturers m
	on t.manufacturer_id = m.id
	join brands b
	on t.brand_id = b.id`
	if key != "any" {
		query = query + " where t.id = $1::int"
		rows, err = db.Query(context.Background(), query, key)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		rows, err = db.Query(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}
	}

	out, err := wrapToJson(rows)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	fmt.Fprint(w, string(out))
}
