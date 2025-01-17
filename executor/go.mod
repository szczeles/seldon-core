module github.com/seldonio/seldon-core/executor

go 1.16

require (
	github.com/cloudevents/sdk-go v1.2.0
	github.com/confluentinc/confluent-kafka-go v1.4.2
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.4.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.1
	github.com/onsi/gomega v1.14.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/common v0.26.0
	github.com/seldonio/seldon-core/operator v0.0.0-00010101000000-000000000000
	github.com/tensorflow/tensorflow/tensorflow/go/core v0.0.0-00010101000000-000000000000
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/zap v1.19.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/grpc v1.37.0
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.21.3
	sigs.k8s.io/controller-runtime v0.9.6
)

replace github.com/tensorflow/tensorflow/tensorflow/go/core => ./proto/tensorflow/core

replace github.com/seldonio/seldon-core/operator => ./_operator

replace k8s.io/client-go => k8s.io/client-go v0.21.3

exclude github.com/go-logr/logr v1.0.0
