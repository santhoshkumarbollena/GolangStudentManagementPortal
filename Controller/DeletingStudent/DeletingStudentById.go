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
