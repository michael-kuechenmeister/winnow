
# Tholian® Winnow Library

This is a golang library that is able to match the CycloneDX SBOM format to
given vulnerabilities in our own format of the [vulnerabilities](https://github.com/tholian-network/vulnerabilities)
repository.

Its purpose is to ease up the validation and analysis of CycloneDX SBOMs,
HBOMs and SaaSBOMs; and to produce as output a list of vulnerabilities
which affect the reported hardware and software.


## Features (work in progress)

- Consume CycloneDX SBOMs
- Consume CycloneDX HBOMs
- Consume the Vulnerabilities JSON format
- Match the BOMs against existing vulnerabilities
- Produce a list of vulnerabilities for affected software and hardware


## License

[AGPL-3.0 license](./LICENSE.txt)

`(c) 2021-2023 Tholian(r) Network`

