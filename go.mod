module github.com/joshqu1985/service-relation

go 1.13

require (
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/joshqu1985/fireman v0.1.7
	github.com/joshqu1985/protos v0.1.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/zap v1.13.0
	google.golang.org/grpc v1.25.1
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace github.com/joshqu1985/fireman => /Users/qulei/mywork/fireman // TODO for debug
