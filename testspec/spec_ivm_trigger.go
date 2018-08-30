package testspec

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
    assetId:
      type: string
      description: "Asset Identifier"
    hostname:
      typoe: string
      description: "Asset Hostname"
    altHostnames:
      type: "[]string"
      description: "The target hostnames for the entities"
    mac:
      type: string
      description: "Asset mac address"
    ip:
      type: string
      description: "IP address of the asset"
    altIps:
      type: "[]string"
      description: "Alternate asset ip addresses"
    cves:
      type: "[]string"
      description: "The cves for the vulns"
    softwareUpdates:
      type: "[]string"
      description: "Software update ids"
    solutions:
      type: "[]string"
      description: "The solution ids"
    timeout:
      type: integer
      description: "Timeout for creating a network containment entry, not sure of unit"
    policy:
      type: string
      description: "Policy to use for network containment"
  entity_set:
    targets:
      type: "[]target"
      description: "Targets to apply to"

triggers:
  ivm_trigger:
    output:
      entities:
        type: entity_set
        desccription: "the entities"
`
