package objects

import (
	"net/http"
	"os"
	"localcloud_go/util/define"
	"strings"
	"git.elenet.me/MaxQ/log"
	"io"
)

func put(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create(os.Getenv(define.STORAGE_ROOT) + "/objects/" +
		strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}
