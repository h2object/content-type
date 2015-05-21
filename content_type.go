package content_type

import (
	"sync"
	"strings"
)

const DefaultContentType = "application/octet-stream"

type ContentTypeHelper struct{
	sync.RWMutex
	suffixes map[string]string
}

func NewContentTypeHelper() *ContentTypeHelper {
	return &ContentTypeHelper{
		suffixes: make(map[string]string),
	}
}

func (helper *ContentTypeHelper) AppendSuffix(suffix, contentType string) {
	helper.Lock()
	defer helper.Unlock()

	helper.suffixes[suffix] = contentType
}

func (helper *ContentTypeHelper) RemoveSuffix(suffix string) {
	helper.Lock()
	defer helper.Unlock()

	delete(helper.suffixes, suffix)
}

func (helper *ContentTypeHelper) ContentTypeByFilename(filename string) string {
	dot := strings.LastIndex(filename, ".")
	if dot == -1 || dot+1 >= len(filename) {
		return DefaultContentType
	}

	extension := filename[dot+1:]

	helper.RLock()
	defer helper.RUnlock()	
	contentType, ok := helper.suffixes[extension]
	if ok {
		if strings.HasPrefix(contentType, "text/") {
			return contentType + "; charset=utf-8"
		}
		return contentType	
	}

	return DefaultContentType
}
