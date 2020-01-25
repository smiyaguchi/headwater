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
If generate history data, please add `from` and `to` field to schema file.

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
Generate test data, use `-m` option.

```
# generate history test data
$ hw gen -m "HISTORY"
```
