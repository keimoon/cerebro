package tacit

type Context map[string]interface{}
type Analyzer func(question string, context Context) (string, error)
type SearchEngine func(keyword string, context Context) (string, error)
type Synthesizer func(document string, context Context) (string, error)
type SearchEngineChooser func(context Context) SearchEngine
type SynthesizerChooser func(context Context) Synthesizer

type Brain struct {
	Context             Context
	Analyzer            Analyzer
	SearchEngineChooser SearchEngineChooser
	SynthesizerChooser  SynthesizerChooser
}

func (b *Brain) Ask(question string) (string, error) {
	b.Context = make(Context)
	keyword, err := b.Analyzer(question, b.Context)
	if err != nil {
		return "", err
	}
	document, err := b.SearchEngineChooser(b.Context)(keyword, b.Context)
	if err != nil {
		return "", err
	}
	return b.SynthesizerChooser(b.Context)(document, b.Context)
}
