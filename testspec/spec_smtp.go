package testspec

const SpecREST = `
plugin_spec_version: v2
name: rest
title: "REST"
description: "REST plugin to make it easy to integrate with RESTful services"
version: 3.0.0
vendor: rapid7
tags: ["rest", "http", "rpc", "microservices", "api"]
icon: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAfQAAAH0CAYAAADL1t+KAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAA4hpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTMyIDc5LjE1OTI4NCwgMjAxNi8wNC8xOS0xMzoxMzo0MCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtcE1NOk9yaWdpbmFsRG9jdW1lbnRJRD0ieG1wLmRpZDowNWZiOGExZS04MWVhLTRmOTItYWMyOS01M2JiMWViM2E5NjciIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6N0JEQUI2RjBFMUFBMTFFNjgyMjY4QUEyN0JGQzIwRTQiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6N0JEQUI2RUZFMUFBMTFFNjgyMjY4QUEyN0JGQzIwRTQiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTUuNSAoTWFjaW50b3NoKSI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAuaWlkOmEwOGJlODk0LTEzZWQtNGMxOC1hNjRlLWQ2MzY2YWYwZGQ5NSIgc3RSZWY6ZG9jdW1lbnRJRD0iYWRvYmU6ZG9jaWQ6cGhvdG9zaG9wOjVhMzZlZDFlLWI1OTgtMTE3OS04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"
enable_cache: true
help: |
  # REST
  ## About
  [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) is a means of communicating over HTTP using verbs and actions to represent operations performed over resources, held in remote locations. It's based off of a document put together by Roy Fielding.
  All actions take 2 pieces of configuration: a route, and an optional series of headers. Any headers set this way will overwrite the default ones in the connection.
  The route can include a query string and arguments as needed, or just be the remaining path, which will be appended to the base URL specified in the connection.
  Additionally, every action except GET will allow you specify a body as part of the request.
  All of the actions return the same set of data:
  * The response, parsed from JSON as an object if possible
  * The response, as a raw string of text
  * The status code
  * The response headers
  ## Actions
  ### PUT
  This action is used to make a PUT request.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |body|object|None|False|Payload to submit to the server when making the REST call|None|
  |headers|object|None|False|Headers to use for the request. These will override any default headers|None|
  |route|string|None|True|The route to append to the base URL e.g. /org/users|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |status|int|False|None|
  |body_object|object|False|None|
  |body_string|string|False|None|
  |headers|object|False|None|
  Example output:
  {
    "body_object": {
      "args": {},
      "data": "{\"hello\": \"world\"}",
      "files": {},
      "form": {},
      "headers": {
        "Accept": "*/*",
        "Accept-Encoding": "gzip, deflate",
        "Connection": "close",
        "Content-Length": "18",
        "Content-Type": "application/json",
        "Host": "httpbin.org",
        "Referer": "https://google.com",
        "User-Agent": "Mozilla/5.0"
      },
      "json": {
        "hello": "world"
      },
      "origin": "73.51.89.6",
      "url": "https://httpbin.org/put"
    },
    "body_string": "{\n  \"args\": {}, \n  \"data\": \"{\\\"hello\\\": \\\"world\\\"}\", \n  \"files\": {}, \n  \"form\": {}, \n  \"headers\": {\n    \"Accept\": \"*/*\", \n    \"Accept-Encoding\": \"gzip, deflate\", \n    \"Connection\": \"close\", \n    \"Content-Length\": \"18\", \n    \"Content-Type\": \"application/json\", \n    \"Host\": \"httpbin.org\", \n    \"Referer\": \"https://google.com\", \n    \"User-Agent\": \"Mozilla/5.0\"\n  }, \n  \"json\": {\n    \"hello\": \"world\"\n  }, \n  \"origin\": \"73.51.89.6\", \n  \"url\": \"https://httpbin.org/put\"\n}\n",
    "status": 200,
    "headers": {
      "Connection": "keep-alive",
      "Server": "gunicorn/19.9.0",
      "Date": "Thu, 13 Sep 2018 15:21:45 GMT",
      "Content-Type": "application/json",
      "Content-Length": "468",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Credentials": "true",
      "Via": "1.1 vegur"
    }
  }
  ### POST
  This action is used to make a POST request.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |body|object|None|False|Payload to submit to the server when making the REST call|None|
  |headers|object|None|False|Headers to use for the request. These will override any default headers|None|
  |route|string|None|True|The route to append to the base URL e.g. /org/users|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |status|int|False|None|
  |body_object|object|False|None|
  |body_string|string|False|None|
  |headers|object|False|None|
  Example output:
  {
    "body_object": {
      "args": {},
      "data": "{\"Best\": \"Komand\"}",
      "files": {},
      "form": {},
      "headers": {
        "Accept": "*/*",
        "Accept-Encoding": "gzip, deflate",
        "Connection": "close",
        "Content-Length": "18",
        "Content-Type": "application/json",
        "Host": "httpbin.org",
        "Referer": "https://google.com",
        "User-Agent": "Mozilla/5.0"
      },
      "json": {
        "Best": "Komand"
      },
      "origin": "73.51.89.6",
      "url": "https://httpbin.org/post"
    },
    "body_string": "{\n  \"args\": {}, \n  \"data\": \"{\\\"Best\\\": \\\"Komand\\\"}\", \n  \"files\": {}, \n  \"form\": {}, \n  \"headers\": {\n    \"Accept\": \"*/*\", \n    \"Accept-Encoding\": \"gzip, deflate\", \n    \"Connection\": \"close\", \n    \"Content-Length\": \"18\", \n    \"Content-Type\": \"application/json\", \n    \"Host\": \"httpbin.org\", \n    \"Referer\": \"https://google.com\", \n    \"User-Agent\": \"Mozilla/5.0\"\n  }, \n  \"json\": {\n    \"Best\": \"Komand\"\n  }, \n  \"origin\": \"73.51.89.6\", \n  \"url\": \"https://httpbin.org/post\"\n}\n",
    "status": 200,
    "headers": {
      "Connection": "keep-alive",
      "Server": "gunicorn/19.9.0",
      "Date": "Thu, 13 Sep 2018 15:20:59 GMT",
      "Content-Type": "application/json",
      "Content-Length": "469",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Credentials": "true",
      "Via": "1.1 vegur"
    }
  }
  ### PATCH
  This action is used to make a PATCH request.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |body|object|None|False|Payload to submit to the server when making the REST call|None|
  |headers|object|None|False|Headers to use for the request. These will override any default headers|None|
  |route|string|None|True|The route to append to the base URL e.g. /org/users|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |status|int|False|None|
  |body_object|object|False|None|
  |body_string|string|False|None|
  |headers|object|False|None|
  Example output:
  {
    "body_object": {
      "args": {},
      "data": "",
      "files": {},
      "form": {},
      "headers": {
        "Accept": "*/*",
        "Accept-Encoding": "gzip, deflate",
        "Connection": "close",
        "Content-Length": "0",
        "Host": "httpbin.org",
        "Referer": "https://google.com",
        "User-Agent": "Mozilla/5.0"
      },
      "json": null,
      "origin": "73.51.89.6",
      "url": "https://httpbin.org/patch"
    },
    "body_string": "{\n  \"args\": {}, \n  \"data\": \"\", \n  \"files\": {}, \n  \"form\": {}, \n  \"headers\": {\n    \"Accept\": \"*/*\", \n    \"Accept-Encoding\": \"gzip, deflate\", \n    \"Connection\": \"close\", \n    \"Content-Length\": \"0\", \n    \"Host\": \"httpbin.org\", \n    \"Referer\": \"https://google.com\", \n    \"User-Agent\": \"Mozilla/5.0\"\n  }, \n  \"json\": null, \n  \"origin\": \"73.51.89.6\", \n  \"url\": \"https://httpbin.org/patch\"\n}\n",
    "status": 200,
    "headers": {
      "Connection": "keep-alive",
      "Server": "gunicorn/19.9.0",
      "Date": "Thu, 13 Sep 2018 15:20:25 GMT",
      "Content-Type": "application/json",
      "Content-Length": "384",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Credentials": "true",
      "Via": "1.1 vegur"
    }
  }
  ### GET
  This action is used to make a GET request.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |headers|object|None|False|Headers to use for the request. These will override any default headers|None|
  |route|string|None|True|The route to append to the base URL e.g. /org/users|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |status|int|False|None|
  |body_object|object|False|None|
  |body_string|string|False|None|
  |headers|object|False|None|
  Example output:
  {
    "body_object": {
      "origin": "73.51.89.6"
    },
    "body_string": "{\n  \"origin\": \"73.51.89.6\"\n}\n",
    "status": 200,
    "headers": {
      "Connection": "keep-alive",
      "Server": "gunicorn/19.9.0",
      "Date": "Thu, 13 Sep 2018 15:19:45 GMT",
      "Content-Type": "application/json",
      "Content-Length": "29",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Credentials": "true",
      "Via": "1.1 vegur"
    }
  }
  ### DELETE
  This action is used to make a DELETE request.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |body|object|None|False|Payload to submit to the server when making the REST call|None|
  |headers|object|None|False|Headers to use for the request. These will override any default headers|None|
  |route|string|None|True|The route to append to the base URL e.g. /org/users|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |status|int|False|None|
  |body_object|object|False|None|
  |body_string|string|False|None|
  |headers|object|False|None|
  Example output:
  {
    "body_object": {
      "args": {},
      "data": "",
      "files": {},
      "form": {},
      "headers": {
        "Accept": "*/*",
        "Accept-Encoding": "gzip, deflate",
        "Connection": "close",
        "Content-Length": "0",
        "Host": "httpbin.org",
        "Referer": "https://google.com",
        "User-Agent": "Mozilla/5.0"
      },
      "json": null,
      "origin": "73.51.89.6",
      "url": "https://httpbin.org/delete"
    },
    "body_string": "{\n  \"args\": {}, \n  \"data\": \"\", \n  \"files\": {}, \n  \"form\": {}, \n  \"headers\": {\n    \"Accept\": \"*/*\", \n    \"Accept-Encoding\": \"gzip, deflate\", \n    \"Connection\": \"close\", \n    \"Content-Length\": \"0\", \n    \"Host\": \"httpbin.org\", \n    \"Referer\": \"https://google.com\", \n    \"User-Agent\": \"Mozilla/5.0\"\n  }, \n  \"json\": null, \n  \"origin\": \"73.51.89.6\", \n  \"url\": \"https://httpbin.org/delete\"\n}\n",
    "status": 200,
    "headers": {
      "Connection": "keep-alive",
      "Server": "gunicorn/19.9.0",
      "Date": "Thu, 13 Sep 2018 15:14:55 GMT",
      "Content-Type": "application/json",
      "Content-Length": "385",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Credentials": "true",
      "Via": "1.1 vegur"
    }
  }
  ## Triggers
  This plugin does not contain any triggers.
  ## Connection
  Configuring a REST connection requires an endpoint to hit, which includes the protocol (http:// or https:// should be explicitly set by the user).
  Additionally, you can set a series of default Headers for every request to use. This would be useful if every request needed to send an auth token along inside of a header, so you don't need to specify it in each individual action.
  The connection configuration accepts the following parameters:
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |default_headers|object|None|False|None|None|
  |base_url|string|None|False|None|None|
  |ssl_verify|boolean|True|False|None|None|
  |basic_auth_credentials|credential_username_password|None|False|None|None|
  ## Troubleshooting
  Any headers set in the action will overwrite the default ones in the connection.
  Any issues connecting to the remote service should be present in the log of the job that ran. If you find any issues that represent bugs in the plugin itself, please contact someone at Komand directly.
  ## Workflows
  Examples
  * Create a request to interact with an in house API
  * Create a request to pull phishing web page for further enrichment
  ## Versions
  * 0.1.0 - Initial plugin
  * 0.1.1 - Update tags
  * 0.1.2 - SSL bug fix in SDK
  * 0.1.3 - Fix post and put actions by using json argument instead of body
  * 0.1.4 - Bug fix for CI tool incorrectly uploading plugins
  * 1.0.0 - Update to v2 Python plugin architecture | Support web server mode
  * 2.0.0 - Update connection to handle SSL verification
  * 3.0.0 - Add basic auth support
  ## References
  * [REST Architecture Style](http://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm)
connection:
  base_url:
    title: "Base URL"
    description: "Base URL e.g. https://httpbin.org"
    type: string
    required: true
  default_headers:
    title: "Default Headers"
    description: "Default headers to include in all requests associated with this connection e.g. { User-Agent: Komand }"
    type: object
    required: false
  ssl_verify:
    title: "SSL Verify"
    type: boolean
    description: "Verify SSL certificate"
    required: true
    default: true
  basic_auth_credentials:
    title: "Basic Auth Credentials"
    type: credential_username_password
    required: false

actions:
  get:
    title: "GET"
    description: "Make a GET request"
    input:
      route:
        title: "Route"
        description: "The route to append to the base URL e.g. /org/users"
        type: string
        required: true
      headers:
        title: "Headers"
        description: "Headers to use for the request. These will override any default headers"
        type: object
        required: false
    output:
      body_object:
        title: "Body Object"
        description: "Response payload from the server as an object"
        type: object
        required: false
      body_string:
        title: "Body String"
        description: "Response payload from the server as a string"
        type: string
        required: false
      status:
        title: "Status"
        description: "Status code of the response from the server"
        type: int
        required: false
      headers:
        title: "Headers"
        description: "Response headers from the server"
        type: object
        required: false
  post:
    title: "POST"
    description: "Make a POST request"
    input:
      route:
        title: "Route"
        description: "The route to append to the base URL e.g. /org/users"
        type: string
        required: true
      headers:
        title: "Headers"
        description: "Headers to use for the request. These will override any default headers"
        type: object
        required: false
      body:
        title: "Body"
        description: "Payload to submit to the server when making the REST call"
        type: object
        required: false
    output:
      body_object:
        title: "Body Object"
        description: "Response payload from the server as an object"
        type: object
        required: false
      body_string:
        title: "Body String"
        description: "Response payload from the server as a string"
        type: string
        required: false
      status:
        title: "Status"
        description: "Status code of the response from the server"
        type: int
        required: false
      headers:
        title: "Headers"
        description: "Response headers from the server"
        type: object
        required: false
  put:
    title: "PUT"
    description: "Make a PUT request"
    input:
      route:
        title: "Route"
        description: "The route to append to the base URL e.g. /org/users"
        type: string
        required: true
      headers:
        title: "Headers"
        description: "Headers to use for the request. These will override any default headers"
        type: object
        required: false
      body:
        title: "Body"
        description: "Payload to submit to the server when making the REST call"
        type: object
        required: false
    output:
      body_object:
        title: "Body Object"
        description: "Response payload from the server as an object"
        type: object
        required: false
      body_string:
        title: "Body String"
        description: "Response payload from the server as a string"
        type: string
        required: false
      status:
        title: "Status"
        description: "Status code of the response from the server"
        type: int
        required: false
      headers:
        title: "Headers"
        description: "Response headers from the server"
        type: object
        required: false
  patch:
    title: "PATCH"
    description: "Make a PATCH request"
    input:
      route:
        title: "Route"
        description: "The route to append to the base URL e.g. /org/users"
        type: string
        required: true
      headers:
        title: "Headers"
        description: "Headers to use for the request. These will override any default headers"
        type: object
        required: false
      body:
        title: "Body"
        description: "Payload to submit to the server when making the REST call"
        type: object
        required: false
    output:
      body_object:
        title: "Body Object"
        description: "Response payload from the server as an object"
        type: object
        required: false
      body_string:
        title: "Body String"
        description: "Response payload from the server as a string"
        type: string
        required: false
      status:
        title: "Status"
        description: "Status code of the response from the server"
        type: int
        required: false
      headers:
        title: "Headers"
        description: "Response headers from the server"
        type: object
        required: false
  delete:
    title: "DELETE"
    description: "Make a DELETE request"
    input:
      route:
        title: "Route"
        description: "The route to append to the base URL e.g. /org/users"
        type: string
        required: true
      headers:
        title: "Headers"
        description: "Headers to use for the request. These will override any default headers"
        type: object
        required: false
      body:
        title: "Body"
        description: "Payload to submit to the server when making the REST call"
        type: object
        required: false
    output:
      body_object:
        title: "Body Object"
        description: "Response payload from the server as an object"
        type: object
        required: false
      body_string:
        title: "Body String"
        description: "Response payload from the server as a string"
        type: string
        required: false
      status:
        title: "Status"
        description: "Status code of the response from the server"
        type: int
        required: false
      headers:
        title: "Headers"
        description: "Response headers from the server"
        type: object
        required: false
`
