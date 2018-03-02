package edit

import (
	"github.com/elves/elvish/edit/ui"
)

// Styles for UI.
var (
	styleForReplacement    = ui.Styles{"underlined"}
	styleForTip            = ui.Styles{}
	styleForSelected       = ui.Styles{"inverse"}
	styleForScrollBarArea  = ui.Styles{"magenta"}
	styleForScrollBarThumb = ui.Styles{"magenta", "inverse"}

	styleForCompilerError = ui.Styles{"white", "bg-red"}

	// Use default style for completion listing
	styleForCompletion = ui.Styles{}
	// Use inverse style for selected completion entry
	styleForSelectedCompletion = ui.Styles{"inverse"}
)
