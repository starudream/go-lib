package annotation

import (
	"strings"
	"sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var (
	ProtoPath = "common/annotation.proto"

	_mom     = map[string]MethodOptions{}
	_momOnce sync.Once
)

func GetMethodOptions(name string) MethodOptions {
	_momOnce.Do(func() {
		filter := func(fd protoreflect.FileDescriptor) bool {
			fileImports := fd.Imports()
			for i := 0; i < fileImports.Len(); i++ {
				if fileImports.Get(i).Path() == ProtoPath {
					return true
				}
			}
			return false
		}
		replace := func(s string) string {
			i := strings.LastIndex(s, ".")
			return "/" + s[:i] + "/" + s[i+1:]
		}
		protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
			if filter(fd) {
				services := fd.Services()
				for i := 0; i < services.Len(); i++ {
					methods := services.Get(i).Methods()
					for j := 0; j < methods.Len(); j++ {
						method := methods.Get(j)
						proto.RangeExtensions(method.Options(), func(_ protoreflect.ExtensionType, v any) bool {
							if opts, ok := v.(MethodOptions); ok {
								_mom[replace(string(method.FullName()))] = opts
							}
							return true
						})
					}
				}
			}
			return true
		})
	})
	if v, ok := _mom[name]; ok {
		return v
	}
	return emptyMethodOptions
}
