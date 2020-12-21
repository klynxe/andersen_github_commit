package setting

import "time"

const SettingConnectName = "connect"

type Connect struct {
	TimeLastConnect time.Time `json:"time" bson:"time"`
}

func NewConnect(timeLastConnect time.Time) *Connect {
	return &Connect{
		TimeLastConnect: timeLastConnect,
	}
}

func (c *Connect) GetTimeLastConnect() time.Time {
	return c.TimeLastConnect
}

func (c *Connect) SetTimeLastConnect(timeLastConnect time.Time) {
	c.TimeLastConnect = timeLastConnect
}
