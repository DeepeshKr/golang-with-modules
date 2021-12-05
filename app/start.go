import (
	"fmt"
	"log"
	"net/http"
	"handlers/greet"
	"handlers/getAllCustomers"
)

func Start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:9000", nil))

}