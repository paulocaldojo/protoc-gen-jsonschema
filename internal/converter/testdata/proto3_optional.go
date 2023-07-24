package testdata

const Proto3Optional = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Proto3Optional",
    "definitions": {
        "Proto3Optional": {
            "required": [
                "query"
            ],
            "properties": {
                "query": {
                    "type": "string"
                },
                "page_number": {
                    "type": "integer"
                },
                "result_per_page": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Proto 3 Optional"
        }
    }
}`

const Proto3OptionalFail = `{
	"page_number": 4
}`

const Proto3OptionalPass = `{
	"query": "what?",
	"page_number": 4
}`
