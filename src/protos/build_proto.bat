protoc -I ./ -I D:/GitRepository/GoPath/src/ --go_out=plugins=grpc:../services/pbs *.proto
protoc -I ./ -I D:/GitRepository/GoPath/src/ --grpc-gateway_out=logtostderr=true:../services/pbs *.proto