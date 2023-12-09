# Project spec: targets

This section is a map of all targets to be built and installed. Most types use
the key as the object's name / output file. Each target must have an `type`
attribute set.

| Type                   | Description                                                             |
|------------------------|-------------------------------------------------------------------------|
| c/executable           | executable program written in C                                         |
| c/library              | library written in C *(building shared/static/pkgconfig per default)*   |
| c++/executable         | executable program written in C++                                       |
| c++/library            | library written in C++ *(building shared/static/pkgconfig per default)* |
| data/desktop           | simple FreeDesktop.org `*.desktop` file                                 |
| data/misc              | misc data *(installed to $datadir)                                      |
| data/pixmaps           | pixmap files *(installed to $datadir/pixmaps)*                          |
| data/lib-script        | install script to arch-independent libdir                               |
| data/lib-script-subdir | install script to arch-independent libdir (subdir by target id)         |
| doc/man                | Unix manual page *(troff/nroff)*                                        |
| doc/misc               | Simple documentation files *(placed under $datadir/doc/...)*            |
| gen/glib-resource      | Generate Glib resource and source code files from XML                   |
| gen/glib-marshal       | Generate Glib marshalling code                                          |
| gen/xdt-csource        | Generate source code for compiling-in XML files *(eg. `*.glade`)*       |
| gen/xxd-csource        | Generate source for compiling in binary data (like xxd -i)              |
| i18n/desktop           | multilingual FreeDesktop.org `*.desktop` file                           |
| i18n/po                | gettext translation files *(building `*.mo` files)*                     |

## Automatic attributes

Some attributes are automatically set in the post-load phase.

| Attribute  | Description                                                    |
|------------|----------------------------------------------------------------|
| @id        | The ID-part of the target key name (used eg. for target names) |
| @id/suffix | Suffix (extension) of the ID, without leading dot              |
| @type      | target type, extracted from target key                         |

## Special key notation

For convenience, the target key (in the yaml struct) may encode the target type:

   `target_id{target_type}`

## Target types

### c/executable: executable program written in C

#### Attributes:

| Attribute       | Default                            | Description                       |
|-----------------|------------------------------------|-----------------------------------|
| c/cflags        |                                    | extra C-flags                     |
| c/defines       |                                    | extra C-defines                   |
| c/ldflags       |                                    | extra linker flags                |
| compiler/lang   | C                                  | compiler language                 |
| file            | ${@@^::name}                       | file name                         |
| headers         |                                    | map of header file bundles        |
| include/dir     |                                    | extra include dirs                |
| install         | ${@@^::@@^::install}               | whether to install                |
| install/dir     | ${buildconf::install-dirs::bindir} | install directory                 |
| install/package | prog                               | install package                   |
| install/perm    | 0755                               | install permissions               |
| install/subdir  |                                    | install subdirectory              |
| link/shared     |                                    | dynamically link internal libs    |
| link/static     |                                    | statically link internal libs     |
| name            | ${@@^::@id}                        | object name                       |
| pkgconf/import  |                                    | IDs of pkgconf-packages to import |
| source          |                                    | source files *(globs)*            |
| source/dir      |                                    | source subdir                     |

#### Example:
```
    settings-dialogs/xfwm4-tweaks-settings:
        type:               c/executable
        pkgconf/import:     [LIBXFCE4KBD_PRIVATE]
        source:             [tweaks-settings.c, range-debouncer.c]
        source/dir:         settings-dialogs
        include/dir:        .
        c/defines:          ${c/defines}
        link/static:        xfwm-common
        install/package:    main
        headers:
            install:        false
            priv:
                source:     [xfwm4-tweaks-dialog_ui.h, range-debouncer.h]
```

### c/library: library written in C

This builds and installs a library written in C, including static _(`*.a`)_, shared _(`*.so.<version>`)_,
pkgconf _(`*.pc`)_, symbolic link _(`*.so`)_ for development, header files, etc.

The attributes are those of `c/executable` plus some more:

| Attribute      | Default        | Description                                               |
|----------------|----------------|-----------------------------------------------------------|
| abi            | 1              | shared object version                                     |
| compiler/lang  | C              | compiler language                                         |
| description    | ${description} | description _(for pkgconf)_                               |
| install        | true           | whether to install                                        |
| install/dir    |                | installation dir                                          |
| install/subdir |                | installation subdir                                       |
| libname        | ${@@^::@id}    | library name _(as used in `-l...` flag)`                  |
| mapfile        |                | linker map file                                           |
| pkgconf        |                | map of pkg-config metadata _(`name:` and `description:`)_ |
| skip/devlink   |                | skip devlink _(to shared object)_                         |
| skip/pkgconf   |                | skip `.pc` file                                           |
| skip/shared    |                | skip shared object                                        |
| skip/static    |                | skip static archive                                       |
| version        | ${version}     | version _(for pkgconf)_                                   |

#### Example:
```
    zlib:
        type:           c/library
        mapfile:        zlib.map
        version:        1
        libname:        z
        c/defines:      ${buildconf::host::flags::c/defines}
        pkgconf:
            name:        zlib
            description: ZLib compression library
        source: "*.c"
```

### data/misc: arbitrary data files (/usr/share/...)

#### Attributes:

| Attribute       | Default                             | Description                                           |
|-----------------|-------------------------------------|-------------------------------------------------------|
| install         | ${@@^::@@^::install}                | whether to install into distro package                |
| install/dir     | ${buildconf::install-dirs::datadir} | install directory                                     |
| install/package | data                                | install package                                       |
| install/perm    | 0064                                | install permissions                                   |
| install/subdir  |                                     | subdirectory _(under standard $datadir) to install to |
| source          |                                     | source files globs                                    |
| source/dir      |                                     | source subdirector                                    |

#### Example:
```
    data/opening:
        type:           data/misc
        install/subdir: lincity/opening
        source:         "*"
        source/dir:     opening
```

### data/lib-script: script libs

Similar to data/misc, but putting it into arch-independent libdir and sets executable flag.

#### Attributes:

| Attribute       | Default                                              | Description                            |
|-----------------|------------------------------------------------------|----------------------------------------|
| install         | ${@@^::@@^::install}                                 | whether to install into distro package |
| install/dir     | ${buildconf::install-dirs::libdir-noarch}/${package} | install directory                      |
| install/package | data                                                 | install package                        |
| install/perm    | 0775                                                 | install permissions                    |
| source          |                                                      | source files globs                     |
| source/dir      |                                                      | source subdirector                     |

#### Example:
```
    foo:
        type:           data/lib-script
        source:         "foo-helper"
```

### data/lib-script-subdir: script libs

Like data/lib-script, but taking `source/dir` from target id.

#### Attributes:

| Attribute       | Default                                              | Description                            |
|-----------------|------------------------------------------------------|----------------------------------------|
| install         | ${@@^::@@^::install}                                 | whether to install into distro package |
| install/dir     | ${buildconf::install-dirs::libdir-noarch}/${package} | install directory                      |
| install/package | data                                                 | install package                        |
| install/perm    | 0775                                                 | install permissions                    |
| source          |                                                      | source files globs                     |
| source/dir      | ${@@^::@id}                                          | source subdirector                     |

#### Example:
```
    plugin-foo:
        type:           data/lib-script-subdir
        source:         "*.sh"
```

### data/desktop

#### Attributes:

| Attribute           | Default                                                    | Description                            |
|---------------------|------------------------------------------------------------|----------------------------------------|
| desktop/type        | Application                                                | `Type=` field                          |
| desktop/name        | ${shortname}                                               | `Name=` field                          |
| desktop/categories  |                                                            | `Categories=` field                    |
| desktop/genericname | ${name}                                                    | `GenericName=` field                   |
| desktop/comment     | ${description}                                             | `Comment=` field                       |
| desktop/icon-file   | ${buildconf::install-dirs::pixmapdir}/${@@^::desktop/icon} | `Icon=` field                          |
| desktop/terminal    | false                                                      | `Terminal=` field                      |
| desktop/exec        |                                                            | `Exec=` field                          |
| desktop/tryexec     |                                                            | `TryExec=` field                       |
| file                | ${@@^::@id}                                                | output file name                       |
| install             | ${@@^::@@^::install}                                       | whether to install into distro package |
| install/dir         | ${buildconf::install-dirs::fdo-appdir}                     | install directory                      |
| install/package     | data                                                       | install package                        |
| install/perm        | 0064                                                       | install permissions                    |
| install/subdir      |                                                            | install subdir                         |

#### Example:
```
    lincity.desktop:
        type:               data/desktop
        desktop/categories: Application;Game;StrategyGame
        desktop/exec:       lincity
        desktop/tryexec:    ${buildconf::install-dirs::bindir}/lincity
```

### data/pixmap

#### Attributes:

| Attribute           | Default                               | Description                            |
|---------------------|---------------------------------------|----------------------------------------|
| file                | ${@@^::@id}                           | output file name                       |
| install             | ${@@^::@@^::install}                  | whether to install into distro package |
| install/dir         | ${buildconf::install-dirs::pixmapdir} | install directory                      |
| install/package     | data                                  | install package                        |
| install/perm        | 0064                                  | install permissions                    |
| install/subdir      |                                       | install subdir                         |
| source              |                                       | source file                            |

#### Example:
```
    data/pixmap:
        type:           data/pixmaps
        source:         debian/lincity.xpm
```

### doc/misc

#### Attributes:

| Attribute           | Default                            | Description                            |
|---------------------|------------------------------------|----------------------------------------|
| source              |                                    | source files _(globs)_                 |
| install             | ${@@^::@@^::install}               | whether to install into distro package |
| install/dir         | ${buildconf::install-dirs::docdir} | install directory                      |
| install/package     | doc                                | install package                        |
| install/perm        | 0064                               | install permissions                    |
| install/subdir      | ${package}                         | install subdir                         |
| compress            | gz                                 | compression method                     |

#### Example:
```
    doc/misc:
        type:           doc/misc
        source:
            - Acknowledgements
            - README
            - TODO
            - CHANGES
            - COPYING
            - COPYRIGHT
```

### doc/man

#### Attributes:

| Attribute           | Default                            | Description                            |
|---------------------|------------------------------------|----------------------------------------|
| source              |                                    | source file                            |
| install             | ${@@^::@@^::install}               | whether to install into distro package |
| install/dir         | ${buildconf::install-dirs::mandir} | install directory                      |
| install/package     | data                               | install package                        |
| install/perm        | 0064                               | install permissions                    |
| install/subdir      |                                    | install subdir                         |
| man/alias           |                                    | manpage alias                          |
| man/compress        | gz                                 | compression method                     |
| man/section         | ${@@^::@id/suffix}                 | manual section                         |
| source              | ${@@^::@id}                        | manual page (nroff/troff) file         |

#### Example:
```
    lincity.6:
        type:           doc/man
        man/alias:      xlincity
```

### i18n/po

This builds/installs gettext `*.mo` files for given linguas. If the `linguas` attribute is missing,
the global variable `${i18n::linguas}` is used. Per default, the `*.po` files are expected in
the `po/` subdirectory.

#### Attributes:

| Attribute       | Default                               | Description                                           |
|-----------------|---------------------------------------|-------------------------------------------------------|
| install         | true                                  | whether to install into distro package                |
| install/dir     | ${buildconf::install-dirs::localedir} | install target directory                              |
| install/package | data                                  | install target package                                |
| install/perm    | 0664                                  | install file permissions                              |
| install/subdir  |                                       | subdirectory _(under standard $datadir) to install to |
| i18n/linguas    | ${i18n::linguas}                      | list of locale names to build for                     |
| i18n/category:  | LC_MESSAGES                           | locale category                                       |
| i18n/domain:    | ${package}                            | locale domain                                         |
| source/dir:     | po                                    | source subdirectory                                   |
| name:           | ${@@^::domain}.mo                     | `*.mo` target file name                               |

#### Example:
```
    po:
        type:    i18n/po
```

### i18n/desktop

#### Attributes:

| Attribute       | Default                                | Description                                           |
|-----------------|----------------------------------------|-------------------------------------------------------|
| i18n/linguas    | ${i18n::linguas}                       | list of locale names to build for                     |
| i18n/po/dir     | po                                     | po file directory                                     |
| install         | true                                   | whether to install into distro package                |
| install/dir     | ${buildconf::install-dirs::fdo-appdir} | install target directory                              |
| install/package | data                                   | install target package                                |
| install/perm    | 0664                                   | install file permissions                              |
| install/subdir  |                                        | subdirectory _(under standard $datadir) to install to |
| output/suffix   | .desktop                               |                                                       |
| source          |                                        | source item names _(not files)_                       |
| source/dir:     |                                        | source subdirectory                                   |
| source/suffix   | .desktop.in                            | input file suffix                                     |

#### Example:
```
    desktop:
        type:           i18n/desktop
        source:         [xfce-wm-settings, xfce-wmtweaks-settings, xfce-workspaces-settings]
        source/dir:     settings-dialogs
```

### gen/glib-resource:

#### Attributes:

| Attribute        | Default                                              | Description               |
|------------------|------------------------------------------------------|---------------------------|
| source           |                                                      | XML source file           |
| source/dir       | .                                                    | source subdir             |
| resource/dir     | ${@@^::source/dir}                                   | resource output directory |
| resource/name    | ${@@^::name}                                         | resource name             |
| output/c/header  | ${@@^::resource/dir}/${@@^::resource/name}.h         | c header output file      |
| output/c/source  | ${@@^::resource/dir}/${@@^::resource/name}.c         | c source output file      |
| output/gresource | ${@@^::resource/dir}/${@@^::resource/name}.gresource | `.gresource` output file  |
| name             | ${@@^::@id}                                          | target name               |

#### Example:
```
    settings-dialogs/workspace-resource:
        type:               gen/glib-resource
        name:               workspace-resource
        source:             workspace.gresource.xml
        source/dir:         settings-dialogs
```

### gen/glib-marshal:

Generate Glib marshalling code from prototype definition file.

#### Attributes:

| Attribute       | Default               | Description             |
|-----------------|-----------------------|-------------------------|
| source          | ${@@^::@id}.list      | prototype list source   |
| source/dir      | .                     | source subdir           |
| resource/name   | ${@@^::@id}           | resource name           |
| output/name     | ${@@^::@id}           | prefix for output files |
| output/c/header | ${@@^::output/name}.h | c header output file    |
| output/c/source | ${@@^::output/name}.c | c source output file    |

#### Example:
```
    src/gq-marshal:
        type:               gen/glib-marshal
        resource/name:      gq_marshal
```

### gen/xdt-csource:

#### Attributes:

| Attribute       | Default           | Description                   |
|-----------------|-------------------|-------------------------------|
| source          | ${@@^::@id}.glade | XML source file               |
| resource/name   |                   | name of resource _(C symbol)_ |
| output/c/header | ${@@^::@id}_ui.h  | c header output file          |

#### Example:
```
    settings-dialogs/workspace-resource:
        type:               gen/glib-resource
        name:               workspace-resource
        source:             workspace.gresource.xml
        source/dir:         settings-dialogs
```

### gen/xxd-csource:

Generate code fragment (header) for compiling in binary data, like `xxd -i`.

#### Attributes:

| Attribute        | Default     | Description          |
|------------------|-------------|----------------------|
| source           |             | binary input file    |
| output/c/header  | ${@@^::@id} | c header output file |

#### Example:
```
    ClayRGB1998_icc.h:
        type:               gen/xxd-csource
        source:             src/ClayRGB1998.icc
```
