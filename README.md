HTTP Content-Type Map
---

this package is a toolkit for HTTP。Just map file to content-type

# Quick Start

````
	import "github.com/h2object/content-type"

	content_type.DefaultContentTypeHelper.ContentTypeByFilename("filename.htm")

````

# Append your own content-type helper

````
	helper := content_type.NewContentTypeHelper()

	helper.AppendSuffix("htm", "text/html")

````
