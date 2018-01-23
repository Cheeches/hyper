package hyper

const (
	// TypeHidden is for hidden parameters
	TypeHidden = "hidden"
	// TypeText is used for parameters of type text/string
	TypeText = "text"
)

const (
	// RenderNone is used on links or items that should not be displayed in the UI
	RenderNone = "none"
	// RenderTransclude is used on links or items to signal that these should be embedded within the current view.
	RenderTransclude = "transclude"
)
