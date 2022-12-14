package main

import (
	"context"
	"time"

	"github.com/starudream/go-lib/app"
	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/httpx"
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/randx"
	"github.com/starudream/go-lib/seq"
)

func main() {
	log.Attach("app", "example-simple")
	app.Init(func() error { log.Info().Msg("init"); return nil })
	app.Add(wrapError(TestAppTime))
	app.Add(wrapError(TestConfig))
	app.Add(wrapError(TestHTTPX))
	app.Add(wrapError(TestRandX))
	app.Add(wrapError(TestSeq))
	app.Defer(TestDefer)
	err := app.OnceGo()
	if err != nil {
		panic(err)
	}
	log.Info().Msg("success")
}

func wrapError(f func()) func(ctx context.Context) error {
	return func(ctx context.Context) error { f(); return nil }
}

func TestAppTime() {
	log.Info().Msgf("startup: %v", app.StartupTime().Format(time.RFC3339Nano))
	log.Info().Msgf("running: %v", app.RunningTime().Format(time.RFC3339Nano))
	log.Info().Msgf("cost: %v", app.CostTime())
}

func TestConfig() {
	log.Warn().Msgf("debug: %v", config.GetBool("debug"))
}

func TestHTTPX() {
	httpx.SetTimeout(3 * time.Second)
	httpx.SetUserAgent("go")
	resp, err := httpx.R().Get("https://www.gstatic.com/generate_204")
	if err != nil {
		panic(err)
	}
	log.Debug().Msgf("response status: %s", resp.Status())
}

func TestRandX() {
	log.Info().Msgf("fake: %s", randx.F().LetterN(16))
}

func TestSeq() {
	log.Info().Msgf("sonyflake: %s", seq.NextId())
	log.Info().Msgf("uuid: %s", seq.UUID())
	log.Info().Msgf("uuid short: %s", seq.UUIDShort())
}

func TestDefer() {
	log.Info().Msg("bye bye")
}
