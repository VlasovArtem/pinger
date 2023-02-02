package pinger

type TestPingProvider struct {
	LastCalledElement string
	Calls             int
	ReturnValue       func() error
}

func NewTestPingProvider(returnValue func() error) PingProvider {
	return &TestPingProvider{
		ReturnValue: returnValue,
	}
}

func (t *TestPingProvider) Ping(element string) error {
	t.Calls++
	t.LastCalledElement = element
	return t.ReturnValue()
}
