package azure_state

import "github.com/spf13/viper"

type stateReaderWriter interface {
	Read(key string) string
	Write(key, value string)
}

type LastContext struct {
	rw stateReaderWriter
}

func NewStateReaderWriter(rw stateReaderWriter) *LastContext {
	return &LastContext{
		rw: rw,
	}
}

func (lc *LastContext) ReadLastContextId() string {
	return lc.rw.Read("lastContextId")

}

func (lc *LastContext) ReadLastContextDisplayName() string {
	return lc.rw.Read("lastContextDisplayName")
}

func (lc *LastContext) WriteLastContext(id string, name string) {
	lc.rw.Write("lastContextId", id)
	lc.rw.Write("lastContextDisplayName", name)
}

type ViperAdapter struct {
	Viper *viper.Viper
}

func (v *ViperAdapter) Read(key string) string {
	return v.Viper.GetString(key)
}

func (v *ViperAdapter) Write(key, value string) {
	v.Viper.Set(key, value)
	if err := v.Viper.WriteConfig(); err != nil {
		panic(err)
	}
}
