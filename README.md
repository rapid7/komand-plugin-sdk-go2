# Komand Go Plugin SDK [![Build Status](https://ci.dev.komand.net/api/badges/komand/plugin-sdk-go2/status.svg)](https://ci.dev.komand.net/komand/plugin-sdk-go2)

## Welcome to the Komand Plugin SDK for Go!

Official documentation is maintained at [docs.komand.com](https://docs.komand.com/docs/plugin-development).

## Generating a Plugin

To generate a plugin using only the SDK alone, and not a part of the Komand Platform, you can use the following steps:

* Review the documentation linked above on how to structure a schema file. 
  * An example is included in this readme as a high level guide, but is not comprehensive.
* Fill out your schema
* Build the binary using `make build`
* `plugin-sdk-go -spec="full/path/to.plugin.spec.example.yaml" -package="github.com/:company/:pluginname/"
* Review the documentation linked above for how to fill out and run the plugin

Note that not all functionality present in the Makefile of the plugin you build will work naturally without the presence of the Komand utilities from our platform. Contact Komand if you run into any issues.

## Example Spec

```yaml
plugin_spec_version: v2
name: example
description: "Example plugin"
version: 1.0.0
vendor: demo
tags: ["example", "demo"]
icon: "data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4KPCEtLSBHZW5lcmF0b3I6IEFkb2JlIElsbHVzdHJhdG9yIDIwLjEuMCwgU1ZHIEV4cG9ydCBQbHVnLUluIC4gU1ZHIFZlcnNpb246IDYuMDAgQnVpbGQgMCkgIC0tPgo8c3ZnIHZlcnNpb249IjEuMSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayIgd2lkdGg9IjE0MC41cHgiCgkgaGVpZ2h0PSIxMDBweCIgdmlld0JveD0iMCAwIDE0MC41IDEwMCIgeG1sOnNwYWNlPSJwcmVzZXJ2ZSI+CjxnIGlkPSJrb21hbmRfaWNvbiIgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTQ4LjEsIC00OS42KSI+Cgk8cmVjdCB4PSIxNjcuNSIgeT0iNTEuMiIgZmlsbD0iI0ZGRkZGRiIgd2lkdGg9IjE5LjMiIGhlaWdodD0iOTguNCIvPgoJPHBvbHlnb24gZmlsbD0iI0ZGRkZGRiIgcG9pbnRzPSIxMjguOSw1MSAxMjEuNSw1MSAxMTQuOSw1MSAxNTAuNCwxMDAuOSAxMTQuOSwxNTEgMTIxLjQsMTUxIDEyOC45LDE1MSAxNjQuNCwxMDAuOSAJIi8+Cgk8cG9seWdvbiBvcGFjaXR5PSIwLjciIGZpbGw9IiNGRkZGRkYiIHBvaW50cz0iMTAyLjksNTEgOTUuNSw1MSA4OC45LDUxIDEyNC40LDEwMC45IDg4LjksMTUxIDk1LjQsMTUxIDEwMi45LDE1MSAxMzguNCwxMDAuOSAJIi8+Cgk8cG9seWdvbiBvcGFjaXR5PSIwLjM1IiBmaWxsPSIjRkZGRkZGIiBwb2ludHM9IjgyLjksNTEgNzUuNSw1MSA2OC45LDUxIDEwNC40LDEwMC45IDY4LjksMTUxIDc1LjQsMTUxIDgyLjksMTUxIDExOC40LDEwMC45IAkiLz4KCTxwb2x5Z29uIG9wYWNpdHk9IjAuMTUiIGZpbGw9IiNGRkZGRkYiIHBvaW50cz0iNjIuOSw1MSA1NS41LDUxIDQ4LjksNTEgODQuNCwxMDAuOSA0OC45LDE1MSA1NS40LDE1MSA2Mi45LDE1MSA5OC40LDEwMC45IAkiLz4KCTxwb2x5Z29uIGZpbGw9IiNGRkZGRkYiIHBvaW50cz0iMTIyLjksNTEgMTE1LjUsNTEgMTA4LjksNTEgMTQ0LjQsMTAwLjkgMTA4LjksMTUxIDExNS40LDE1MSAxMjIuOSwxNTEgMTU4LjQsMTAwLjkgCSIvPgo8L2c+Cjwvc3ZnPgo="
help: |
  ## About
  Example plugin for testing.

  ## Actions
  ### Say Goodbye
  This action is used to say goodbye to a name.

  #### Input

  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |name|string|None|False|Name to say goodbye to|None|

  #### Output

  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |message|string|False|None|
  |time|date|False|None|

  ### Wave
  This action is used to wave a user specified number of times.
  #### Input

  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |count|integer|None|False|None|None|

  #### Output

  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |message|string|False|None|

  ## Triggers
  ### Emit Greeting
  This trigger is used to triggers a greeting every interval seconds (by default, 15 seconds).
  #### Input

  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |interval|integer|15|False|How often (in seconds) to trigger a greeting|None|

  #### Output

  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |greeting|string|False|None|
  |time|date|False|None|

  ## Connection
  The connection configuration accepts the following parameters:

  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |hostname|string|None|False|None|None|

  ## Troubleshooting
  This plugin does not contain any troubleshooting information.

  ## Workflows
  Examples:
  * Testing

  ## References
  * [Komand](https://www.komand.com/)

connection:
  hostname:
    type: string

triggers:
  emit_greeting:
    description: "Triggers a greeting every Interval seconds (by default, 15 seconds)"
    input:
      interval:
        type: integer
        description: "How often (in seconds) to trigger a greeting"
        default: 15
    output:
      greeting:
        type: string
      time:
        type: date

actions:
  wave:
    description: "Wave"
    input:
      count:
        type: integer
    output:
      message:
        type: string
  say_goodbye:
    description: "Say goodbye"
    input:
      name:
        type: string
        description: "Name to say goodbye to"
    output:
      message:
        type: string
      time:
        type: date
```

This would generate a plugin with 2 actions: Wave, and SayGoodbye

Additionally, it would come with 1 trigger: EmitGreeting.

For more information on Actions, Triggers, and plugins themselves, please refer to the linked documentation above.