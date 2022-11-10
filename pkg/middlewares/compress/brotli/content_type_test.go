package brotli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseContentType_equals(t *testing.T) {
	testCases := []struct {
		desc      string
		pct       parsedContentType
		mediaType string
		params    map[string]string
		expect    assert.BoolAssertionFunc
	}{
		{
			desc:   "empty parsed content type",
			expect: assert.True,
		},
		{
			desc: "simple content type",
			pct: parsedContentType{
				mediaType: "plain/text",
			},
			mediaType: "plain/text",
			expect:    assert.True,
		},
		{
			desc: "content type with params",
			pct: parsedContentType{
				mediaType: "plain/text",
				params: map[string]string{
					"charset": "utf8",
				},
			},
			mediaType: "plain/text",
			params: map[string]string{
				"charset": "utf8",
			},
			expect: assert.True,
		},
		{
			desc: "different content type",
			pct: parsedContentType{
				mediaType: "plain/text",
			},
			mediaType: "application/json",
			expect:    assert.False,
		},
		{
			desc: "content type with params",
			pct: parsedContentType{
				mediaType: "plain/text",
				params: map[string]string{
					"charset": "utf8",
				},
			},
			mediaType: "plain/text",
			params: map[string]string{
				"charset": "latin-1",
			},
			expect: assert.False,
		},
		{
			desc: "different number of parameters",
			pct: parsedContentType{
				mediaType: "plain/text",
				params: map[string]string{
					"charset": "utf8",
				},
			},
			mediaType: "plain/text",
			params: map[string]string{
				"charset": "utf8",
				"q":       "0.8",
			},
			expect: assert.False,
		},
	}

	for _, test := range testCases {
		test := test

		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			test.expect(t, test.pct.equals(test.mediaType, test.params))
		})
	}
}
