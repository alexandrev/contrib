package datetimeExt

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
	"github.com/tkuchiki/parsetime"
)

type IsoDateToUnix struct {
}

func init() {
	function.Register(&IsoDateToUnix{})
}

func (s *IsoDateToUnix) Name() string {
	return "isodatetounix"
}

func (s *IsoDateToUnix) GetCategory() string {
	return "datetimeExt"
}

func (s *IsoDateToUnix) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *IsoDateToUnix) Eval(params ...interface{}) (interface{}, error) {
	date, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("Format date first argument must be string")
	}

	p, err := parsetime.NewParseTime()
	if err != nil {
		log.RootLogger().Errorf("New date parser %s error %s", date, err.Error())
		return date, err
	}
	t, err := p.US(date)
	if err != nil {
		//Try toresolved it with just parse
		t, err = p.Parse(date)
		if err != nil {
			log.RootLogger().Errorf("Parsing date %s error %s", date, err.Error())
			return date, err
		}
	}
	return t.Unix(), nil
}
