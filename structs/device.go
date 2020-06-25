package structs

const (
	TYPE_OUTLET     string = "action.devices.types.OUTLET"
	TYPE_THERMOSTAT string = "action.devices.types.THERMOSTAT"
)

const (
	TRAITS_ON_OFF              string = "action.devices.traits.OnOff"
	TRAITS_TEMPERATURE_SETTING string = "action.devices.traits.TemperatureSetting"
)

// SYNC
type Device struct {
	Id              string                 `json:"id"`
	Type            string                 `json:"type"`
	Traits          []string               `json:"traits"`
	Name            DeviceName             `json:"name"`
	WillReportState bool                   `json:"willReportState"`
	RoomHint        string                 `json:"roomHint"`
	DeviceInfo      DeviceInfo             `json:"deviceInfo"`
	Attributes      map[string]interface{} `json:"attributes"`
}

type DeviceName struct {
	DefaultNames []string `json:"defaultNames"`
	Name         string   `json:"name"`
	Nicknames    []string `json:"nicknames"`
}

type DeviceInfo struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	HwVersion    string `json:"hwVersion"`
	SwVersion    string `json:"swVersion"`
}

func (d *Device) Copy(Id string, Name string) Device {
	return Device{
		Id:     Id,
		Type:   d.Type,
		Traits: d.Traits,
		Name: DeviceName{
			DefaultNames: []string{Name},
			Name:         Name,
		},
		WillReportState: d.WillReportState,
		RoomHint:        d.RoomHint,
		DeviceInfo:      d.DeviceInfo,
	}
}

// EXECUTE
type Command struct {
	Devices   []DeviceCommand `json:"devices"`
	Execution []Execution     `json:"execution"`
}

type DeviceCommand struct {
	Id         string                 `json:"id"`
	CustomData map[string]interface{} `json:"customData"`
}

type Execution struct {
	Command string                 `json:"command"`
	Params  map[string]interface{} `json:"params"`
}
