package DeletingStudent

import (
	"fmt"
	"log"
	"net/http"

	Model "../../Model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

type Student Model.Student

func DeleteAStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, OPTIONS, POST, DELETE")
	//w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "acheron"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	vars := mux.Vars(r)
	ID := vars["id"]

	if err := session.Query(`DELETE FROM student WHERE id =?`, ID).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Endpoint Hit: DeleteASTudent")

}
