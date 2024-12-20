package proto

//go:generate protoc --go_opt=paths=source_relative --go_out=. --go-plugin_opt=paths=source_relative,disable_pb_gen=true --go-plugin_out=. plugin.proto
