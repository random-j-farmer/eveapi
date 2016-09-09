#! /bin/bash

exec > solarsystems.go

DB=$HOME/tmp/sqlite-latest.sqlite

cat <<EOF
package mapdata

// this file has been generated from the latest sqlite sde
// do not edit

import "github.com/random-j-farmer/eveapi/types"

var systemNameByID = map[uint64]types.SolarSystem{
EOF

sqlite3 $DB <<EOF
select solarSystemID || ': types.SolarSystem{"' ||  solarSystemName || '", ' || security || ', ' || regionID || '}, ' from mapSolarSystems;
EOF

echo '}'
