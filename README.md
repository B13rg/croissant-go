# 🥐 croissant-go

A lightweight Go library for working with the Croissant data format, designed to provide formatting, validation, querying, and resource management capabilities.

<details>
<summary><b>Table of Contents</b></summary>
<p>

- [🥐 croissant-go](#-croissant-go)
  - [Key Features](#key-features)
  - [Technical Overview](#technical-overview)
  - [Installation](#installation)
    - [Getting Started](#getting-started)
    - [Library Use](#library-use)
  - [Documentation and Additional Resources](#documentation-and-additional-resources)
    - [Links](#links)
    - [Alternatives](#alternatives)
  - [Development](#development)
    - [Setup](#setup)
    - [Tests](#tests)
  - [Contributing](#contributing)
  - [License](#license)

</p>
</details>

## Key Features

Work-in-Progress.

## Technical Overview

croissant-go is implemented in idiomatic Go, focusing on modular design, clear interfaces, and support for schema validation and dataset integration.

## Installation

```sh
git clone https://github.com/B13rg/croissant-go.git
cd croissant-go
go build ./...
```

### Getting Started

See available commands:

```
croissant-go help
```

### Library Use

```go
import github.com/B13rg/croissant-go/pkg/croissant

func loadFile(path string) error {
  // Load file
  dataset, err := croissant.NewDataSetFromPath(path)
  if err != nil {
    // error loading file
    return err
  }

  fmt.Print("dataset name: ", dataset.Name)
}
```

## Documentation and Additional Resources

### Links

- [Project Docs](docs/)
- [Croissant Data Format](https://github.com/mlcommons/croissant)

### Alternatives

- [croissant-python](https://github.com/mlcommons/croissant)
- [Other MLCommons libraries](https://github.com/mlcommons)

## Development

### Setup

Install `go-task`:

```
brew install go-task
```

Setup dependencies:

```
task setup
```

### Tests

```
task test
```

## Contributing

Contributions are welcome! Please open issues or submit pull requests with your improvements.

## License

MIT License. See `LICENSE` for details.