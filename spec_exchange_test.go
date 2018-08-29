package sdk

const SpecExchange = `
plugin_spec_version: v2
name: microsoft_exchange
title: "Microsoft Exchange"
description: "Microsoft Exchange is a mail and calendaring service"
version: 5.0.0
vendor: komand
tags: ["microsoft", "exchange", "email", "calendar"]
icon: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAfQAAAH0CAYAAADL1t+KAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAA4ZpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTM4IDc5LjE1OTgyNCwgMjAxNi8wOS8xNC0wMTowOTowMSAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtcE1NOk9yaWdpbmFsRG9jdW1lbnRJRD0ieG1wLmRpZDowNWZiOGExZS04MWVhLTRmOTItYWMyOS01M2JiMWViM2E5NjciIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6NjUyNzAxRkU4NUNCMTFFNzkwODdCOERDQkZFREU4RTAiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6NjUyNzAxRkQ4NUNCMTFFNzkwODdCOERDQkZFREU4RTAiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTcgKE1hY2ludG9zaCkiPiA8eG1wTU06RGVyaXZlZEZyb20gc3RSZWY6aW5zdGFuY2VJRD0ieG1wLmlpZDo3NmRmYTM0OS1mNmVhLTQ4ZTQtYTMyOS0xNTMxMGJlODIxZDciIHN0UmVmOmRvY3VtZW50SUQ9ImFkb2JlOmRvY2lkOnBob3Rvc2hvcDo1YTM2ZWQxZS1iNTk4LTExNzktODQx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"
language: python
help: |
  ## About
  [Microsoft Exchange](https://products.office.com/en-us/exchange/email) is a mail and calendaring service.
  ## Actions
  ### Delete Email
  This action is used to delete a specified email.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |email_id|string|None|True|ID of the email to delete|None|
  |account|string|None|False|Account to delete the email from. The account configured in your Connection must have Impersonation access|None|
  |deletion_method|string|None|True|Deletes an email. Emails can be hard deleted, soft deleted (kept in recoverable folder), or moved to trash|['Hard', 'Soft', 'Trash']|
  |flatten_attachments|boolean|False|False|Flattens attachments to a single list|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |email|email|False|Email that has been deleted|
  Example output:
  {
    "email": {
        "account": "bob@test.com",
        "attached_emails": [],
        "attached_files": [],
        "body": "Test email",
        "flattened_attached_emails": [],
        "flattened_attached_files": [],
        "headers": [{
            "name": "Received",
            "value": ""
        }],
        "id": "AQMkAGRlNjhiMTkwLTJhMjUtNGI1ZS1hMGVlAC01MjUyZjgxZjM4Y2IARgAAAxV9OuuaYO1GuNrmyowDcygHAOrvwEx5H1ZNgslFeFvir+EAAAIBDAAAAOrvwEx5H1ZNgslFeFvir+EAAAIFaAAAAA==",
        "is_read": false,
        "sender": "joe@test.com",
        "subject": "Check this out!"
    }
  }
  ### Set Email Categories
  This action is used to set a defined category on an email.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |email_id|string|None|True|ID of the email to move|None|
  |account|string|None|False|Account to set the email category on. The account configured in your Connection must have Impersonation access|None|
  |categories|[]string|None|True|Categories to set on the email. Existing categories on the email will be overwritten|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |success|boolean|True|Whether or not the category set was successful|
  Example output:
  {
    "success": true
  }
  ### Search Inboxes
  This action is used to search account inboxes for emails using an EWS QueryString. Currently, searches are limited to the inbox folder and subfolders are skipped.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |accounts|[]string|None|True|List of accounts to be searched|None|
  |querystring|string|None|True|EWS QueryString to used for the search. Please refer to the EWS QueryString documentation available through Microsoft for assistance if needed|None|
  |flatten_attachments|boolean|False|False|Flattens attachments to a single list|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |match_count|integer|True|Number of matches found|
  |matches|[]email|True|Emails that match the EWS QueryString|
  Example output:
  {
    "match_count": 1,
    "matches": [{
        "account": "bob@test.com",
        "attached_emails": [],
        "attached_files": [],
        "body": "Test email",
        "flattened_attached_emails": [],
        "flattened_attached_files": [],
        "headers": [{
            "name": "Received",
            "value": ""
        }],
        "id": "AQMkAGRlNjhiMTkwLTJhMjUtNGI1ZS1hMGVlAC01MjUyZjgxZjM4Y2IARgAAAxV9OuuaYO1GuNrmyowDcygHAOrvwEx5H1ZNgslFeFvir+EAAAIBDAAAAOrvwEx5H1ZNgslFeFvir+EAAAIFaAAAAA==",
        "is_read": false,
        "sender": "joe@test.com",
        "subject": "Check this out!"
    }]
  }

  ### Move Email
  This action is used to move an email from one folder to another.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |email_id|string|None|True|ID of the email to move|None|
  |account|string|None|False|Account to move the email on. Emails can only be moved within the account specified here. The account configured in your Connection must have Impersonation access|None|
  |destination_folder|string|None|True|Common values are Calendar, Trash, Drafts, Inbox, Outbox, Sent, Junk, Tasks, Contacts. You can also use a completely custom value, for example python_mailing_list. Folder names are case-sensitive|None|
  |flatten_attachments|boolean|False|False|Flattens attachments to a single list|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |email|email|False|Email that was moved|
  Example output:
  {
    "email": {
        "account": "bob@test.com",
        "attached_emails": [],
        "attached_files": [],
        "body": "Test email",
        "flattened_attached_emails": [],
        "flattened_attached_files": [],
        "headers": [{
            "name": "Received",
            "value": ""
        }],
        "id": "AQMkAGRlNjhiMTkwLTJhMjUtNGI1ZS1hMGVlAC01MjUyZjgxZjM4Y2IARgAAAxV9OuuaYO1GuNrmyowDcygHAOrvwEx5H1ZNgslFeFvir+EAAAIBDAAAAOrvwEx5H1ZNgslFeFvir+EAAAIFaAAAAA==",
        "is_read": false,
        "sender": "joe@test.com",
        "subject": "Check this out!"
    }
  }
  ## Triggers
  ### Email Received
  This trigger is used to poll a mailbox for new emails.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |poll_interval|integer|15|True|How often to poll the specified mailbox for new emails, in seconds. Default value is 15 seconds|None|
  |subject_query|string||False|Query to search for in subject (regex capable). Only these emails will be marked as read|None|
  |mailbox_name|string|None|True|Common values are Calendar, Trash, Drafts, Inbox, Outbox, Sent, Junk, Tasks, Contacts. You can also use a completely custom value, for example python_mailing_list. Mailbox names are case-sensitive|None|
  |flatten_attachments|boolean|False|False|Flattens attachments to a single list|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |email|email|False|Email|
  Example output:
  {
    "email": {
        "account": "bob@test.com",
        "attached_emails": [],
        "attached_files": [],
        "body": "Test email",
        "flattened_attached_emails": [],
        "flattened_attached_files": [],
        "headers": [{
            "name": "Received",
            "value": ""
        }],
        "id": "AQMkAGRlNjhiMTkwLTJhMjUtNGI1ZS1hMGVlAC01MjUyZjgxZjM4Y2IARgAAAxV9OuuaYO1GuNrmyowDcygHAOrvwEx5H1ZNgslFeFvir+EAAAIBDAAAAOrvwEx5H1ZNgslFeFvir+EAAAIFaAAAAA==",
        "is_read": false,
        "sender": "joe@test.com",
        "subject": "Check this out!"
    }
  }
  ## Connection
  The connection configuration accepts the following parameters:
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |username|string|None|False|Usually in WINDOMAIN\username format, where WINDOMAIN is the name of the domain your username is connected to. Some servers accept usernames in PrimarySMTPAddress ('myusername@example.com') format. UPN format is also supported|None|
  |primary_smtp_address|string|None|False|Example\: john@example.com|None|
  |exchange_server_version|string|None|False|Exchange server version. If autodiscover is set to false then this option must be specified|['Exchange 2007', 'Exchange 2010', 'Exchange 2010 SP2', 'Exchange 2013', 'Exchange 2016']|
  |access_type|string|Delegate|True|Access type to use for login. See differences between Delegate and Impersonation here\: https\://blogs.msdn.microsoft.com/exchangedev/2009/06/15/exchange-impersonation-vs-delegate-access/|['Delegate', 'Impersonation']|
  |server|string|None|False|If autodiscover is set to false then this option must be specified. Example\: mail.example.com or exchange.example.com/EWS/Exchange.asmx|None|
  |autodiscover|boolean|True|True|Allow the plugin to auto-detect your exchange server settings. This can fail in certain Exchange server configurations|None|
  |verify_ssl|boolean|True|True|Whether or not to verify SSL connection. Try setting to false if you are having issues connecting|None|
  |use_ssl|boolean|True|False|Whether or not to use SSL|None|
  |password|password|None|False|Password for authenticating to the Exchange server|None|
  ## Troubleshooting
  * The ID for an email changes as actions are taken on the email. Because of this, actions that return an email must have the ID of that email used for subsequent actions.
  * Actions that have an input for account support Impersonation. This means that if the account configured in your plugin Connection has Impersonation rights then you will be able to execute those actions against other accounts within your Exchange environment.
  If you receive errors with this, double-check that your account has Impersonation rights and that they are over the specified accounts.
  * If you receive errors when trying to receive mail, double-check the mailbox name to verify that it is valid and using the proper casing.
  ## Workflows
  Examples
  * Email analysis
  ## Versions
  * 0.1.0 - Initial plugin
  * 0.1.2 - Add access type as input option in connection
  * 0.1.3 - New action for deleting a message
  * 1.0.0 - Overhaul of connection code and more configuration inputs to support more varied server configurations
  * 1.1.0 - Add action: Move Email
  * 1.1.1 - Add input to connection: Disable SSL verification
  * 1.2.0 - Add action: Set Email Categories. Add categories as a property on the email type
  * 1.2.1 - Fix connection test so it fails when invalid credentials are used
  * 1.2.2 - SSL bug fix in SDK
  * 1.3.0 - Add attachments of attachments. All attachments are base64 encoded. Added more error handling.
  * 2.0.0 - Deleting emails no longer iterates all email, Changekey required for action
  * 2.0.1 - Fix delete bug, Fix flatten argument always evaluating as true
  * 2.1.0 - Subject query search | move to v2 architecture
  * 2.1.1 - Bug fix for CI tool incorrectly uploading plugins
  * 3.0.0 - Overhaul message type (email headers, separated attachment types) | Support web server mode
  * 3.0.1 - Disable exponential backoff in connection attempts, add attached_emails and attached_files properties to email attachments
  * 4.0.0 - Add action: Search Inboxes | bugfix: potential for invalid ID/changekey when working with email | unify terminology | hardened code and more descriptive error logging | execute actions across impersonated accounts
  * 5.0.0 - Remove changekey input to allow for easier email plugin swapping in workflows
  ## References
  * [Microsoft Exchange](https://products.office.com/en-us/exchange/email)
  * [EWS QueryString Syntax](https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/querystring-querystringtype#remarks)
  * [Delegate vs. Impersonation Access](https://blogs.msdn.microsoft.com/exchangedev/2009/06/15/exchange-impersonation-vs-delegate-access/)
types:
  header:
    name:
      type: string
      required: true
    value:
      type: string
      required: true

  attachment_file:
    name:
      type: string
      required: false
    content:
      type: bytes
      required: false

  attachment_email:
    id:
      type: string
      title: "ID"
      required: false
    sender:
      type: string
      required: false
    subject:
      type: string
      required: false
    body:
      type: string
      required: false
    categories:
      type: "[]string"
      required: false
    date_received:
      type: date
      required: false
    headers:
      type: "[]header"
      required: true
    attached_files:
      type: "[]attachment_file"
      required: false
    attached_emails:
      type: "[]object"
      required: false

  email:
    id:
      type: string
      title: "ID"
      required: true
    account:
      type: string
      title: "Account"
      required: false
      description: "Primary SMTP Address of the account for which the email was found on"
    is_read:
      type: boolean
      title: "Is Read"
      required: false
      description: "Whether or not the email has been read"
    sender:
      type: string
      required: true
    subject:
      type: string
      required: false
    body:
      type: string
      required: false
    attached_files:
      type: "[]attachment_file"
      required: false
    attached_emails:
      type: "[]attachment_email"
      required: false
    flattened_attached_emails:
      type: "[]attachment_email"
      required: false
    flattened_attached_files:
      type: "[]attachment_file"
      required: false
    categories:
      type: "[]string"
      required: false
    date_received:
      type: date
      required: false
    headers:
      type: "[]header"
      required: true


connection:
  user_pass:
    type: credential_username_password
    title: "Username and password"
    description: "Username and password for the account. Username is usually in domain\\\\username format, where domain is the name of the domain your username is connected to. Some servers accept usernames in PrimarySMTPAddress ('joe@example.com') format. UPN format is also supported"
    required: true
    order: 1
  primary_smtp_address:
    type: string
    title: "Primary SMTP Address"
    description: "Example: john@example.com"
    order: 3
    required: false
  autodiscover:
    title: "Autodiscover"
    type: boolean
    description: "Allow the plugin to auto-detect your exchange server settings. This can fail in certain Exchange server configurations"
    default: true
    required: true
    order: 4
  access_type:
    title: "Access Type"
    type: string
    enum:
    - "Delegate"
    - "Impersonation"
    default: "Delegate"
    required: true
    description: "Access type to use for login. See differences between Delegate and Impersonation here: https://blogs.msdn.microsoft.com/exchangedev/2009/06/15/exchange-impersonation-vs-delegate-access/"
    order: 5
  use_ssl:
    title: "Use SSL"
    type: boolean
    description: "Whether or not to use SSL"
    default: true
    order: 6
    required: true
  verify_ssl:
    title: "Verify SSL"
    type: boolean
    description: "Whether or not to verify SSL connection. Try setting to false if you are having issues connecting"
    required: true
    default: true
  server:
    title: "Server"
    type: string
    description: "If autodiscover is set to false then this option must be specified. Example: mail.example.com or exchange.example.com/EWS/Exchange.asmx"
    order: 7
    required: false
  exchange_server_version:
    title: "Exchange Server Version"
    type: string
    description: "Exchange server version. If autodiscover is set to false then this option must be specified"
    default: ""
    enum:
    - ""
    - "Exchange 2007"
    - "Exchange 2010"
    - "Exchange 2010 SP2"
    - "Exchange 2013"
    - "Exchange 2016"
    order: 8
    required: false

triggers:
  email_received:
    title: "Email Received in Mailbox"
    description: "Poll mailbox for new emails"
    input:
      mailbox_name:
        title: "Mailbox Name"
        type: string
        description: "Common values are Calendar, Trash, Drafts, Inbox, Outbox, Sent, Junk, Tasks, Contacts. You can also use a completely custom value, for example python_mailing_list. Mailbox names are case-sensitive"
        required: true
      poll_interval:
        title: "Poll Interval"
        type: integer
        description: "How often to poll the specified mailbox for new emails, in seconds. Default value is 15 seconds"
        default: 15
        required: true
      flatten_attachments:
        title: "Flatten Attachments"
        type: boolean
        description: "Flattens attachments to a single list"
        default: false
        required: false
      subject_query:
        title: "Subject Query"
        type: string
        description: "Query to search for in subject (regex capable). Only these emails will be marked as read"
        default: ""
        required: false
    output:
      email:
        title: "Email"
        type: email
        required: false
        description: "Email"

actions:
  delete_email:
    title: "Delete Email"
    description: "Deletes a specified email"
    input:
      email_id:
        type: string
        title: "Email ID"
        description: "ID of the email to delete"
        required: true
      deletion_method:
        title: "Deletion Method"
        type: string
        description: "Deletes an email. Emails can be hard deleted, soft deleted (kept in recoverable folder), or moved to trash"
        enum:
        - "Hard"
        - "Soft"
        - "Trash"
        required: true
      flatten_attachments:
        title: "Flatten Attachments"
        type: boolean
        description: "Flattens attachments to a single list"
        default: false
        required: false
      account:
        title: "Account"
        type: string
        description: "Account to delete the email from. The account configured in your Connection must have Impersonation access"
        required: false
    output:
      email:
        title: "Email"
        type: email
        required: false
        description: "Email that has been deleted"

  move_email:
    title: "Move Email"
    description: "Moves an email from one folder to another"
    input:
      email_id:
        type: string
        title: "Email ID"
        description: "ID of the email to move"
        required: true
      destination_folder:
        type: string
        title: "Destination Folder"
        description: "Common values are Calendar, Trash, Drafts, Inbox, Outbox, Sent, Junk, Tasks, Contacts. You can also use a completely custom value, for example python_mailing_list. Folder names are case-sensitive"
        required: true
      flatten_attachments:
        title: "Flatten Attachments"
        type: boolean
        description: "Flattens attachments to a single list"
        default: false
        required: false
      account:
        title: "Account"
        type: string
        description: "Account to move the email on. Emails can only be moved within the account specified here. The account configured in your Connection must have Impersonation access"
        required: false
    output:
      email:
        title: "Email"
        type: email
        required: false
        description: "Email that was moved"

  set_email_categories:
    title: "Set Email Categories"
    description: "Sets the defined categories on an email"
    input:
      email_id:
        type: string
        title: "Email ID"
        description: "ID of the email to set categories on"
        required: true
      categories:
        type: "[]string"
        title: "Categories"
        description: "Categories to set on the email. Existing categories on the email will be overwritten"
        required: true
      account:
        type: string
        description: "Account to set the email category on. The account configured in your Connection must have Impersonation access"
        required: false
        title: "Account"
    output:
      success:
        title: "Success"
        type: boolean
        required: true
        description: "Whether or not the category set was successful"

  search_inboxes:
    title: "Search Inboxes"
    description: "Search account inboxes for emails using an EWS QueryString"
    input:
      accounts:
        type: "[]string"
        title: "Accounts"
        required: true
        description: "List of accounts to be searched"
      querystring:
        type: string
        title: "EWS QueryString"
        required: true
        description: "EWS QueryString to used for the search. Please refer to the EWS QueryString documentation available through Microsoft for assistance if needed"
      all_folders:
        type: boolean
        title: "Search All Folders"
        description: "Whether or not to search all Inbox subfolders. This will not search other folders at the root level of the mailbox and may increase processing time"
        required: true
        default: false
      flatten_attachments:
        title: "Flatten Attachments"
        type: boolean
        description: "Flattens attachments to a single list"
        default: false
        required: false
    output:
      match_count:
        type: integer
        title: "Match Count"
        description: "Number of matches found"
        required: true
      matches:
        type: "[]email"
        title: "Matches"
        description: "Emails that match the EWS QueryString"
        required: true
`
