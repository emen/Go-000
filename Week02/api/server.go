package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/emen/Go-000/Week02/model"
	"github.com/emen/Go-000/Week02/model/mysql"
)

func main() {
	mysqlClient := &mysql.MySQLClient{}
	userService := model.NewUserService(mysqlClient)

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		paths := strings.Split(r.URL.Path, "/")
		id, err := strconv.Atoi(paths[len(paths)-1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := userService.Get(id)
		if errors.Is(err, model.ErrNoRecord) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// no no no
		encoder := json.NewEncoder(w)
		encoder.Encode(user)
	})

	http.ListenAndServe(":3000", nil)
}
