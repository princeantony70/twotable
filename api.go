package main

import (

  "net/http"
  "html/template"
  "database/sql"
  "log"
  _ "github.com/go-sql-driver/mysql"

)


var db *sql.DB
var err error
var tpl *template.Template





type question struct {
ID  int
Question string
Qid  int
Answer string
Section string
Viewtype string

}





func checkErr(err error) {
    if err != nil {
   log.Fatalln(err)
        }
    }



func init(){
  db, err = sql.Open("mysql", "root:nfn@/shiftpixy")
	checkErr(err)
	err = db.Ping()
	checkErr(err)
}



func main(){

  http.HandleFunc("/tab1",tab1)
  http.HandleFunc("/tab2",tab2 )
}

func tab1(w http.ResponseWriter, req *http.Request){
if req.Method == http.MethodPost{
//req.ParseForm()
qus := question{}
qus.Question = req.FormValue("id")
qus.Answer1  = req.FormValue("question")
checkErr(err)
_,err = db.Exec(
  "INSERT INTO  tab1 (id, question) VALUES (?, ?)",
  qus.ID,
  qus.Question,
)
checkErr(err)
http.Redirect(w, req, "/thankyou", http.StatusSeeOther)
return
}
http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
}





func tab2(w http.ResponseWriter, req *http.Request){
if req.Method == http.MethodPost{
//req.ParseForm()
qus := question{}
qus.Question = req.FormValue("qid")
qus.Answer1  = req.FormValue("answer")
qus.Answer1  = req.FormValue("section")
qus.Answer1  = req.FormValue("viewtype")
checkErr(err)
_,err = db.Exec(
  "INSERT INTO  tab1 (id, question) VALUES (?, ?,?,?)",
  qus.Qid,
  qus.Answer,
  qus.Section,
  qus.Viewtype,
)
checkErr(err)
http.Redirect(w, req, "/thankyou", http.StatusSeeOther)
return
}
http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
}





func thankyou(w http.ResponseWriter, req *http.Request){
tpl,err :=template.ParseFiles("thankyou.html")

if err != nil{
  log.Fatalln("error parsing template thankyou",err)
}

err =tpl.ExecuteTemplate(w,"thankyou.html",nil)
if err !=nil{
  log.Fatalln("error executing template thank you",err)
}
checkErr(err)
}




