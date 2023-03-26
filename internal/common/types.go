package common

type ConfigResponse struct {
	Meta Meta   `json:"meta"`
	Data []Data `json:"data"`
}

type Message struct {
	Title       any    `json:"title"`
	Description any    `json:"description"`
	Show        bool   `json:"show"`
	Type        string `json:"type"`
	Page        string `json:"page"`
}

type Meta struct {
	Timestamp       int     `json:"timestamp"`
	Message         Message `json:"message"`
	GeoIPServiceURL string  `json:"geoIPServiceURL"`
}

type Specifics struct {
	CooperationRefusedORIURL string `json:"cooperationRefusedORIUrl"`
}

type Data struct {
	CountryCode       string    `json:"countryCode"`
	CountryName       string    `json:"countryName"`
	RegistryURL       string    `json:"registryUrl"`
	CustomRegistryURL any       `json:"customRegistryUrl"`
	Specifics         Specifics `json:"specifics"`
	ProxyURL          string    `json:"proxyUrl"`
	IgnoreURL         any       `json:"ignoreUrl"`
}

type Refused struct {
	CooperationRefused bool   `json:"cooperationRefused"`
	Url                string `json:"url"`
}
