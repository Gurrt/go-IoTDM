package IoTDM

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type RuntimeInfo struct {
	DB *DatabaseHandle
}

type ErrorJSON struct {
	Error string `json:"error"`
}

type genericResultJSON struct {
	Result interface{} `json:"result"`
}

func (runtime *RuntimeInfo) StartAPI() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/profiles/{id}", runtime.getProfile).Methods("GET")
	router.HandleFunc("/profiles/{id}/{field}/convert", runtime.convert).Methods("GET")
	fmt.Println("Starting API Server @ port 8081")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS()(router)))
}

func (runtime *RuntimeInfo) getProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	profile, err := runtime.DB.GetProfile(id)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	json.NewEncoder(w).Encode(profile)
}

func (runtime *RuntimeInfo) convert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	profile, err := runtime.DB.GetProfile(id)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	f := vars["field"]
	field, found := profile.Fields[f]
	if !found {
		w.WriteHeader(404)
		return
	}

	query := r.URL.Query()
	valueArr := query["value"]
	fromArr := query["from"]
	toArr := query["to"]
	if (fromArr == nil && toArr == nil) || valueArr == nil {
		json.NewEncoder(w).Encode(ErrorJSON{Error: "Expecting query parameter 'value' and atleast one of the following: 'from', 'to'"})
		return
	}
	value, from, to := getValueFromto(valueArr, fromArr, toArr)

	result, err := field.convert(from, to, value)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorJSON{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(genericResultJSON{Result: result})
}

func getValueFromto(valueArr []string, fromArr []string, toArr []string) (string, string, string) {
	var val string
	var from string
	var to string

	if fromArr == nil {
		from = "domain"
	} else {
		from = fromArr[0]
	}

	if toArr == nil {
		to = "domain"
	} else {
		to = toArr[0]
	}

	val = valueArr[0]

	return val, from, to
}
