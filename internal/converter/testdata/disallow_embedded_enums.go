package testdata

const DisallowEmbeddedEnums = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Money",
    "definitions": {
        "Money": {
            "properties": {
                "amount": {
                    "type": "string"
                },
                "currency": {
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        }
                    ],
                    "title": "Currency"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Money"
        }
    }
}`
