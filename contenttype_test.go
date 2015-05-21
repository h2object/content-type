package content_type

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestContentType(t *testing.T) {
	assert.Equal(t, "text/html; charset=utf-8", DefaultContentTypeHelper.ContentTypeByFilename("avi.htm"))
	assert.Equal(t, "video/quicktime", DefaultContentTypeHelper.ContentTypeByFilename("avi.mov"))
}