package testdata

const OptionalTypesArrayEnums = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "required": [
        "required_some_enum"
    ],
    "properties": {
        "required_some_enum": {
            "enum": [
                "FLAT",
                "NESTED_OBJECT",
                "NESTED_MESSAGE",
                "ARRAY_OF_TYPE",
                "ARRAY_OF_OBJECT",
                "ARRAY_OF_MESSAGE"
            ],
            "type": "string",
            "title": "Some Enum"
        },
        "opt_some_enum": {
            "enum": [
                "FLAT",
                "NESTED_OBJECT",
                "NESTED_MESSAGE",
                "ARRAY_OF_TYPE",
                "ARRAY_OF_OBJECT",
                "ARRAY_OF_MESSAGE"
            ],
            "title": "Some Enum",
            "type": [
                "string",
                "null"
            ]
        }
    },
    "additionalProperties": true,
    "type": "object",
    "title": "Optional Types Array Enums"
}`
