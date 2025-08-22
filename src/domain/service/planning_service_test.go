package service

import (
	"testing"

	"assessment/domain/entity"
	"assessment/dto"
	"assessment/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// PlanningServiceTestSuite holds mocks and service under test.
type PlanningServiceTestSuite struct {
	suite.Suite

	eqRepo *mocks.MockEquipmentRepo
	plRepo *mocks.MockPlanningRepo

	svc *PlanningService
}

func (s *PlanningServiceTestSuite) SetupTest() {
	// create mocks once per test (clean state)
	s.eqRepo = new(mocks.MockEquipmentRepo)
	s.plRepo = new(mocks.MockPlanningRepo)

	// create service using those mocks
	s.svc = NewPlanningService(s.eqRepo, s.plRepo)
}

func (s *PlanningServiceTestSuite) TearDownTest() {
	// Assert that all expectations were met for each test
	s.eqRepo.AssertExpectations(s.T())
	s.plRepo.AssertExpectations(s.T())
}

// --- Tests ---

func (s *PlanningServiceTestSuite) TestIsAvailable_Available() {
	// equipment with stock 9
	s.eqRepo.On("GetEquipment", mock.Anything, int64(1)).
		Return(&entity.Equipment{Id: 1, Stock: 9}, nil)

	// plannings overlapping produce maxConcurrent = 5 (2 + 3)
	plannings := []entity.Planning{
		{Id: 1, Equipment: 1, Quantity: 2, StartAt: mustParseTime(s.T(), "2019-05-30 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-01 00:00:00")},
		{Id: 2, Equipment: 1, Quantity: 3, StartAt: mustParseTime(s.T(), "2019-05-31 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-02 00:00:00")},
	}
	plReq := dto.GetPlanningsRequest{Equipment: 1, StartAt: mustParseTime(s.T(), "2019-05-30 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-03 00:00:00")}
	s.plRepo.On("GetPlannings", mock.Anything, plReq).Return(plannings, nil)

	ok, err := s.svc.IsAvailable(1, 2, plReq.StartAt, plReq.EndAt)
	s.Require().NoError(err)
	s.Require().True(ok)
}

func (s *PlanningServiceTestSuite) TestIsAvailable_NotAvailable() {
	s.eqRepo.On("GetEquipment", mock.Anything, int64(1)).
		Return(&entity.Equipment{Id: 1, Stock: 9}, nil)

	plannings := []entity.Planning{
		{Id: 1, Equipment: 1, Quantity: 4, StartAt: mustParseTime(s.T(), "2019-05-29 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-02 00:00:00")},
		{Id: 2, Equipment: 1, Quantity: 7, StartAt: mustParseTime(s.T(), "2019-05-31 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-04 00:00:00")},
	}
	plReq := dto.GetPlanningsRequest{Equipment: 1, StartAt: mustParseTime(s.T(), "2019-05-30 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-03 00:00:00")}
	s.plRepo.On("GetPlannings", mock.Anything, plReq).Return(plannings, nil)

	ok, err := s.svc.IsAvailable(1, 0, plReq.StartAt, plReq.EndAt) // request 0 more -> still over capacity
	s.Require().NoError(err)
	s.Require().False(ok)
}

func (s *PlanningServiceTestSuite) TestIsAvailable_EquipmentNotFound() {
	// return nil, nil to simulate not found (service expects nil => error)
	s.eqRepo.On("GetEquipment", mock.Anything, int64(99)).Return(nil, nil)

	_, err := s.svc.IsAvailable(99, 1, mustParseTime(s.T(), "2019-05-30 00:00:00"), mustParseTime(s.T(), "2019-06-03 00:00:00"))
	s.Require().Error(err)
}

func (s *PlanningServiceTestSuite) TestGetShortages_NoShortages() {
	s.eqRepo.On("GetEquipments", mock.Anything).Return([]entity.Equipment{
		{Id: 1, Stock: 10},
		{Id: 2, Stock: 5},
	}, nil)

	allReq := dto.GetPlanningsBetweenRequest{StartAt: mustParseTime(s.T(), "2019-05-29 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-05 00:00:00")}
	s.plRepo.On("GetPlanningsBetween", mock.Anything, allReq).Return([]entity.Planning{
		{Id: 1, Equipment: 1, Quantity: 2, StartAt: mustParseTime(s.T(), "2019-05-30 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-01 00:00:00")},
		{Id: 2, Equipment: 2, Quantity: 1, StartAt: mustParseTime(s.T(), "2019-05-30 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-02 00:00:00")},
	}, nil)

	res, err := s.svc.GetShortages(allReq.StartAt, allReq.EndAt)
	s.Require().NoError(err)
	s.Require().Empty(res)
}

func (s *PlanningServiceTestSuite) TestGetShortages_WithShortages() {
	s.eqRepo.On("GetEquipments", mock.Anything).Return([]entity.Equipment{
		{Id: 100, Stock: 9},
	}, nil)

	allReq := dto.GetPlanningsBetweenRequest{StartAt: mustParseTime(s.T(), "2019-05-29 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-05 00:00:00")}
	s.plRepo.On("GetPlanningsBetween", mock.Anything, allReq).Return([]entity.Planning{
		{Id: 1, Equipment: 100, Quantity: 4, StartAt: mustParseTime(s.T(), "2019-05-29 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-02 00:00:00")},
		{Id: 2, Equipment: 100, Quantity: 5, StartAt: mustParseTime(s.T(), "2019-05-31 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-04 00:00:00")},
		{Id: 3, Equipment: 100, Quantity: 2, StartAt: mustParseTime(s.T(), "2019-06-01 00:00:00"), EndAt: mustParseTime(s.T(), "2019-06-03 00:00:00")},
	}, nil)

	res, err := s.svc.GetShortages(allReq.StartAt, allReq.EndAt)
	s.Require().NoError(err)
	s.Require().Len(res, 1)
	s.Require().Equal(int64(-2), res[100])
}

// entry point for the suite
func TestPlanningServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PlanningServiceTestSuite))
}
