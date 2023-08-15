package testdata

const Proto2OptionalTypesArray = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "required": [
        "query"
    ],
    "properties": {
        "query": {
            "type": "string"
        },
        "options": {
            "properties": {
                "page_number": {
                    "type": [
                        "null",
                        "integer"
                    ]
                },
                "result_per_page": {
                    "type": [
                        "null",
                        "integer"
                    ]
                }
            },
            "additionalProperties": true,
            "type": [
                "null",
                "object"
            ]
        },
        "message": {
            "required": [
                "query"
            ],
            "properties": {
                "query": {
                    "type": "string"
                },
                "options": {
                    "type": [
                        "null",
                        "integer"
                    ]
                }
            },
            "additionalProperties": true,
            "type": [
                "null",
                "object"
            ]
        },
        "nested_enum": {
            "enum": [
                "FLAT",
                0,
                "NESTED_OBJECT",
                1,
                "NESTED_MESSAGE",
                2,
                "ARRAY_OF_TYPE",
                3,
                "ARRAY_OF_OBJECT",
                4,
                "ARRAY_OF_MESSAGE",
                5
            ],
            "title": "Nested Enum",
            "type": [
                "string",
                "integer",
                "null"
            ]
        }
    },
    "additionalProperties": true,
    "type": "object",
    "title": "Proto 2 Optional Types Array"
}`

const Proto2OptionalTypesArrayMessage = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "required": [
        "query"
    ],
    "properties": {
        "query": {
            "type": "string"
        },
        "options": {
            "type": [
                "null",
                "integer"
            ]
        }
    },
    "additionalProperties": true,
    "type": "object",
    "title": "Proto 2 Optional Types Array Message"
}`

const Proto2OptionalTypesArrayFail = `{
    "options": {
        "page_number": 4
    }
}`

const Proto2OptionalTypesArrayPass = `{
    "query": "what?",
    "options": {
        "page_number": 4
    }
}`
