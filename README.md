# qa-clickup-api
REST api for creating Issues in clickup.com

**API URL**: ```http://84.252.128.7:4000/v1/```

# Routes

## List of tags

### Request

```GET /clickup/tags```

### Response

```json
{
  "code": 0,
  "message": "",
  "tags": [
    {
      "name": "devops",
      "tag_fg": "#EA80FC",
      "tag_bg": "#EA80FC",
      "creator": 3807198
    },
    {
      "name": "design",
      "tag_fg": "#E50000",
      "tag_bg": "#E50000",
      "creator": 3807198
    }
  ]
}
```

## Create Issue

### Request

```POST /clickup/issues```

**Body**

```
multipart form:

title string
description string
priority int
tags string (splitted by space)
file file

```

### Response

```json
{
  "code": 0,
  "message": "task created"
}
```
