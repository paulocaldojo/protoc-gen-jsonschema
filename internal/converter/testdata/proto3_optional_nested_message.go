package testdata

const Proto3OptionalNestedMessage = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Proto3OptionalNestedMessage",
    "definitions": {
        "Proto3OptionalNestedMessage": {
            "required": [
                "query"
            ],
            "properties": {
                "query": {
                    "type": "string"
                },
                "options": {
                    "$ref": "#/definitions/samples.Proto3NestedMessageOptions",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Proto 3 Optional Nested Message"
        },
        "samples.Proto3NestedMessageOptions": {
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
            "title": "Proto 3 Nested Message Options"
        }
    }
}`

const Proto3NestedMessageOptions = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Proto3NestedMessageOptions",
    "definitions": {
        "Proto3NestedMessageOptions": {
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
            "title": "Proto 3 Nested Message Options"
        }
    }
}`

const Proto3OptionalNestedMessageFail = `{
	"options": {
		"page_number": 4
	}
}`

const Proto3OptionalNestedMessagePass = `{
	"query": "what?",
	"options": {
		"page_number": 4
	}
}`
