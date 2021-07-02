package group_test

import (
	"encoding/json"
	"testing"

	"github.com/wtask-go/mixpanel/ingestion/group"
	"github.com/wtask-go/mixpanel/internal/assets"
)

func Test_json_schema(t *testing.T) {
	cases := []struct {
		data interface{}
		ok   bool
	}{
		{
			&group.Set{}, false,
		},
		{
			&group.Set{Set: map[string]interface{}{
				"key": "value",
			}}, true,
		},
		{
			&group.SetOnce{}, false,
		},
		{
			&group.SetOnce{SetOnce: map[string]interface{}{
				"key": "value",
			}}, true,
		},
		{
			&group.ListAppend{}, false,
		},
		{
			&group.ListAppend{Append: map[string]interface{}{
				"list_key": "value",
			}}, true,
		},
		{
			&group.ListRemove{}, false,
		},
		{
			&group.ListRemove{Remove: map[string]interface{}{
				"list_key": "value",
			}}, true,
		},
		{
			&group.Unset{}, false,
		},
		{
			&group.Unset{Unset: []string{
				"numeric_key",
			}}, true,
		},
	}

	schema := assets.MustCompileSchema("openapi/engage.schema.json")

	for i, c := range cases {
		data, err := json.Marshal(c.data)
		if err != nil {
			t.Fatalf("[#%d] %s", i, err)
		}

		err = assets.ValidateJSON(schema, data)

		switch {
		case err != nil && c.ok:
			t.Fatalf("[#%d] unexpected error: %s, data: %+v, json: %s", i, err, c.data, data)
		case err == nil && !c.ok:
			t.Fatalf("[#%d] unexpectedly valid json %s", i, data)
		}
	}
}
