# xodr - Work in Progress
[![Go Reference](https://pkg.go.dev/badge/github.com/christianbirchler-org/go-xodr.svg)](https://pkg.go.dev/github.com/christianbirchler-org/go-xodr)

This small package provide functions to converts road descriptions in the OpenDRIVE format to sequence of xy coordinates.

## Links
- [Overview of OpenDRIVE XML elements](https://publications.pages.asam.net/standards/ASAM_OpenDRIVE/ASAM_OpenDRIVE_Specification/latest/specification/06_general_architecture/06_05_overview_elements.html)

## Contribution
### Generate OpenDRIVE structs
```{go}
go generate
```

### TODO
- Generate structs
- Define attributes

## License
[MIT License](./LICENSE)