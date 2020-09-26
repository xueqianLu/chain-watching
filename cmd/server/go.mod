module github.com/hpb-project/chain-watching/cmd/server

go 1.14

require (
	github.com/hpb-project/chain-watching v0.0.0-00010101000000-000000000000
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
)

replace github.com/hpb-project/chain-watching => ../../
