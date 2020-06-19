package replacement

import (
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// ReplacementMask is to mask with a value from another field
type MaskEngine struct {
	Field string
}

// NewMask return a CommandMask from a value
func NewMask(field string) MaskEngine {
	return MaskEngine{field}
}

// Mask masks a value with another field of the json
func (remp MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	return context[0][remp.Field], nil
}

func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if len(conf.Mask.Replacement) != 0 {
		return config.WithEntry(conf.Selector.Jsonpath, NewMask(conf.Mask.Replacement)), true, nil
	}
	return nil, false, nil
}