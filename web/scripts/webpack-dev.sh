#!/usr/bin/env bash
set -eux

basePath="$GOPATH/src/github.com/dnote/dnote"
appPath="$basePath"/web

(
  cd "$appPath" &&
  PUBLIC_PATH=$PUBLIC_PATH \
  COMPILED_PATH=$COMPILED_PATH \
    "$appPath"/scripts/setup.sh &&

  BUNDLE_BASE_URL=$BUNDLE_BASE_URL
  ASSET_BASE_URL=$ASSET_BASE_URL \
  COMPILED_PATH=$COMPILED_PATH \
  PUBLIC_PATH=$PUBLIC_PATH \
  IS_TEST=true \
    node "$appPath"/scripts/placeholder.js &&

  "$appPath"/node_modules/.bin/webpack-dev-server\
    --env.isTest="$IS_TEST"\
    --config "$appPath"/webpack/dev.config.js
)
