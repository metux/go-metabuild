# Model Driven Build System and Distro packaging

## Motivation

Over the decades, lots of build systems have been invented, solving lots of
problems occuring on building complex software. Much work has been done on
making software builds portable, e.g. by probing platform specific aspects
of the target system and driving platform build tools based on platform
agnostic build scripts. But one major aspect in software lifecycle is still
left out of scope: Building *distro specific packages*.

Traditionally, the responsibilities have been split between upstreams and
distros: the upstream teams just care about the software in general, while
distro maintainers handle the integration of thousands of individual packages
into the distro's ecosystem, providing an easy to install and operate
environment. While this approach has served us well for several decades now,
it still leaves a lot of individual work to distro maintainers and slows down
delivery timelines. One of the major problems is distros having different
policies _(for good reasons)_ on various aspects, e.g. file placement,
package splitting, dependency handling, and many more.

In recent years, tensions between upstreams _(interested in fast delivery)_
and distros _(caring about long term stability)_ have been increasing.
Numerous attempts have been made for some universal packaging approach,
but none having a fundamental breakthrough yet. Most of those basically
trying to take distros out of the loop, thus creating lots of other problems,
e.g. risking badly maintained bundled 3rdparty libraries, massive bloating
up of installation sizes, lack of integration between various software components.

The `metabuild` build system is an approach for attacking the problem from
entirely different vector: using a high level model for declaring the software's
*structure*, instead of rules or script code for implementing actual *build process*.
That way, it is possible to deduce the complete build process, down to distro
specific packages, and so relieve both upstreams as well as distro package
maintainers from a lot of manual work and increase overall delivery speed.

Such model can also serve as project metadata for IDE's and so improve
integration between build system and IDE.

## Modeling software structure

*Most* software packages are made of a quite limited set of artifacts, which
can be easily described by a few attributes _(e.g. type, name, sources, etc)_
where everything else _(e.g. compiler commands, file names and install
locations, target packages, etc)_ can deduced from, with the help of target
platform specific rules and policies.

Dependencies, target specific configuration and optional features also can
easily be described by simple attribute lists, so that actual probing 
process can be deduced automatically.

Choosing YAML as file format, with variable interpolation and defaults
seems a natural choice.

Note that `metabuild` in no way is an attempt to catch *all* possible scenarios,
there'll still be complex cases like browsers, office suites or the Linux kernel,
which won't easily fit into the model or would require huge refactoring.
Supporting a large enough portion of existing packages, as found in usual distros,
is good enough in scope of this project.

### Probing build / target environment

Before actual build can start, often several aspects of the target system need
to be checked, e.g. certain headers, types, libraries, etc. The functionality
is somewhat similar to `autoconf` - just written in a more compact and easier-to-parse
YAML notation.

#### Example: target distro and toolchain check

In almost all cases, the target distro needs to be probed, so we know which
distro specific policies have to be applied _(e.g. install pathes, sub-package
splitting, etc)_. For C-based projects, we also have to check for a C-Compiler.

```
configure:
    checks:
        - type:           target-distro
        - type:           c/compiler
          mandatory:      true
```

#### Example: some typical C checks

This example showing how to probe for certain C headers and functions,
also adding C defines on success. Corresponds to various `autoconf`'s 
`AC_CHECK_*` macros along with calling `AC_DEFINE()`.

```
        [ ... ]
        - c/header:       unistd.h
          mandatory:      false
          yes/c/defines:  HAVE_UNISTD_H
        - c/function:     fseeko
          yes/c/defines:  HAVE_FSEEKO
          mandatory:      true
        - c/type:         size_t
          c/header:       [stdio.h, stdlib.h]
          yes/c/defines:  HAVE_SIZE_T
        - c/type:         off64_t
          c/header:       sys/types.h
        - c/function:     strerror
          yes/c/defines:  HAVE_STRERROR
        - c/function:     vsnprintf
          c/header:       [stdio.h, stdarg.h]
          yes/c/defines:  HAVE_VSNPRINTF
        - c/type:         size_t
          c/header:       [stdio.h, stdlib.h]
          yes/c/defines:  HAVE_SIZE_T
```

#### Example: probing external libraries via pkg-config

In the following example several external libraries are probed via `pkgconf`.
Each of the map entries representing one `pkgconf` query, the map keys can
be used to indentify the library on import statements by individual targets.

The first statement sets a C-define on success, while the second
aborts the build if one dependency is missing.

```
        [ ... ]
        - pkgconf:
            XRES:                 xres
          yes/c/define: HAVE_XRES

        - pkgconf:
            GLIB:                 glib-2.0 >= 2.66.0
            GTK:                  gtk+-3.0 >= 3.24.0
            LIBXFCE4UTIL:         libxfce4util-1.0 >= 4.8.0
            LIBXFCE4UI:           libxfce4ui-2 >= 4.12.0
            LIBXFCE4KBD_PRIVATE:  libxfce4kbd-private-3 >= 4.12.0
            LIBXFCONF:            libxfconf-0 >= 4.13.0
            LIBWNCK:              libwnck-3.0 >= 3.14
            XINERAMA:             xinerama
            X11:                  x11
          mandatory:    true
```

### Optional features

Many software has optional features, which can be chosen at build time. Most of
such feature switches are pretty simple _(either on/off or a plain scalar value)_,
but depend on external libraries.

#### Example

The following example declares simple option switch:

* requires a library probed via `pkgconf` _(aborts if the library is missing)_
* off by default
* defines some C defines and adds importing the library, if enabled

```
startup_notification:
    pkgconf/require: STARTUP_NOTIFICATION
    default: n
    set@y:
        sym: [HAVE_STARTUP_NOTIFICATION, HAVE_LIBSTARTUP_NOTIFICATION]
        pkg: STARTUP_NOTIFICATION
```

If the option is enabled _(value is "y"), the elements of the `set@y` entry are
merged into the `buildconf::host::flags` configuration map: existing entries
are appended, so the same symbol can be reused, in order to collect. These collected
values can be referenced in other places, e.g. for target properties.

Similarily, more `set@*` entries can be defined for option other values, e.g.
for acting on disabled options or multiple choices.

### Target objects

#### Library example

The following example shows how a library can be expressed in simple yaml,
including all it's sources, headers, pkg-config metadata, ABI version, etc,
even with _(compile-time)_ optional features.

By that data we can deduce everything needed for building all library targets
_(static, shared, pkg-config, etc)_, as well as all needed information for
linking the library to other targets within the same project.

```
zlib{c/library}:
    version:              1
    library/name:         z
    library/mapfile:      zlib.map
    c/defines:            ${buildconf::host::flags::c/defines}
    pkgconf:
        name:             zlib
        description:      ZLib compression library
    source:               [adler32.c, crc32.c, deflate.c, infback.c, inffast.c,
                           inflate.c, inftrees.c, trees.c, zutil.c]
    source@option/gz=y:   [compress.c, uncompr.c, gzclose.c, gzlib.c,
                           gzread.c, gzwrite.c]
    headers:
        pub:
            source:       [zlib.h, zconf.h]
        priv:
            install:      false
            source:       [crc32.h, deflate.h, gzguts.h, inffast.h, inffixed.h,
                           inflate.h, inftrees.h, trees.h, zutil.h]
```

#### Executable example

This example shows a minimal C program, explicitly statically linking a library
from the same project.

```
minigzip{c/executable}:
    source:         test/minigzip.c
    link/static:    zlib
```

A more sophisticated example that's importing external libraries _(via pkg-config)_
and using variable interpolation to automatically compute source directory from
target name. It also put's the executable into a different location, as well as
into the `main` sub-package _(default is `progs` sub-package)_

```
helper-dialog/helper-dialog{c/executable}:
    install/dir:        ${helperpath}
    source:             ${@@^::@basename}.c
    source/dir:         ${@@^::@dirname}
    c/defines:          ${c/defines}
    install/package:    main
    pkgconf/import:     [GTK, LIBXFCE4UTIL, X11]
```

## Proof-Of-Concept implementation

### Design decisions

* statically linked: since `metabuild` is only used at build-time
  _(developer- or CI-machines), and upgrading those should be as easy as possible,
  the pro's outweight the con's
* golang: easy hi-level and typesafe language with garbage collection, good
  cross-platform support and good library ecosystem, and a fast compiler
* use [go-magicdict](https://github.com/metux/go-magicdict): provides easy YAML
  data access with variable interpolation and defaults
* defaults are supposed to be customized for individual platforms / distros,
  while the project metadata (`metabuild.yml`) is kept platform agnostic.

### Prototype implementation

The implementation is split into several layers:

* spec: reading the model and encapsulate model access
* engine: driving the build process
* * autoconf: probing and configuration
* * builder: the actual builder implementations for individual target types
* * packager: distro specific packaging

#### Spec layer

The main project model is represented by the `Global` struct, which is loaded
from the project's `metabuild.yml` and the _(potentially platform specific)_ defaults.
From here, all sub-nodes _(e.g. targets, probing, etc)_ can be retrieved.


