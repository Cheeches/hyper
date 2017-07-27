package hyper

// Action .
type Action struct {
	Label        string     `json:"label,omitempty"`
	Description  string     `json:"description,omitempty"`
	Rel          string     `json:"rel"`
	Href         string     `json:"href,omitempty"`
	Encoding     string     `json:"encoding,omitempty"`
	Method       string     `json:"method,omitempty"`
	Template     string     `json:"template,omitempty"`
	Parameters   Parameters `json:"parameters,omitempty"`
	Context      string     `json:"context,omitempty"`
	OK           string     `json:"ok,omitempty"`
	Cancel       string     `json:"cancel,omitempty"`
	Confirmation string     `json:"confirmation,omitempty"`
	Render       string     `json:"render,omitempty"`
}

// Actions .
type Actions []Action

func (as Actions) Len() int {
	return len(as)
}

func (as Actions) Less(i, j int) bool {
	return as[i].Label < as[j].Label
}

func (as Actions) Swap(i, j int) {
	as[i], as[j] = as[j], as[i]
}

// FindByRel .
func (as Actions) FindByRel(rel string) (Action, bool) {
	for _, l := range as {
		if l.Rel == rel {
			return l, true
		}
	}
	return Action{}, false
}

// Filter .
func (as Actions) Filter(f func(Action) bool) Actions {
	filtered := []Action{}
	for _, a := range as {
		if f(a) {
			filtered = append(filtered, a)
		}
	}
	return filtered
}

const (
	MethodPOST   = "POST"
	MethodPATCH  = "PATCH"
	MethodDELETE = "DELETE"
)
