package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/scuti/golang-github-fetch-user-repos-example/model"
)

type Error struct {
	Message string `json:"message"`
}

func (o *Error) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func getUserRepositoriesAction(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	reqUrl := "https://api.github.com/users/" + userId + "/repos"

	if resp, err := http.Get(reqUrl); err != nil {
		fmt.Println("err: ", err)
	} else {
		var bodyBytes []byte
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		var content []*model.GitRepository
		var result []*model.GitRepositoryResponse

		json.NewDecoder(resp.Body).Decode(&content)
		if content != nil && len(content) > 0 {
			w.WriteHeader(200)

			for _, repo := range content {
				result = append(result, repo.ToClientResponseItem())
			}
		}

		w.WriteHeader(200)
		w.Write([]byte(model.GitRepositoryResponseListToJson(result)))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{userId:[A-Za-z0-9]+}/repositories", http.HandlerFunc(getUserRepositoriesAction))

	port := "5000"
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println(fmt.Sprintf("Server started on port: %s", port))
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error when start server. Error: %s", err.Error()))
	}
}
