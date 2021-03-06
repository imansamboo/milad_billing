package main

import (
    excelreader "./excelreader"
    file_handler "./file_handler"
    orm "./orm"
    _ "bufio"
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    _ "os"
    "strings"
)

type Article struct {
    Id      string `json:"Id"`
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
    // NOTE: Ordering is important here! This has to be defined before
    // the other `/article` endpoint.
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/excel", insertExcel2Database).Methods("POST")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    // Loop over all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func updateArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]
    reqBody, _ := ioutil.ReadAll(r.Body)

    for index, article := range Articles {
        // if our id path parameter matches one of our
        // articles
        if article.Id == id {
            // updates our Articles array to remove the
            // article
            var article Article
            json.Unmarshal(reqBody, &article)

            Articles = append(Articles[:index], Articles[index+1:]...)
            Articles = append(Articles, article)
            json.NewEncoder(w).Encode(article)
        }

    }

}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    // once again, we will need to parse the path parameters
    vars := mux.Vars(r)
    // we will need to extract the `id` of the article we
    // wish to delete
    id := vars["id"]

    // we then need to loop through all our articles
    for index, article := range Articles {
        // if our id path parameter matches one of our
        // articles
        if article.Id == id {
            // updates our Articles array to remove the
            // article
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    Articles = []Article{
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}

func insertExcel2Database(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Excel Page!")
    err := r.ParseMultipartForm(32 << 20) // limit your max input length!
    if err != nil {
        panic(err)
    }
    //var buf bytes.Buffer
    // in your case file would be fileupload
    file, header, err := r.FormFile("attachment")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    name := strings.Split(header.Filename, ".")
    fmt.Printf("File name %s\n", name[0])
    // Copy the file data to my buffer
    //io.Copy(&buf, file)
    // do something with the contents...
    // I normally have a struct defined and unmarshal into a struct, but this will
    // work as an example
    //reqBody, _ := ioutil.ReadAll(r.MultipartForm.File)
	fileName := file_handler.RandStringRunes(8)
	fileName += ".xlsx"
	path := "report_files/" + fileName
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	io.Copy(f, file)
	err, cdrList := excelreader.ConvertExcel2EntityList(path)
    if err != nil {
    	panic(err)
	}
	err = orm.BatchInsert(cdrList)
	if err != nil {
		panic(err)
	}
    orm.Insert()
    fmt.Println("Endpoint Hit: Excel Page saved in database")
}