#!/bin/sh
set -eu

## func

die() { echo "$*" >&2; exit 2; }  # complain to STDERR and exit with error
needs_arg() { if [ -z "$OPTARG" ]; then die "$( basename $0 ):option requires an argument -- --$OPT"; fi; }

logger() { command logger --socket-errors=off $@ ; }
