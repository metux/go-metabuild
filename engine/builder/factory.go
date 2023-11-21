package builder

import (
	"fmt"

	"github.com/metux/go-metabuild/engine/builder/c"
	"github.com/metux/go-metabuild/engine/builder/data"
	"github.com/metux/go-metabuild/engine/builder/doc"
	"github.com/metux/go-metabuild/engine/builder/i18n"
	"github.com/metux/go-metabuild/spec/target"
	"github.com/metux/go-metabuild/util/jobs"
)

func CreateBuilder(o target.TargetObject, id string) (jobs.Job, error) {
	switch t := o.Type(); t {

	/* plain C */
	case target.TypeCExecutable:
		return c.MakeBuilderCExecutable(o, id), nil
	case target.TypeCLibrary:
		return c.MakeBuilderCLibrary(o, id), nil

	/* C++ */
	case target.TypeCxxExecutable:
		return c.MakeBuilderCExecutable(o, id), nil
	case target.TypeCxxLibrary:
		return c.MakeBuilderCLibrary(o, id), nil

	/* data files */
	case target.TypeDataMisc:
		return data.MakeDataMisc(o, id), nil
	case target.TypeDataDesktop:
		return data.MakeDataDesktop(o, id), nil
	case target.TypeDataPixmaps:
		return data.MakeDataMisc(o, id), nil

	/* locales */
	case target.TypeI18nPo:
		return i18n.MakeI18nPo(o, id), nil

	/* documentation */
	case target.TypeDocMan:
		return doc.MakeManPages(o, id), nil
	case target.TypeDocMisc:
		return doc.MakeDocMisc(o, id), nil

	default:
		return nil, fmt.Errorf("unsupported target type: %s", t)
	}
}
