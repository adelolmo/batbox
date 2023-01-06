# batbox
The Terror of Mechagodzilla - The True Story by Fredrik Ramsberg

## Build

To build the project:

    make

The build process supports cross platform compilation.
To set the target architecture add the argument `ARCH` with the value when calling `make`.
Supported architectures:
* amd64
* i386
* armhf
* arm64


    make ARCH=armhf

### Debian package

As an additional step after building, you can create a debian package.

    make debian

To set the target architecture add the argument `ARCH` with the value when calling `make`.
Supported architectures:
* amd64
* i386
* armhf
* arm64


    make debian ARCH=armhf

## Install

### Package
After building the package you can install it with `dpkg` command.

    sudo dpkg -i batbox_1.0.0_armhf.deb

### Binary only
You can also install the binary without building the package.

    sudo make install

## Run
Execute the program running the following command:

    batbox

## Run a game

    batbox -d mechagodzilla
    > RESTART
