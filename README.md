# trail

CLI to run your dataframes against different services.

## Dependencies

TODO: Add `***REMOVED***` installation here

## Usage

```
go install trail/*.go
```

```
Usage:
  trail [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  follow      Follow a path for a dataframe
  help        Help about any command

Flags:
  -h, --help   help for trail
```

#### Follow call dataframe against SLU service

```
trail follow --input-proto records.pb
```
