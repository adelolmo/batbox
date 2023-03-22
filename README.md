# batbox
**Hobbit -- The True Story**

**The Terror of Mecha Godzilla -- The True Story**

by Fredrik Ramsberg

---


I found these two conversational adventures here http://www.spagmag.org/archives/h.html#hobbit

The Terror of Mecha Godzilla is working correctly. The Hobbit is not yet fully tested.

    NAME: Hobbit - The True Story
    AUTHOR: Fredrik Ramsberg and Johan Berntsson
    EMAIL: d91frera SP@G und.ida.liu.se (Ramsberg -- no address provided for Berntsson)
    DATE: Monday 19 April 1993
    PARSER: Very strange and not very good
    SUPPORTS: DOS
    AVAILABILITY: Shareware -- $10, but I think it's a joke
    URL: ftp://ftp.ifarchive.org/if-archive/games/pc/hobbit.zip

    NAME: The Terror of Mecha Godzilla - The True Story
    AUTHOR: Fredrik Ramsberg
    EMAIL: d91frera SP@G und.ida.liu.se
    DATE: Monday 4 October 1993
    PARSER: Very strange and not very good
    SUPPORTS: DOS
    AVAILABILITY: Shareware -- $10, but I think it's a joke
    URL: ftp://ftp.ifarchive.org/if-archive/games/pc/mecha.zip

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

## Install

    sudo make install

## Run
Execute the program running the following command:

    batbox

## Run a game

    batbox mechagodzilla
    > RESTART
    
To enable debugging add `-d` before the game:

    batbox -d mechagodzilla
    
