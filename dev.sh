#!/bin/sh
# using github.com/githubnemo/CompileDaemon 
# for the development cycle
CompileDaemon -build "gb build" -command "./bin/astralboot -vvv " -directory ./src/astralboot/ -pattern "(.+\\.go|.+\\.js|.+\\.html)$"
