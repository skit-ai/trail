---
# Feel free to add content and custom Front Matter to this file.
# To modify the layout, see https://jekyllrb.com/docs/themes/#overriding-theme-defaults

layout: default
title: Home
nav_order: 1
---

# Trail

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/skit-ai/trail?style=flat-square)
![Binary release](https://github.com/skit-ai/trail/actions/workflows/release.yml/badge.svg)

[Trail][trail] is a CLI to run tagged and untagged [dataframes][dataframes]
against different services (currently, [SLU service][slu-service]).

It works on calls/turn data as defined in
[dataframes](https://github.com/skit-ai/dataframes) that has standard datatype
definitions.

## Installation

Get the latest binaries from the releases [here][binaries]. Choose the
appropriate binary based on your machine's OS and architecture.

## Usage

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
      --type string                  Type of record. One of: [tagged, untagged] (optional) (default "untagged")
```

## Usage patterns

##### Export both intents and entities labels in CSV

```
trail follow --input-csv records.csv --slu-host http://localhost:6969 --slu-language en --slu-client client_slu --output-intents-csv ./intents.csv --output-entities-csv ./entities.csv --concurrency 30
```

##### Export either of intents or entities labels in CSV

Provide `--output-intents-csv` or `--output-entities-csv` as appropriate.

```
trail follow --input-csv records.csv --slu-host http://localhost:6969 --slu-language en --slu-client client_slu --output-intents-csv ./intents.csv --output-entities-csv ./entities.csv --concurrency 30
```

##### Using tagged or untagged dataframe

Use `--type` flag to specify type of input dataframe. Type is one of `tagged`
or `untagged`. Default: `untagged`

```
trail follow --input-csv records.csv --slu-host http://localhost:6969 --slu-language en --slu-client client_slu --output-intents-csv ./intents.csv --output-entities-csv ./entities.csv --concurrency 30 --type tagged
```

##### Export SLU service response and generate your own label files

```
trail follow --input-csv records.csv --slu-host http://localhost:6969 --slu-language en --slu-client client_slu --output-intents-csv ./intents.csv --output-entities-csv ./entities.csv --concurrency 30 | ./scripts/gen_label_files.py
```

Sample Python script to generate label files by piping response from the command at [gen_label_files.py][gen-labels]


## Development

- Install dependencies

```
go mod download
```

- Build binary

```
go install trail/*.go
```

[trail]: https://github.com/skit-ai/trail
[dataframes]: https://github.com/skit-ai/dataframes
[slu-service]: https://github.com/skit-ai/slu-service/
[binaries]: https://github.com/skit-ai/trail/releases/
[gen-labels]: ./scripts/gen_label_files.py
