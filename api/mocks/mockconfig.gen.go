// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/fabric-sdk-go/api (interfaces: Config)

package mock_api

import (
	x509 "crypto/x509"
	gomock "github.com/golang/mock/gomock"
	api "github.com/hyperledger/fabric-sdk-go/api"
	factory "github.com/hyperledger/fabric/bccsp/factory"
	viper "github.com/spf13/viper"
)

// MockConfig is a mock of Config interface
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return _m.recorder
}

// GetCAClientCertFile mocks base method
func (_m *MockConfig) GetCAClientCertFile(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetCAClientCertFile", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCAClientCertFile indicates an expected call of GetCAClientCertFile
func (_mr *MockConfigMockRecorder) GetCAClientCertFile(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCAClientCertFile", arg0)
}

// GetCAClientKeyFile mocks base method
func (_m *MockConfig) GetCAClientKeyFile(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetCAClientKeyFile", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCAClientKeyFile indicates an expected call of GetCAClientKeyFile
func (_mr *MockConfigMockRecorder) GetCAClientKeyFile(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCAClientKeyFile", arg0)
}

// GetCAConfig mocks base method
func (_m *MockConfig) GetCAConfig(_param0 string) (*api.CAConfig, error) {
	ret := _m.ctrl.Call(_m, "GetCAConfig", _param0)
	ret0, _ := ret[0].(*api.CAConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCAConfig indicates an expected call of GetCAConfig
func (_mr *MockConfigMockRecorder) GetCAConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCAConfig", arg0)
}

// GetCAKeyStorePath mocks base method
func (_m *MockConfig) GetCAKeyStorePath() string {
	ret := _m.ctrl.Call(_m, "GetCAKeyStorePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCAKeyStorePath indicates an expected call of GetCAKeyStorePath
func (_mr *MockConfigMockRecorder) GetCAKeyStorePath() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCAKeyStorePath")
}

// GetCAServerCertFiles mocks base method
func (_m *MockConfig) GetCAServerCertFiles(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "GetCAServerCertFiles", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCAServerCertFiles indicates an expected call of GetCAServerCertFiles
func (_mr *MockConfigMockRecorder) GetCAServerCertFiles(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCAServerCertFiles", arg0)
}

// GetCSPConfig mocks base method
func (_m *MockConfig) GetCSPConfig() *factory.FactoryOpts {
	ret := _m.ctrl.Call(_m, "GetCSPConfig")
	ret0, _ := ret[0].(*factory.FactoryOpts)
	return ret0
}

// GetCSPConfig indicates an expected call of GetCSPConfig
func (_mr *MockConfigMockRecorder) GetCSPConfig() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCSPConfig")
}

// GetCryptoConfigPath mocks base method
func (_m *MockConfig) GetCryptoConfigPath() string {
	ret := _m.ctrl.Call(_m, "GetCryptoConfigPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCryptoConfigPath indicates an expected call of GetCryptoConfigPath
func (_mr *MockConfigMockRecorder) GetCryptoConfigPath() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCryptoConfigPath")
}

// GetFabricClientViper mocks base method
func (_m *MockConfig) GetFabricClientViper() *viper.Viper {
	ret := _m.ctrl.Call(_m, "GetFabricClientViper")
	ret0, _ := ret[0].(*viper.Viper)
	return ret0
}

// GetFabricClientViper indicates an expected call of GetFabricClientViper
func (_mr *MockConfigMockRecorder) GetFabricClientViper() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFabricClientViper")
}

// GetKeyStorePath mocks base method
func (_m *MockConfig) GetKeyStorePath() string {
	ret := _m.ctrl.Call(_m, "GetKeyStorePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetKeyStorePath indicates an expected call of GetKeyStorePath
func (_mr *MockConfigMockRecorder) GetKeyStorePath() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetKeyStorePath")
}

// GetMspID mocks base method
func (_m *MockConfig) GetMspID(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetMspID", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMspID indicates an expected call of GetMspID
func (_mr *MockConfigMockRecorder) GetMspID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMspID", arg0)
}

// GetNetworkConfig mocks base method
func (_m *MockConfig) GetNetworkConfig() (*api.NetworkConfig, error) {
	ret := _m.ctrl.Call(_m, "GetNetworkConfig")
	ret0, _ := ret[0].(*api.NetworkConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkConfig indicates an expected call of GetNetworkConfig
func (_mr *MockConfigMockRecorder) GetNetworkConfig() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNetworkConfig")
}

// GetOrdererConfig mocks base method
func (_m *MockConfig) GetOrdererConfig(_param0 string) (*api.OrdererConfig, error) {
	ret := _m.ctrl.Call(_m, "GetOrdererConfig", _param0)
	ret0, _ := ret[0].(*api.OrdererConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdererConfig indicates an expected call of GetOrdererConfig
func (_mr *MockConfigMockRecorder) GetOrdererConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOrdererConfig", arg0)
}

// GetPeersConfig mocks base method
func (_m *MockConfig) GetPeersConfig(_param0 string) ([]api.PeerConfig, error) {
	ret := _m.ctrl.Call(_m, "GetPeersConfig", _param0)
	ret0, _ := ret[0].([]api.PeerConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeersConfig indicates an expected call of GetPeersConfig
func (_mr *MockConfigMockRecorder) GetPeersConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPeersConfig", arg0)
}

// GetRandomOrdererConfig mocks base method
func (_m *MockConfig) GetRandomOrdererConfig() (*api.OrdererConfig, error) {
	ret := _m.ctrl.Call(_m, "GetRandomOrdererConfig")
	ret0, _ := ret[0].(*api.OrdererConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandomOrdererConfig indicates an expected call of GetRandomOrdererConfig
func (_mr *MockConfigMockRecorder) GetRandomOrdererConfig() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRandomOrdererConfig")
}

// GetSecurityAlgorithm mocks base method
func (_m *MockConfig) GetSecurityAlgorithm() string {
	ret := _m.ctrl.Call(_m, "GetSecurityAlgorithm")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSecurityAlgorithm indicates an expected call of GetSecurityAlgorithm
func (_mr *MockConfigMockRecorder) GetSecurityAlgorithm() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSecurityAlgorithm")
}

// GetSecurityLevel mocks base method
func (_m *MockConfig) GetSecurityLevel() int {
	ret := _m.ctrl.Call(_m, "GetSecurityLevel")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetSecurityLevel indicates an expected call of GetSecurityLevel
func (_mr *MockConfigMockRecorder) GetSecurityLevel() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSecurityLevel")
}

// GetTLSCACertPool mocks base method
func (_m *MockConfig) GetTLSCACertPool(_param0 string) (*x509.CertPool, error) {
	ret := _m.ctrl.Call(_m, "GetTLSCACertPool", _param0)
	ret0, _ := ret[0].(*x509.CertPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTLSCACertPool indicates an expected call of GetTLSCACertPool
func (_mr *MockConfigMockRecorder) GetTLSCACertPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTLSCACertPool", arg0)
}

// GetTLSCACertPoolFromRoots mocks base method
func (_m *MockConfig) GetTLSCACertPoolFromRoots(_param0 [][]byte) (*x509.CertPool, error) {
	ret := _m.ctrl.Call(_m, "GetTLSCACertPoolFromRoots", _param0)
	ret0, _ := ret[0].(*x509.CertPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTLSCACertPoolFromRoots indicates an expected call of GetTLSCACertPoolFromRoots
func (_mr *MockConfigMockRecorder) GetTLSCACertPoolFromRoots(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTLSCACertPoolFromRoots", arg0)
}

// IsSecurityEnabled mocks base method
func (_m *MockConfig) IsSecurityEnabled() bool {
	ret := _m.ctrl.Call(_m, "IsSecurityEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSecurityEnabled indicates an expected call of IsSecurityEnabled
func (_mr *MockConfigMockRecorder) IsSecurityEnabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsSecurityEnabled")
}

// IsTLSEnabled mocks base method
func (_m *MockConfig) IsTLSEnabled() bool {
	ret := _m.ctrl.Call(_m, "IsTLSEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTLSEnabled indicates an expected call of IsTLSEnabled
func (_mr *MockConfigMockRecorder) IsTLSEnabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsTLSEnabled")
}

// TcertBatchSize mocks base method
func (_m *MockConfig) TcertBatchSize() int {
	ret := _m.ctrl.Call(_m, "TcertBatchSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// TcertBatchSize indicates an expected call of TcertBatchSize
func (_mr *MockConfigMockRecorder) TcertBatchSize() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TcertBatchSize")
}
