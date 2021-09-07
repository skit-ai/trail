# trail

CLI to run your dataframes against different services (currently, SLU service).

## Dependencies

- `go mod download`

TODO: Make it easier to setup and generate binaries

## Usage

```
go install trail/*.go
```

`trail follow --help`

```
Follow a dataframe and get predicted dataframes against your service config

Usage:
  trail follow <args> [flags]

Flags:
      --concurrency int              Max concurrent requests to SLU service (optional) (default 30)
  -h, --help                         help for follow
      --input-csv string             input csv file (required)
      --output-entities-csv string   output entities csv file
      --output-intents-csv string    output intents csv file
      --slu-client string            Name of the client (required)
      --slu-host string              http://host:port for SLU service (required)
      --slu-language string          Language code. Example: en, hi (required)
```

## Binaries

Get the latest binaries [here][binaries]

### Usage patterns

#### Follow call dataframe (CSV) against SLU service and export

##### Export both intents and entities labels in CSV

```
trail follow --input-csv records.csv --slu-host http://localhost:6969 --slu-language en --slu-client client_slu --output-intents-csv ./intents.csv --output-entities-csv ./entities.csv --concurrency 30
```

##### Export either of intents or entities labels in CSV

provide `--output-intents-csv` or `--output-entities-csv` as appropriate.

```
trail follow --input-csv records.csv --slu-host http://localhost:6969 --slu-language en --slu-client client_slu --output-intents-csv ./intents.csv --output-entities-csv ./entities.csv --concurrency 30
```


##### Export SLU service response and generate your own label files

```
trail follow --input-csv records.csv --slu-host http://localhost:6969 --slu-language en --slu-client client_slu --output-intents-csv ./intents.csv --output-entities-csv ./entities.csv --concurrency 30 | ./scripts/gen_label_files.py
```

Sample Python script to generate label files by piping response from the command at [gen_label_files.py][gen-labels]


[gen-labels]: ./scripts/gen_label_files.py
[binaries]: https://github.com/skit-ai/trail/releases/

