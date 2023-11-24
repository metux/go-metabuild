# Project spec: configure section

## Build configuration checks

Various checks can be run before configuring and building a package. These are
listed in the `configure::checks` section.

Each check needs a `type` attribute, but in many cases this can be automatically
deduced by presence of other fields _(thus can be omitted)_: `c/function`, `c/type`, `c/header`, `pkgconf`

### Generic attributes _(for all types)_:

| Attribute      | Description                                                                             |
|----------------|-----------------------------------------------------------------------------------------|
| type           | type of the check                                                                       |
| mandatory      | if true, abort if the check fails                                                       |
| build          | if true, check the *build* system instead of target *host* _(eg. internal build tools)_ |
| yes/c/defines  | C defines to be set when check succeeded                                                |
| no/c/defines   | C defines to be set when check failed                                                   |
| yes/c/cflags   | C compiler flags to be set when check succeeded                                         |
| no/c/cflags    | C linker flags to be set when check failed                                              |
| yes/c/ldflags  | C linker flags to be set when check succeeded                                           |
| no/c/ldflags   | C linker flags to be set when check failed                                              |

### Check types

#### c/header

Compile test checking for whether the C-Compiler can include the given header file(s).

Example:

```
configue:
  checks:
    - c/header: stdio.h
      yes/c/defines: HAVE_STDIO_H
```

#### c/function

Compile test checking whether C-Compiler/Linker finds some function.

Example:

```
configure:
  checks:
    - c/function: gettimeofday
      yes/c/defines: HAVE_GETTIMEOFDAY
```

#### c/type

Compile test checking whether C-Compiler knows some type. Can also include extra headers for the test.

Example:

```
configure:
  checks:
    - c/type:         size_t
      c/header:       [stdio.h, stdlib.h]
      yes/c/defines:  HAVE_SIZE_T
```

#### c/compiler

Looks for the C-Compiler and stores it's command, target architecture, etc.
Should be done before any c/* checks.

Example:

```
configure:
  checks:
    - type: c/compiler
```

#### c++/compiler

Looks for the C-Compiler and stores it's command, target architecture, etc, in *buildconf*.
Should be done before any c/* checks.

Example:

```
configure:
  checks:
    - type: c++/compiler
```

#### pkgconf

Check for imported packages via pkgconf and stores retrieved data in the *buildconf*.

Each listed package query *(map value)* is probed separately and recorded in *buildconf* with given ID *(map key)*
The check fails if one of the queries fails, but only after all queries had been done.

```
configure:
  checks:
    - pkgconf:
        GTK:       gtk+-3.0 >= 3.24.0
        XINERAMA:  xinerama
        X11:       x11
```

## Config file generators

Config header generators are listed in configure::generate section.
The "type" attribute may be omitted if it can be detected by other attributes.

### Available generators:

| Type     | Description |
|----------|--------------|
| config.h | create a C header file w/ #define's per config item (similar to autoconf) |
| kconf    | create a make include, similar to Linux kernel's ".config" file |

### Attributes:

| Attribute | Description |
|-----------|--------------|
| type      | generator type |
| output    | file to generate |
| template  | input for templating (optional) |
| marker    | string pattern to be replaced in template, with generated values (optional) |
| config.h  | shortcut for setting output and type=config.h |
| kconf     | shortcut for setting output and type=kconf |
