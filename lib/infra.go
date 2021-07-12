package lib

type simpleEngineColumn struct {
	Title   string `json:"title"`
	Index   int    `json:"index"`
	Content string `json:"content"`
	URL     string `json:"url"`
}

type EngineResultInfra struct {
	Engine []struct {
		Google []simpleEngineColumn `json:"google"`
		Sogou  []simpleEngineColumn `json:"sogou"`
		Baidu  []simpleEngineColumn `json:"baidu"`
		Yandex []simpleEngineColumn `json:"yandex"`
	} `json:"engine"`
}
