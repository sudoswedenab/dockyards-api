package main

//go:generate go tool oapi-codegen --config=cfg/types.yaml spec/types.yaml
//go:generate go tool oapi-codegen --config=cfg/backend.yaml spec/backend.yaml
//go:generate go tool oapi-codegen --config=cfg/client.yaml spec/backend.yaml
