package main

import (
   "net/http"
   "fmt"
   "log"
)

type helloJSON struct {
   UserName string `json:"user_name"`
   Content string `json:"content"`
}

func main() {
   http.HandleFunc("/", helloHandler)
   http.HandleFunc("/hoge", hogeHandler)
   http.HandleFunc("/todo", todoHandler)

   log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler (w http.ResponseWriter, r *http.Request) {
   fmt.Fprint(w, "Hello World\n")
}

func hogeHandler (w http.ResponseWriter, r *http.Request) {
   fmt.Fprint(w, "hoge\n")
}

func todoHandler (w http.ResponseWriter, r *http.Request) {
   switch r.Method {
   case http.MethodGet:
      w.WriteHeader(http.StatusOK)
      name := r.URL.Query().Get("name")
      fmt.Fprint(w, "GET todo\n")
   case http.MethodPost:
      body := r.Body
      defer body.Close()

      buf := new(bytes.Buffer)
      io.Copy(buf, body)

      var hello helloJSON
      json.Unmarshal(buf.bytes(), &hello)

      w.WriteHeader(http.StatusCreated)
      fmt.Fprint(w, "POST todo\n")
   default:
      w.WriteHeader(http.StatusMethodNotAllowed)
      fmt.Fprint(w, "Method not allowed.\n")
   }
}
