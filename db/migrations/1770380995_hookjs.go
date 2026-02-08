package migrations

import (
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
	"github.com/shynome/err0"
	"github.com/shynome/err0/try"
	"remoon.net/well/db"
)

func init() {
	migrations.Register(func(app core.App) (err error) {
		defer err0.Then(&err, nil, nil)

		hookjs := core.NewBaseCollection(db.TableHookJS, ID(db.TableHookJS))
		hookjs.System = true
		hookjs.Fields.Add(
			&core.TextField{
				Name: "name", Id: ID("name"), System: true,
				Required: false, Presentable: true,
			},
			&core.NumberField{
				Name: "order", Id: ID("order"), System: true,
				OnlyInt: true,
			},
			&core.BoolField{
				Name: "disabled", Id: ID("disabled"), System: true,
			},
			&core.EditorField{
				Name: "hookjs", Id: ID("hookjs"), System: true,
				Required: true,
			},
		)
		try.To(app.Save(hookjs))

		return nil
	}, func(app core.App) error {
		return fmt.Errorf("no rollback")
	})
}
