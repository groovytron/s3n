# s3n (Swiss Social Security Number)

`s3n` is a  golang library to validate and format swiss social security numbers (aka. *AVS* in french and *AHV* in german).

## Install

To install s3n, use `go get`:

```
go get github.com/groovytron/s3n
```

## Usage

```go
// Validate a swiss social security number
valid := s3n.IsValid("756.1234.5678.97")

// Not dotted numbers are validated
valid := s3n.IsValid("7561234567897")

dotted, err := s3n.DottedFormat("7561234567897") // "7561234567897" => "756.1234.5678.97"
dotless, err := s3n.DotlessFormat("756.1234.5678.97") // "756.1234.5678.97" => "7561234567897"
```
