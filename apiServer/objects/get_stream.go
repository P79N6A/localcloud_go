package objects

import (
	"io"
	"localcloud_go/apiServer/locate"
	"fmt"
	"localcloud_go/util/objectstream"
)

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return objectstream.NewGetStream(server, object)
}