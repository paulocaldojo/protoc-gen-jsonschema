package testdata

const DisallowEmbeddedEnumsAsStringsOnly = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Money",
    "definitions": {
        "Money": {
            "properties": {
                "amount": {
                    "type": "string"
                },
                "currency": {
                    "type": "string",
                    "title": "Currency"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Money"
        }
    }
}`
