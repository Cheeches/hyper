package hyper

// Property represents part of a domain state
type Property struct {
	Label       string      `json:"label,omitempty"`
	Description string      `json:"description,omitempty"`
	Render      string      `json:"render,omitempty"`
	Name        string      `json:"name"`
	Value       interface{} `json:"value"`
	Type        string      `json:"type,omitempty"`
	Unit        string      `json:"unit,omitempty"`
	Display     string      `json:"display,omitempty"`
}

// Properties is a collection of Property
type Properties []Property

// Find returns a Property by name
func (ps Properties) Find(name string) (Property, bool) {
	for _, p := range ps {
		if p.Name == name {
			return p, true
		}
	}
	return Property{}, false
}

// KeyBy calculates a map keyed by the result of the extractKey funktion.
func (ps Properties) KeyBy(extractKey func(Property) string) map[string]Property {
	if len(ps) == 0 {
		return nil
	}
	m := map[string]Property{}
	for _, p := range ps {
		key := extractKey(p)
		m[key] = p
	}
	return m
}

// KeyByName returns a map of Properties keyed by name
func (ps Properties) KeyByName() map[string]Property {
	return ps.KeyBy(func(p Property) string {
		return p.Name
	})
}
