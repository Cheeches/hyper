package hyper

import (
	"strings"
)

func Query(root Item, q string) interface{} {
	switch q {
	case "":
		return nil
	case ".":
		return root
	case ".label":
		return root.Label
	case ".description":
		return root.Description
	case ".render":
		return root.Render
	case ".id":
		return root.ID
	case ".rel":
		return root.Rel
	case ".type":
		return root.Type
	case ".properties":
		return root.Properties
	case ".data":
		return root.Data
	case ".links":
		return root.Links
	case ".actions":
		return root.Actions
	case ".items":
		return root.Items
	case ".errors":
		return root.Errors
	default:
		if strings.HasPrefix(q, "#") {
			if item, ok := Search(root, q[1:]); ok {
				return item
			}
		}
		return nil
	}
}
