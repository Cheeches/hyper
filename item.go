package hyper

import "encoding/json"

// Item has properties, links, actions and (sub-)actions.
type Item struct {
	Label       string     `json:"label,omitempty"`
	Description string     `json:"description,omitempty"`
	Rel         string     `json:"rel,omitempty"`
	ID          string     `json:"id,omitempty"`
	Type        string     `json:"type,omitempty"`
	Properties  Properties `json:"properties,omitempty"`
	Data        Data       `json:"data,omitempty"`
	Links       Links      `json:"links,omitempty"`
	Actions     Actions    `json:"actions,omitempty"`
	Items       Items      `json:"items,omitempty"`
	Errors      Errors     `json:"errors,omitempty"`
	Render      string     `json:"render,omitempty"`
}

// Search performs a DFS with the goal to find an Item by the specified id
func (i Item) Search(id string) (Item, bool) {
	frontier := []Item{i}
	for len(frontier) > 0 {
		var next Item
		next, frontier = frontier[0], frontier[1:]
		if next.ID == id {
			return next, true
		}
		frontier = append(frontier, next.Items...)
	}
	return Item{}, false
}

// AddProperty add a Property to this Item
func (i *Item) AddProperty(p Property) {
	i.Properties = append(i.Properties, p)
}

// AddProperties adds many Properties to this Item
func (i *Item) AddProperties(ps Properties) {
	i.Properties = append(i.Properties, ps...)
}

// AddItem adds a (sub-)Item to this Item
func (i *Item) AddItem(sub Item) {
	i.Items = append(i.Items, sub)
}

// AddItems adds many (sub-)Items to this Item
func (i *Item) AddItems(subs []Item) {
	i.Items = append(i.Items, subs...)
}

// AddLink adds a link to this Item
func (i *Item) AddLink(l Link) {
	i.Links = append(i.Links, l)
}

// AddLinks adds many links to this Item
func (i *Item) AddLinks(ls Links) {
	i.Links = append(i.Links, ls...)
}

// AddAction adds an Action to this Item
func (i *Item) AddAction(a Action) {
	i.Actions = append(i.Actions, a)
}

// AddActions adds many Actions to this Item
func (i *Item) AddActions(as Actions) {
	i.Actions = append(i.Actions, as...)
}

// String returns string representation
func (i Item) String() string {
	bs, _ := json.MarshalIndent(i, "", "  ")
	return string(bs)
}

// Items represents a collection of Item
type Items []Item

// Find returns an Item that satifies the specification
func (is Items) Find(isSatisfiedBy func(Item) bool) (Item, bool) {
	for _, i := range is {
		if isSatisfiedBy(i) {
			return i, true
		}
	}
	return Item{}, false
}

func (is Items) KeyBy(extractKey func(Item) string) map[string]Item {
	if len(is) == 0 {
		return nil
	}
	m := map[string]Item{}
	for _, i := range is {
		key := extractKey(i)
		m[key] = i
	}
	return m
}

// FindByID returns an Item that has a specific id
func (is Items) FindByID(id string) (Item, bool) {
	return is.Find(ItemIDEquals(id))
}

// KeyByID returns a map of Items keyed by the Item ids
func (is Items) KeyByID() map[string]Item {
	return is.KeyBy(func(i Item) string {
		return i.ID
	})
}

// FindByRel returns an Item that has a specific rel
func (is Items) FindByRel(rel string) (Item, bool) {
	return is.Find(ItemRelEquals(rel))
}

// KeyByRel returns a map of Items keyed by the Item rel
func (is Items) KeyByRel() map[string]Item {
	return is.KeyBy(func(i Item) string {
		return i.Rel
	})
}

// Filter returns a collection of Items that conform to the profided specification
func (is Items) Filter(isSatisfiedBy func(Item) bool) Items {
	filtered := []Item{}
	for _, i := range is {
		if isSatisfiedBy(i) {
			filtered = append(filtered, i)
		}
	}
	return filtered
}

// ItemIDEquals is used to Filter a collection of Items by id
func ItemIDEquals(id string) func(Item) bool {
	return func(i Item) bool {
		return id == i.ID
	}
}

// ItemRelEquals is used to Filter a collection of Items by rel
func ItemRelEquals(rel string) func(Item) bool {
	return func(i Item) bool {
		return rel == i.Rel
	}
}

const (
	// RenderNone is used on links or items that should not be displayed in the UI
	RenderNone = "none"
	// RenderTransclude is used on links or items to signal that these should be embedded within the current view.
	RenderTransclude = "transclude"
)

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

// Property represents part of a domain state
type Property struct {
	Name        string      `json:"name"`
	Value       interface{} `json:"value"`
	Type        string      `json:"type,omitempty"`
	Unit        string      `json:"unit,omitempty"`
	Label       string      `json:"label,omitempty"`
	Description string      `json:"description,omitempty"`
	Display     string      `json:"display,omitempty"`
	Render      string      `json:"render,omitempty"`
}

// Data .
type Data interface{}

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

// Parameter .
type Parameter struct {
	Label       string        `json:"label,omitempty"`
	Description string        `json:"description,omitempty"`
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	Value       interface{}   `json:"value,omitempty"`
	Placeholder string        `json:"placeholder,omitempty"`
	Options     SelectOptions `json:"options,omitempty"`
	Related     string        `json:"related,omitempty"`
	Components  interface{}   `json:"components,omitempty"`
	Pattern     string        `json:"pattern,omitempty"`
	Min         interface{}   `json:"min,omitempty"`
	Max         interface{}   `json:"max,omitempty"`
	MaxLength   interface{}   `json:"max-length,omitempty"`
	Size        interface{}   `json:"size,omitempty"`
	Step        interface{}   `json:"step,omitempty"`
	Cols        interface{}   `json:"cols,omitempty"`
	Rows        interface{}   `json:"rows,omitempty"`
	Required    bool          `json:"required,omitempty"`
	ReadOnly    bool          `json:"read-only,omitempty"`
	Multiple    bool          `json:"multiple,omitempty"`
}

// Parameters .
type Parameters []Parameter

// FindByName .
func (as Parameters) FindByName(name string) (Parameter, bool) {
	for _, l := range as {
		if l.Name == name {
			return l, true
		}
	}
	return Parameter{}, false
}

// SelectOption .
type SelectOption struct {
	Label       string         `json:"label,omitempty"`
	Description string         `json:"description,omitempty"`
	Value       interface{}    `json:"value,omitempty"`
	Options     []SelectOption `json:"options,omitempty"`
}

// SelectOptions .
type SelectOptions []SelectOption

func (s SelectOptions) Len() int {
	return len(s)
}

func (s SelectOptions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SelectOptions) Less(i, j int) bool {
	return s[i].Label < s[j].Label
}

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

const (
	MethodPOST   = "POST"
	MethodPATCH  = "PATCH"
	MethodDELETE = "DELETE"
)

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

// Error .
type Error struct {
	Label       string `json:"label,omitempty"`
	Description string `json:"label,omitempty"`
	Message     string `json:"message"`
	Code        string `json:"code,omitempty"`
}

// Errors .
type Errors []Error

const (
	// TypeHidden is for hidden Parameters
	TypeHidden = "hidden"
	TypeText   = "text"
)
