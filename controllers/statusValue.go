package controllers

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"path"
)

func Value(w http.ResponseWriter, r *http.Request) {
	water := rand.Intn(100)
	wind := rand.Intn(100)

	var (
		file          = path.Join("view", "index.html")
		template, err = template.ParseFiles(file)
		waterStatus   string
		windStatus    string
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if water < 5 {
		waterStatus = "aman"
	} else if water > 5 && water <= 15 {
		waterStatus = "siaga"
	} else {
		waterStatus = "bahaya"
	}

	if wind < 6 {
		windStatus = "aman"
	} else if wind > 6 && wind <= 15 {
		windStatus = "siaga"
	} else {
		windStatus = "bahaya"
	}

	msg1 := fmt.Sprintf("Water %d\n", water)
	msg2 := fmt.Sprintf("Wind %d\n", wind)

	data := map[string]interface{}{
		"water":        msg1,
		"wind":         msg2,
		"status_water": waterStatus,
		"status_wind":  windStatus,
	}

	err = template.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
