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
      type: string
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
  entity:
    cves:
      type: "[]string"
      description: "The cves for the vulns"
    policy:
      type: string
      description: "Policy to use for network containment"
    softwareUpdates:
      type: "[]string"
      description: "Software update ids"
    timeout:
      type: integer
      description: "Timeout for creating a network containment entry, not sure of unit"
    targets:
      type: "[]target"
      description: "The assets that are part of the workflow"

triggers:
  ivm_trigger:
    output:
      timestamp:
        type: date
        description: "The time the event was triggered"
      type:
        type: string
        description: "type of workflow"
      name:
        type: string
        description: "The human-readable name of the alert"
      description:
        type: string
        description: "The description of the alert"
      link:
        type: string
        description: "The deep link to the triggered workflow"
      automationId:
        type: string
        description: "The unique identifier of an automated workflow instance"
      entities:
        type: "[]entity"
        description: "Entities that are involved in triggered workflow"
`
