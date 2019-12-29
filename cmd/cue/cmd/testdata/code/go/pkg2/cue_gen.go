// Code generated by gogen.Generate; DO NOT EDIT.

package pkgtwo

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode/gocodec"
)

var cuegenvalBarzer = cuegenMake("Barzer", &Barzer{})

// Validate validates x.
func (x *Barzer) Validate() error {
	return cuegenCodec.Validate(cuegenvalBarzer, x)
}

var cuegenCodec, cuegenInstance = func() (*gocodec.Codec, *cue.Instance) {
	var r *cue.Runtime
	r = &cue.Runtime{}
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0]
}()

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	v := cuegenInstance.Lookup(name)
	if !v.Exists() {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 229 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xffD\xce\xc1J\xc3@\x10\xc6\xf1\xf96\x11u\xa9\xbe\x81\x10\x8a\a{\xe9\x82\x17!\xe0E\u0123\x88\x1e\xc5\u00f0;\u0190\xb6\x1b\x9a\t\x82\xa2\xa0\xd6\xea[\xafD\x85\x9e\x06~\xf0\x1f\xbe\xbd\xf4e`\xd27!\xbd\x13\x9d\xa4\xb7\f\x18\u054bNy\xe1\u5715\aG\x86\xfc:F\x85!\xe4W\xac\x0f\x18\x11\xb6.\xea\x99tHk\":H\x9f\x06\u063f\xbd\xf3\xbdL\xef\xeb\xd9\u007f\xb9&\xa4\x15\xd1Q\xfa\u0200\x9d\x8d\xaf\b\x06\xf9%\xcfex\x94\xff\xa2%\xa2z\xd8\x01`\xbbl\x9bJ\x1f#\x80\x89\x9f\a\xe7{q\xc3U\xe94\xb0\xb2\xf31\x88\xab\xa2k\x9b\xea\xd8q\bS\xdf\v&-\xfb\x86+)\xfebk\xcfx\xf9$\u02f2x\xb6\xbb7eq\xfa:\xf6\xac\x87c\xfbb\x89~\x02\x00\x00\xff\xffo/}\x10\xf7\x00\x00\x00")