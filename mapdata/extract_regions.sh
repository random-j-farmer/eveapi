#! /bin/bash

exec > regions.go

DB=$HOME/tmp/sqlite-latest.sqlite

cat <<EOF
package mapdata

// this file has been generated from the latest sqlite sde
// do not edit

var regionNameByID = map[uint64]string{
EOF

sqlite3 $DB <<EOF
select regionID || ': "' || regionName || '", ' from mapRegions;
EOF

echo '}'

