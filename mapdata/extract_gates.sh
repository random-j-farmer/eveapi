#! /bin/bash

DB=$HOME/tmp/sqlite-latest.sqlite

doTheExtract() {
    cat <<EOF
package mapdata

// this file has been generated from the latest sqlite sde
// do not edit

import "github.com/random-j-farmer/eveapi/types"

var gatesBySolarSystemID = map[uint64][]types.Gate{
EOF

    sqlite3 $DB > /tmp/systems.txt <<EOF 
select solarSystemID from mapSolarSystems;
EOF

    while read id ; do
        echo -n "$id: {"
            sqlite3 $DB <<EOF
select '{"' || nm.itemName || '", '  || x || ', ' || y || ', ' || z || '},' from mapDenormalize md, invNames nm where solarSystemID=$id and groupID=10 and md.itemID=nm.itemID;
EOF
    echo "},"
    done < /tmp/systems.txt
    echo '}'
}


doTheExtract | perl -ne 'print $_ if ! /\{\}/;' > gates.go
