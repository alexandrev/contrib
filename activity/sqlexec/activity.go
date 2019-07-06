package sqlexec

import (
	"database/sql"

	"github.com/project-flogo/contrib/activity/sqlquery/util"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

const (
	ovResults = "results"
)

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{MaxIdleConns: 2}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	dbHelper, err := util.GetDbHelper(s.DbType)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("DB: '%s'", s.DbType)

	// todo move this to a shared connection object
	db, err := getConnection(s)
	if err != nil {
		return nil, err
	}

	sqlStatement, err := util.NewSQLStatement(dbHelper, s.Query)
	if err != nil {
		return nil, err
	}

	act := &Activity{db: db, dbHelper: dbHelper, sqlStatement: sqlStatement}

	if !s.DisablePrepared {
		ctx.Logger().Debugf("Using PreparedStatement: %s", sqlStatement.PreparedStatementSQL())
		act.stmt, err = db.Prepare(sqlStatement.PreparedStatementSQL())
		if err != nil {
			return nil, err
		}
	}

	return act, nil
}

// Activity is a Counter Activity implementation
type Activity struct {
	dbHelper       util.DbHelper
	db             *sql.DB
	sqlStatement   *util.SQLStatement
	stmt           *sql.Stmt
	labeledResults bool
}

// Metadata implements activity.Activity.Metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func (a *Activity) Cleanup() error {
	if a.stmt != nil {
		err := a.stmt.Close()
		log.RootLogger().Warnf("error cleaning up SQL Query activity: %v", err)
	}

	log.RootLogger().Tracef("cleaning up SQL Query activity")

	return a.db.Close()
}

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	in := &Input{}
	err = ctx.GetInputObject(in)
	if err != nil {
		return false, err
	}

	results, err := a.doExec(in.Params)
	if err != nil {
		return false, err
	}

	err = ctx.SetOutput(ovResults, results)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *Activity) doExec(params map[string]interface{}) (interface{}, error) {

	var err error

	var result interface{}

	if a.stmt != nil {
		args := a.sqlStatement.GetPreparedStatementArgs(params)
		result, err = a.stmt.Exec(args...)
	} else {
		result, err = a.db.Exec(a.sqlStatement.ToStatementSQL(params))
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

//todo move to shared connection
func getConnection(s *Settings) (*sql.DB, error) {

	db, err := sql.Open(s.DriverName, s.DataSourceName)
	if err != nil {
		return nil, err
	}

	if s.MaxOpenConns > 0 {
		db.SetMaxOpenConns(s.MaxOpenConns)
	}

	if s.MaxIdleConns != 2 {
		db.SetMaxIdleConns(s.MaxIdleConns)
	}

	return db, err
}
