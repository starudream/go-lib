package cron

import (
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/core/v2/utils/signalutil"
)

type (
	Cron    = cron.Cron
	Entry   = cron.Entry
	EntryID = cron.EntryID

	Job = cron.Job
)

var (
	_c     *Cron
	_cem   sync.Map // key: name
	_cnm   sync.Map // key: EntryID, value: name
	_cOnce sync.Once
)

func C() *Cron {
	_cOnce.Do(func() {
		_c = cron.New(
			cron.WithLocation(time.Local),
			cron.WithSeconds(),
			cron.WithLogger(_logger),
		)
	})
	return _c
}

func AddJob(spec, name string, job func()) error {
	if name == "" {
		return fmt.Errorf("name is empty")
	}

	if _, loaded := _cem.LoadOrStore(name, true); loaded {
		return fmt.Errorf("job %s already exists", name)
	}

	var entry Entry

	cmd := cron.FuncJob(func() {
		var err error
		defer func() {
			if r := recover(); r != nil && err == nil {
				err = fmt.Errorf("[%s] %v", osutil.CallerString(1), r)
			}
			if err != nil {
				_logger.E("failed", "entry", entry.ID, "err", err)
			}
		}()
		job()
	})

	id, err := C().AddJob(spec, cmd)
	if err != nil {
		_cem.Delete(name)
		return err
	}

	entry = C().Entry(id)

	_cnm.Store(id, name)

	return nil
}

func Remove(name string) {
	_cnm.Range(func(_id, _name any) bool {
		if _name.(string) == name {
			C().Remove(_id.(EntryID))
			_cnm.Delete(_id)
			_cem.Delete(_name)
			return false
		}
		return true
	})
}

func Run() {
	C().Start()
	<-signalutil.Done()
	C().Stop()
}
