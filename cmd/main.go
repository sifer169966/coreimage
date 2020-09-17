package main

import (
	"encoding/json"
	"fl/coreimage/playground"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	TAGBOX_URL string = "api for connect to tags box service"
)

func httpHandle(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		request := map[string]interface{}{
			"p_number": "",
			"n_number": "",
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			w.Write([]byte("error message"))
			return
		}
		playground.AppendData(request["p_number"].(string), request["n_number"].(string))
		w.Write([]byte("Finish"))
		return
	}).Methods("POST")
	r.HandleFunc("/tags/box", HandleTagsImage).Methods("POST")
}
func HandleTagsImage(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Fatal(err)
	}
	multipartFile := map[string]*multipart.FileHeader{}
	_, multipartFile["file"], err = r.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	for file, values := range multipartFile {
		var setFile *multipart.FileHeader
		if file == "file" {
			setFile = values
		}
		checkType := strings.Split(setFile.Header.Get("Content-Type"), "/")
		if checkType[0] != "image" {
			log.Fatal("invalid image data")
		}
		switch checkType[1] {
		case "jpg":
		case "png":
		case "jpeg":
		default:
			log.Fatal("invalid image type")
		}
		file, err := setFile.Open()
		if err != nil {
			log.Fatal(err)
		}
		// byteContainer, err := ioutil.ReadAll(file)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer file.Close()
		// buff := new(bytes.Buffer)
		// buff.Write()
		//reader := bytes.NewReader(buff.Bytes())
		PostToTagsImage(TAGBOX_URL, file)
	}
}
func PostToTagsImage(url string, file io.Reader) (string, []byte) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Status, content
}

func main() {
	router := mux.NewRouter()
	httpHandle(router)
	fmt.Println("Hello World")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"Accept", "multipart/form-data", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
