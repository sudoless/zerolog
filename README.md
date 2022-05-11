# Zero Allocation JSON Logger

Forked from [rs/zerolog](https://github.com/rs/zerolog).

## Changes:

- Remove all dependencies
- Remove non-source files (images, CNAME, etc)
- Remove syslog, journald, hlog, cbor
- Restructure package directories
- Remove go112 build tags
- Remove _unused_ (subjective) code
- Review and fix gosec, gocyclo and linter warnings
- Bump go.mod version to 1.17
- Align struct memory
- Release as v2
