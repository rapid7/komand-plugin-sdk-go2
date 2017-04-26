# Komand Go Plugin SDK [![Build Status](https://ci.dev.komand.net/api/badges/komand/plugin-sdk-go2/status.svg)](https://ci.dev.komand.net/komand/plugin-sdk-go2)

## Welcome to the Komand Plugin SDK for Go!

Official documentation is maintained at [docs.komand.com](https://docs.komand.com/docs/plugin-development).

## Generating a Plugin

To generate a plugin using only the SDK alone, and not a part of the Komand Platform, you can use the following steps:

* Review the documentation linked above on how to structure a schema file. An example is included in this repo as a high level guide, but is not comprehensive.
* Fill out your schema
* Build the binary using `make build`
* `plugin-sdk-go -spec="full/path/to.plugin.spec.example.yaml" -package="github.com/:company/:pluginname/"
* Review the documentation linked above for how to fill out and run the plugin