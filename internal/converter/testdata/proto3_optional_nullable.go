package testdata

const Proto3OptionalNullable = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Proto3OptionalNullable",
    "definitions": {
        "Proto3OptionalNullable": {
            "required": [
                "query",
                "option"
            ],
            "properties": {
                "query": {
                    "type": "string"
                },
                "options": {
                    "additionalProperties": true,
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "$ref": "#/definitions/samples.Proto3OptionalNullable.Options"
                        }
                    ]
                },
                "option": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Proto 3 Optional Nullable"
        },
        "samples.Proto3OptionalNullable.Options": {
            "properties": {
                "page_number": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "integer"
                        }
                    ]
                },
                "result_per_page": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "integer"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Options"
        }
    }
}`

const Proto3OptionalNullableFail = `{
    "options": {
        "page_number": 4
    }
}`

const Proto3OptionalNullablePass = `{
    "query": "what?",
    "options": {
        "page_number": 4
    }
}`
