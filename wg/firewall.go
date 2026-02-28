package wg

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"remoon.net/well/db"
	"remoon.net/well/wg/firewall"
)

var protoAllows = firewall.NewProtoAllows()

type empty = struct{}

func InitFirewall(app core.App) error {
	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		protoAllows.Reset()
		records, err := app.FindAllRecords(db.TableExports, dbx.HashExp{"disabled": false})
		if err != nil {
			return err
		}
		for _, r := range records {
			addAllows(r)
		}
		return e.Next()
	})
	app.OnRecordAfterCreateSuccess(db.TableExports).BindFunc(func(e *core.RecordEvent) error {
		addAllows(e.Record)
		return e.Next()
	})
	app.OnRecordUpdate(db.TableExports).BindFunc(func(e *core.RecordEvent) error {
		r, err := e.App.FindRecordById(db.TableExports, e.Record.Id)
		if err != nil {
			return err
		}
		removeAllows(r)
		addAllows(e.Record)
		return e.Next()
	})
	app.OnRecordAfterDeleteSuccess(db.TableExports).BindFunc(func(e *core.RecordEvent) error {
		removeAllows(e.Record)
		return e.Next()
	})
	return nil
}

func addAllows(r *core.Record) {
	if r.GetBool("disabled") {
		return
	}
	p := r.GetString("protocol")
	port := uint16(r.GetInt("port"))
	switch p {
	case "TCP":
		protoAllows.TCP.Set(port, empty{})
	case "UDP":
		protoAllows.UDP.Set(port, empty{})
	}
}

func removeAllows(r *core.Record) {
	p := r.GetString("protocol")
	port := uint16(r.GetInt("port"))
	switch p {
	case "TCP":
		protoAllows.TCP.Remove(port)
	case "UDP":
		protoAllows.UDP.Remove(port)
	}
}
