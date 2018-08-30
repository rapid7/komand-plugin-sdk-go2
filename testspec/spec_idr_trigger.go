package testspec

const SpecIDRTrigger = `
plugin_spec_version: v2
name: idr_trigger
title: "IDR Trigger generator plugin"
description: "IDR Trigger generator plugin"
version: 1.0.0
vendor: rpaid7
tags: ["ivm", "trigger", "generator"]

types:
  asset:
    shortname:
      type: string
      description: "shortname of the asset"
      required: true
    fqdn:
      type: string
      description: "The fully qualified domain name"
  user:
    name:
      type: string
      description: "The name of the user, as '<firstname> <lastname>' if available, or account name otherwise"
      required: true
    distinguishedName:
      type: string
      description: "The Active Directory distinguished name of the user"
    emails:
      type: "[]string"
      description: "The email addresses associated with the user"
      required: true
  actor:
    users:
      type: "[]user"
      description: "The users that are part of the investigation"
    assets:
      type: "[]asset"
      description: "The assets that are part of the investigation"

triggers:
  idr_trigger:
    output:
      timestamp:
        type: string
        description: "The time the alert was triggered in ISO 8601 extended timestamp format"
      type:
        type: string
        description: "The type of alert"
      name:
        type: string
        description: "The human-readable name of the alert"
      description:
        type: string
        description: "The description of the alert"
      link:
        type: string
        description: "The deep link to the investigation in insightIDR"
      actors:
        type: actor
        description: "The insightIDR actors that are part of the investigation"
`
