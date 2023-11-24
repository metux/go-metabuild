# Project specification YAML file

## Global settings

| Attribute | Description |
|-----------|--------------|
| package   | package name |
| version   | eg. used in .pc files |
| srcdir    | will chdir here before build |

## Sections

### configure: project configuration and file generators

The `configure:` section contains settings for project configuration
*(eg. running various checks, look for imported packages)* as well as
generating config files.

See [here](configure.md) for details.

### targets: target objects to build and install

The `target:`section holds a map of target objects to be built and installed.

See [here](targets.md) for details.
