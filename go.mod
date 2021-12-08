module github.com/bitsongofficial/go-bitsong

go 1.17

require (
	github.com/RobotsAndPencils/goviz v0.0.0-20161006195526-3924fe917ccd // indirect
	github.com/cosmos/cosmos-sdk v0.44.5
	github.com/cosmos/ibc-go/v2 v2.0.0
	github.com/gorilla/mux v1.8.0
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.2.1
	github.com/strangelove-ventures/packet-forward-middleware v1.0.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.4
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/cosmos/cosmos-sdk v0.44.5
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
