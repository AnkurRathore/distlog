# distlog
Distributed Log commit using Go

## Running the program.
Main is located under cmd/server directory and can be run as follows:

go run main.go

## Testing it out
$ curl -X POST localhost:8000 -d \
'{"record": {"value": "TGV0J3MgR28gIzEK"}}'
$ curl -X POST localhost:8000 -d \
'{"record": {"value": "TGV0J3MgR28gIzIK"}}'
$ curl -X POST localhost:8000 -d \
'{"record": {"value": "TGV0J3MgR28gIzMK"}}'