package templatemask

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByTemplate(t *testing.T) {
	template := "{{.name}}.{{.surname}}@gmail.com"
	tempMask, err := NewMask(template)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	config := model.NewMaskConfiguration().
		WithEntry("mail", tempMask)
	maskingEngine := model.MaskingEngineFactory(config)

	data := model.Dictionary{"name": "Jean", "surname": "Bonbeur", "mail": "jean44@outlook.com"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"name": "Jean", "surname": "Bonbeur", "mail": "Jean.Bonbeur@gmail.com"}
	assert.Equal(t, waited, result, "Should create the right mail")
}

func TestMaskingShouldReplaceSensitiveValueByTemplateInNested(t *testing.T) {
	template := "{{.customer.identity.name}}.{{.customer.identity.surname}}@gmail.com"
	tempMask, err := NewMask(template)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	config := model.NewMaskConfiguration().
		WithEntry("mail", tempMask)
	maskingEngine := model.MaskingEngineFactory(config)

	data := model.Dictionary{"customer": model.Dictionary{"identity": model.Dictionary{"name": "Jean", "surname": "Bonbeur"}}, "mail": "jean44@outlook.com"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"customer": model.Dictionary{"identity": model.Dictionary{"name": "Jean", "surname": "Bonbeur"}}, "mail": "Jean.Bonbeur@gmail.com"}
	assert.Equal(t, waited, result, "Should create the right mail")
}

func TestRegistryMaskToConfigurationShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Template: "{{.name}}.{{.surname}}@gmail.com"}}
	config, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	assert.Nil(t, err, "error should be nil")
	mask, _ := NewMask("{{.name}}.{{.surname}}@gmail.com")
	waitedConfig := model.NewMaskConfiguration().WithEntry("", mask)
	assert.Equal(t, waitedConfig, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestRegistryMaskToConfigurationShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}

func TestRegistryMaskToConfigurationShouldReturnAnErrorInWrongConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Template: "{{.name}.{{.surname}}@gmail.com"}}
	mask, present, err := RegistryMaskToConfiguration(maskingConfig, model.NewMaskConfiguration(), 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be true")
	assert.NotNil(t, err, "error shouldn't be nil")
}
