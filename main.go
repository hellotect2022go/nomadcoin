package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/hellotect2022go/nomadcoin/blockchain"
)

const port string = ":4000"
const templateDir string = "template/"

var templates *template.Template

type HomeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	//tmpl := template.Must(template.ParseFiles("template/pages/home.gohtml"))
	data := HomeData{"Home", blockchain.AllBlocks()}
	//tmpl.Execute(w, data)
	templates.ExecuteTemplate(w, "home", data)
}

func add(w http.ResponseWriter, r *http.Request) {
	fmt.Println("????", r.Method)
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		fmt.Println(data)
		blockchain.GetBlockChain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}

}

func main() {
	fmt.Println(blockchain.AllBlocks()[0].Data)
	//tmpl := template.Must(template.ParseFiles("template/pages/home.gohtml"))
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // 프로그램이 Exit(1) : {error 로 인해 종료시킴} 로그를 출력해줌

}
