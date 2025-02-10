module github.com/riweston/aztx/pkg/profile

go 1.21

require (
	github.com/google/uuid v1.6.0
	github.com/riweston/aztx/pkg/errors v0.0.0
	github.com/riweston/aztx/pkg/state v0.0.0
	github.com/riweston/aztx/pkg/storage v0.0.0
	github.com/riweston/aztx/pkg/subscription v0.0.0
	github.com/riweston/aztx/pkg/tenant v0.0.0
	github.com/riweston/aztx/pkg/types v0.0.0
	github.com/spf13/viper v1.19.0
)

replace (
	github.com/riweston/aztx/pkg/errors => ../errors
	github.com/riweston/aztx/pkg/state => ../state
	github.com/riweston/aztx/pkg/storage => ../storage
	github.com/riweston/aztx/pkg/subscription => ../subscription
	github.com/riweston/aztx/pkg/tenant => ../tenant
	github.com/riweston/aztx/pkg/types => ../types
)

require (
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/gdamore/tcell/v2 v2.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/ktr0731/go-ansisgr v0.1.0 // indirect
	github.com/ktr0731/go-fuzzyfinder v0.8.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nsf/termbox-go v1.1.1 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/term v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
