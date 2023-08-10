package testdata

const Proto2OptionalNullable = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Proto2OptionalNullable",
    "definitions": {
        "Proto2OptionalNullable": {
            "properties": {
                "payload": {
                    "additionalProperties": true,
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "$ref": "#/definitions/samples.Proto2OptionalNullable.NestedPayload"
                        }
                    ]
                },
                "description": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Proto 2 Optional Nullable"
        },
        "samples.Proto2OptionalNullable.NestedPayload": {
            "required": [
                "name",
                "id"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ]
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "number"
                        }
                    ]
                },
                "complete": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "boolean"
                        }
                    ]
                },
                "topology": {
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
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        },
                        {
                            "type": "null"
                        }
                    ],
                    "title": "Topology"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Nested Payload"
        }
    }
}`

const Proto2OptionalNullableFail = `{
    "payload": {
        "topology": "FLAT"
    }
}`

const Proto2OptionalNullablePass = `{
    "payload": {
        "name": "something",
        "timestamp": "1970-01-01T00:00:00Z",
        "id": 1,
        "rating": 100,
        "complete": true
    }
}`
