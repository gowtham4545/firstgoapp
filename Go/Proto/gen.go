package rides

//go:generate mkdir -p pb
//go:generate protoc --go_out=pb --go_opt=paths=source_relative ./rides.proto
