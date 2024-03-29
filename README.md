# gorm-rapier

[gorm-rapier](https://github.com/thinkgos/gorm-rapier) is an assist rapier for gorm.

[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/thinkgos/gorm-rapier?tab=doc)
[![codecov](https://codecov.io/gh/thinkgos/gorm-rapier/graph/badge.svg?token=aHu5wq1m6i)](https://codecov.io/gh/thinkgos/gorm-rapier)
[![Tests](https://github.com/thinkgos/gorm-rapier/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/thinkgos/gorm-rapier/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/gorm-rapier)](https://goreportcard.com/report/github.com/thinkgos/gorm-rapier)
[![Licence](https://img.shields.io/github/license/thinkgos/gorm-rapier)](https://raw.githubusercontent.com/thinkgos/gorm-rapier/main/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/thinkgos/gorm-rapier)](https://github.com/thinkgos/gorm-rapier/tags)

## Overview

- Idiomatic and Reusable API from Dynamic Raw SQL
- 100% Type-safe API without interface{}
- Almost supports all features, plugins, DBMS that GORM supports
- Almost same behavior as gorm you used.

## Usage

Use go get.

```bash
go get github.com/thinkgos/gorm-rapier
```

Then import the package into your own code.

```go
import "github.com/thinkgos/gorm-rapier"
```

## Guide

- [gorm Guide](https://gorm.io/docs/)
- [gorm rapier Guide](https://thinkgos.github.io/gorm-rapier)

## Example

- [create](./example_create_test.go): example create
- [query](./example_query_test.go): example query
- [advance query](./example_advance_query_test.go): example advance query
- [update](./example_update_test.go): example update
- [delete](./example_delete_test.go): example delete

## Reference

- [gorm](https://github.com:go-gorm/gorm)
- [sea-orm](https://github.com/SeaQL/sea-orm)
- [ent](https://github.com/ent/ent)

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.
