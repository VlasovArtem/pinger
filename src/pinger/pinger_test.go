package pinger

//type PingerTestSuite struct {
//	suite.Suite
//}
//
//func TestPingerTestSuite(t *testing.T) {
//	tesingSuite := &PingerTestSuite{}
//
//	suite.Run(t, tesingSuite)
//}
//
//func (p *PingerTestSuite) TestNewPinger() {
//	pinger, processorMock, _ := createPinger()
//	defer pinger.Stop()
//
//	assert.Equal(p.T(), processorMock, pinger.processor)
//	assert.Nil(p.T(), pinger.lastPingResult)
//
//	processorMock.AssertCalled(p.T(), "GetDefaultConfig")
//}
//
//func (p *PingerTestSuite) TestPinger_AddIp_NotTrustedAndPingSuccessful() {
//	pinger, _, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	ip := "192.168.1.1"
//
//	pingProviderMock.On("Ping", mock.Anything).Return(nil)
//
//	err := pinger.AddIp(ip, false)
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), []string{ip}, pinger.currentConfig.Ips)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", ip)
//}
//
//func (p *PingerTestSuite) TestPinger_AddIp_NotTrustedAndPingFailed() {
//	pinger, _, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	ip := "192.168.1.1"
//
//	pingProviderMock.On("Ping", mock.Anything).Return(errors.New("error"))
//
//	err := pinger.AddIp(ip, false)
//
//	assert.ErrorContains(p.T(), err, "error")
//	assert.Equal(p.T(), []string{}, pinger.currentConfig.Ips)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", ip)
//}
//
//func (p *PingerTestSuite) TestPinger_AddIp_WithTrustedAndPingFailed() {
//	pinger, _, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	ip := "192.168.1.1"
//
//	err := pinger.AddIp(ip, true)
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), []string{ip}, pinger.currentConfig.Ips)
//
//	pingProviderMock.AssertNotCalled(p.T(), "Ping", ip)
//}
//
//func (p *PingerTestSuite) TestPinger_SetConsensus_WithValidConsensus() {
//	type consensus struct {
//		income   string
//		expected config.Quorum
//	}
//
//	tests := []consensus{
//		{
//			income:   "any",
//			expected: config.ANY,
//		},
//		{
//			income:   "ANY",
//			expected: config.ANY,
//		},
//		{
//			income:   "all",
//			expected: config.ALL,
//		},
//		{
//			income:   "ALL",
//			expected: config.ALL,
//		},
//	}
//
//	pinger, _, _ := createPinger()
//	defer pinger.Stop()
//
//	for _, test := range tests {
//		err := pinger.SetQuorum(test.income)
//
//		assert.Nil(p.T(), err)
//		assert.Equal(p.T(), test.expected, pinger.currentConfig.Quorum)
//	}
//}
//
//func (p *PingerTestSuite) TestPinger_SetConsensus_WithInvalidConsensus() {
//	pinger, _, _ := createPinger()
//	defer pinger.Stop()
//
//	err := pinger.SetQuorum("invalid")
//
//	assert.ErrorContains(p.T(), err, "consensus is not valid. Should be 'ALL' or ANY")
//}
//
//func (p *PingerTestSuite) TestPinger_Reset() {
//	pinger, processorMock, _ := createPinger()
//	defer pinger.Stop()
//
//	pinger.activationFlag.Store(true)
//	initConfig := *pinger.currentConfig
//
//	processorMock.On("GetDefaultConfig").Return(createTestConfig())
//
//	pinger.Reset()
//
//	assert.NotEqual(p.T(), initConfig, pinger.currentConfig)
//	assert.False(p.T(), pinger.activationFlag.Load())
//
//	processorMock.AssertCalled(p.T(), "GetDefaultConfig")
//}
//
//func (p *PingerTestSuite) TestPinger_SetTimeout_WithValidValues() {
//	type timeoutFields struct {
//		timeout     string
//		timeoutType string
//		expected    time.Duration
//	}
//
//	tests := []timeoutFields{
//		{
//			timeout:     "1",
//			timeoutType: "seconds",
//			expected:    time.Second * 1,
//		},
//		{
//			timeout:     "10",
//			timeoutType: "SECONDS",
//			expected:    time.Second * 10,
//		},
//		{
//			timeout:     "5",
//			timeoutType: "minutes",
//			expected:    time.Minute * 5,
//		},
//		{
//			timeout:     "3",
//			timeoutType: "MINUTES",
//			expected:    time.Minute * 3,
//		},
//	}
//
//	pinger, _, _ := createPinger()
//	defer pinger.Stop()
//
//	for _, test := range tests {
//		p.Run(fmt.Sprintf("Valid set timeout. Timeout: %s, Timeout Type: %s", test.timeout, test.timeoutType), func() {
//
//			err := pinger.SetTimeout(test.timeout, test.timeoutType)
//
//			assert.Nil(p.T(), err)
//			assert.Equal(p.T(), test.expected, pinger.currentConfig.Timeout)
//		})
//	}
//}
//
//func (p *PingerTestSuite) TestPinger_SetTimeout_WithInvalidValues() {
//	type timeoutFields struct {
//		timeout     string
//		timeoutType string
//		err         error
//		description string
//	}
//
//	tests := []timeoutFields{
//		{
//			timeout:     "1.00",
//			timeoutType: "seconds",
//			err:         errors.New("timeout is not integer"),
//			description: "Invalid timeout value",
//		},
//		{
//			timeout:     "0",
//			timeoutType: "seconds",
//			err:         errors.New("timeout should be greater then zero"),
//			description: "Zero timeout value",
//		},
//		{
//			timeout:     "-10",
//			timeoutType: "seconds",
//			err:         errors.New("timeout should be greater then zero"),
//			description: "Negative timeout value",
//		},
//		{
//			timeout:     "1",
//			timeoutType: "",
//			err:         errors.New("timeout type could not be empty"),
//			description: "empty timeoutType value",
//		},
//		{
//			timeout:     "1",
//			timeoutType: "invalid",
//			err:         errors.New(fmt.Sprintf("timeout type is not exists. Valid types: '%s', '%s'", config.MINUTES, config.SECONDS)),
//			description: "invalid timeoutType value",
//		},
//	}
//
//	pinger, _, _ := createPinger()
//	defer pinger.Stop()
//
//	for _, test := range tests {
//		p.Run(test.description, func() {
//			err := pinger.SetTimeout(test.timeout, test.timeoutType)
//
//			assert.EqualError(p.T(), err, test.err.Error())
//			assert.Equal(p.T(), time.Second, pinger.currentConfig.Timeout)
//		})
//	}
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithoutIps() {
//	pinger, _, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	_, err := pinger.Start()
//
//	assert.EqualError(p.T(), err, "ips are not exists")
//
//	pingProviderMock.AssertNotCalled(p.T(), "Ping", mock.Anything)
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAnyConsensusAndPingSuccessful() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", mock.Anything).Return(nil)
//	processorMock.On("OnSuccess").Return()
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		return checkProcessorMockCalled(processorMock, "OnSuccess")
//	}, time.Second*5, time.Second)
//
//	assert.True(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	pingProviderMock.AssertNumberOfCalls(p.T(), "Ping", 1)
//	processorMock.AssertCalled(p.T(), "OnSuccess")
//	processorMock.AssertNotCalled(p.T(), "OnError")
//	processorMock.AssertNotCalled(p.T(), "GetTrigger")
//
//	pinger.Stop()
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAnyConsensusAndPingUnsuccessful() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.currentConfig.SetQuorum(config.ANY)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", mock.Anything).Return(errors.New("error"))
//	processorMock.On("OnError").Return()
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		for _, call := range processorMock.Calls {
//			if call.Method == "OnError" {
//				return true
//			}
//		}
//		return false
//	}, time.Second*5, time.Second)
//
//	assert.False(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.2")
//	processorMock.AssertCalled(p.T(), "OnError")
//
//	processorMock.AssertNotCalled(p.T(), "OnSuccess")
//	processorMock.AssertNotCalled(p.T(), "GetTrigger")
//
//	pinger.Stop()
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAnyConsensusAndFirstPingUnsuccessful() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.currentConfig.SetQuorum(config.ANY)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", "192.168.1.1").Return(errors.New("first"))
//	pingProviderMock.On("Ping", "192.168.1.2").Return(nil)
//	processorMock.On("OnSuccess").Return()
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		if pinger.lastPingResult == nil {
//			return false
//		}
//		return true
//	}, time.Second*5, time.Second)
//
//	assert.True(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.2")
//	processorMock.AssertCalled(p.T(), "OnSuccess")
//
//	processorMock.AssertNotCalled(p.T(), "OnError")
//	processorMock.AssertNotCalled(p.T(), "GetTrigger")
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAnyConsensusAndLastPingResultNotChanges() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.lastPingResult = helper.Ptr(true)
//
//	pinger.currentConfig.SetQuorum(config.ANY)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", mock.Anything).Return(nil)
//	processorMock.On("GetTrigger").Return(ON_CHANGE)
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		for _, call := range processorMock.Calls {
//			if call.Method == "GetTrigger" {
//				return true
//			}
//		}
//		return false
//	}, time.Second*5, time.Second)
//
//	assert.True(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	processorMock.AssertCalled(p.T(), "GetTrigger")
//
//	processorMock.AssertNotCalled(p.T(), "OnSuccess")
//	processorMock.AssertNotCalled(p.T(), "OnError")
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAnyConsensusAndLastPingResultNotChangesAndPingError() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.lastPingResult = helper.Ptr(false)
//
//	pinger.currentConfig.SetQuorum(config.ANY)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", mock.Anything).Return(errors.New("error"))
//	processorMock.On("GetTrigger").Return(ON_CHANGE)
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		for _, call := range processorMock.Calls {
//			if call.Method == "GetTrigger" {
//				return true
//			}
//		}
//		return false
//	}, time.Second*5, time.Second)
//
//	assert.False(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	processorMock.AssertCalled(p.T(), "GetTrigger")
//
//	processorMock.AssertNotCalled(p.T(), "OnSuccess")
//	processorMock.AssertNotCalled(p.T(), "OnError")
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAnyConsensusAndLastPingResultNotChangesAndTriggerConstantly() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.lastPingResult = helper.Ptr(true)
//
//	pinger.currentConfig.SetQuorum(config.ANY)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", mock.Anything).Return(nil)
//	processorMock.On("OnSuccess").Return()
//	processorMock.On("GetTrigger").Return(CONSTANTLY)
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		for _, call := range processorMock.Calls {
//			if call.Method == "GetTrigger" {
//				return true
//			}
//		}
//		return false
//	}, time.Second*5, time.Second)
//
//	assert.True(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	processorMock.AssertCalled(p.T(), "GetTrigger")
//	processorMock.AssertCalled(p.T(), "OnSuccess")
//
//	processorMock.AssertNotCalled(p.T(), "OnError")
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAnyConsensusAndLastPingResultNotChangesAndTriggerConstantlAndPingError() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.lastPingResult = helper.Ptr(true)
//
//	pinger.currentConfig.SetQuorum(config.ANY)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", mock.Anything).Return(errors.New("error"))
//	processorMock.On("OnError").Return()
//	processorMock.On("GetTrigger").Return(CONSTANTLY)
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		for _, call := range processorMock.Calls {
//			if call.Method == "GetTrigger" {
//				return true
//			}
//		}
//		return false
//	}, time.Second*5, time.Second)
//
//	assert.False(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.2")
//	processorMock.AssertCalled(p.T(), "GetTrigger")
//	processorMock.AssertCalled(p.T(), "OnError")
//
//	processorMock.AssertNotCalled(p.T(), "OnSuccess")
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAllConsensusAndPingSuccessful() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.currentConfig.SetQuorum(config.ALL)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", mock.Anything).Return(nil)
//	processorMock.On("OnSuccess").Return()
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		return checkProcessorMockCalled(processorMock, "OnSuccess")
//	}, time.Second*5, time.Second)
//
//	assert.True(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.2")
//	pingProviderMock.AssertNumberOfCalls(p.T(), "Ping", 2)
//	processorMock.AssertCalled(p.T(), "OnSuccess")
//
//	processorMock.AssertNotCalled(p.T(), "OnError")
//	processorMock.AssertNotCalled(p.T(), "GetTrigger")
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAllConsensusAndPingUnsuccessful() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.currentConfig.SetQuorum(config.ALL)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", mock.Anything).Return(errors.New("error"))
//	processorMock.On("OnError").Return()
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		return checkProcessorMockCalled(processorMock, "OnError")
//	}, time.Second*5, time.Second)
//
//	assert.False(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.2")
//	processorMock.AssertCalled(p.T(), "OnError")
//
//	processorMock.AssertNotCalled(p.T(), "OnSuccess")
//	processorMock.AssertNotCalled(p.T(), "GetTrigger")
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAllConsensusAndFirstPingUnsuccessful() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.currentConfig.SetQuorum(config.ALL)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", "192.168.1.1").Return(errors.New("first"))
//	processorMock.On("OnError").Return()
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		if pinger.lastPingResult == nil {
//			return false
//		}
//		return true
//	}, time.Second*5, time.Second)
//
//	assert.False(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	processorMock.AssertCalled(p.T(), "OnError")
//
//	processorMock.AssertNotCalled(p.T(), "OnSuccess")
//	processorMock.AssertNotCalled(p.T(), "GetTrigger")
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAllConsensusAndLastPingResultNotChanges() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.lastPingResult = helper.Ptr(true)
//
//	pinger.currentConfig.SetQuorum(config.ALL)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", mock.Anything).Return(nil)
//	processorMock.On("GetTrigger").Return(ON_CHANGE)
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		for _, call := range processorMock.Calls {
//			if call.Method == "GetTrigger" {
//				return true
//			}
//		}
//		return false
//	}, time.Second*5, time.Second)
//
//	assert.True(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.2")
//	processorMock.AssertCalled(p.T(), "GetTrigger")
//
//	processorMock.AssertNotCalled(p.T(), "OnSuccess")
//	processorMock.AssertNotCalled(p.T(), "OnError")
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAllConsensusAndLastPingResultNotChangesAndPingError() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.lastPingResult = helper.Ptr(false)
//
//	pinger.currentConfig.SetQuorum(config.ALL)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", mock.Anything).Return(errors.New("error"))
//	processorMock.On("GetTrigger").Return(ON_CHANGE)
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		for _, call := range processorMock.Calls {
//			if call.Method == "GetTrigger" {
//				return true
//			}
//		}
//		return false
//	}, time.Second*5, time.Second)
//
//	assert.False(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.2")
//	processorMock.AssertCalled(p.T(), "GetTrigger")
//
//	processorMock.AssertNotCalled(p.T(), "OnSuccess")
//	processorMock.AssertNotCalled(p.T(), "OnError")
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAllConsensusAndLastPingResultNotChangesAndTriggerConstantly() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.lastPingResult = helper.Ptr(true)
//
//	pinger.currentConfig.SetQuorum(config.ALL)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", mock.Anything).Return(nil)
//	processorMock.On("OnSuccess").Return()
//	processorMock.On("GetTrigger").Return(CONSTANTLY)
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		for _, call := range processorMock.Calls {
//			if call.Method == "GetTrigger" {
//				return true
//			}
//		}
//		return false
//	}, time.Second*5, time.Second)
//
//	assert.True(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.2")
//	processorMock.AssertCalled(p.T(), "GetTrigger")
//	processorMock.AssertCalled(p.T(), "OnSuccess")
//
//	processorMock.AssertNotCalled(p.T(), "OnError")
//}
//
//func (p *PingerTestSuite) TestPinger_Start_WithAllConsensusAndLastPingResultNotChangesAndTriggerConstantlyAndPingError() {
//	pinger, processorMock, pingProviderMock := createPinger()
//	defer pinger.Stop()
//
//	pinger.lastPingResult = helper.Ptr(false)
//
//	pinger.currentConfig.SetQuorum(config.ALL)
//	pinger.currentConfig.AddIp("192.168.1.1")
//	pinger.currentConfig.AddIp("192.168.1.2")
//
//	pingProviderMock.On("Ping", mock.Anything).Return(errors.New("error"))
//	processorMock.On("OnError").Return()
//	processorMock.On("GetTrigger").Return(CONSTANTLY)
//
//	response, err := pinger.Start()
//
//	assert.Nil(p.T(), err)
//	assert.Equal(p.T(), struct {
//		message       string
//		currentConfig config.Config
//	}{
//		message:       "Pinger started",
//		currentConfig: *pinger.currentConfig,
//	}, response)
//	assert.Eventually(p.T(), func() bool {
//		for _, call := range processorMock.Calls {
//			if call.Method == "GetTrigger" {
//				return true
//			}
//		}
//		return false
//	}, time.Second*5, time.Second)
//
//	assert.False(p.T(), *pinger.lastPingResult)
//
//	pingProviderMock.AssertCalled(p.T(), "Ping", "192.168.1.1")
//	processorMock.AssertCalled(p.T(), "GetTrigger")
//	processorMock.AssertCalled(p.T(), "OnError")
//
//	processorMock.AssertNotCalled(p.T(), "OnSuccess")
//}
//
//func createTestConfig() *config.Config {
//	return config.NewConfig(config.ANY, time.Second)
//}
//
//func createPinger() (*Pinger, *mocks.Processor, *mocks.PingProvider) {
//	processorMock := new(mocks.Processor)
//	pingProviderMock := new(mocks.PingProvider)
//
//	processorMock.On("GetDefaultConfig").Return(createTestConfig())
//
//	return NewPinger(processorMock, pingProviderMock), processorMock, pingProviderMock
//}
//
//func checkProcessorMockCalled(processorMock *mocks.Processor, methodName string) bool {
//	for _, call := range processorMock.Calls {
//		if call.Method == methodName {
//			return true
//		}
//	}
//	return false
//}
