package testspec

const SpecChatBot = `
plugin_spec_version: v2
name: chatbot_slack
title: "Chatbot for Slack"
description: "Chatbot integrations for Slack (build version: 1.0.0.7)"
version: 1.0.0
vendor: komand_chatbot
tags: ["slack", "chat", "chatbot"]
icon: "data:image/png;base64,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"
help: |
  ## About
  [Chatbot for Slack](https://slack.com/) is an internally used plugin with Komand. It will be downloaded automatically when creating Chatbots within Komand, so no need to manually install :)
  This plugin utilizes the [Slack API](https://api.slack.com/).
  ## Actions
  ### Search
  This action is used to search the message archive.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |count|integer|100|False|Max message count|None|
  |sort|string|score|False|Sort|None|
  |highlight|boolean|None|False|Highlight|None|
  |query|string|None|True|Query|None|
  |page|integer|1|False|Page|None|
  |sort_dir|string|desc|False|Sort Direction (asc or desc)|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |matches|[]slack_search_message|False|Matches|
  |count|int|False|Count|
  |total|int|False|Total|
  |page|int|False|Page|
  |pages|int|False|Pages|
  Example output:

  ### Post Message
  This action is used to post a message to the Slack channel.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |message|string|None|False|Message to send|None|
  |recipient|string|None|True|A channel (e.g. #channel) or username|None|
  |attachments|[]object|None|False|JSON array of attachments - see https\://api.slack.com/docs/message-attachments (Advanced)|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |channel_id|string|False|ChannelID of successful message|
  |timestamp|string|False|Timestamp of sucessful message|
  Example output:

  ### Set Channel Topic
  This action is used to set the topic of a channel.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |topic|string|None|True|Topic|None|
  |channel|string|None|False|Channel whose topic you want to set|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |topic|string|False|Topic set|
  Example output:

  ### Post Interactive Message
  This action is used to post a message and prompt for a response.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |message|string|None|False|Message to send|None|
  |recipient|string|None|True|A channel (e.g. #channel) or username|None|
  |attachments|bytes|None|False|JSON array of attachments - see https\://api.slack.com/docs/message-attachments (Advanced)|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |channel_id|string|False|ChannelID of successful message|
  |timestamp|string|False|Timestamp of sucessful message|
  Example output:

  ### Bulk Lookup User
  This action is used to bulk lookup user by ID.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |user_id|[]string|None|True|User ID|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |user_info|[]bulk_user|False|User's information|
  Example output:
  {
    [
      {
        "display_name": "ross",
        "email": "ross@komand.com",
        "first_name": "Ross",
        "last_name": "Nanopoulos",
        "phone": "",
        "real_name": "Ross Nanopoulos",
        "status_text": "",
        "team": "",
        "title": ""
      }
    ]
  }
  ### Lookup User
  This action is used to lookup a user by ID.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |user_id|string|None|True|User ID|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |user_info|user|False|User's information|
  Example output:
  {
    "display_name": "ross",
    "email": "ross@komand.com",
    "first_name": "Ross",
    "last_name": "Nanopoulos",
    "phone": "",
    "real_name": "Ross Nanopoulos",
    "status_text": "",
    "team": "",
    "title": ""
  }
  ### Get Groups
  This action is used to get all groups.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |exclude_archived|boolean|True|True|Exclude archived groups from results|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |channels|[]group|False|Array of groups|
  Example output:
  ### Upload File
  This action is used to upload a file to Slack.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |content|bytes|None|False|File content (base64 encoded)|None|
  |channels|string|None|False|Channels which the file should be shared to (comma-separated)|None|
  |title|string|None|False|Title of Attachment|None|
  |filename|string|None|False|File name|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |file|slack_file|False|File|
  Example output:

  ### Upload Snippet
  This action is used to upload a snippet to Slack.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |content|string|None|True|Snippet text|None|
  |channels|string|None|False|Channels which the file should be shared to (comma-separated)|None|
  |filetype|string|None|False|File type|None|
  |title|string|None|False|Title of Snippet|None|
  |filename|string|None|False|File name|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |file|slack_file|False|File|
  Example output:

  ### Get Channels
  This action is used to get all channels.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |exclude_archived|boolean|True|True|Exclude archived channels from results|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |channels|[]channel|False|Array of channels|
  Example output:

  ## Triggers
  ### Message With File
  This trigger is used to trigger on a message with a file attachment.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |match_text|string|None|False|Regexp match on message text|None|
  |match_filename|string|None|False|Regexp match on a specific filename|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |timestamp|string|False|Timestamp|
  |message|message|False|Message|
  |type|string|False|Message type|
  |file|file|False|File|
  Example output:

  ### Message
  This trigger is used to trigger on message.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |match_text|string|None|False|Regex match on message content e.g. hi or .*hi.* matches john says hi in the message text|None|
  |match_channel|string|None|False|Regex match on channel. Otherwise listen on any channel|None|
  |type|string|any|True|Trigger on direct messages, group chats, or any|['any', 'direct', 'group']|
  |match_user|string|None|False|Regex match on username sending the message. Otherwise listen for any private message|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |timestamp|string|False|Timestamp|
  |message|message|False|Message|
  |type|string|False|Message Type|
  Example output:

  ### Event
  This trigger is used to monitor for events.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |subtype|string|None|True|Event subtype to monitor|['bot_message', 'channel_archive', 'channel_join', 'channel_leave', 'channel_name', 'channel_purpose', 'channel_topic', 'channel_unarchive', 'file_comment', 'file_mention', 'file_share', 'group_archive', 'group_join', 'group_leave', 'group_name', 'group_purpose', 'group_topic', 'group_unarchive', 'me_message', 'message_changed', 'message_deleted', 'pinned_item', 'unpinned_item']|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |timestamp|string|False|Timestamp|
  |event|message|False|Event|
  Example output:

  ## Connection
  This plugin requires a Slack API token to authenticate.
  The connection configuration accepts the following parameters:
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |bot_user_oAuth_access_token|string|None|False|Bot User OAuth Access Token|None|
  |verification_token|string|None|False|Slack App Verification Token|None|
  *Token* is the Bot User OAuth Access Token which can be found on the Install App tab of the Slack app management page.
  *Verification Token* is the Verification Token found on the Basic Information tab of the Slack app management page. For interactive messages, this token is used to verify that requests are actually coming from Slack.
  ## versions
  * 1.0.0 - Initial plugin
  ## Troubleshooting
  This plugin does not contain any troubleshooting information.
  ## Workflows
  Examples:
  * Notification
  ## References
  * [Slack](https://slack.com/)
  * [Slack API](https://api.slack.com/)
  * [Chatops Example: Investigating URLs](https://komand.zendesk.com/hc/en-us/articles/115000954607-Use-Case-Chatops-Example-Investigating-URLs)
connection:
  bot_user_oAuth_access_token:
    type: password
    description: "Bot User OAuth Access Token"
    required: true
  verification_token:
    type: password
    description: "Slack App Verification Token"
    required: true

types:
  ctx_channel:
    id:
      title: "ID"
      type: string
    name:
      type: string
  ctx_message:
    user:
      type: string
    username:
      type: string
    text:
      type: string
    timestamp:
      type: string
    type:
      type: string
  slack_search_message:
    type:
      type: string
    channel:
      type: ctx_channel
    user:
      type: string
    username:
      type: string
    text:
      type: string
    permalink:
      type: string
    previous:
      type: ctx_message
    previous2:
      type: ctx_message
    next:
      type: ctx_message
    next2:
      type: ctx_message
  attachment_field:
    title:
      type: string
    value:
      type: string
    short:
      type: boolean
  attachment_action:
    name:
      type: string
    text:
      type: string
    type:
      type: string
    value:
      type: string
    style:
      type: string
  attachment:
    color:
      type: string
    fallback:
      type: string
    author_name:
      type: string
    author_subname:
      type: string
    author_icon:
      type: string
    author_link:
      type: string
    title:
      type: string
    title_link:
      type: string
    pretext:
      type: string
    text:
      type: string
    image_url:
      type: string
    thumb_url:
      type: string
    fields:
      type: "[]attachment_field"
    markdown_in:
      type: "[]string"
    callback_id:
      title: "Callback ID"
      type: string
    attachment_type:
      type: string
    actions:
      type: "[]attachment_action"
  item_reaction:
    name:
      type: string
    count:
      type: integer
    users:
      type: "[]string"
  comment:
    id:
      title: "ID"
      type: string
    created:
      type: integer
    timestamp:
      type: integer
    user:
      type: string
    comment:
      type: string
  icon:
    iconurl:
      type: string
    iconemoji:
      type: string
  edited:
    user:
      type: string
    timestamp:
      type: string
  sdk_file:
    filename:
      type: string
    content:
      type: bytes
  slack_file:
    id:
      title: "ID"
      type: string
    created:
      type: int
    timestamp:
      type: int
    name:
      type: string
    title:
      type: string
    minetype:
      type: string
    image_exif_rotation:
      type: int
    filetype:
      type: string
    pretty_type:
      type: string
    user:
      type: string
    mode:
      type: string
    editable:
      type: boolean
    is_external:
      type: boolean
    external_type:
      type: string
    size:
      type: int
    url:
      type: string
    url_download:
      type: string
    url_private:
      type: string
    url_private_download:
      type: string
    permalink:
      type: string
    permalink_public:
      type: string
    edit_link:
      type: string
    preview:
      type: string
    preview_highlight:
      type: string
    lines:
      type: int
    lines_more:
      type: int
    is_public:
      type: boolean
    public_url_shared:
      type: boolean
    channels:
      type: "[]string"
    groups:
      type: "[]string"
    ims:
      type: "[]string"
    initial_comment:
      type: comment
    comments_count:
      type: int
    num_stars:
      type: int
    is_starred:
      type: boolean
  message:
    user:
      type: "string"
      description: "User name"
    user_id:
      title: "User ID"
      type: "string"
      description: "Slack User ID"
    channel:
      type: "string"
      description: "Channel name"
    channel_id:
      title: "Channel ID"
      type: "string"
      description: "Slack channel ID"
    text:
      type: "string"
      description: "Text"
    ts:
      type: "string"
      description: "Timestamp"
    attachments:
      type: "[]attachment"
      description: "Message attachment"
  channel:
    channel_id:
      title: "Channel ID"
      type: string
    channel_name:
      type: string
      description: "Channel name"
    is_channel:
      type: boolean
      description: "Is channel"
    created:
      type: integer
      description: "Time of channel creation"
    creator:
      type: string
      description: "User ID of channel creator"
    is_archived:
      type: boolean
      description: "Is channel archived"
    is_general:
      type: boolean
      description: "Is general"
    name_normalized:
      type: string
      description: "Name normalized"
    is_shared:
      type: boolean
      description: "Is general"
    is_org_shared:
      type: boolean
      description: "Is org shared"
    is_member:
      type: boolean
      description: "Is member of channel"
    is_private:
      type: boolean
      description: "Is private channel"
    is_mpim:
      type: boolean
      description: "Is mpim"
    members:
      type: "[]string"
      description: "Array of ID's of users in channel"
    topic:
      type: object
      description: "Topic object"
    purpose:
      type: object
      description: "Purpose object"
    previous_names:
      type: "[]string"
      description: "Array of previous channel names"
    num_members:
      description: "Number of members in channel"
      type: integer
  group:
    channel_id:
      title: "Channel ID"
      type: string
    channel_name:
      type: string
      description: "Channel name"
    created:
      type: integer
      description: "Time of channel creation"
    creator:
      type: string
      description: "User ID of channel creator"
    is_archived:
      type: boolean
      description: "Is channel archived"
    members:
      type: "[]string"
      description: "Array of ID's of users in channel"
    topic:
      type: object
      description: "Topic object"
    purpose:
      type: object
      description: "Purpose object"
  user:
    first_name:
      type: string
    last_name:
        type: string
    real_name:
      type: string
    display_name:
      type: string
    email:
      type: string
    phone:
      type: string
    title:
      type: string
    status_text:
      type: string
    team:
      type: string
  bulk_user:
    first_name:
      type: string
    last_name:
        type: string
    real_name:
      type: string
    display_name:
      type: string
    email:
      type: string
    phone:
      type: string
    title:
      type: string
    status_text:
      type: string
    team:
      type: string
    userID:
      title: "User ID"
      type: string


triggers:
  message:
    name: "New Message"
    description: "Trigger on message"
    input:
      match_user:
        type: string
        title: "Match User"
        description: "Regex match on username sending the message. Otherwise listen for any private message"
        required: false
      match_channel:
        type: string
        title: "Match Channel"
        description: "Regex match on channel. Otherwise listen on any channel"
        required: false
      match_text:
        type: string
        title: "Match Text"
        description: "Regex match on message content e.g. hi or .*hi.* matches john says hi in the message text"
        required: false
      type:
        type: string
        title: "Type"
        description: "Trigger on direct messages, group chats, or any"
        enum:
          - "any"
          - "direct"
          - "group"
        default: "any"
        required: true
    output:
      message:
        type: message
        description: "Message"
        title: "Message"
        required: false
      type:
        type: string
        description: "Message Type"
        title: "Type"
        required: false
      timestamp:
        type: string
        description: "Timestamp"
        required: false
        title: "Timestamp"
  event:
    name: "Slack Event"
    description: "Trigger on event"
    input:
      subtype:
        type: string
        title: "Subtype"
        description: "Event subtype to monitor"
        enum:
          - "bot_message"
          - "channel_archive"
          - "channel_join"
          - "channel_leave"
          - "channel_name"
          - "channel_purpose"
          - "channel_topic"
          - "channel_unarchive"
          - "file_comment"
          - "file_mention"
          - "file_share"
          - "group_archive"
          - "group_join"
          - "group_leave"
          - "group_name"
          - "group_purpose"
          - "group_topic"
          - "group_unarchive"
          - "me_message"
          - "message_changed"
          - "message.im"
          - "message_deleted"
          - "pinned_item"
          - "unpinned_item"
        required: true
    output:
      event:
        type: message
        description: "Event"
        title: "Event"
        required: false
      timestamp:
        type: string
        description: "Timestamp"
        title: "Timestamp"
        required: false
  message_with_file:
    name: "New Message With File"
    description: "Trigger on a message with a file attachment"
    input:
      match_text:
        type: string
        title: "Match Text"
        description: "Regexp match on message text"
        required: false
      match_filename:
        title: "Match Filename"
        type: string
        description: "Regexp match on a specific filename"
        required: false
    output:
      message:
        type: message
        description: "Message"
        title: "Message"
        required: false
      type:
        type: string
        description: "Message type"
        title: "Type"
        required: false
      timestamp:
        description: "Timestamp"
        required: false
        title: "Timestamp"
        type: string
      file:
        type: file
        description: "File"
        title: "File"
        required: false

actions:
  upload_snippet:
    description: "Upload a snippet to Slack"
    input:
      title:
        type: "string"
        description: "Title of Snippet"
        required: false
        title: "Title"
      filetype:
        type: "string"
        description: "File type"
        required: false
        title: "File Type"
      filename:
        type: "string"
        description: "File name"
        required: false
        title: "File Name"
      content:
        type: "string"
        description: "Snippet text"
        required: true
        title: "Content"
      channels:
        type: "string"
        description: "Channels which the file should be shared to (comma-separated)"
        required: false
        title: "Channels"
    output:
      file:
        type: slack_file
        required: false
        title: "File"
        description: "File"

  upload_file:
    description: "Upload a file to Slack"
    input:
      title:
        type: "string"
        description: "Title of Attachment"
        required: false
        title: "Title"
      filename:
        type: "string"
        description: "File name"
        title: "File Name"
        required: false
      content:
        type: "bytes"
        description: "File content (base64 encoded)"
        title: "Content"
        required: false
      channels:
        type: "string"
        description: "Channels which the file should be shared to (comma-separated)"
        required: false
        title: "Channels"
    output:
      file:
        type: slack_file
        required: false
        title: "File"
        description: "File"

  post_message:
    description: "Post a message to the Slack channel"
    input:
      recipient:
        type: string
        description: "A channel (e.g. #channel) or username"
        title: "Select a Channel or Username"
        required: true
      message:
        type: string
        description: "Message to send"
        required: false
        title: "Message"
      attachments:
        type: "[]object"
        description: "JSON array of attachments - see https://api.slack.com/docs/message-attachments (Advanced)"
        required: false
        title: "Attachments"
    output:
      channel_id:
        title: "Channel ID"
        type: string
        description: "ChannelID of successful message"
        required: false
      timestamp:
        type: string
        description: "Timestamp of sucessful message"
        required: false
        title: "Timestamp"

  post_interactive_message:
    title: "Response Prompt"
    description: "Post a message and prompt for a response"
    input:
      message:
        type: string
        description: "Message to send"
        title: "Message"
        required: false
      recipient:
        type: string
        description: "A channel (e.g. #channel) or username"
        required: true
      attachments:
        type: bytes
        description: "JSON array of attachments - see https://api.slack.com/docs/message-attachments (Advanced)"
        required: false
        title: "Attachments"
    output:
      channel_id:
        title: "Channel ID"
        type: string
        description: "ChannelID of successful message"
        required: false
      timestamp:
        type: string
        description: "Timestamp of sucessful message"
        required: false

  search:
    title: "Search messages"
    description: Search message archive
    input:
      count:
        default: 100
        required: false
        type: integer
        description: "Max message count"
        title: "Count"
      highlight:
        required: false
        type: boolean
        description: "Highlight"
        title: "Highlight"
      page:
        default: 1
        required: false
        type: integer
        description: "Page"
        title: "Page"
      query:
        required: true
        type: string
        description: "Query"
        title: "Query"
      sort:
        default: score
        required: false
        type: string
        description: "Sort"
        title: "Sort"
      sort_dir:
        default: desc
        required: false
        type: string
        description: "Sort Direction (asc or desc)"
        title: "Sort Direction"
    output:
      matches:
        description: "Matches"
        type: "[]slack_search_message"
        required: false
        title: "Matches"
      count:
        description: "Count"
        type: int
        required: false
        title: "Count"
      total:
        description: "Total"
        type: int
        required: false
        title: "Total"
      page:
        description: "Page"
        type: int
        required: false
        title: "Page"
      pages:
        description: "Pages"
        type: int
        required: false
        title: "Pages"

  set_channel_topic:
    title: "Set Channel Topic"
    description: "Sets the topic of a channel"
    input:
      channel:
        type: string
        description: "Channel whose topic you want to set"
        title: "Channel"
        required: false
      topic:
        type: string
        description: "Topic"
        required: true
        title: "Topic"
    output:
      topic:
        type: string
        description: "Topic set"
        required: false

  get_channels:
    title: "Get Channels"
    description: "Get all channels"
    input:
      exclude_archived:
        title: "Exclude Archived"
        type: boolean
        required: true
        default: true
        description: "Exclude archived channels from results"
    output:
      channels:
        type: "[]channel"
        required: false
        description: "Array of channels"

  get_groups:
    title: "Get Groups"
    description: "Get all groups"
    input:
      exclude_archived:
        title: "Exclude Archived"
        type: boolean
        required: true
        default: true
        description: "Exclude archived groups from results"
    output:
      channels:
        type: "[]group"
        required: false
        description: "Array of groups"

  lookup_user:
    title: "User Lookup"
    description: "Lookup user by ID"
    input:
      user_id:
        title: "User ID"
        type: string
        required: true
        description: "User ID"
    output:
      user_info:
        title: "User Info"
        type: user
        required: false
        description: "User's information"

  bulk_lookup_user:
    title: "Bulk User Lookup"
    description: "Bulk lookup user by ID"
    input:
      user_id:
        title: "User ID"
        type: "[]string"
        required: true
        description: "User ID"
    output:
      user_info:
        title: "User Info"
        type: "[]bulk_user"
        required: false
        description: "User's information"

  add_reaction:
    title: "Add Reaction"
    description: "Add reaction to slack post"
    input:
      reaction:
        type: string
        description: "Reaction to use"
        title: "React"
        required: true
        default: "smile"
      channelID: 
        type: string
        title: "Channel ID"
        description: "Channel ID"
        required: true
      timestamp:
        type: string
        title: "Timestamp"
        description: "Timestamp"
        required: true
    output:
      success:
        type: boolean
        description: "Success or failue"
        required: false
`
