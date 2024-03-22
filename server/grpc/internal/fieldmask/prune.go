package fieldmask

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Prune(msg proto.Message, paths []string) {
	NestedMaskFromPaths(paths).Prune(msg)
}

type NestedMask map[string]NestedMask

func NestedMaskFromPaths(paths []string) NestedMask {
	mask := make(NestedMask)
	for _, path := range paths {
		curr := mask
		var letters []rune
		for _, letter := range path {
			if letter == '.' {
				if len(letters) == 0 {
					continue
				}
				key := string(letters)
				c, ok := curr[key]
				if !ok {
					c = make(NestedMask)
					curr[key] = c
				}
				curr = c
				letters = nil
				continue
			}
			letters = append(letters, letter)
		}
		if len(letters) != 0 {
			key := string(letters)
			if _, ok := curr[key]; !ok {
				curr[key] = make(NestedMask)
			}
		}
	}

	return mask
}

func (mask NestedMask) Prune(msg proto.Message) {
	if len(mask) == 0 {
		return
	}

	rft := msg.ProtoReflect()
	rft.Range(func(fd protoreflect.FieldDescriptor, _ protoreflect.Value) bool {
		m, ok := mask[string(fd.Name())]
		if ok {
			if len(m) == 0 {
				rft.Clear(fd)
				return true
			}
			if fd.IsMap() {
				mp := rft.Get(fd).Map()
				mp.Range(func(mk protoreflect.MapKey, mv protoreflect.Value) bool {
					if mi, ok := m[mk.String()]; ok {
						if i, ok := mv.Interface().(protoreflect.Message); ok && len(mi) > 0 {
							mi.Prune(i.Interface())
						} else {
							mp.Clear(mk)
						}
					}

					return true
				})
			} else if fd.IsList() {
				list := rft.Get(fd).List()
				for i := 0; i < list.Len(); i++ {
					m.Prune(list.Get(i).Message().Interface())
				}
			} else if fd.Kind() == protoreflect.MessageKind {
				m.Prune(rft.Get(fd).Message().Interface())
			}
		}
		return true
	})
}
