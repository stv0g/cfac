package all

import (
	_ "github.com/stv0g/cfac/pkg/mobility/apag"
	_ "github.com/stv0g/cfac/pkg/occupancy/carolus"
	_ "github.com/stv0g/cfac/pkg/occupancy/cccac"
	_ "github.com/stv0g/cfac/pkg/occupancy/terminmanagement"

	_ "github.com/stv0g/cfac/pkg/freifunk"
	_ "github.com/stv0g/cfac/pkg/health/blutspende"
	_ "github.com/stv0g/cfac/pkg/occupancy/gyms/fitx"
	_ "github.com/stv0g/cfac/pkg/occupancy/gyms/mcfit"
	_ "github.com/stv0g/cfac/pkg/occupancy/gyms/medaix"
	_ "github.com/stv0g/cfac/pkg/occupancy/gyms/wof"
	// _ "github.com/stv0g/cfac/pkg/occupancy/gyms/hochschulsport"
	// _ "github.com/stv0g/cfac/pkg/occupancy/spielbank"
)
