package testdata

const EmbeddedNestedMessage = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "required": [
        "payload"
    ],
    "properties": {
        "payload": {
            "required": [
                "name",
                "id"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
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
                        }
                    ],
                    "title": "Topology"
                }
            },
            "additionalProperties": false,
            "type": "object"
        },
        "description": {
            "type": "string"
        }
    },
    "additionalProperties": true,
    "type": "object",
    "title": "Embedded Nested Message"
}`

const EmbeddedNestedMessagePayload = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "required": [
        "name",
        "id"
    ],
    "properties": {
        "name": {
            "type": "string"
        },
        "timestamp": {
            "type": "string"
        },
        "id": {
            "type": "integer"
        },
        "rating": {
            "type": "number"
        },
        "complete": {
            "type": "boolean"
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
                }
            ],
            "title": "Topology"
        }
    },
    "additionalProperties": true,
    "type": "object",
    "title": "Embedded Message Payload"
}`

const EmbeddedNestedMessageFail = `{
    "payload": {
        "topology": "FLAT"	
    }
}`

const EmbeddedNestedMessagePass = `{
    "payload": {
        "id": 1,
        "name": "something",
        "topology": "FLAT"
    }
}`
