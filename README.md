# go-xodr
[![Go Reference](https://pkg.go.dev/badge/github.com/christianbirchler-org/go-xodr.svg)](https://pkg.go.dev/github.com/christianbirchler-org/go-xodr)

This package provides APIs for road network specifications in the OpenDRIVE format.

## Links
- [Overview of OpenDRIVE XML elements](https://publications.pages.asam.net/standards/ASAM_OpenDRIVE/ASAM_OpenDRIVE_Specification/latest/specification/06_general_architecture/06_05_overview_elements.html)

## Contributing
### Generate OpenDRIVE structs
```{go}
go generate
```

### Outlook
- generate functions
- generate XML tags on structs
- test package with sample xodr files from ASAM

## License
[MIT](./LICENSE)
