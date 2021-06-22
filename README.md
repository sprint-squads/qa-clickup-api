# qa-clickup-api
REST api for creating Issues in clickup.com and uploading attachment files. Needs minio client for files to be uploaded 

**API URL**: ```http://84.252.128.7:4000/v1/```

There is a sample google-chrome extension using this api in the following repository:

```https://github.com/sprint-squads/qa-clickup-chrome-ext```


# Setup
Add the appropriate values to your .env file
```bash
HTTP_PORT=<server port>
HOST=<server host>
CLICKUP_URL=<full path to clickup api, default is https://api.clickup.com/api/v2>
CLICKUP_SPACE_ID=<space id>
CLICKUP_LIST_ID=<list id>
CLICKUP_ACCESS_TOKEN=<clickup access token>
MINIO_URL=<url of minio client, where to upload files>
MINIO_ACCESS_KEY=<minio client access key>
MINIO_SECRET_KEY=<minio client secret key>
MINIO_USE_SSL=<if use SSL>
MINIO_BUCKET_NAME=<minio bucket name>
```
# Usage

Run the following command to start web server
```bash
$ go run main.go serve
```

# Routes

## List of tags

List of tags of the space
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

Creates task with the optional fields and files from multipart form body
### Request

```POST /clickup/issues```

**Body**
#### Multipart form:
```bash
title string
description string
priority int (1-Low, 2-Normal, 3-High, 4-Urgent)
tags string (splitted by space)
file file (files to be uploaded into your minio client)
```

### Response

```json
{
  "code": 0,
  "message": "task created"
}
```
