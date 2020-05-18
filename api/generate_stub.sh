#!/usr/bin/env bash

swagger-codegen generate \
  --lang go-server \
  --output .. \
  --input-spec swagger.yaml \
  --config config.json \
  --git-user-id marcomicera \
  --git-repo-id sayhi
