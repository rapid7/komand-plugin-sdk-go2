package sdk

const SpecIVMTrigger = `
plugin_spec_version: v2
name: idr_trigger
title: "IDR Trigger generator plugin"
description: "IDR Trigger generator plugin"
version: 1.0.0
vendor: rpaid7
tags: ["ivm", "trigger", "generator"]

types:
  target:
    hostnames:
      type: "[]string"
      description: "The target hostnames for the entities"
    cves:
      type: "[]string"
      description: "The cves for the vulns"
    solutions:
      type: "[]string"
      description: "The solution ids"
  entity_set:
    targets:
      type: "[]target"
      description: "Targets to apply to"

triggers:
  idr_trigger:
    output:
      entities:
        type: entity_set
        desccription: "the entities"
`
