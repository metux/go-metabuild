# go-metabuild

Experimental model based build system with packaging

1. High level modeling of SW's structure, following
   convention-over-code, where everyhing else can be
   automatically deduced from. Simple YAML file.
   Describing what you wanna get, instead of how to do it.
3. Automatic checks/configuration, similar to autoconf
   et al, but much simpler config.
4. Produces ready-to-deploy distro specific packages
   directly from source, w/o huge extra tool stacks.
5. Platform/distro specific policies can be centrally
   customized by platform maintainer/operator, w/o
   having to touch per package build scripts.

See [examples/pkg/zlib.yaml](examples/pkg/zlib.yaml) or
[examples/pkg/lincity.yaml](examples/pkg/lincity.yaml) for
an examples of the per package configuration.

Features
--------

* purely declarative description of SW structure instead of rules for actual build process
* automatic toolchain and library detection
* pkgconfig and distro packages as 1st class citizen
* builds native distro packages
* automatic runtime dependences for shared libraries and .pc files
* respects distro policies (eg. artifact placements, compression, stripping, ...)

Project status
--------------

Right now, yet purely research. No releases yet.
Will switch to normal development, once a few practical example packages can be fully built/packaged (inc. docs etc) for at least Debian-based distros.
After that, more languages and target distros will be added.
