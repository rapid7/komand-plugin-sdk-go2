package testspec

const SpecSimpleTypes = `
plugin_spec_version: v2
name: chatbot_slack
title: "Chatbot for Slack"
description: "Chatbot integrations for Slack (build version: 1.0.0.7)"
version: 1.0.0
vendor: komand_chatbot
tags: ["slack", "chat", "chatbot"]

types:
  thing:
    larg:
      type: int
  ctx_channel:
    id:
      title: "ID"
      type: string
    name:
      type: string
    date:
      type: date
    things_n_such:
      type: "[]thing"
actions:
  my_action:
    title: "ACTION!"
    description: "Lights, Camera"
    input:
      stuff:
        type: thing
        title: "This is some STUFFFFFF"
        description: "Things can be stuff, it turns out, and ints?"
`
