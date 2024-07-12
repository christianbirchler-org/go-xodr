# xodr - Work in Progress
This small package provide functions to converts road descriptions in the OpenDRIVE format to sequence of xy coordinates.

## Links
- [Overview of OpenDRIVE XML elements](https://publications.pages.asam.net/standards/ASAM_OpenDRIVE/ASAM_OpenDRIVE_Specification/latest/specification/06_general_architecture/06_05_overview_elements.html)
- [OpenDRIVE Parametric Cubic Curve](https://publications.pages.asam.net/standards/ASAM_OpenDRIVE/ASAM_OpenDRIVE_Specification/latest/specification/09_geometries/09_06_param_poly3.html)

## Contribution
### Generate OpenDRIVE structs
```{go}
go generate
```

### TODO
- Write template files
- Generate structs
- Define attributes

## License
[MIT License](./LICENSE)