package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)


func homePage(w http.ResponseWriter, r *http.Request){
	t := time.Now()
	fmt.Fprintf(w,t.Format("15h04"))
}


func writeTofile(data string){
	
	if _, err := os.Stat("data.txt"); os.IsNotExist(err) {
		os.OpenFile("data.txt", os.O_RDONLY|os.O_CREATE, 0666)
	}	
	file, err := os.OpenFile("data.txt", os.O_WRONLY|os.O_APPEND, 0644)


	fileContent := ""
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	fileContent += data
	fileContent += "\n"
	defer file.Close()

	file.WriteString(fileContent)

}

func returnAllAPayloads(w http.ResponseWriter, r *http.Request){
	var author string
	var entry string
	var words string

	switch r.Method {
		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w,"Something went bad")
				return
			}
			for key, value := range r.PostForm {
				if key == "author"{
					author = value[0]
				}else if key == "entry"{
					entry = value[0]
					words += entry
					
				}
			}	
			writeTofile(words)
			fmt.Fprintf(w,author,":", entry)
	}
}

func getFromFile(w http.ResponseWriter, r *http.Request){
	content, err := ioutil.ReadFile("data.txt")
	if err != nil {
		 log.Fatal(err)
	}

   fmt.Fprintf(w,string(content))
}

func handleRequests() {
    http.HandleFunc("/", homePage)
	http.HandleFunc("/hello", returnAllAPayloads)
	http.HandleFunc("/entries", getFromFile)
    log.Fatal(http.ListenAndServe(":4567", nil))
}

func main() {
    handleRequests()
}