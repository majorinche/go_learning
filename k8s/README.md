mkdir -p /root/go/src/k8s.io
cd /root/go/src/k8s.io
git clone https://github.com/kubernetes/apimachinery

mkdir -p /root/go/src/github.com/gogo
cd /root/go/src/github.com/gogo
git clone https://github.com/gogo/protobuf

mkdir -p /root/go/src/github.com/google/
cd /root/go/src/github.com/google/
git clone https://github.com/google/gofuzz

mkdir -p /root/go/src/k8s.io/
cd /root/go/src/k8s.io
git clone https://github.com/kubernetes/client-go.git
git clone https://github.com/kubernetes/api.git
git clone https://github.com/kubernetes/klog.git
git clone https://github.com/kubernetes/utils.git

mkdir -p /root/go/src/github.com/davecgh
cd /root/go/src/github.com/davecgh
git clone https://github.com/davecgh/go-spew
mkdir -p /root/go/src/github.com/go-logr
cd /root/go/src/github.com/go-logr
git clone https://github.com/go-logr/logr

mkdir -p /root/go/src/github.com/golang
cd /root/go/src/github.com/golang
git clone https://github.com/golang/protobuf

mkdir -p /root/go/src/golang.org/x
cd /root/go/src/golang.org/x
git clone https://github.com/golang/text
git clone https://github.com/golang/net
git clone https://github.com/golang/oauth2
git clone https://github.com/golang/time
git clone https://github.com/golang/crypto
git clone https://github.com/golang/sys

mkdir -p /root/go/src/github.com/googleapis
cd /root/src/github.com/googleapis
git clone https://github.com/googleapis/gnostic
mkdir -p /root/go/src/github.com/imdario
cd /root/go/src/github.com/imdario
git clone https://github.com/imdario/mergo
mkdir -p /root/go/src/github.com/modern-go
cd /root/go/src/github.com/modern-go
git clone https://github.com/modern-go/reflect2
git clone https://github.com/modern-go/concurrent
mkdir -p /root/go/src/github.com/spf13
cd /root/go/src/github.com/spf13
git clone https://github.com/spf13/pflag


mkdir -p /root/go/src/google.golang.org/
cd /root/go/src/google.golang.org
git clone https://github.com/protocolbuffers/protobuf-go
mv protobuf-go protobuf

mkdir -p /root/go/src/github.com/json-iterator
cd /root/go/src/github.com/json-iterator
git clone https://github.com/json-iterator/go


mkdir -p /root/go/src/sigs.k8s.io/
cd /root/go/src/sigs.k8s.io/
git clone https://github.com/kubernetes-sigs/structured-merge-diff
git clone https://github.com/kubernetes-sigs/yaml



go get gopkg.in/inf.v0
go get gopkg.in/yaml.v3

