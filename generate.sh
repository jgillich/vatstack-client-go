#!/bin/sh

curl -O https://raw.githubusercontent.com/vatstack/openapi/master/vatstack.json
openapi-generator-cli generate --skip-validate-spec --git-repo-id vatstack-client-go --git-user-id jgillich -g go -o vatstack -i vatstack.json  --additional-properties=packageName=vatstack
