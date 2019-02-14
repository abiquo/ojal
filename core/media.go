package core

import (
	"fmt"
	"strings"
)

func getMedia(media string) string {
	if media != "" && !strings.ContainsAny(media, "/") {
		media = fmt.Sprintf("application/vnd.abiquo.%v+json", media)
	}
	return media
}

// Media ...
type Media string

// Media resolves m Media string
func (m Media) Media() string {
	return getMedia(string(m))
}

// String ...
func (m Media) String() string {
	return m.Media()
}
