package internal

type Update struct {
	InstalledVersion int    `json:"installed_version"`
	LatestVersion    int    `json:"latest_version"`
	State            string `json:"state"`
}

type LixeeState struct {
	MOTDETAT                    interface{} `json:"MOTDETAT"`
	ActiveRegisterTierDelivered string      `json:"active_register_tier_delivered"`
	ApparentPower               int         `json:"apparent_power"`
	AvailablePower              int         `json:"available_power"`
	CurrentSummDelivered        int         `json:"current_summ_delivered"`
	CurrentTarif                string      `json:"current_tarif"`
	Linkquality                 int         `json:"linkquality"`
	MeterSerialNumber           string      `json:"meter_serial_number"`
	MotDEtat                    string      `json:"mot_d_etat"`
	RmsCurrent                  int         `json:"rms_current"`
	RmsCurrentMax               int         `json:"rms_current_max"`
	Update                      Update      `json:"update"`
	UpdateAvailable             interface{} `json:"update_available"`
	WarnDPS                     int         `json:"warn_d_p_s"`
}
