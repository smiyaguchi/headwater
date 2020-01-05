#!/bin/bash

go build -o hw

json=$(cat << EOS
[
  {
    "column":"column1",
    "type":"STRING",
    "unique":"true",
    "nullable":"false"
  },
  {
    "column":"column2",
    "type":"NUMERIC",
    "precision":"3",
    "scale":"1",
    "unique":"false",
    "nullable":"false"
  }
]
EOS
)

echo "$json" > schema.json
