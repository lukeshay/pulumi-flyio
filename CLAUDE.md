# Pulumi Fly.io Provider Development Guide

## Build Commands
- Full build: `make build`
- Provider only: `make provider`
- Generate Fly.io API code: `make gen`
- Run all tests: `make test_all`
- Run provider tests: `make test_provider`
- Lint code: `make lint`
- Install local build: `make install`

## Testing
- Run a specific test: `cd tests && go test -v -count=1 -run TestNameHere`
- Test with examples: `cd examples/yaml && pulumi up`

## Code Style Guidelines
- **Formatting**: Go standard format (`go fmt`)
- **Naming**: Use camelCase for variables, PascalCase for functions and types
- **Error Handling**: Always check errors and return meaningful messages
- **Annotations**: All resources, args, and state structs must use `infer.Annotated` with descriptive comments
- **Dependencies**: Keep external dependencies to a minimum
- **Imports**: Group standard library, external packages, and internal packages

## Pulumi Resource Pattern
- Define resource struct with args and state
- Implement required interfaces (Create, Delete, etc.)
- Add proper annotations for schema documentation
- Handle preview mode correctly in resource implementations