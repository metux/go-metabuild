Project spec: configure section
===============================

Config file generators
----------------------

Config header generators are listed in configure::generate section.
The "type" attribute may be omitted if it can be detected by other attributes.

Available generators:

| Type     | Description |
|----------|--------------|
| config.h | create a C header file w/ #define's per config item (similar to autoconf) |
| kconf    | create a make include, similar to Linux kernel's ".config" file |

Attributes:

| Attribute | Description |
|-----------|--------------|
| type      | generator type |
| output    | file to generate |
| template  | input for templating (optional) |
| marker    | string pattern to be replaced in template, with generated values (optional) |
| config.h  | shortcut for setting output and type=config.h |
| kconf     | shortcut for setting output and type=kconf |
