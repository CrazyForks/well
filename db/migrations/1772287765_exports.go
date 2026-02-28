package migrations

import (
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
	"remoon.net/well/db"
)

func init() {
	migrations.Register(func(app core.App) (err error) {
		defer err0.Then(&err, nil, nil)

		exports := core.NewBaseCollection(db.TableExports, ID(db.TableExports))
		exports.Fields.Add(
			&core.TextField{
				Name: "name", Id: ID("name"), System: true,
				Required: false, Presentable: true,
			},
			&core.BoolField{
				Name: "disabled", Id: ID("disabled"), System: true,
			},
			&core.BoolField{
				Name: "disabled", Id: ID("disabled"), System: true,
			},
			&core.SelectField{
				Name: "protocol", Id: ID("protocol"), System: true,
				Required: true,
				Values:   []string{"TCP", "UDP"}, MaxSelect: 1,
			},
			&core.NumberField{
				Name: "port", Id: ID("port"), System: true,
				Required: true,
				OnlyInt:  true,
				Min:      types.Pointer[float64](0),
				Max:      types.Pointer[float64](65535),
			},
		)
		try.To(app.Save(exports))

		return nil
	}, func(app core.App) error {
		return fmt.Errorf("no rollback")
	})
}
