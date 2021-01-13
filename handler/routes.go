package handler

import (
	"net/http"
	"../db"
)

func SetUpRouting() *http.ServeMux  {
	todoHandler := &todoHandler{
		 &db.Sample{},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/samples" , todoHandler.GetSamples)

	return mux
}
