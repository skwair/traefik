package brotli

// parsedContentType is the parsed representation of one of the inputs to ContentTypes.
// From https://github.com/klauspost/compress/blob/master/gzhttp/compress.go#L401.
type parsedContentType struct {
	mediaType string
	params    map[string]string
}

// equals returns whether this content type matches another content type.
func (p parsedContentType) equals(mediaType string, params map[string]string) bool {
	if p.mediaType != mediaType {
		return false
	}

	// if p has no params, don't care about other's params
	if len(p.params) == 0 {
		return true
	}

	// if p has any params, they must be identical to other's.
	if len(p.params) != len(params) {
		return false
	}

	for k, v := range p.params {
		if w, ok := params[k]; !ok || v != w {
			return false
		}
	}

	return true
}
