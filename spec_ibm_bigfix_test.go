package sdk

const SpecBigFix = `
plugin_spec_version: v2
name: ibm_bigfix
title: "IBM BigFix"
description: "IBM BigFix is a systems-management software product developed by IBM for managing large groups of computers"
version: 2.1.2
vendor: rapid7
tags: [ "ibm", "TEM", "endpoint", "patch" ]
icon: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAfQAAAH0CAYAAADL1t+KAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAA4hpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTM4IDc5LjE1OTgyNCwgMjAxNi8wOS8xNC0wMTowOTowMSAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtcE1NOk9yaWdpbmFsRG9jdW1lbnRJRD0ieG1wLmRpZDowNWZiOGExZS04MWVhLTRmOTItYWMyOS01M2JiMWViM2E5NjciIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6QTA0ODlGN0E5NTkxMTFFNzgyQzBERjIxMEJFODNEQUIiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6QTA0ODlGNzk5NTkxMTFFNzgyQzBERjIxMEJFODNEQUIiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTUuNSAoTWFjaW50b3NoKSI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAuaWlkOjkzN2RiNmMxLWUxNGMtNDE5Zi1hZmRjLTZkZmJjMDgzMjYxYSIgc3RSZWY6ZG9jdW1lbnRJRD0iYWRvYmU6ZG9jaWQ6cGhvdG9zaG9wOjVhMzZlZDFlLWI1OTgtMTE3OS04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"

help: |
  ## About
  [IBM BigFix](https://www.ibm.com/security/bigfix/) is a systems-management software product developed by IBM for managing large groups of computers.
  This plugin allows the user to interact with their IBM BigFix instance for managing Actions, Computers, Computer Groups, Fixlets and Sites.
  Users can create and deploy Actions to specific Sites, Computers or Computer Groups for vulnerability and patch remediation.
  The connection requires a valid URL to the users BigFix instance, as well as a username and password.
  This plugin contains actions for interacting with the following BigFix objects:
  - Actions
  - Computers
  - Computer Groups
  - Fixlets
  - Sites
  It has a "query" action, which will return the result of a provided BigFix Relevance string.
  It also has a custom endpoint action, which allows GET / POST / DELETE / PUT requests to be made against a provided endpoint.
  This plugin utilizes the [IBM BigFix API](https://developer.bigfix.com/rest-api/)
  This plugin uses the [lxml Python library](http://lxml.de/) for parsing the XML responses from BigFix.
  ## Actions
  ### Get Computer
  This action is used to get core properties of a computer.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |id|string|None|False|The ID of the computer|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |computer|ComputerCoreProperties|True|The core properties of the computer|
  Example output:
  {
    "computer": {
      "Resource": "https://10.1.1.1:52311/api/computer/1840504",
      "Name": "WIN7"
    }
  }
  ### Custom Endpoint
  This action is used to execute a request against a custom endpoint.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |endpoint|string|None|True|The endpoint against which the request will be executed e.g. 'computers'|None|
  |data|string|None|False|The XML data which will be attached to the request|None|
  |method|string|None|True|The method to use for the request|['DELETE', 'GET', 'PATCH', 'POST']|
  |parameters|object|None|False|The URL parameters to include in the request|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |response|string|True|The XML response returned from the endpoint|
  Example output:
  
  {
    "response": "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<BESAPI xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:noNamespaceSchemaLocation=\"BESAPI.xsd\">\n\t<Computer Resource=\"https://10.1.1.1:52311/api/computer/745243\">\n\t\t<LastReportTime>Mon, 23 Oct 2017 14:39:48 +0000</LastReportTime>\n\t\t<ID>745243</ID>\n\t</Computer>\n\t<Computer Resource=\"https://10.1.1.1:52311/api/computer/1840504\">\n\t\t<LastReportTime>Fri, 20 Oct 2017 11:42:45 +0000</LastReportTime>\n\t\t<ID>1840504</ID>\n\t</Computer>\n\t<Computer Resource=\"https://10.1.1.1:52311/api/computer/5177265\">\n\t\t<LastReportTime>Mon, 23 Oct 2017 14:40:02 +0000</LastReportTime>\n\t\t<ID>5177265</ID>\n\t</Computer>\n\t<Computer Resource=\"https://10.1.1.1:52311/api/computer/14122952\">\n\t\t<LastReportTime>Mon, 23 Oct 2017 14:37:55 +0000</LastReportTime>\n\t\t<ID>14122952</ID>\n\t</Computer>\n</BESAPI>\n"
  }
  
  ### Create Computer Group
  This action is used to create a computer group in a specified site.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |computer_group|ComputerGroup|None|True|The description of the Computer Group to create in BigFix|None|
  |chosen_site|SiteDefinition|None|True|Details of the site where the Computer Group will be created|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |computer_group|ComputerGroupDefinition|True|Information about the newly created Computer Group|
  Example output:
  
  {
    "computer_group": {
      "LastModified": "2017-10-23T14:40:13+00:00",
      "Name": "Test Computer Group",
      "ID": 738,
      "Resource": "https://10.1.1.1:52311/api/computergroup/master/738"
    }
  }
  
  ### Get Computers For Fixlet
  This action is used to fetch a list of computers which are relevant to a particular Fixlet.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |site_details|SiteDefinition|None|True|The details of the site to which the Fixlet belongs|None|
  |fixlet_id|string|None|True|The ID for the Fixlet|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |computers|[]string|True|The computer IDs which are relevant to the Fixlet|
  Example output:
  
  {
    "type": "action_event",
    "body": {
      "meta": {},
      "output": {
        "computers": [
          "745243",
          ...
          "14122952"
        ]
      },
      "status": "ok"
    },
    "version": "v1"
  }
  
  ### Create Action
  This action is used to create an action in BigFix.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |SourcedFixletAction|SourcedFixletAction|None|False|The Action (Site, Fixlet ID, Action, Settings) to create|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |action|ActionDefinition|True|Details of the newly created Action|
  Example output:
  
  {
    "action": {
      "Resource": "https://10.1.1.1:52311/api/action/737",
      "Name": "BES Quick Reference - Production",
      "LastModified": "2017-10-23T14:40:11+00:00",
      "ID": 737
    }
  }
  
  ### Stop Action
  This action is used to stop an action in BigFix.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |action_id|integer|None|True|The Action ID|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |response|string|True|The response returned from BigFix|
  Example output:
  
  {
    "response": "ok"
  }
  
  ### Get Fixlet
  This action is used to fetch a specific Fixlet from a site.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |id|string|None|True|Identifier for the Fixlet|None|
  |chosen_site|SiteDefinition|None|True|Details of the site from which to retrieve a specific Fixlet|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |fixlet|FixletWithActions|True|The information retrieved about the Fixlet|
  Example output:
  
  {
    "fixlet": {
      "SourceReleaseDate": "2002-04-11",
      "MIMEFields": [
        {
          "Name": "x-fixlet-first-propagation",
          "Value": "Thu, 31 Jan 2002 16:49:06 +0000"
        },
        {
          "Name": "x-fixlet-domain_attributes",
          "Value": "BES Warn Critical"
        },
        {
          "Name": "x-fixlet-modification-time",
          "Value": "Tue, 02 Oct 2012 23:13:00 +0000"
        }
      ],
      "Category": "Licensing",
      "Description": "\n<TABLE><TBODY><TR><TD><FONT size=2>You appear to have more BES Clients on your network than your license allows. The computers listed are operating in \"grace mode\", which means that their license number is higher than allowed by your license but lower than the threshold at which they will stop functioning. <BR><BR>This situation may occur if you are in the process of uninstalling and reinstalling BES Clients, in which case the license numbers for the uninstalled BES Clients may not yet have been released, or it may be because you are actively running more clients than allowed by your license. If the latter is the case, you should contact BigFix immediately (<A href=\"mailto:TEM@dk.ibm.com\">TEM@dk.ibm.com</A>)&nbsp;to obtain a larger license. </FONT></TD></TR></TBODY></TABLE>\n",
      "DownloadSize": 0,
      "Delay": "PT1S",
      "SourceSeverity": "Critical",
      "Source": "BigFix",
      "Relevance": [
        "evaluation of client license = false",
        "seat count state of client license = \"Grace\""
      ],
      "Title": "BES Clients in Seat Count Grace Mode",
      "SourceID": "<Unspecified>"
    }
  }
  
  ### Get Action Status
  This action is used to get an action status from BigFix.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |action_id|integer|None|True|The Action ID|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |status|string|True|The status of the Action|
  Example output:
  
  {
    "status": "Expired"
  }
  
  ### Fetch Computers
  This action is used to fetch a list of computers.
  #### Input
  This action does not contain any inputs.
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |computers|[]Computer|True|List of computer identifiers returned from BigFix|
  Example output:
  
  {
    "computers": [
      {
        "ID": "745243",
        "Resource": "https://10.1.1.1:52311/api/computer/745243",
        "LastReportTime": "2017-10-23T14:39:48+00:00"
      },
      ...
      {
        "ID": "14122952",
        "Resource": "https://10.1.1.1:52311/api/computer/14122952",
        "LastReportTime": "2017-10-23T14:37:55+00:00"
      }
    ]
  }
  
  ### Delete Computer Group
  This action is used to delete a computer group in a specified site.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |id|string|None|True|The ID of the Computer Group to delete|None|
  |chosen_site|SiteDefinition|None|True|Details of the site containing the Computer Group to be deleted|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |message|string|False|The message returned from the BigFix console|
  |result|boolean|True|The result of the operation|
  Example output:
  
  {
    "result": true
  }
  
  ### Fetch Sites
  This action is used to fetch a list of sites and their types available to the authenticated operator.
  #### Input
  This action does not contain any inputs.
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |sites|[]SiteDefinition|True|The current list of Sites on the Server|
  Example output:
  
  {
    "sites": [
      {
        "Resource": "https://10.1.1.1:52311/api/site/external/BES%20Support",
        "GatherURL": "http://sync.bigfix.com/cgi-bin/bfgather/bessupport",
        "Name": "BES Support",
        "Type": "external"
      },
      ...
      {
        "Resource": "https://10.1.1.1:52311/api/site/master",
        "GatherURL": "http://R7_BigFix:52311/cgi-bin/bfgather.exe/actionsite",
        "Name": "ActionSite",
        "Type": "master"
      }
    ]
  }
  
  ### Delete Action
  This action is used to delete an action from BigFix.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |action_id|integer|None|True|The Action ID|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |response|string|True|The response returned from BigFix|
  Example output:
  
  {
    "response": "ok"
  }
  
  ### Fetch Relevant Fixlets
  This action is used to fetch a list of Fixlets which are relevant to provided criteria.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |computer_groups|[]string|None|False|The computer groups to which the Fixlets are relevant|None|
  |titles|[]string|None|False|The titles to which the Fixlets must match|None|
  |site|string|None|False|The site to which the Fixlets belong|None|
  |cves|[]string|None|False|The CVEs to which the Fixlets are relevant|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |results|[]FixletActionSummary|True|The relevant Fixlets retrieved from BigFix|
  Example output:
  
  {
  "results": [
      {
        "SiteName": "BES Support",
        "Name": "Automatically Restart Stopped BES Clients Using TaskScheduler",
        "DefaultActionId": "Action1",
        "Id": "250"
      },
      ...
      {
        "SiteName": "BES Support",
        "Name": "Install/Update IBM BigFix Client Deploy Tool (Version 9.5.7)",
        "DefaultActionId": "Action1",
        "Id": "3139"
      }
    ]
  }
  
  ### Query
  This action is used to evaluate a relevance expression and get the result.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |query_string|string|None|True|The Relevance expression to evaluate|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |results|QueryResponse|True|The results from the evaluated expression|
  Example output:
  
  {
    "results": {
      "Resource": "(id of it, name of site of it) of unique values of relevant Fixlets of BES computers",
      "Result": [
        {
          "Answers": [
            {
              "Type": "integer",
              "Value": "38"
            },
            {
              "Type": "string",
              "Value": "ActionSite"
            }
          ]
        },
        ...
        {
          "Answers": [
            {
              "Type": "integer",
              "Value": "404376715"
            },
            {
              "Type": "string",
              "Value": "Enterprise Security"
            }
          ]
        }
      ],
      "Evaluation": {
        "Plurality": "Plural",
        "Time": 41
      }
    }
  }
  
  ### Login
  This action is used to create or refresh connection.
  #### Input
  This action does not contain any inputs.
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |success|boolean|True|Whether the login request succeeded|
  Example output:
  
  {
    "success": true
  }
  
  ### Get Site
  This action is used to fetch a specific site.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |chosen_site|SiteDefinition|None|True|Details for the site|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |site|SiteDescription|True|Details of the selected Site|
  Example output:
  
  {
      "site": {
        "Name": "ActionSite",
        "GatherURL": "http://10.1.1.1:52311/cgi-bin/bigfix/actionsite",
        "Description": "",
        "GlobalReadPermission": false,
        "Subscription": {
          "Mode": "All"
        }
  }
  
  ### Get Site Content
  This action is used to get identifiers for all content from a site - Fixlets, tasks, baselines, analysis and computer groups.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |chosen_site|SiteDefinition|None|True|Details of the site from which to retrieve the content|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |fixlets|[]FixletDefinition|True|The current list of Fixlets on the Site|
  |analyses|[]AnalysisDefinition|True|The current list of Analyses on the Site|
  |tasks|[]TaskDefinition|True|The current list of Tasks on the Site|
  |baselines|[]BaselineDefinition|True|The current list of Baselines on the Site|
  |computergroups|[]ComputerGroupDefinition|True|The current list of Computer Groups on the Site|
  Example output:
  
  {
    "computergroups": [],
    "analyses": [
      {
        "Resource": "https://10.1.1.1:52311/api/analysis/external/BES%20Support/204",
        "ID": 204,
        "LastModified": "2017-10-19T20:07:16+00:00",
        "Name": "BES Component Versions"
      },
      ...
      {
        "Resource": "https://10.1.1.1:52311/api/task/external/BES%20Support/3161",
        "ID": 3161,
        "LastModified": "2017-10-19T20:07:16+00:00",
        "Name": "Install IBM BigFix Relay (Version 9.2.12)"
      }
    ],
    "baselines": []
  }
  
  ### Get Actions
  This action is used to get all action summaries from BigFix.
  #### Input
  This action does not contain any inputs.
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |actions|[]ActionDefinition|True|Information on the Actions present in BigFix|
  Example output:
  
  {
    "actions": [
      {
        "LastModified": "2017-09-18T11:16:09+00:00",
        "ID": 37,
        "Name": "Change '__Group_0_Man group 1' Setting",
        "Resource": "https://10.1.1.1:52311/api/action/37"
      },
     ...
      {
        "LastModified": "2017-10-23T14:40:11+00:00",
        "ID": 737,
        "Name": "BES Quick Reference - Production",
        "Resource": "https://10.1.1.1:52311/api/action/737"
      }
    ]
  }
  
  ### Create Baseline
  This action is used to create a Baseline in BigFix.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |baseline|Baseline|None|False|The Site, Actions, Fixlet IDs and Settings to create a Baseline|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |baseline_details|BaselineDefinition|True|Details of the newly created Baseline|
  Example output:
  
  {
      "baseline_details": {
        "LastModified": "2018-06-18T14:45:45+00:00",
        "Name": "Testing the create_baseline action",
        "ID": 2978,
        "Resource": "<url_of_bigfix_console>/api/baseline/master/2978"
      }
  }
  
  ### Get Matching Fixlets
  This action is used to return a list of Fixlets from a site with properties matching input criteria.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |chosen_site|SiteDefinition|None|True|Details of the site from which to retrieve matching Fixlets|None|
  |criteria|[]FieldMatcher|None|True|The criteria used to match relevant Fixlets. All provided criteria must match for a Fixlet ID to be returned. Each criterion should be in the form {"FieldName"\: "a", "Matcher"\: "b", "Value"\: ["c", "d"]}, creating a JSON array|None|
  Example Criteria input:
  
  [
    {
      "FieldName": "Title",
      "Matcher": "contains",
      "Value": ["Windows 11", "Visual Basic 6.0"]
    },
    {
      "FieldName": "CVENames",
      "Matcher": "doesNotEqual",
      "Value": ["CVE-1234", "Unspecified"]
    }
  ]
  
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |matching_fixlet_ids|[]string|True|The IDs of Fixlets matching the provided criteria|
  Example output:
  
  {
    "matching_fixlet_ids": [
      "800815",
      "800817",
      "807006",
      "807008",
      "807011",
      "807012",
      "807016",
      "807017",
      "807021",
      "807022",
      "807026",
      "807027",
      "807031",
      "807032",
      "807041",
      "807046",
      "1600493"
    ]
  }
  
  ### Fetch Fixlets
  This action is used to fetch information on all Fixlets for a provided site.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |chosen_site|SiteDefinition|None|True|Details of the site from which to retrieve the Fixlets|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |fixlets|[]FixletDefinition|True|List of Fixlets from the provided Site|
  Example output:
  
  {
    "fixlets": [
      {
        "ID": 24,
        "Resource": "https://10.5.1.92:52311/api/fixlet/external/Patches%20for%20Mac%20OS%20X/24",
        "Name": "MAINTENANCE: Repair Permissions",
        "LastModified": "2017-10-20T08:11:48+00:00"
      },
      ...
      {
        "ID": 98140754,
        "Resource": "https://10.5.1.92:52311/api/fixlet/external/Patches%20for%20Mac%20OS%20X/98140754",
        "Name": "UPDATE: Safari 11.0 - MacOS Sierra (10.12 Client)",
        "LastModified": "2017-10-20T08:11:48+00:00"
      }
    ]
  }
  
  ### Fetch Group Members
  This action is used to fetch a list of computers relevant to a computer group.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |id|string|None|False|The ID of the computer group|None|
  |chosen_site|SiteDefinition|None|True|Details of the site to which the Computer Group belongs|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |computers|[]string|True|The resource links for the computers within the computer group|
  Example output:
  
  {
    "computers": [
      "https://10.5.1.92:52311/api/computer/1840504"
    ]
  }
  
  ### Create Multiaction Group
  This action is used to create a multi-action group in BigFix.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |multiaction_group|MultiActionGroup|None|False|The Actions (Site, Fixlet ID, Action) to create along with the Target and Settings|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |action|ActionDefinition|True|Details of the newly created MultiAction Group|
  Example output:
  
  {
    "action":
    {
      "ID": 4884,
      "Name": "Local Test Create MultiAction Group",
      "Resource": "https://10.5.1.92:52311/api/action/4884",
      "LastModified": "2017-11-21T21:30:49+00:00"
    }
  }
  
  ### Fetch Computer Groups
  This action is used to fetch a list of computer groups for a provided site.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |chosen_site|SiteDefinition|None|True|Details of the site from which to retrieve the Computer Groups|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |computer_groups|[]ComputerGroupDefinition|True|The computer groups belonging to the specified site|
  Example output:
  
  {
    "computer_groups": [
      {
        "ID": 36,
        "Resource": "https://10.5.1.92:52311/api/computergroup/master/36",
        "Name": "Man group 1",
        "LastModified": "2017-09-20T09:29:24+00:00"
      },
      ...
      {
        "ID": 38,
        "Resource": "https://10.5.1.92:52311/api/computergroup/master/38",
        "Name": "Windows Machines",
        "LastModified": "2017-09-18T11:17:52+00:00"
      }
    ]
  }
  
  ### Get Settings
  This action is used to get a computer's settings.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |computer_id|string|None|True|The computer ID|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |settings|[]Setting|True|List of settings for the chosen Computer|
  Example output:
  
  {
    "settings": [
      {
        "Value": "Service manager shutdown request",
        "Name": "_BESClient_LastShutdown_Reason"
      },
      ...
      {
        "Value": "",
        "Name": "__Relay_Control_Server2"
      }
    ]
  }
  
  ### Query Host For Specific Software
  This action is used to query a host to see if a specific software package is installed.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |hostname|string|None|True|Hostname of the host to query|None|
  |software|string|None|True|Installed software to query the host for|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |found|boolean|True|Whether or not the host has the specified software installed|
  Example output:
  
  {
    "found": true
  }
  
  ### Query Host For Installed Software
  This action is used to query a host for all installed software.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |hostname|string|None|True|Hostname of the host to query|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |results|[]string|False|Software packages installed on the host|
  Example output:
  
  {
      "results": ["Adobe Acrobat 5.0 | 5.0", "Adobe Flash Player 28 NPAPI | 28.0.0.161"]
  }
  
  ### Query Host Inventory
  This action is used to get an inventory of hosts, including their hostnames, IP addresses, and operating system.
  #### Input
  This action does not contain any inputs.
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |results|[]inventory_host|False|Hosts found in the BigFix inventory|
  Example output:
  
  {
  	"results": [{
  		"os": "Win2012 6.2.9200",
  		"ip": "10.4.23.69",
  		"hostname": "BIGFIX-SRV"
  	}, {
  		"os": "Win2012 6.2.9200",
  		"ip": "10.4.23.46",
  		"hostname": "BIGFIX-CLT-W12"
  	}]
  }
  
  ## Triggers
  ### Poll Finished Actions
  This trigger is used to poll for recently stopped or expired actions.
  #### Input
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |pattern|string|None|False|Pattern to match against Action name|None|
  |interval|integer|60|True|Poll interval (in minutes)|None|
  #### Output
  |Name|Type|Required|Description|
  |----|----|--------|-----------|
  |action_ids|[]string|True|IDs of stopped or expired actions|
  Example output:
  
  {
    "action_ids": ["1122"]
  }
  
  ## Connection
  The connection configuration accepts the following parameters:
  |Name|Type|Default|Required|Description|Enum|
  |----|----|-------|--------|-----------|----|
  |url|string|None|True|Host URL E.g. https\://<bigfix_server>\:52311|None|
  |username|string|None|True|Basic Auth Username|None|
  |password|password|None|True|Basic Auth Password|None|
  |ssl_verify|boolean|True|True|Validate certificate|None|
  ## Troubleshooting
  This plugin does not contain any troubleshooting information.
  ## Versions
  * 0.1.0 - Initial plugin
  * 1.0.0 - Adding trigger
  * 2.0.0 - Support web server mode
  * 2.1.0 - Add actions: Query Host Inventory, Query Host for Installed Software, Query Host for Specific Software
  * 2.1.1 - Update help
  * 2.1.2 - Add Baseline action
  ## References
  * [IBM BigFix Website](https://www.ibm.com/security/bigfix/)
  * [IBM BigFix REST API](https://developer.bigfix.com/rest-api/).
  * [lxml Website](http://lxml.de/)
# https://github.com/komand/plugins/issues/718

connection:
  url:
    title: "URL"
    type: string
    description: "Host URL e.g. https://<bigfix_server>:52311"
    required: true
  credentials:
    title: "Credentials"
    type: credential_username_password
    description: "Username and password"
    required: true
  ssl_verify:
    title: "SSL Verify"
    type: boolean
    description: "Validate certificate"
    default: true
    required: true

triggers:
  poll_finished_actions:
    description: "Poll for recently stopped or expired Actions"
    input:
      interval:
        type: integer
        default: 60
        description: "Poll interval (in minutes)"
        required: true
      pattern:
        type: string
        description: "Pattern to match against Action name"
        required: false
    output:
      action_ids:
        type: "[]string"
        description: "IDs of stopped or expired actions"
        required: true

types:
  # Sub-types for main XML types
  ComputerCoreProperties:
    Resource:
      type: string
      description: "The resource URI of the computer group"
    Name:
      type: string
      description: "The name of the computer"

  BESActionTarget:
    ComputerName:
      title: "Computer Names"
      type: "[]string"
      description: "A list of target BigFix Computer Names"
    ComputerID:
      title: "Computer IDs"
      type: "[]integer"
      description: "A list of target BigFix Computer IDs"
    CustomRelevance:
      title: "Custom Relevance"
      type: string
      description: "A Relevance string describing the target Computers e.g by specific Computer Group"
    AllComputers:
      type: boolean
      description: "Selects all computers as the target"

  BESActionSourceFixlet:
    GatherURL:
      type: string
      required: false
    SiteName:
      title: "Site Name"
      type: string
      required: false
      description: "The name of the target Site"
    SiteID:
      title: "Site ID"
      type: integer
      required: false
      description: "The ID of the target Site"
    FixletID:
      title: "Fixlet ID"
      type: integer
      required: true
      description: "The ID of the target Fixlet"
    Action:
      type: string
      required: false
      description: "The Action to perform"

  ComputerGroupSearchComponentRelevance:
    Relevance:
      description: "The Relevance string linking Computers"
      type: string
      required: true
    Comparison:
      description: "Whether the target Computers match or do not match the provided Relevance string"
      type: string
      enum:
      - "IsTrue"
      - "IsFalse"

  ComputerGroupSearchComponentPropertyReference:
    SearchText:
      title: "Search Text"
      description: "The text to the match the provided Property"
      type: string
      required: true
    Relevance:
      description: "The Relevance string"
      type: string
      required: true
    PropertyName:
      title: "Property Name"
      description: "The property to match Computers with"
      type: string
      required: true
    Comparison:
      description: "How the provided text should be related to the Property chosen"
      type: string
      enum:
      - "Contains"
      - "DoesNotContain"
      - "Equals"
      - "DoesNotEqual"

  ComputerGroupSearchComponentGroupReference:
    GroupName:
      title: "Computer Group Name"
      description: "The Computer Group"
      type: string
      required: true
    Comparison:
      description: "How the Computers should be related to the Computer Group"
      type: string
      enum:
      - "IsMember"
      - "IsNotMember"

  ActionSuccessCriteria:
    Option:
      type: string
      description: "How BigFix determines the success of the Action"
      required: true
      enum:
      - "RunToCompletion"
      - "OriginalRelevance"
      - "CustomRelevance"
    Value:
      type: string
      description: "The Relevance string describing the success criteria"
      required: true

  LocalGroup:
    Name:
      type: string
      required: true
    Value:
      type: string

  DomainGroup:
    Name:
      type: string
      required: true
    Sid:
      type: string
      required: true
    Value:
      type: string
      required: true

  ActionReapplyInterval:
    Value:
      description: "The time interval before an attempt is made to reapply the action"
      type: string
      enum:
      - "PT15M"
      - "PT30M"
      - "PT1H"
      - "PT2H"
      - "PT4H"
      - "PT6H"
      - "PT8H"
      - "PT12H"
      - "P1D"
      - "P2D"
      - "P3D"
      - "P5D"
      - "P7D"
      - "P15D"
      - "P30D"

  ActionMessageMaxPostponementInterval:
    Value:
      description: "The chosen maximum postponement time interval"
      type: string
      enum:
      - "PT15M"
      - "PT30M"
      - "PT1H"
      - "PT2H"
      - "PT4H"
      - "PT6H"
      - "PT8H"
      - "PT12H"
      - "P1D"
      - "P2D"
      - "P3D"
      - "P5D"
      - "P7D"
      - "P15D"
      - "P30D"

  ActionMessageTimeInterval:
    Value:
      description: "The time interval before sending the Action message"
      type: string
      enum:
      - "PT1M"
      - "PT2M"
      - "PT3M"
      - "PT4M"
      - "PT5M"
      - "PT10M"
      - "PT15M"
      - "PT30M"
      - "PT1H"
      - "PT2H"
      - "PT4H"
      - "PT6H"
      - "PT8H"
      - "PT12H"
      - "P1D"
      - "P2D"
      - "P3D"
      - "P5D"
      - "P7D"
      - "P15D"
      - "P30D"

  ActionDeadlineBehavior:
    Value:
      description: "The behaviour of the Action when the Deadline is reached"
      type: string
      enum:
      - "ForceToRun"
      - "RunAutomatically"

  RetryWaitInterval:
    Value:
      description: "The interval to wait before retrying the Action"
      type: string
      enum:
      - "PT10M"
      - "PT15M"
      - "PT30M"
      - "PT1H"
      - "PT2H"
      - "PT4H"
      - "PT6H"
      - "PT8H"
      - "PT12H"
      - "P1D"
      - "P2D"
      - "P3D"
      - "P5D"
      - "P7D"
      - "P15D"
      - "P30D"

  MIMEField:
    Name:
      description: "The name / identifier of the MIMEField"
      type: string
      required: true
    Value:
      description: "The value of the MIMEField"
      type: string
      required: true

  ActionSettingsLockMessage:
    Title:
      description: "If the Action Message has a title"
      type: boolean
      required: false
    Text:
      description: "If the Action Message has text"
      type: boolean
      required: false
    ShowActionButton:
      title: "Show Action Button"
      description: "If the Action Message shows an Action button"
      type: boolean
      required: false
    ShowCancelButton:
      title: "Show Cancel Button"
      description: "If the Action Message shows a Cancel button"
      type: boolean
      required: false
    Postponement:
      description: "If the Action Message contains Postponement options"
      type: boolean
      required: false
    Timeout:
      description: "If the Action Message contains Timeout options"
      type: boolean
      required: false

  ActionSettingsLockPreAction:
    Title:
      description: "If the Action PreAction has a title"
      type: boolean
      required: false
    Text:
      description: "If the Action PreAction has text"
      type: boolean
      required: false
    AskToSaveWork:
      title: "Ask To Save Work"
      description: "If the Action PreAction contains an Ask to Save work option"
      type: boolean
      required: false
    ShowActionButton:
      title: "Show Action Button"
      description: "If the Action PreAction contains an Action button"
      type: boolean
      required: false
    ShowCancelButton:
      title: "Show Cancel Button"
      description: "If the Action PreAction contains a Cancel button"
      type: boolean
      required: false
    DeadlineBehavior:
      title: "Deadline Behaviour"
      description: "If the Action PreAction contains Deadline behaviour"
      type: boolean
      required: false
    ShowConfirmation:
      description: "If the Action PreAction shows the Confirmation message"
      type: boolean
      required: false

  ActionSettingsLockRunningMessage:
    Title:
      description: "If the Action Running Message has a title"
      type: boolean
      required: false
    Text:
      description: "If the Action Running Message has text"
      type: boolean
      required: false

  ActionSettingsLockPostActionBehavior:
    Title:
      description: "If the Action Settings PostAction has a title"
      type: boolean
      required: false
    Text:
      description: "If the Action Settings PostAction has text"
      type: boolean
      required: false
    Behavior:
      description: "If the Action Settings PostAction has Behaviour options"
      type: boolean
      required: false
    AllowCancel:
      title: "Allow Cancellation"
      description: "If the Action Settings PostAction has options to allow cancelling"
      type: boolean
      required: false
    Postponement:
      description: "If the Action Settings PostAction has Postponement options"
      type: boolean
      required: false
    Timeout:
      description: "If the Action Settings PostAction has Timeout options"
      type: boolean
      required: false
    Deadline:
      description: "If the Action Settings PostAction has Deadline options"
      type: boolean
      required: false

  ActionSettingsMessage:
    Title:
      description: "The title of the Action Message"
      type: string
      required: false
    Text:
      description: "The text of the Action Message"
      type: string
      required: false
    ShowActionButton:
      title: "Show Action Button"
      description: "If the Action Message shows the Action button"
      type: boolean
      required: false
    ShowCancelButton:
      title: "Show Cancel Button"
      description: "If the Action Message shows the Cancel button"
      type: boolean
      required: false
    AllowPostponement:
      title: "Allow Postponement"
      description: "If the Action Message allows Postponement"
      type: boolean
      required: false
    MaxPostponementInterval:
      title: "Maximum Postponement Interval"
      description: "The Action Message maximum Postponement interval"
      type: ActionMessageMaxPostponementInterval
      required: false
    PostponementDeadlineOffset:
      title: "Postponement Deadline Offset"
      description: "The Action Message offset for the Postponement Deadline. The value must follow the ISO-8601 standard for duration format e.g PT15M for 15 minutes or P2D for 2 days"
      type: string
      required: false
    HasTimeout:
      title: "Has Timeout"
      description: "If the Action Message has a timeout"
      type: boolean
      required: false
    TimeoutInterval:
      title: "Timeout Interval"
      description: "The time interval of the Timeout"
      type: ActionMessageTimeInterval
      required: false

  ActionSettingsPreAction:
    Text:
      description: "The text of the Action PreAction"
      type: string
      required: false
    AskToSaveWork:
      title: "Ask To Save Work"
      description: "If the Action PreAction asks to save work"
      type: boolean
      required: false
    ShowActionButton:
      title: "Show Action Button"
      description: "If the Action PreAction shows the Action button"
      type: boolean
      required: false
    ShowCancelButton:
      title: "Show Cancel Button"
      description: "If the Action PreAction shows the Cancel button"
      type: boolean
      required: false
    DeadlineBehavior:
      title: "Deadline Behaviour"
      description: "The Action PreAction deadline behaviour"
      type: ActionDeadlineBehavior
      required: false
    DeadlineType:
      title: "Deadline Type"
      description: "The Action PreAction deadline type"
      type: string
      enum:
      - "Interval"
      - "Absolute"
    DeadlineInterval:
      title: "Deadline Interval"
      description: "The Action PreAction deadline interval"
      type: ActionMessageTimeInterval
      required: false
    DeadlineOffset:
      title: "Deadline Offset"
      description: "The Action PreAction deadline offset. The value must follow the ISO-8601 standard for duration format e.g PT15M for 15 minutes or P2D for 2 days"
      type: string
      required: false
    DeadlineLocalOffset:
      title: "Deadline Local Offset"
      description: "The Action PreAction deadline local offset. The value must follow the ISO-8601 standard for duration format e.g PT15M for 15 minutes or P2D for 2 days"
      type: string
      required: false
    ShowConfirmation:
      title: "Show Confirmation"
      description: "If the Action PreAction shows confirmation"
      type: boolean
      required: false
    Confirmation:
      description: "The Action PreAction confirmation message"
      type: string
      required: false

  ActionSettingsRunningMessage:
    Title:
      description: "The title of the Action Running Message"
      type: string
      required: false
    Text:
      description: "The text of the Action Running Message"
      type: string
      required: false

  ActionSettingsTimeRange:
    StartTime:
      description: "The Action start time"
      type: string
      required: false
    EndTime:
      description: "The Action end time"
      type: string
      required: false

  ActionSettingsDayOfWeekConstraint:
    Sun:
      description: "If the Action is restricted to running on a Sunday"
      type: boolean
      required: false
    Mon:
      description:  "If the Action is restricted to running on a Monday"
      type: boolean
      required: false
    Tue:
      description: "If the Action is restricted to running on a Tuesday"
      type: boolean
      required: false
    Wed:
      description: "If the Action is restricted to running on a Wednesday"
      type: boolean
      required: false
    Thu:
      description: "If the Action is restricted to running on a Thursday"
      type: boolean
      required: false
    Fri:
      description: "If the Action is restricted to running on a Friday"
      type: boolean
      required: false
    Sat:
      description: "If the Action is restricted to running on a Saturday"
      type: boolean
      required: false

  ActionSettingsUIGroupConstraints:
    Win9xGroups:
      description: "A collection of UI constraints for Win9xGroups *"
      #type: "[]Win9xGroup"
      type: "[]string"
      required: false
    WinNTGroups:
      description: "A collection of UI constraints for WinNTGroups *"
      #type: "[]WinNTGroup"
      type: "[]string"
      required: false
    LocalGroups:
      title: "Local Groups"
      description: "A collection of UI constraints for LocalGroups *"
      type: "[]LocalGroup"
      required: false
    DomainGroups:
      title: "Domain Groups"
      description: "A collection of UI constraints for DomainGroups *"
      type: "[]DomainGroup"
      required: false

  ActionSettingsWhose:
    Property:
      description: "The property to check"
      type: string
      required: false
    Relation:
      description: "How the targets property should match the provided value"
      type: string
    Value:
      description: "The property value to match"
      type: string
      required: false

  ActionSettingsRetryWait:
    Value:
      description: "The interval to wait before retrying the Action"
      type: RetryWaitInterval
      required: true
    Behavior:
      default: "WaitForInterval"
      description: "The RetryWait behaviour of the Action"
      type: string
      enum:
      - "WaitForReboot"
      - "WaitForInterval"

  ActionSettingsPostActionBehavior:
    Title:
      description: "The title of the PostAction Behaviour"
      type: string
      required: false
    Text:
      description: "The text of the PostAction Behaviour"
      type: string
      required: false
    AllowCancel:
      description: "If the PostAction allows the user to cancel"
      type: boolean
      required: false
    AllowPostponement:
      title: "Allow Postponement"
      description: "If the PostAction allows postponement"
      type: boolean
      required: false
    MaxPostponementInterval:
      title: "Max Postponement Interval"
      description: "The maximum postponement interval"
      type: ActionMessageMaxPostponementInterval
      required: false
    HasTimeout:
      title: "Has Timeout"
      description: "If the PostAction has a timeout"
      type: boolean
      required: false
    TimeoutInterval:
      title: "Timeout Interval"
      description: "The interval of the PostAction timeout"
      type: ActionMessageTimeInterval
      required: false
    PostActionDeadlineBehavior:
      title: "PostAction Deadline Behaviour"
      description: "The deadline behaviour of the PostAction"
      type: ActionDeadlineBehavior
      required: false
    PostActionDeadlineInterval:
      title: "PostAction Deadline Interval"
      description: "The interval to wait"
      type: ActionMessageTimeInterval
      required: false
    Behavior:
      default: "Nothing"
      description: "The PostAction behaviour"
      type: string
      enum:
      - "Nothing"
      - "Restart"
      - "Shutdown"

  FixletActionDescription:
    PreLink:
      type: string
    Link:
      type: string
      description: "The Fixlet Action Link"
    PostLink:
      type: string

  FixletActionActionScript:
    Value:
      type: string
      description: "The value of the Fixlet Actions ActionScript - the script run"
      required: true
    MIMEType:
      type: string
      description: "The type of ActionScript provided"
      default: "application/x-Fixlet-Windows-Shell"

  ComputerGroupRelevance:
    SearchComponentRelevance:
      title: "SearchComponent Relevance"
      type: ComputerGroupSearchComponentRelevance
      description: "This Computer Group Relevance groups computers based on if they match or do not match the provided Relevance string"
    SearchComponentPropertyReference:
      title: "SearchComponent Property Reference"
      type: ComputerGroupSearchComponentPropertyReference
      description: "This Computer Group Relevance groups computers if they have a Property that matches an expected value. This grouping can be further restricted by applying a Relevance string"
    SearchComponentGroupReference:
      title: "SearchComponent Group Reference"
      type: ComputerGroupSearchComponentGroupReference
      description: "This Computer Group Relevance groups computers on whether they are in (or not) a specific existing Computer Group"

  FixletWithActionsGroupRelevance:
    Groups:
      description: "The collection of Computer Group Relevances that describe the Computers that should receive the Fixlet"
      type: "[]ComputerGroupRelevance"
    JoinByIntersection:
      title: "Join By Intersection"
      description: "If the various Relevances each describe different groupings, or if a Computer needs to be present in all groups"
      type: boolean
      required: true

  FixletWithActionsWizardData:
    Name:
      description: "The name displayed in the Wizard"
      type: string
      required: false
    DataStore:
      description: "The associated Fixlet datastore *"
      type: string
      required: false
    StartURL:
      description: "The starting URL for the Fixlet Wizard *"
      type: string
      required: false
    SkipUI:
      description: "If the Wizard skips the UI section"
      type: string
      required: false

  SubscriptionCustomGroup:
    ComputerGroups:
      title: "Computer Groups"
      type: "[]ComputerGroupRelevance"
      description: "The collection of Computer Group Relevances describing the subscription to the site by different Computers"
    JoinByIntersection:
      title: "Join By Intersection"
      type: boolean
      description: "If the various Relevances each describe different groupings, or if a Computer needs to be present in all groups"

  Subscription:
    Mode:
      type: string
      description: "The type of Subscription - does it include all Computers, none, an AdHoc grouping or a CustomGroup decided by the User"
      required: true
      enum:
      - "All"
      - "None"
      - "Adhoc"
      - "Custom"
    CustomGroup:
      title: "Custom Group"
      type: SubscriptionCustomGroup
      description: "The Custom Grouping linking Computers together by a set of shared traits or values"

  ActionSettingsLocks:
    HasMessage:
      title: "Has Message"
      description: "If the Action Settings has a message option"
      type: boolean
      required: false
    Message:
      description: "The availability of settings for the Action message"
      type: ActionSettingsLockMessage
    ActionUITitle:
      title: "Action UI Title"
      description: "If the Action settings have a UI title option"
      type: boolean
      required: false
    PreActionShowUI:
      title: "PreAction Show UI"
      description: "If the Action settings have an option to show the PreAction UI element"
      type: boolean
      required: false
    PreAction:
      description: "The availability of settings for the Action PreAction"
      type: ActionSettingsLockPreAction
    HasRunningMessage:
      title: "Has Running Message"
      description: "If the Action has a Running Message option"
      type: boolean
      required: false
    RunningMessage:
      title: "Running Message"
      description: "The availability of settings for the Action Running Message"
      type: ActionSettingsLockRunningMessage
    TimeRange:
      title: "Time Range"
      description: "If the Action has a Time Range option"
      type: boolean
      required: false
    StartDateTimeOffset:
      title: "StartDate Time Offset"
      description: "If the Action has a Start Date Time Offset option"
      type: boolean
      required: false
    EndDateTimeOffset:
      title: "EndDate Time Offset"
      description: "If the Action has a End Date Time Offset option"
      type: boolean
      required: false
    DayOfWeekConstraint:
      title: "Day of Week Constraint"
      description: "If the Action has a Day of the Week Constraint option"
      type: boolean
      required: false
    ActiveUserRequirement:
      title: "Active User Requirement"
      description: "If the Action has an Active User Requirement option"
      type: boolean
      required: false
    ActiveUserType:
      title: "Active User Type"
      description: "If the Action has an Active User Type option"
      type: boolean
      required: false
    Whose:
      description: "If the Action has a Whose setting option"
      type: boolean
      required: false
    PreActionCacheDownload:
      title: "PreAction Cache Download"
      description: "If the Action has a PreAction Cache Download setting option"
      type: boolean
      required: false
    Reapply:
      description: "If the Action has a Reapply setting option"
      type: boolean
      required: false
    ReapplyLimit:
      title: "Reapply Limit"
      description: "If the Action has a Reapply Limit setting option"
      type: boolean
      required: false
    ReapplyInterval:
      title: "Reapply Interval"
      description: "If the Action has a Reapply Interval setting option"
      type: boolean
      required: false
    RetryCount:
      title: "Retry Count"
      description: "If the Action has a Retry Count setting option"
      type: boolean
      required: false
    RetryWait:
      title: "Retry Wait"
      description: "If the Action has a Retry Wait setting option"
      type: boolean
      required: false
    TemporalDistribution:
      title: "Temporal Distribution"
      description: "If the Action has a Temporal Distribution setting option"
      type: boolean
      required: false
    ContinueOnErrors:
      title: "Continue On Errors"
      description: "If the Action has a Continue On Errors setting option"
      type: boolean
      required: false
    PostActionBehaviour:
      title: "PostAction Behaviour"
      description: "The availability of settings for the PostAction Behaviour"
      type: ActionSettingsLockPostActionBehavior
    IsOffer:
      title: "Is Offer"
      description: "If the Action has an IsOffer setting option"
      type: boolean
      required: false
    AnnounceOffer:
      title: "Announce Offer"
      description: "If the Action has an Announce Offer setting option"
      type: boolean
      required: false
    OfferTitle:
      title: "Offer Title"
      description: "The Offer Title of the Action" 
      type: string
      required: false
    OfferCategory:
      title: "Offer Category"
      description: "The Offer Category of the Action" 
      type: string
      required: false
    OfferDescriptionHTML:
      title: "Offer Description HTML"
      description: "The HTML of the Offer Description for the Action" 
      type: string
      required: false

  ActionSettings:
    HasMessage:
      title: "Has Message"
      description: "If the Action has a message" 
      type: boolean
      required: false
    Message:
      description: "The settings for the message of the Action"
      type: ActionSettingsMessage
    ActionUITitle:
      title: "Action UI Title"
      description: "The Action UI title" 
      type: string
      required: false
    PreActionShowUI:
      title: "PreAction Show UI"
      description: "If the PreAction shows UI elements" 
      type: boolean
      required: false
    PreAction:
      description: "The settings for the PreAction of the Action"
      type: ActionSettingsPreAction
    HasRunningMessage:
      title: "Has Running Message"
      description: "If the Action has a Running message" 
      type: boolean
      required: false
    RunningMessage:
      title: "Running Message"
      description: "The settings for the Running message of the Action"
      type: ActionSettingsRunningMessage
    HasTimeRange:
      title: "Has Time Range"
      description: "If the Action has a time range" 
      type: boolean
      required: false
    TimeRange:
      title: "Time Range"
      description: "The settings for the time range of the action"
      type: ActionSettingsTimeRange
    HasStartTime:
      title: "Has Start Time"
      description: "If the Action has a start time" 
      type: boolean
      required: false
    StartDateTimeOffset:
      title: "StartDate Time Offset"
      description: "The offset for the Action start date / time. The value must follow the ISO-8601 standard for duration format e.g PT15M for 15 minutes or P2D for 2 days"
      type: string
      required: false
    StartDateTimeLocalOffset:
      title: "StartDate Time Local Offset"
      description: "The local offset for the Action start date / time. The value must follow the ISO-8601 standard for duration format e.g PT15M for 15 minutes or P2D for 2 days"
      type: string
      required: false
    HasEndTime:
      title: "Has End Time"
      description: "If the Action has an end time" 
      type: boolean
      required: false
    EndDateTimeOffset:
      title: "EndDate Time Offset"
      description: "The offset for the Action end date / time. The value must follow the ISO-8601 standard for duration format e.g PT15M for 15 minutes or P2D for 2 days"
      type: string
      required: false
    EndDateTimeLocalOffset:
      title: "EndDate Time Local Offset"
      description: "The local offset for the Action end date / time. The value must follow the ISO-8601 standard for duration format e.g PT15M for 15 minutes or P2D for 2 days"
      type: string
      required: false
    HasDayOfWeekConstraint:
      title: "Has Day of Week Constraint"
      description: "If the Action has a Day of the Week Constraint" 
      type: boolean
      required: false
    DayOfWeekConstraint:
      title: "Day of Week Constraint"
      description: "The Day of the Week Constraint(s) for the Action"
      type: ActionSettingsDayOfWeekConstraint
    UseUTCTime:
      title: "Use UTC Time"
      description: "If the Action uses UTC time" 
      type: boolean
      required: false
    ActiveUserRequirement:
      title: "Active User Requirement"
      description: "The Active User Requirement setting for the Action" 
      type: string
      enum:
      - "NoRequirement"
      - "RequireUser"
      - "RequireNoUser"
    ActiveUserType:
      title: "Active User Type"
      description: "The Active User Type of the Action" 
      type: string
      enum:
      - "AllUsers"
      - "LocalUsers"
      - "UsersOfGroups"
    UIGroupConstraints:
      title: "UI Group Constraints"
      description: "The UI Group Constraints for the Action"
      type: ActionSettingsUIGroupConstraints
    HasWhose:
      title: "Has Whose"
      description: "If the Action has a Whose setting" 
      type: boolean
      required: false
    Whose:
      description: "The Whose setting for the Action"
      type: ActionSettingsWhose
    PreActionCacheDownload:
      title: "PreAction Cache Download"
      description: "If the Action has a PreAction cache download" 
      type: boolean
      required: false
    Reapply:
      description: "If the Action can be reapplied"
      type: boolean
      required: false
    HasReapplyLimit:
      title: "Has Reapply Limit"
      description: "If the Action has a Reapply limit" 
      type: boolean
      required: false
    ReapplyLimit:
      title: "Reapply Limit"
      description: "The Reapply limit of the Action" 
      type: integer
      required: false
    HasReapplyInterval:
      title: "Has Reapply Interval"
      description: "If the Action has a Reapply interval" 
      type: boolean
      required: false
    ReapplyInterval:
      title: "Reapply Interval"
      description: "The Reapply interval of the Action" 
      type: ActionReapplyInterval
      required: false
    HasRetry:
      title: "Has Retry"
      description: "If the Action can be retried" 
      type: boolean
      required: false
    RetryCount:
      title: "Retry Count"
      description: "How many times the Action can be retried"
      type: integer
      required: false
    RetryWait:
      title: "Retry Wait"
      description: "The Action Retry Wait settings"
      type: ActionSettingsRetryWait
    HasTemporalDistribution:
      title: "Has Temporal Distribution"
      description: "If the Action has a Temporal Distribution setting"
      type: boolean
      required: false
    TemporalDistribution:
      title: "Temporal Distribution"
      description: "The Temporal Distribution of the Action. The value must be non-negative and follow the ISO-8601 standard for duration format e.g PT15M for 15 minutes or P2D for 2 days"
      type: string
      required: false
    ContinueOnErrors:
      title: "Continue On Errors"
      description: "If the Action will continue after an error" 
      type: boolean
      required: false
    PostActionBehavior:
      title: "PostAction Behaviour"
      description: "The Action PostAction Behaviour settings" 
      type: ActionSettingsPostActionBehavior
    IsOffer:
      title: "Is Offer"
      description: "If the Action is an Offer" 
      type: boolean
      required: false
    AnnounceOffer:
      title: "Announce Offer"
      description: "If the Action Offer is announced"
      type: boolean
      required: false
    OfferCategory:
      title: "Offer Category"
      description: "The Offer Category of the Action"
      type: string
      required: false
    OfferDescriptionHTML:
      title: "Offer Description HTML"
      description: "The HTML description of the Actions offer" 
      type: string
      required: false

  FixletAction:
    Description:
      description: "The description of the Fixlet Action"
      type: FixletActionDescription
    ActionScript:
      type: FixletActionActionScript
      description: "The ActionScript of the Fixlet Action"
      required: true
    SuccessCriteria:
      title: "Success Criteria"
      type: ActionSuccessCriteria
      description: "The success criteria for the Fixlet Action - how the Fixlet determines a successful application of the Action"
    SuccessCriteriaLocked:
      title: "Success Criteria Locked"
      type: boolean
      description: "If the success criteria for the Fixlet Action is locked or changeable"
    Settings:
      type: ActionSettings
      description: "The settings of the Fixlet Action"
    SettingsLocks:
      title: "Settings Locks"
      description: "The available settings for the Fixlet Action"
      type: ActionSettingsLocks
    ID:
      type: string
      description: "The ID of the Fixlet Action"
      required: true

  # Main XML Types

  SourcedFixletAction:
    SourceFixlet:
      title: "Source Fixlet"
      type: BESActionSourceFixlet
    Target:
      type: BESActionTarget
    Settings:
      type: ActionSettings

  Computer:
    ID:
      type: string
    LastReportTime:
      title: "Last Report Time"
      type: date

  Setting:
    Name:
      type: string
    Value:
      type: string

  Property:
    Name:
      type: string
    Value:
      type: string

  Evaluation:
    Time:
      type: integer
    Plurality:
      type: string

  Answer:
    Type:
      type: string
    Value:
      type: string

  Answers:
    Answers:
      type: "[]Answer"

  QueryResponse:
    Resource:
      type: string
    Result:
      type: "[]Answers"
    Evaluation:
      type: Evaluation

  SiteDescription:
    Name:
      type: string
    Type:
      description: "The type of the site e.g master; external; operator"
      type: string
    GatherURL:
      type: string
    Description:
      type: string
    Domain:
      type: string
    GlobalReadPermission:
      type: boolean
    Subscription:
      type: Subscription
    Masthead:
      type: string

  FixletWithActions:
    Title:
      type: string
      required: true
    Description:
      type: string
      required: true
    Relevance:
      description: "The Fixlet Relevance string for matching with suitable Computers"
      type: "[]string"
      required: false
    GroupRelevance:
      title: "Group Relevance"
      description: "The Fixlet Group Relevance for matching with groups of Computers" 
      type: FixletWithActionsGroupRelevance
    Category:
      description: "The category of the Fixlet e.g Security Update" 
      type: string
      required: false
    WizardData:
      title: "Wizard Data"
      type: FixletWithActionsWizardData
    DownloadSize:
      title: "Download Size"
      type: integer
      required: false
    Source:
      type: string
      required: false
    SourceID:
      title: "Source ID"
      type: string
      required: false
    SourceReleaseDate:
      title: "Source Release Date"
      type: string
      required: false
    SourceSeverity:
      title: "Source Severity"
      type: string
      required: false
    CVENames:
      title: "CVE Names"
      description: "The CVEs associated with this Fixlet" 
      type: string
      required: false
    SANSID:
      type: string
      required: false
    MIMEFields:
      type: "[]MIMEField"
      required: false
    Domain:
      type: string
      required: false
    Delay:
      description: "The value must be non-negative and follow the ISO-8601 standard for duration format e.g PT15M for 15 minutes or P2D for 2 days"
      type: string
      required: false
    SkipUI:
      type: boolean
      required: false
    DefaultAction:
      type: FixletAction
    Action:
      type: "[]FixletAction"

  ComputerGroup:
    Title:
      type: string
      description: "The title of the Computer Group"
      required: true
    Domain:
      type: string
      description: "The domain of the Computer Group in BigFix"
    JoinByIntersection:
      title: "Join By Intersection"
      type: boolean
      description: "When provided with more than one Relevance, determines how the results are chosen"
      required: true
    IsDynamic:
      title: "Is Dynamic"
      type: boolean
      description: "Determines if the size or makeup of the Computer Group can changed based on a condition"
    EvaluateOnClient:
      title: "Evaluate On Client"
      type: boolean
      description:
    ComputerGroupRelevance:
      title: "Computer Group Relevances"
      type: "[]ComputerGroupRelevance"
      description: "The collection of Computer Group Relevances describing how Computers in this Computer Group are related"
      required: true
    SkipUI:
      type: boolean
      description:

  SiteDefinition:
    Type:
      type: string
      description: "The type of site in BigFix e.g external, master or custom. These are set values within BigFix and can also be found in the sites URL"
      required: true
      enum:
      - "master"
      - "custom"
      - "external"
      - "operator"
      - "action"
    Name:
      type: string
      required: true
    Resource:
      type: string
      required: false
    GatherURL:
      type: string

  FixletDefinition:
    Name:
      type: string
      required: false
    ID:
      type: integer
      required: false
    Resource:
      type: string
      required: true
    LastModified:
      type: date
      required: true

  TaskDefinition:
    Name:
      type: string
      required: false
    ID:
      type: integer
      required: false
    Resource:
      type: string
      required: true
    LastModified:
      type: date
      required: true

  BaselineDefinition:
    Name:
      type: string
      required: false
    ID:
      type: integer
      required: false
    Resource:
      type: string
      required: true
    LastModified:
      type: date
      required: true

  AnalysisDefinition:
    Name:
      type: string
      required: false
    ID:
      type: integer
      required: false
    Resource:
      type: string
      required: true
    LastModified:
      type: date
      required: true

  ComputerGroupDefinition:
    Name:
      type: string
      required: false
    ID:
      type: integer
      required: false
    Resource:
      type: string
      required: true
    LastModified:
      type: date
      required: true

  ActionDefinition:
    Resource:
      type: string
      required: true
    LastModified:
      type: date
      required: true
    Name:
      type: string
      required: false
    ID:
      type: integer
      required: false

  FieldMatcher:
    FieldName:
      title: "Field Name"
      type: string
      description: "The field name to match the provided values on. Must match the field names returned from the BigFix API"
      required: true
    Matcher:
      type: string
      description: "The matcher describing the relationship between the field name and the values provided"
      required: true
      enum:
      - "contains"
      - "doesNotContain"
      - "equals"
      - "doesNotEqual"
      - "exists"
      - "doesNotExist"
      - "isEmpty"
      - "isNotEmpty"
    Value:
      type: "[]string"
      description: "The values that a Fixlet's chosen field must match to be returned from the Action. Each field can be matched to a set of values. Only one value in the array needs to match. If not provided, the matcher should be one of isEmpty / isNotEmpty / doesNotExist"
      required: false

  FixletActionSummary:
    Name:
      type: string
      description: "The name of the Fixlet"
    SiteName:
      title: "Site Name"
      type: string
      description: "The site to which the Fixlet belongs"
    Id:
      type: string
      description: "The Fixlet's numerical ID"
      title: "ID"
    DefaultActionId:
      title: "Default Action ID"
      type: string
      description: "The ID of the Fixlet's default action"

  MultiActionGroup:
    Title:
      type: string
      description: "The title of the MultiAction Group to be displayed in BigFix"
    SourceFixlets:
      title: "Source Fixlets"
      type: "[]FixletActionSummary"
      description: "The Fixlets to use as part of the MultiAction Group"
    Target:
      type: BESActionTarget
      description: "The Target(s) of the MultiAction Group"
    Settings:
      type: ActionSettings
      description: "The Action settings applied to the whole of the MultiAction Group"

  Baseline:
    Title:
      type: string
      description: "The title of the Baseline to be displayed in BigFix"
    Description:
      type: string
      description: "The description of the Baseline to be displayed in BigFix"
    Relevance:
      type: string
      description: "The relevance string that will add Fixlets to the Baseline"
    SourceFixlets:
      title: "Source Fixlets"
      type: "[]FixletActionSummary"
      description: "The Fixlets to use as part of the Baseline"

  inventory_host:
    os:
      type: string
      title: "OS"
      description: "Operating system information"
    ip:
      type: string
      title: "IP Address"
      description: "IP address of host"
    hostname:
      type: string
      title: "Hostname"
      description: "Hostname"

actions:
  get_computers_for_fixlet:
    name: "Get Computers for Fixlet"
    description: "Fetch a list of computers which are relevant to a particular Fixlet"
    input:
      fixlet_id:
       title: "Fixlet ID"
       type: string
       description: "The ID for the Fixlet"
       required: true
      site_details:
        title: "Site Details"
        type: SiteDefinition
        description: "The details of the site to which the Fixlet belongs"
        required: true
    output:
      computers:
        type: "[]string"
        description: "The computer IDs which are relevant to the Fixlet"
        required: true

  fetch_relevant_fixlets:
    name: "Fetch Relevant Fixlets"
    description: "Fetch a list of Fixlets which are relevant to provided criteria"
    input:
      site:
        type: string
        description: "The site to which the Fixlets belong"
        required: false
      titles:
        type: "[]string"
        description: "The titles to which the Fixlets must match"
        required: false
      computer_groups:
        type: "[]string"
        description: "The computer groups to which the Fixlets are relevant"
        required: false
      cves:
        title: "CVEs"
        type: "[]string"
        description: "The CVEs to which the Fixlets are relevant"
        required: false
    output:
      results:
        type: "[]FixletActionSummary"
        description: "The relevant Fixlets retrieved from BigFix"
        required: true

  login:
    name: "Login"
    description: "Create or refresh connection"
    input:
    output:
      success:
        type: boolean
        description: "Whether the login request succeeded"
        required: true

  fetch_computers:
    name: "Fetch Computers"
    description: "Fetch a list of computers"
    input:
    output:
      computers:
        type: "[]Computer"
        description: "List of computer identifiers returned from BigFix"
        required: true

  get_settings:
    name: "Get Settings"
    description: "Get a computer's settings"
    input:
      computer_id:
        title: "Computer ID"
        type: string
        description: "The computer ID"
        required: true
    output:
      settings:
        type: "[]Setting"
        description: "List of settings for the chosen Computer"
        required: true

  query:
    name: "Query"
    description: "Evaluate a Relevance expression and get the result"
    input:
      query_string:
        title: "Query"
        type: string
        description: "The Relevance expression to evaluate"
        required: true
    output:
      results:
        title: "Query Results"
        type: QueryResponse
        description: "The results from the evaluated expression"
        required: true

  query_host_inventory:
    name: "Query Host Inventory"
    description: "Get an inventory of hosts, including their hostnames, IP addresses, and operating system"
    output:
      results:
        title: "Hosts"
        type: "[]inventory_host"
        description: "Hosts found in the BigFix inventory"
        required: false

  query_host_for_installed_software:
    name: "Query Host for Installed Software"
    description: "Query a host for all installed software"
    input:
      hostname:
        title: "Hostname"
        description: "Hostname of the host to query"
        type: string
        required: true
    output:
      results:
        title: "Software"
        type: "[]string"
        description: "Software packages installed on the host"
        required: false

  query_host_for_specific_software:
    name: "Query Host for Specific Software"
    description: "Query a host to see if a specific software package is installed"
    input:
      hostname:
        title: "Hostname"
        description: "Hostname of the host to query"
        type: string
        required: true
      software:
        title: "Software"
        description: "Installed software to query the host for"
        type: string
        required: true
    output:
      found:
        title: "Found"
        type: boolean
        description: "Whether or not the host has the specified software installed"
        required: true



  custom_endpoint:
    name: "Custom Endpoint Request"
    description: "Execute a request against a custom endpoint"
    input:
      endpoint:
        type: string
        description: "The endpoint against which the request will be executed e.g. 'computers'"
        required: true
      method:
        type: string
        description: "The method to use for the request"
        required: true
        enum:
          - "DELETE"
          - "GET"
          - "PATCH"
          - "POST"
      parameters:
        type: object
        description: "The URL parameters to include in the request"
        required: false
      data:
        type: string
        description: "The XML data which will be attached to the request"
        required: false
    output:
      response:
        type: string
        description: "The XML response returned from the endpoint"
        required: true

  fetch_computer_groups:
    name: "Fetch Computer Groups"
    description: "Fetch a list of Computer Groups for a provided Site"
    input:
      chosen_site:
        title: "Chosen Site"
        type: SiteDefinition
        description: "Details of the site from which to retrieve the Computer Groups"
        required: true
    output:
      computer_groups:
        type: "[]ComputerGroupDefinition"
        description: "The computer groups belonging to the specified site"
        required: true

  fetch_group_members:
    name: "Fetch Group Members"
    description: "Fetch a list of computers relevant to a computer group"
    input:
      chosen_site:
        title: "Chosen Site"
        type: SiteDefinition
        description: "Details of the site to which the Computer Group belongs"
        required: true
      id:
        title: "ID"
        type: string
        description: "The ID of the computer group"
        required: false
    output:
      computers:
        type: "[]string"
        description: "The resource links for the computers within the computer group"
        required: true

  get_computer:
    name: "Get Computer"
    description: "Get core properties of a computer"
    input:
      id:
        type: string
        title: "ID"
        description: "The ID of the computer"
        required: false
    output:
      computer:
        type: ComputerCoreProperties
        description: "The core properties of the computer"
        required: true

  fetch_fixlets:
    name: "Fetch Fixlets"
    description: "Fetch information on all Fixlets for a provided Site"
    input:
      chosen_site:
        title: "Chosen Site"
        type: SiteDefinition
        description: "Details of the site from which to retrieve the Fixlets"
        required: true
    output:
      fixlets:
        title: "Fixlet Definitions"
        type: "[]FixletDefinition"
        description: "List of Fixlets from the provided Site"
        required: true

  fetch_sites:
    name: "Fetch Sites"
    description: "Fetch a list of Sites and their types available to the authenticated operator"
    input:
    output:
      sites:
        title: "Site Definitions"
        description: "The current list of Sites on the Server"
        type: "[]SiteDefinition"
        required: true

  get_site_content:
    name: "Get Site Content"
    description: "Gets identifiers for all content from a site - Fixlets, Tasks, Baselines, Analysis and Computer Groups"
    input:
      chosen_site:
        title: "Chosen Site"
        type: SiteDefinition
        description: "Details of the site from which to retrieve the content"
        required: true
    output:
      fixlets:
        type: "[]FixletDefinition"
        description: "The current list of Fixlets on the Site"
        required: true
      tasks:
        type: "[]TaskDefinition"
        description: "The current list of Tasks on the Site"
        required: true
      baselines:
        type: "[]BaselineDefinition"
        description: "The current list of Baselines on the Site"
        required: true
      analyses:
        type: "[]AnalysisDefinition"
        description: "The current list of Analyses on the Site"
        required: true
      computergroups:
        title: "Computer Groups"
        type: "[]ComputerGroupDefinition"
        description: "The current list of Computer Groups on the Site"
        required: true

  get_site:
    name: "Get Site"
    description: "Fetch a specific site"
    input:
      chosen_site:
        title: "Chosen Site"
        type: SiteDefinition
        description: "Details for the site"
        required: true
    output:
      site:
        title: "Returned Site"
        type: SiteDescription
        description: "Details of the selected Site"
        required: true

  get_fixlet:
    name: "Get Fixlet"
    description: "Fetches a specific Fixlet from a Site"
    input:
      chosen_site:
        title: "Chosen Site"
        type: SiteDefinition
        description: "Details of the site from which to retrieve a specific Fixlet"
        required: true
      id:
        title: "Fixlet ID"
        type: string
        description: "Identifier for the Fixlet"
        required: true
    output:
      fixlet:
        title: "Returned Fixlet"
        type: FixletWithActions
        description: "The information retrieved about the Fixlet"
        required: true

  get_matching_fixlets:
    name: "Get Matching Fixlets"
    description: "Return a list of Fixlets from a Site with properties matching input criteria"
    input:
      chosen_site:
        title: "Chosen Site"
        type: SiteDefinition
        description: "Details of the site from which to retrieve matching Fixlets"
        required: true
      criteria:
        title: "Search Criteria"
        type: "[]FieldMatcher"
        description: 'The criteria used to match relevant Fixlets. All provided criteria must match for a Fixlet ID to be returned. Each criterion should be in the form {"FieldName": "a", "Matcher": "b", "Value": ["c", "d"]}, creating a JSON array'
        required: true
    output:
      matching_fixlet_ids:
        title: "Matching Fixlet IDs"
        type: "[]string"
        description: "The IDs of Fixlets matching the provided criteria"
        required: true

  create_computer_group:
    name: "Create Computer Group"
    description: "Create a Computer group in a specified site"
    input:
      chosen_site:
        title: "Chosen Site"
        type: SiteDefinition
        description: "Details of the site where the Computer Group will be created"
        required: true
      computer_group:
        title: "ComputerGroup"
        type: ComputerGroup
        description: "The description of the Computer Group to create in BigFix"
        required: true
    output:
      computer_group:
        title: "Computer Group"
        type: ComputerGroupDefinition
        description: "Information about the newly created Computer Group"
        required: true

  delete_computer_group:
    name: "Delete Computer Group"
    description: "Delete a Computer group in a specified site"
    input:
      chosen_site:
        title: "Chosen Site"
        type: SiteDefinition
        description: "Details of the site containing the Computer Group to be deleted"
        required: true
      id:
        title: "Computer Group ID"
        type: string
        description: "The ID of the Computer Group to delete"
        required: true
    output:
      result:
        type: boolean
        description: "The result of the operation"
        required: true
      message:
        type: string
        description: "The message returned from the BigFix console"
        required: false

  get_actions:
    name: "Get Actions"
    description: "Get all Action summaries from BigFix"
    input:
    output:
      actions: 
        type: "[]ActionDefinition"
        description: "Information on the Actions present in BigFix"
        required: true

  get_action_status:
    name: "Get Action Status"
    description: "Get an Action status from BigFix"
    input:
      action_id:
        title: "Action ID"
        type: integer
        description: "The Action ID"
        required: true
    output:
      status:
        title: "Action Status"
        type: string
        description: "The status of the Action"
        required: true

  delete_action:
    name: "Delete Action"
    description: "Delete an Action from BigFix"
    input:
      action_id:
        title: "Action ID"
        type: integer
        description: "The Action ID"
        required: true
    output:
        response:
          title: "Response Message"
          type: string
          description: "The response returned from BigFix"
          required: true

  stop_action:
    name: "Stop Action"
    description: "Stop an Action in BigFix"
    input:
      action_id:
        title: "Action ID"
        type: integer
        description: "The Action ID"
        required: true
    output:
        response:
          title: "Response Message"
          type: string
          description: "The response returned from BigFix"
          required: true

  create_action:
    name: "Create an Action"
    description: "Create an Action in BigFix"
    input:
      SourcedFixletAction:
        title: "Sourced Fixlet Action"
        type: SourcedFixletAction
        description: "The Action (Site, Fixlet ID, Action, Settings) to create"
        required: false
    output:
      action:
        type: ActionDefinition
        description: "Details of the newly created Action"
        required: true

  create_multiaction_group:
    name: "Create MultiAction Group"
    description: "Create a MultiAction Group in BigFix"
    input:
      multiaction_group:
        title: "Multi-Action Group"
        type: MultiActionGroup
        description: "The Actions (Site, Fixlet ID, Action) to create along with the Target and Settings"
        required: false
    output:
      action:
        type: ActionDefinition
        description: "Details of the newly created MultiAction Group"
        required: true

  create_baseline:
    name: "Create Baseline"
    description: "Create a Baseline in BigFix"
    input:
      baseline:
        title: "Baseline"
        type: Baseline
        description: "The Site, Actions, Fixlet IDs and Settings to create a Baseline"
        required: false
    output:
      baseline_details:
        type: BaselineDefinition
        description: "Details of the newly created Baseline"
        required: true
`
