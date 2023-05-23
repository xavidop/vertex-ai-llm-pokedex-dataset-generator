# Vertex AI Dataset Generator to fine-tune Palm 2 model

## Prerequisites

You will need to have the following items before starting:
1. Golang 1.19+

Execute the following command to download all the dependencies:
```shell
go mod tidy
```

## Generate Dataset

To generate the dataset, execute the following command:
```shell
go run main.go
```

The dataset will be generated in the `root` folder.

## Parameters

You can modify the following parameters in the `main.go` file:
```go
const pokemons = 151
const datasetFile = "dataset.jsonl"
```