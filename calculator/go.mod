module calculador

go 1.16

// indicate in which folder the module really stay
// doesn't matter the name of the module file
replace mathOps => ../_mathOps

replace utilities => ../_utilities

// all the files in the module folder are considered

// launch go mod tidy to generate requires
require (
	mathOps v0.0.0-00010101000000-000000000000
	utilities v0.0.0-00010101000000-000000000000
)
