//go:generate protoc -I=proto --go_out=. --go-grpc_out=. accounts.proto pomelo.proto common.proto
package eosn_base_api
