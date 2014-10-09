package cerebro

import (
	"github.com/keimoon/cerebro/analyzer"
	"github.com/keimoon/cerebro/search"
	"github.com/keimoon/cerebro/synthesizer"
	"github.com/keimoon/cerebro/tacit"
)

var Cerebro = &tacit.Brain{
	Analyzer:            analyzer.Simple,
	SearchEngineChooser: search.SimpleChooser,
	SynthesizerChooser:  synthesizer.SimpleChooser,
}
