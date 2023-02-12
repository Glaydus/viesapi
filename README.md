# VIES API Client for Go

[![Release](https://img.shields.io/github/v/release/glaydus/viesapi)](https://github.com/glaydus/viesapi/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/glaydus/viesapi.svg)](https://pkg.go.dev/github.com/glaydus/viesapi)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/glaydus/viesapi)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This is the unofficial repository for VIES API Client for Go: https://viesapi.eu

Viesapi.eu service provides selected entrepreneurs data using i.a. web services, programming libraries and dedicated applications.
By using the available software (libraries, applications and Excel add-in) your customers will be able to:

* check contractors EU VAT number status in VIES system,
* download company details from VIES system,
* automatic fill in the invoice forms,

in the fastest possible way.

## Documentation

The documentation and samples are available at https://viesapi.eu/docs/

### Go modules

If your application uses Go modules for dependency management (recommended), add an import for each service that you use in your application.

Example:

```go
import (
  "github.com/glaydus/viesapi"
)
```

Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's `go.mod` file.

### `go get` command

Alternatively, you can use the `go get` command to download and install the appropriate packages that your application uses:

```sh
go get -u github.com/glaydus/viesapi
```


## License

This project is delivered under the Apache 2.0 license. You can find the license's full text in [LICENSE](LICENSE).