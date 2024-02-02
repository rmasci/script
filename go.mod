module github.com/rmasci/script

go 1.21.4

require (
	github.com/bitfield/script v0.22.0
	github.com/google/go-cmp v0.5.9
	github.com/itchyny/gojq v0.12.13
	github.com/rmasci/csvtable v0.7.0
	github.com/rogpeppe/go-internal v1.11.0
	mvdan.cc/sh/v3 v3.7.0
)

require (
	github.com/itchyny/timefmt-go v0.1.5 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rmasci/gotabulate v1.2.8 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/tools v0.11.0 // indirect
)

replace github.com/rmasci/csvtable => ../csvtable
