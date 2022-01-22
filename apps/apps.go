// Package apps provides a clean way for Tidbyt to be able to get a list of all
// community apps.
package apps

// Code generated by tools/generator. DO NOT EDIT.

import (
	"tidbyt.dev/community/apps/bigclock"
	"tidbyt.dev/community/apps/clockbyhenry"
	"tidbyt.dev/community/apps/datetimeclock"
	"tidbyt.dev/community/apps/daynightmap"
	"tidbyt.dev/community/apps/digitalrain"
	"tidbyt.dev/community/apps/dvdlogo"
	"tidbyt.dev/community/apps/fuzzyclock"
	"tidbyt.dev/community/apps/golfhandicap"
	"tidbyt.dev/community/apps/manifest"
	"tidbyt.dev/community/apps/nightscout"
	"tidbyt.dev/community/apps/nyancat"
	"tidbyt.dev/community/apps/pokedex"
	"tidbyt.dev/community/apps/sbbtimetable"
	"tidbyt.dev/community/apps/showthis"
	"tidbyt.dev/community/apps/theysaidso"
	"tidbyt.dev/community/apps/twitterfollows"
)

// GetManifests returns a list of all apps in the this repository. Add your applet
// below to include it in the Tidbyt Mobile app for all Tidbyt users.
func GetManifests() []manifest.Manifest {
	return []manifest.Manifest{
		bigclock.New(),
		clockbyhenry.New(),
		datetimeclock.New(),
		daynightmap.New(),
		digitalrain.New(),
		dvdlogo.New(),
		fuzzyclock.New(),
		golfhandicap.New(),
		nightscout.New(),
		nyancat.New(),
		pokedex.New(),
		sbbtimetable.New(),
		showthis.New(),
		theysaidso.New(),
		twitterfollows.New(),
	}
}
