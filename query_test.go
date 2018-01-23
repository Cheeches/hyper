package hyper_test

import (
	"reflect"
	"testing"

	"github.com/cognicraft/hyper"
)

func TestQuery(t *testing.T) {

	root := hyper.Item{
		Label:       "ROOT",
		Description: "This is the root item.",
		ID:          "root",
		Rel:         "item",
		Type:        "test-item",
		Properties: hyper.Properties{
			{
				Name:  "foo",
				Value: "foo-val",
			},
		},
		Data: 3,
		Links: hyper.Links{
			{
				Rel: "details",
			},
		},
		Actions: hyper.Actions{
			{
				Rel: "create",
			},
		},
		Items: hyper.Items{
			{
				ID: "root:1",
			},
		},
		Errors: hyper.Errors{
			{
				Code:    "33",
				Message: "something bad happened",
			},
		},
		Render: hyper.RenderNone,
	}

	tests := []struct {
		root   hyper.Item
		q      string
		result interface{}
	}{
		{
			root:   root,
			q:      "",
			result: nil,
		},
		{
			root:   root,
			q:      "foo",
			result: nil,
		},
		{
			root:   root,
			q:      ".label",
			result: "ROOT",
		},
		{
			root:   root,
			q:      ".description",
			result: "This is the root item.",
		},
		{
			root:   root,
			q:      ".render",
			result: hyper.RenderNone,
		},
		{
			root:   root,
			q:      ".",
			result: root,
		},
		{
			root:   root,
			q:      ".id",
			result: "root",
		},
		{
			root:   root,
			q:      ".rel",
			result: "item",
		},
		{
			root:   root,
			q:      ".type",
			result: "test-item",
		},
		{
			root: root,
			q:    ".properties",
			result: hyper.Properties{
				{
					Name:  "foo",
					Value: "foo-val",
				},
			},
		},
		{
			root:   root,
			q:      ".data",
			result: 3,
		},
		{
			root: root,
			q:    ".links",
			result: hyper.Links{
				{
					Rel: "details",
				},
			},
		},
		{
			root: root,
			q:    ".actions",
			result: hyper.Actions{
				{
					Rel: "create",
				},
			},
		},
		{
			root: root,
			q:    ".items",
			result: hyper.Items{
				{
					ID: "root:1",
				},
			},
		},
		{
			root: root,
			q:    ".errors",
			result: hyper.Errors{
				{
					Code:    "33",
					Message: "something bad happened",
				},
			},
		},
		{
			root: root,
			q:    "#root:1",
			result: hyper.Item{
				ID: "root:1",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.q, func(t *testing.T) {
			got := hyper.Query(test.root, test.q)
			if !reflect.DeepEqual(test.result, got) {
				t.Errorf("want: %s, got: %s", hyper.JSONString(test.result), hyper.JSONString(got))
			}
		})
	}
}
