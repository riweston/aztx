package state

import "github.com/spf13/viper"

// StateManager handles all state operations
type StateManager interface {
	GetLastContext() (id string, name string)
	SetLastContext(id string, name string) error
}

type ViperStateManager struct {
	viper *viper.Viper
}

func NewViperStateManager(v *viper.Viper) *ViperStateManager {
	return &ViperStateManager{viper: v}
}

func (v *ViperStateManager) GetLastContext() (string, string) {
	return v.viper.GetString("lastContextId"),
		v.viper.GetString("lastContextDisplayName")
}

func (v *ViperStateManager) SetLastContext(id string, name string) error {
	v.viper.Set("lastContextId", id)
	v.viper.Set("lastContextDisplayName", name)
	return v.viper.WriteConfig()
}
