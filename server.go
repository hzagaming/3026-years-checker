// 3026 Years Checker — 静态站 + JSON 接口
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const target = 3026

type result struct {
	Year  int    `json:"year"`
	Delta int    `json:"delta"`
	Text  string `json:"text"`
	Error string `json:"error,omitempty"`
}

func describe(year int) result {
	delta := year - target
	text := ""
	switch {
	case delta == 0:
		text = "就是 3026 年"
	case delta > 0:
		text = "在 3026 年之后 " + strconv.Itoa(delta) + " 年"
	default:
		text = "在 3026 年之前 " + strconv.Itoa(-delta) + " 年"
	}
	return result{Year: year, Delta: delta, Text: strconv.Itoa(year) + " " + text}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web")))
	mux.HandleFunc("/api/check", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		year, err := strconv.Atoi(r.URL.Query().Get("year"))
		if err != nil {
			json.NewEncoder(w).Encode(result{Error: "year 需要是整数"})
			return
		}
		json.NewEncoder(w).Encode(describe(year))
	})
	log.Println("3026 Years Checker on http://127.0.0.1:8086")
	log.Fatal(http.ListenAndServe("127.0.0.1:8086", mux))
}
