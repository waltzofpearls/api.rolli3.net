package main

type version2 struct {
	*version1
}

func newVersion2(a *api) *version2 {
	return &version2{
		newVersion1(a),
	}
}

func (api *version2) attach() {
	v2 := api.engine.Group("/v2")
	{
		v2.GET("/status", api.version1.getStatus)
	}
}
