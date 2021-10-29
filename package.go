//go:generate protoc -I=./proto/v1 --go_out=./proto/v1 --go-grpc_out=./proto/v1 accounts.proto pomelo.proto common.proto
package eosn_base_api
