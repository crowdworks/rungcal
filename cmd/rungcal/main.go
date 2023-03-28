package main

import (
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/ciloholic/rungcal"
)

var (
	app           = kingpin.New("rungcal", "A command-line rungcal application")
	cliTargetDate = app.Flag("target-date", "Targate date(yyyy-mm-dd)").Required().String()
	cliProject    = app.Flag("project", "Filter by Project Name").String()
	cliVerbose    = app.Flag("verbose", "Set verbose mode").Bool()
	cliDryRun     = app.Flag("dry-run", "Set dry run mode").Bool()

	insertCommand  = app.Command("insert", "Insert function")
	insertRecreate = insertCommand.Flag("recreate", "recreate").Default("false").Bool()

	deleteCommand = app.Command("delete", "Delete function")
)

func main() {
	os.Exit(_main())
}

func _main() int {
	command := kingpin.MustParse(app.Parse(os.Args[1:]))
	insertOption := rungcal.InsertOption{
		Option:   rungcal.Option{TargetDate: *cliTargetDate, Project: *cliProject, Verbose: *cliVerbose, DryRun: *cliDryRun},
		ReCreate: *insertRecreate,
	}
	deleteOption := rungcal.DeleteOption{
		Option: rungcal.Option{TargetDate: *cliTargetDate, Project: *cliProject, Verbose: *cliVerbose, DryRun: *cliDryRun},
	}

	switch command {
	case insertCommand.FullCommand():
		return rungcal.Insert(insertOption)
	case deleteCommand.FullCommand():
		// TODO: カレンダーのイベント削除が手間なので、サブコマンドで実行出来るようにしたい
		return rungcal.Delete(deleteOption)
	}

	return 0
}
