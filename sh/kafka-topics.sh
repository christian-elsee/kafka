#!/bin/sh
set -eu

## lib
. "$( dirname $0 )/lib.sh"

## arg
: ${KAFKACLI:=/usr/local/bin/kafka}
flags=""
server=""

while getopts b:tc:p:r:-: OPT; do
  # support long options: https://stackoverflow.com/a/28466267/519360
  if [ "$OPT" = "-" ]; then   # long option: reformulate OPT and OPTARG
    OPT="${OPTARG%%=*}"       # extract long option name
    OPTARG="${OPTARG#$OPT}"   # extract long option argument (may be empty)
    OPTARG="${OPTARG#=}"      # if long option argument, remove assigning `=`
  fi
  case "$OPT" in
    h | help               ) echo "Usage:$( basename $0 ) - filename ... " >&2;;

    c | create             ) action=create ;;
    d | delete             ) action=delete ;;
    l | list               ) action=list ;;
    b | bootstrap-server   ) needs_arg; server="$OPTARG" ;;
    t | topic              ) needs_arg; flags="$flags --topic $OPTARG" ;;
    p | partitions         ) needs_arg; flags="$flags --partitions $OPTARG" ;;
    r | replication-factor ) needs_arg; flags="$flags --replication-factor $OPTARG" ;;
    ??* ) die "Illegal option --$OPT" ;;  # bad long option
    ?   ) exit 2 ;;  # bad short option (error reported via getopts)
  esac
done
shift $((OPTIND-1)) # remove parsed options and args from $@ list

: ${action:?must be defined as create, delete or list}
: ${server:?bootstrap-server must be defined}

## main
logger -sp DEBUG -- "Enter" \
  :: "trace=$0" \
     "action=$action" \
     "server=$server" \
  :: "$flags"

echo "$KAFKACLI" topic "$action" "$flags" \
  | xargs \
  | tee /dev/stderr \
  | xargs sh -c 'exec $@' _
