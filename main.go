package main

import(
	"fmt"
	"strings"
	"io/ioutil"
	"time"
	"net/http"
	"encoding/json"
	"chi"
)
	

type Article struct{
	Timestamp string `json:"timestamp"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

type ArticleRequest struct{
	Key       string `json:"key"`
	Value     string `json:"value"`
}

var Articles=[]Article{
	// Article{Timestamp: "2019-12-02T06:63:32Z", Key: "a", Value: "some value"},
	// Article{Timestamp: "2019-12-02T06:53:35Z", Key: "asdf", Value: "some other value"},	
}

func listAllArticles(w http.ResponseWriter, r *http.Request){
	for i:=0;i<len(Articles);i++{
		for j:=len(Articles)-1;j>i;j--{
			if(strings.Compare(Articles[j].Timestamp,Articles[j-1].Timestamp)<0){
				temp:=Articles[j]
				Articles[j]=Articles[j-1]
				Articles[j-1]=temp
			}
		}
	}
	json.NewEncoder(w).Encode(Articles)
}

func createArticle(w http.ResponseWriter, r *http.Request){
	str,_:= ioutil.ReadAll(r.Body)
	s:=new(ArticleRequest)
	json.Unmarshal(str, &s)

	var temp Article
	temp=Article{Timestamp: time.Now().String(), Key: s.Key, Value: s.Value}
	Articles=append(Articles,temp)
}


func main(){
	r := chi.NewRouter()

	r.Post("/add",createArticle)
	r.Get("/list",listAllArticles)

	http.ListenAndServe(":80", r)
}