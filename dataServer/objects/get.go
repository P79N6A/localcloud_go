package objects

import (
	"net/http"
	"os"
	"localcloud_go/util/define"
	"strings"
	"git.elenet.me/MaxQ/log"
	"io"
)

func get(w http.ResponseWriter, r *http.Request) {
	f, e := os.Open(os.Getenv(define.STORAGE_ROOT) + "/objects/" +
		strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
