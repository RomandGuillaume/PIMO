package regex

import (
	"math/rand"

	regen "github.com/zach-klippenstein/goregen"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

//MaskEngine is a value that mask thanks to a regular expression
type MaskEngine struct {
	generator regen.Generator
}

// NewMask return a RegexMask from a regexp
func NewMask(exp string, seed int64) (MaskEngine, error) {
	generator, err := regen.NewGenerator(exp, &regen.GeneratorArgs{RngSource: rand.NewSource(seed)})
	return MaskEngine{generator}, err
}

//Mask returns a string thanks to a regular expression
func (rm MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	out := rm.generator.Generate()
	return out, nil
}

func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if len(conf.Mask.Regex) != 0 {
		mask, err := NewMask(conf.Mask.Regex, seed)
		if err != nil {
			return nil, true, err
		}
		return config.WithEntry(conf.Selector.Jsonpath, mask), true, err
	}
	return nil, false, nil
}