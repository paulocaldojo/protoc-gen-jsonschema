package testdata

const Proto3OptionalNestedObject = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Proto3OptionalNestedObject",
    "definitions": {
        "Proto3OptionalNestedObject": {
            "required": [
                "query"
            ],
            "properties": {
                "query": {
                    "type": "string"
                },
                "options": {
                    "$ref": "#/definitions/samples.Proto3OptionalNestedObject.Options",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Proto 3 Optional Nested Object"
        },
        "samples.Proto3OptionalNestedObject.Options": {
            "properties": {
                "page_number": {
                    "type": "integer"
                },
                "result_per_page": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Options"
        }
    }
}`

const Proto3OptionalNestedObjectFail = `{
	"options": {
		"page_number": 4
	}
}`

const Proto3OptionalNestedObjectPass = `{
	"query": "what?",
	"options": {
		"page_number": 4
	}
}`
