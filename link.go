package hyper

// Link .
type Link struct {
	Label          string     `json:"label,omitempty"`
	Description    string     `json:"description,omitempty"`
	Rel            string     `json:"rel"`
	Href           string     `json:"href,omitempty"`
	Type           string     `json:"type,omitempty"`
	Language       string     `json:"language,omitempty"`
	Template       string     `json:"template,omitempty"`
	Parameters     Parameters `json:"parameters,omitempty"`
	Context        string     `json:"context,omitempty"`
	Render         string     `json:"render,omitempty"`
	Accept         string     `json:"accept,omitempty"`
	AcceptLanguage string     `json:"accept-language,omitempty"`
}

// Links .
type Links []Link

// FindByRel .
func (ls Links) FindByRel(rel string) (Link, bool) {
	for _, l := range ls {
		if l.Rel == rel {
			return l, true
		}
	}
	return Link{}, false
}

//FindRelated .
func (ls Links) FindRelated(rel string) Links {
	result := Links{}
	for _, l := range ls {
		if l.Rel == rel {
			result = append(result, l)
		}
	}
	return result
}

// Filter .
func (ls Links) Filter(f func(Link) bool) Links {
	filtered := []Link{}
	for _, l := range ls {
		if f(l) {
			filtered = append(filtered, l)
		}
	}
	return filtered
}
