package apierrors

import(
	"fmt"
	"net/http"
)

func PrintError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "%v", err)
}
