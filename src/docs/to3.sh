#!/bin/bash

old_version_v2=$(< docs/v2/v2_swagger.json)

echo "$old_version_v2" | curl -X 'POST' 'https://converter.swagger.io/api/convert' -H 'accept: application/json' -H 'Content-Type: application/json' --data-binary @- > docs/v2/openapi3.json