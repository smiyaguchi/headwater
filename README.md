# headwater
`headwater` is generate test data for big data

## Building

```
git clone https://github.com/smiyaguchi/headwater
go get -u github.com/brianvoe/gofakeit
cd $GOPATH/src/github.com/smiyaguchi/headwater
make
```

## Usage
### Generate normal test data
`headwater` uses json format schema file. Default schema filename is `schema.json`.<br>
To change schema file, please use `-s` option.

```
# generate test data
$ hw gen

# generate test data by specifying the number of records
$ hw gen -c 100

# generate test data from specified schema file (default schema file is "schema.json")
$ hw gen -s schema_test.json
```

### Generate history test data
If generate history data, please add `history` field. Or `from` and `to`field.
When `history` field added, create history data for one column.

```
[
  {
    "name": "from_date",
    "description": "from date column",
    "type": "TIMESTAMP",
    "mode": "NULLABLE",
    "history": "true"
  }
]
```

```
# generate history test data for one column
$ hw gen -m "HISTORY"
```

On the other hand, when `from` and `to` field are added, the history data is created in the From-To format.

```
[
  {
    "name": "from_date",
    "description": "from date column",
    "type": "TIMESTAMP",
    "mode": "NULLABLE",
    "from": "true"
  },
  {
    "name": "to_date",
    "description": "to date column",
    "type": "TIMESTAMP",
    "mode": "NULLABLER",
    "to": "true"
  }
]
```
```
# generate history test data from-to format
$ hw gen -m "HISTORY"
```
