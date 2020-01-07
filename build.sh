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
    "type":"INTEGER",
    "precision":"3",
    "scale":"1",
    "unique":"false",
    "nullable":"false"
  },
  {
    "column":"column3",
    "type":"FLOAT",
    "precision":"3",
    "scale":"1",
    "unique":"false",
    "nullable":"false"
  },
  {
    "column":"column4",
    "type":"NUMERIC",
    "precision":"3",
    "scale":"1",
    "unique":"false",
    "nullable":"false"
  },
  {
    "column":"column5",
    "type":"BOOLEAN",
    "precision":"3",
    "scale":"1",
    "unique":"false",
    "nullable":"false"
  },
  {
    "column":"column6",
    "type":"TIMESTAMP",
    "precision":"3",
    "scale":"1",
    "unique":"false",
    "nullable":"false"
  },
  {
    "column":"column7",
    "type":"TIMESTAMP",
    "precision":"3",
    "scale":"1",
    "unique":"false",
    "nullable":"false"
  },
  {
    "column":"column8",
    "type":"DATE",
    "precision":"3",
    "scale":"1",
    "unique":"false",
    "nullable":"false"
  },
  {
    "column":"column9",
    "type":"TIME",
    "precision":"3",
    "scale":"1",
    "unique":"false",
    "nullable":"false"
  },
  {
    "column":"column10",
    "type":"TIME",
    "precision":"3",
    "scale":"1",
    "unique":"false",
    "nullable":"false"
  },
  {
    "column":"column11",
    "type":"DATETIME",
    "precision":"3",
    "scale":"1",
    "unique":"false",
    "nullable":"false"
  }
]
EOS
)

echo "$json" > schema.json
