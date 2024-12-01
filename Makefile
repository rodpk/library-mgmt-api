.PHONY: gen-oapi-spec
gen-oapi-spec:
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -package openapi -config api/oapi-gen.cfg.yml api/api-v1.yml