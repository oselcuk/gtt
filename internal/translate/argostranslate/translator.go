package argostranslate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/eeeXun/gtt/internal/translate/core"
)

const (
	textURL = "https://translate.argosopentech.com/translate"
)

type ArgosTranslate struct {
	*core.Language
	*core.TTSLock
	core.EngineName
}

func NewArgosTranslate() *ArgosTranslate {
	return &ArgosTranslate{
		Language:   core.NewLanguage(),
		TTSLock:    core.NewTTSLock(),
		EngineName: core.NewEngineName("ArgosTranslate"),
	}
}

func (t *ArgosTranslate) GetAllLang() []string {
	return lang
}

func (t *ArgosTranslate) Translate(message string) (translation *core.Translation, err error) {
	translation = new(core.Translation)
	var data map[string]interface{}

	res, err := http.PostForm(textURL,
		url.Values{
			"q":      {message},
			"source": {langCode[t.GetSrcLang()]},
			"target": {langCode[t.GetDstLang()]},
		})
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if len(data) <= 0 {
		return nil, errors.New("Translation not found")
	}

	translation.TEXT = fmt.Sprintf("%v", data["translatedText"])

	return translation, nil
}

func (t *ArgosTranslate) PlayTTS(lang, message string) error {
	defer t.ReleaseLock()

	return errors.New(t.GetEngineName() + " does not support text to speech")
}
