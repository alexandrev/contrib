<!--
title: Log
weight: 4615
-->

#Flow Context

Returns the context of running flow
## Installation

### Flogo CLI
```bash
flogo install github.com/alexandrev/contrib/activity/flowcontext
```

## Configuration

### Input:
| Name       | Type   | Description
|:---        | :---   | :---    
| message    | string | The message to log
| addDetails | bool   | Append contextual execution information to the log message

## Examples
The below example logs a message 'test message':

```json
{
  "id": "log_message",
  "name": "Log Message",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/flowcontext",
    "input": {
      "message": "test message",
      "addDetails": "false"
    }
  }
}
```