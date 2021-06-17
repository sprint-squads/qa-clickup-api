# qa-clickup-api
REST api for creating Issues in clickup.com

routes:
1. Get tags. http://84.252.128.7:4000/v1/clickup/tags

header:
Authorization: 3851228_2087f0167d551169256f2f86e8b21fec4bc90075

2. Create issues. http://84.252.128.7:4000/v1/clickup/issues

header:
Authorization: 3851228_2087f0167d551169256f2f86e8b21fec4bc90075

body (multipart form): 
title string;
description string;
priority int;
tags string (splitted by space);
file file;