// Code generated by MockGen. DO NOT EDIT.
// Source: streams.go

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/joschahenningsen/TUM-Live/model"
)

// MockStreamsDao is a mock of StreamsDao interface.
type MockStreamsDao struct {
	ctrl     *gomock.Controller
	recorder *MockStreamsDaoMockRecorder
}

// MockStreamsDaoMockRecorder is the mock recorder for MockStreamsDao.
type MockStreamsDaoMockRecorder struct {
	mock *MockStreamsDao
}

// NewMockStreamsDao creates a new mock instance.
func NewMockStreamsDao(ctrl *gomock.Controller) *MockStreamsDao {
	mock := &MockStreamsDao{ctrl: ctrl}
	mock.recorder = &MockStreamsDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamsDao) EXPECT() *MockStreamsDaoMockRecorder {
	return m.recorder
}

// AddVodView mocks base method.
func (m *MockStreamsDao) AddVodView(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVodView", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVodView indicates an expected call of AddVodView.
func (mr *MockStreamsDaoMockRecorder) AddVodView(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVodView", reflect.TypeOf((*MockStreamsDao)(nil).AddVodView), id)
}

// ClearWorkersForStream mocks base method.
func (m *MockStreamsDao) ClearWorkersForStream(stream model.Stream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearWorkersForStream", stream)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearWorkersForStream indicates an expected call of ClearWorkersForStream.
func (mr *MockStreamsDaoMockRecorder) ClearWorkersForStream(stream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearWorkersForStream", reflect.TypeOf((*MockStreamsDao)(nil).ClearWorkersForStream), stream)
}

// CreateStream mocks base method.
func (m *MockStreamsDao) CreateStream(stream *model.Stream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStream", stream)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStream indicates an expected call of CreateStream.
func (mr *MockStreamsDaoMockRecorder) CreateStream(stream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStream", reflect.TypeOf((*MockStreamsDao)(nil).CreateStream), stream)
}

// DeleteSilences mocks base method.
func (m *MockStreamsDao) DeleteSilences(streamID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSilences", streamID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSilences indicates an expected call of DeleteSilences.
func (mr *MockStreamsDaoMockRecorder) DeleteSilences(streamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSilences", reflect.TypeOf((*MockStreamsDao)(nil).DeleteSilences), streamID)
}

// DeleteStream mocks base method.
func (m *MockStreamsDao) DeleteStream(streamID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteStream", streamID)
}

// DeleteStream indicates an expected call of DeleteStream.
func (mr *MockStreamsDaoMockRecorder) DeleteStream(streamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStream", reflect.TypeOf((*MockStreamsDao)(nil).DeleteStream), streamID)
}

// DeleteStreamsWithTumID mocks base method.
func (m *MockStreamsDao) DeleteStreamsWithTumID(ids []uint) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteStreamsWithTumID", ids)
}

// DeleteStreamsWithTumID indicates an expected call of DeleteStreamsWithTumID.
func (mr *MockStreamsDaoMockRecorder) DeleteStreamsWithTumID(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStreamsWithTumID", reflect.TypeOf((*MockStreamsDao)(nil).DeleteStreamsWithTumID), ids)
}

// DeleteUnit mocks base method.
func (m *MockStreamsDao) DeleteUnit(id uint) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUnit", id)
}

// DeleteUnit indicates an expected call of DeleteUnit.
func (mr *MockStreamsDaoMockRecorder) DeleteUnit(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUnit", reflect.TypeOf((*MockStreamsDao)(nil).DeleteUnit), id)
}

// GetAllStreams mocks base method.
func (m *MockStreamsDao) GetAllStreams() ([]model.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllStreams")
	ret0, _ := ret[0].([]model.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllStreams indicates an expected call of GetAllStreams.
func (mr *MockStreamsDaoMockRecorder) GetAllStreams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllStreams", reflect.TypeOf((*MockStreamsDao)(nil).GetAllStreams))
}

// GetCurrentLive mocks base method.
func (m *MockStreamsDao) GetCurrentLive(ctx context.Context) ([]model.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentLive", ctx)
	ret0, _ := ret[0].([]model.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentLive indicates an expected call of GetCurrentLive.
func (mr *MockStreamsDaoMockRecorder) GetCurrentLive(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentLive", reflect.TypeOf((*MockStreamsDao)(nil).GetCurrentLive), ctx)
}

// GetCurrentLiveNonHidden mocks base method.
func (m *MockStreamsDao) GetCurrentLiveNonHidden(ctx context.Context) ([]model.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentLiveNonHidden", ctx)
	ret0, _ := ret[0].([]model.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentLiveNonHidden indicates an expected call of GetCurrentLiveNonHidden.
func (mr *MockStreamsDaoMockRecorder) GetCurrentLiveNonHidden(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentLiveNonHidden", reflect.TypeOf((*MockStreamsDao)(nil).GetCurrentLiveNonHidden), ctx)
}

// GetDuePremieresForWorkers mocks base method.
func (m *MockStreamsDao) GetDuePremieresForWorkers() []model.Stream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDuePremieresForWorkers")
	ret0, _ := ret[0].([]model.Stream)
	return ret0
}

// GetDuePremieresForWorkers indicates an expected call of GetDuePremieresForWorkers.
func (mr *MockStreamsDaoMockRecorder) GetDuePremieresForWorkers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDuePremieresForWorkers", reflect.TypeOf((*MockStreamsDao)(nil).GetDuePremieresForWorkers))
}

// GetDueStreamsForWorkers mocks base method.
func (m *MockStreamsDao) GetDueStreamsForWorkers() []model.Stream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDueStreamsForWorkers")
	ret0, _ := ret[0].([]model.Stream)
	return ret0
}

// GetDueStreamsForWorkers indicates an expected call of GetDueStreamsForWorkers.
func (mr *MockStreamsDaoMockRecorder) GetDueStreamsForWorkers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDueStreamsForWorkers", reflect.TypeOf((*MockStreamsDao)(nil).GetDueStreamsForWorkers))
}

// GetLiveStreamsInLectureHall mocks base method.
func (m *MockStreamsDao) GetLiveStreamsInLectureHall(lectureHallId uint) ([]model.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLiveStreamsInLectureHall", lectureHallId)
	ret0, _ := ret[0].([]model.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLiveStreamsInLectureHall indicates an expected call of GetLiveStreamsInLectureHall.
func (mr *MockStreamsDaoMockRecorder) GetLiveStreamsInLectureHall(lectureHallId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLiveStreamsInLectureHall", reflect.TypeOf((*MockStreamsDao)(nil).GetLiveStreamsInLectureHall), lectureHallId)
}

// GetStreamByID mocks base method.
func (m *MockStreamsDao) GetStreamByID(ctx context.Context, id string) (model.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStreamByID", ctx, id)
	ret0, _ := ret[0].(model.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStreamByID indicates an expected call of GetStreamByID.
func (mr *MockStreamsDaoMockRecorder) GetStreamByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStreamByID", reflect.TypeOf((*MockStreamsDao)(nil).GetStreamByID), ctx, id)
}

// GetStreamByKey mocks base method.
func (m *MockStreamsDao) GetStreamByKey(ctx context.Context, key string) (model.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStreamByKey", ctx, key)
	ret0, _ := ret[0].(model.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStreamByKey indicates an expected call of GetStreamByKey.
func (mr *MockStreamsDaoMockRecorder) GetStreamByKey(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStreamByKey", reflect.TypeOf((*MockStreamsDao)(nil).GetStreamByKey), ctx, key)
}

// GetStreamByTumOnlineID mocks base method.
func (m *MockStreamsDao) GetStreamByTumOnlineID(ctx context.Context, id uint) (model.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStreamByTumOnlineID", ctx, id)
	ret0, _ := ret[0].(model.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStreamByTumOnlineID indicates an expected call of GetStreamByTumOnlineID.
func (mr *MockStreamsDaoMockRecorder) GetStreamByTumOnlineID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStreamByTumOnlineID", reflect.TypeOf((*MockStreamsDao)(nil).GetStreamByTumOnlineID), ctx, id)
}

// GetStreamsByIds mocks base method.
func (m *MockStreamsDao) GetStreamsByIds(ids []uint) ([]model.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStreamsByIds", ids)
	ret0, _ := ret[0].([]model.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStreamsByIds indicates an expected call of GetStreamsByIds.
func (mr *MockStreamsDaoMockRecorder) GetStreamsByIds(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStreamsByIds", reflect.TypeOf((*MockStreamsDao)(nil).GetStreamsByIds), ids)
}

// GetStreamsWithWatchState mocks base method.
func (m *MockStreamsDao) GetStreamsWithWatchState(courseID, userID uint) ([]model.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStreamsWithWatchState", courseID, userID)
	ret0, _ := ret[0].([]model.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStreamsWithWatchState indicates an expected call of GetStreamsWithWatchState.
func (mr *MockStreamsDaoMockRecorder) GetStreamsWithWatchState(courseID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStreamsWithWatchState", reflect.TypeOf((*MockStreamsDao)(nil).GetStreamsWithWatchState), courseID, userID)
}

// GetUnitByID mocks base method.
func (m *MockStreamsDao) GetUnitByID(id string) (model.StreamUnit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitByID", id)
	ret0, _ := ret[0].(model.StreamUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitByID indicates an expected call of GetUnitByID.
func (mr *MockStreamsDaoMockRecorder) GetUnitByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitByID", reflect.TypeOf((*MockStreamsDao)(nil).GetUnitByID), id)
}

// GetWorkersForStream mocks base method.
func (m *MockStreamsDao) GetWorkersForStream(stream model.Stream) ([]model.Worker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkersForStream", stream)
	ret0, _ := ret[0].([]model.Worker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkersForStream indicates an expected call of GetWorkersForStream.
func (mr *MockStreamsDaoMockRecorder) GetWorkersForStream(stream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkersForStream", reflect.TypeOf((*MockStreamsDao)(nil).GetWorkersForStream), stream)
}

// SaveCAMURL mocks base method.
func (m *MockStreamsDao) SaveCAMURL(stream *model.Stream, url string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SaveCAMURL", stream, url)
}

// SaveCAMURL indicates an expected call of SaveCAMURL.
func (mr *MockStreamsDaoMockRecorder) SaveCAMURL(stream, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCAMURL", reflect.TypeOf((*MockStreamsDao)(nil).SaveCAMURL), stream, url)
}

// SaveCOMBURL mocks base method.
func (m *MockStreamsDao) SaveCOMBURL(stream *model.Stream, url string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SaveCOMBURL", stream, url)
}

// SaveCOMBURL indicates an expected call of SaveCOMBURL.
func (mr *MockStreamsDaoMockRecorder) SaveCOMBURL(stream, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCOMBURL", reflect.TypeOf((*MockStreamsDao)(nil).SaveCOMBURL), stream, url)
}

// SaveEndedState mocks base method.
func (m *MockStreamsDao) SaveEndedState(streamID uint, hasEnded bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveEndedState", streamID, hasEnded)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveEndedState indicates an expected call of SaveEndedState.
func (mr *MockStreamsDaoMockRecorder) SaveEndedState(streamID, hasEnded interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveEndedState", reflect.TypeOf((*MockStreamsDao)(nil).SaveEndedState), streamID, hasEnded)
}

// SavePRESURL mocks base method.
func (m *MockStreamsDao) SavePRESURL(stream *model.Stream, url string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SavePRESURL", stream, url)
}

// SavePRESURL indicates an expected call of SavePRESURL.
func (mr *MockStreamsDaoMockRecorder) SavePRESURL(stream, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePRESURL", reflect.TypeOf((*MockStreamsDao)(nil).SavePRESURL), stream, url)
}

// SavePauseState mocks base method.
func (m *MockStreamsDao) SavePauseState(streamID uint, paused bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePauseState", streamID, paused)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePauseState indicates an expected call of SavePauseState.
func (mr *MockStreamsDaoMockRecorder) SavePauseState(streamID, paused interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePauseState", reflect.TypeOf((*MockStreamsDao)(nil).SavePauseState), streamID, paused)
}

// SaveStream mocks base method.
func (m *MockStreamsDao) SaveStream(vod *model.Stream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveStream", vod)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveStream indicates an expected call of SaveStream.
func (mr *MockStreamsDaoMockRecorder) SaveStream(vod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveStream", reflect.TypeOf((*MockStreamsDao)(nil).SaveStream), vod)
}

// SaveWorkerForStream mocks base method.
func (m *MockStreamsDao) SaveWorkerForStream(stream model.Stream, worker model.Worker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveWorkerForStream", stream, worker)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveWorkerForStream indicates an expected call of SaveWorkerForStream.
func (mr *MockStreamsDaoMockRecorder) SaveWorkerForStream(stream, worker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveWorkerForStream", reflect.TypeOf((*MockStreamsDao)(nil).SaveWorkerForStream), stream, worker)
}

// SetLectureHall mocks base method.
func (m *MockStreamsDao) SetLectureHall(streamIDs []uint, lectureHallID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLectureHall", streamIDs, lectureHallID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLectureHall indicates an expected call of SetLectureHall.
func (mr *MockStreamsDaoMockRecorder) SetLectureHall(streamIDs, lectureHallID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLectureHall", reflect.TypeOf((*MockStreamsDao)(nil).SetLectureHall), streamIDs, lectureHallID)
}

// SetStreamNotLiveById mocks base method.
func (m *MockStreamsDao) SetStreamNotLiveById(streamID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStreamNotLiveById", streamID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStreamNotLiveById indicates an expected call of SetStreamNotLiveById.
func (mr *MockStreamsDaoMockRecorder) SetStreamNotLiveById(streamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStreamNotLiveById", reflect.TypeOf((*MockStreamsDao)(nil).SetStreamNotLiveById), streamID)
}

// UnsetLectureHall mocks base method.
func (m *MockStreamsDao) UnsetLectureHall(streamIDs []uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsetLectureHall", streamIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsetLectureHall indicates an expected call of UnsetLectureHall.
func (mr *MockStreamsDaoMockRecorder) UnsetLectureHall(streamIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsetLectureHall", reflect.TypeOf((*MockStreamsDao)(nil).UnsetLectureHall), streamIDs)
}

// UpdateSilences mocks base method.
func (m *MockStreamsDao) UpdateSilences(silences []model.Silence, streamID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSilences", silences, streamID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSilences indicates an expected call of UpdateSilences.
func (mr *MockStreamsDaoMockRecorder) UpdateSilences(silences, streamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSilences", reflect.TypeOf((*MockStreamsDao)(nil).UpdateSilences), silences, streamID)
}

// UpdateStream mocks base method.
func (m *MockStreamsDao) UpdateStream(stream model.Stream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStream", stream)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStream indicates an expected call of UpdateStream.
func (mr *MockStreamsDaoMockRecorder) UpdateStream(stream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStream", reflect.TypeOf((*MockStreamsDao)(nil).UpdateStream), stream)
}

// UpdateStreamFullAssoc mocks base method.
func (m *MockStreamsDao) UpdateStreamFullAssoc(vod *model.Stream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStreamFullAssoc", vod)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStreamFullAssoc indicates an expected call of UpdateStreamFullAssoc.
func (mr *MockStreamsDaoMockRecorder) UpdateStreamFullAssoc(vod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStreamFullAssoc", reflect.TypeOf((*MockStreamsDao)(nil).UpdateStreamFullAssoc), vod)
}
