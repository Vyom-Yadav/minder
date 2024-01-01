#!/usr/bin/env bash

set -e

PREV_YEAR=$1
CURR_YEAR=$2
DIR=$3

OLD_VALUE="Copyright $PREV_YEAR Stacklok, Inc."
NEW_VALUE="Copyright $CURR_YEAR Stacklok, Inc."

find "$DIR" -type f -exec sed -i "s/$OLD_VALUE/$NEW_VALUE/g" {} \;

BASEDIR=$(pwd)
echo "Distinct Diff in $DIR is:"
cd "$DIR"
git diff | grep -Eh "^\+"  | grep -v "+++ b" | sort | uniq
cd "$BASEDIR"
