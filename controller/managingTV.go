package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gitlab.com/TheChertila/REST_API_TV_Market/datastore"
)

func RemoveTV(key string, w http.ResponseWriter) {
	db := datastore.NewDBConnection()
	defer db.Close(context.Background())

	query := "DELETE FROM tv WHERE id = $1::int"

	res, err := db.Exec(context.Background(), query, key)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	count := res.RowsAffected()
	if count == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 OK"))

}

func AddTV(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		tv  tvSpecs
	)
	db := datastore.NewDBConnection()
	defer db.Close(context.Background())
	t := tvs{}

	w.Header().Set("Content-Type", "application/json")

	query := `insert into tv (brand_id, manufacturer_id, model, "year")
	select b.id as b_id, m.id as m_id, o.model, o."year"
	from (select id from brands where brand_name = $1::varchar) b
		,(select id from manufacturers where manufacturer_name = $2::varchar ) m
		,(select $3::varchar as model, $4::smallint as "year") o;`

	fmt.Fprintf(w, "Post from website! r.PostForm = %v\n", r.PostForm)
	tv.Brand = r.FormValue("brand")
	tv.Manufacturer = r.FormValue("manufacturer")
	tv.Model = r.FormValue("model")
	tv.Year, err = strconv.ParseInt(r.FormValue("year"), 10, 16)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query(context.Background(), query, tv.Brand, tv.Manufacturer, tv.Model, tv.Year)
	if err != nil {
		log.Fatal((err))
	}

	t.TVs = append(t.TVs, tv)

	json.NewEncoder(w).Encode(tv)
}
