package service

import (
	"errors"
	"math"
	"time"

	"go_be_enrollment/internal/modules/examroomassignment/dto"
	"go_be_enrollment/internal/modules/examroomassignment/entity"
	"go_be_enrollment/internal/modules/examroomassignment/repository"
	app_repo "go_be_enrollment/internal/modules/application/repository"
	room_repo "go_be_enrollment/internal/modules/examroom/repository"
)

type ExamRoomAssignmentService interface {
	GetList(filter *dto.ExamRoomAssignmentFilter) (*dto.PaginatedExamRoomAssignmentRes, error)
	GetDetail(id uint) (*dto.ExamRoomAssignmentRes, error)
	Create(req *dto.ExamRoomAssignmentCreateReq) (*dto.ExamRoomAssignmentRes, error)
	Update(id uint, req *dto.ExamRoomAssignmentUpdateReq) (*dto.ExamRoomAssignmentRes, error)
	Delete(id uint) error
	GetListByRoomID(roomID uint) ([]dto.ExamRoomAssignmentRes, error)
}

type examRoomAssignmentService struct {
	repo     repository.ExamRoomAssignmentRepository
	appRepo  app_repo.ApplicationRepository
	roomRepo room_repo.ExamRoomRepository
}

func NewExamRoomAssignmentService(repo repository.ExamRoomAssignmentRepository, appRepo app_repo.ApplicationRepository, roomRepo room_repo.ExamRoomRepository) ExamRoomAssignmentService {
	return &examRoomAssignmentService{repo: repo, appRepo: appRepo, roomRepo: roomRepo}
}

func (s *examRoomAssignmentService) GetList(filter *dto.ExamRoomAssignmentFilter) (*dto.PaginatedExamRoomAssignmentRes, error) {
	assignments, total, err := s.repo.GetList(filter)
	if err != nil {
		return nil, err
	}

	limit := filter.Limit
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}

	page := filter.Page
	if page <= 0 {
		page = 1
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	var resData []dto.ExamRoomAssignmentRes
	for _, a := range assignments {
		cName := ""
		nID := ""
		rName := ""
		if a.Application != nil {
			cName = a.Application.CandidateFullName
			nID = a.Application.NationalID
		}
		if a.ExamRoom != nil {
			rName = a.ExamRoom.RoomName
		}
		resData = append(resData, dto.ExamRoomAssignmentRes{
			ID:                a.ID,
			ApplicationID:     a.ApplicationID,
			CandidateFullName: cName,
			NationalID:        nID,
			ExamRoomID:        a.ExamRoomID,
			RoomName:          rName,
			SeatNumber:        a.SeatNumber,
			CreatedAt:         a.CreatedAt.Format(time.RFC3339),
			UpdatedAt:         a.UpdatedAt.Format(time.RFC3339),
		})
	}
	
	if resData == nil {
		resData = []dto.ExamRoomAssignmentRes{}
	}

	return &dto.PaginatedExamRoomAssignmentRes{
		Data:       resData,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

func (s *examRoomAssignmentService) GetDetail(id uint) (*dto.ExamRoomAssignmentRes, error) {
	a, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("dữ liệu phân phòng không tồn tại")
	}

	cName := ""
	nID := ""
	rName := ""
	if a.Application != nil {
		cName = a.Application.CandidateFullName
		nID = a.Application.NationalID
	}
	if a.ExamRoom != nil {
		rName = a.ExamRoom.RoomName
	}

	return &dto.ExamRoomAssignmentRes{
		ID:                a.ID,
		ApplicationID:     a.ApplicationID,
		CandidateFullName: cName,
		NationalID:        nID,
		ExamRoomID:        a.ExamRoomID,
		RoomName:          rName,
		SeatNumber:        a.SeatNumber,
		CreatedAt:         a.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         a.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *examRoomAssignmentService) Create(req *dto.ExamRoomAssignmentCreateReq) (*dto.ExamRoomAssignmentRes, error) {
	app, err := s.appRepo.GetAdminDetail(req.ApplicationID)
	if err != nil {
		return nil, errors.New("hồ sơ không tồn tại")
	}

	if app.ApplicationStatus != "Approved" {
		return nil, errors.New("chỉ phần phòng cho hồ sơ đã được phê duyệt (Approved)")
	}

	room, err := s.roomRepo.FindByID(req.ExamRoomID)
	if err != nil {
		return nil, errors.New("phòng thi không tồn tại")
	}

	existingAppAssigned, _ := s.repo.FindByApplicationID(req.ApplicationID)
	if existingAppAssigned != nil {
		return nil, errors.New("hồ sơ này đã được phân phòng rồi")
	}

	existingSeat, _ := s.repo.FindByRoomAndSeat(req.ExamRoomID, req.SeatNumber)
	if existingSeat != nil {
		return nil, errors.New("số báo danh/số ghế này đã được sử dụng trong phòng thi")
	}

	count, err := s.repo.CountByRoomID(req.ExamRoomID)
	if err != nil {
		return nil, err
	}
	if room.Capacity > 0 && int(count) >= room.Capacity {
		return nil, errors.New("phòng thi đã đạt số lượng tối đa")
	}

	assignment := &entity.ExamRoomAssignment{
		ApplicationID: req.ApplicationID,
		ExamRoomID:    req.ExamRoomID,
		SeatNumber:    req.SeatNumber,
	}

	if err := s.repo.Create(assignment); err != nil {
		return nil, err
	}

	return s.GetDetail(assignment.ID)
}

func (s *examRoomAssignmentService) Update(id uint, req *dto.ExamRoomAssignmentUpdateReq) (*dto.ExamRoomAssignmentRes, error) {
	assignment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("dữ liệu phân phòng không tồn tại")
	}

	room, err := s.roomRepo.FindByID(req.ExamRoomID)
	if err != nil {
		return nil, errors.New("phòng thi không tồn tại")
	}

	if assignment.ExamRoomID != req.ExamRoomID || assignment.SeatNumber != req.SeatNumber {
		existingSeat, _ := s.repo.FindByRoomAndSeat(req.ExamRoomID, req.SeatNumber)
		if existingSeat != nil && existingSeat.ID != assignment.ID {
			return nil, errors.New("số báo danh/số ghế này đã được sử dụng trong phòng thi")
		}
	}

	if assignment.ExamRoomID != req.ExamRoomID {
		count, err := s.repo.CountByRoomID(req.ExamRoomID)
		if err != nil {
			return nil, err
		}
		if room.Capacity > 0 && int(count) >= room.Capacity {
			return nil, errors.New("phòng thi đã đạt số lượng tối đa")
		}
	}

	assignment.ExamRoomID = req.ExamRoomID
	assignment.SeatNumber = req.SeatNumber

	if err := s.repo.Update(assignment); err != nil {
		return nil, err
	}

	return s.GetDetail(assignment.ID)
}

func (s *examRoomAssignmentService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("dữ liệu phân phòng không tồn tại")
	}
	return s.repo.Delete(id)
}

func (s *examRoomAssignmentService) GetListByRoomID(roomID uint) ([]dto.ExamRoomAssignmentRes, error) {
	_, err := s.roomRepo.FindByID(roomID)
	if err != nil {
		return nil, errors.New("phòng thi không tồn tại")
	}

	assignments, err := s.repo.GetListByRoomID(roomID)
	if err != nil {
		return nil, err
	}

	var resData []dto.ExamRoomAssignmentRes
	for _, a := range assignments {
		cName := ""
		nID := ""
		rName := ""
		if a.Application != nil {
			cName = a.Application.CandidateFullName
			nID = a.Application.NationalID
		}
		if a.ExamRoom != nil {
			rName = a.ExamRoom.RoomName
		}
		resData = append(resData, dto.ExamRoomAssignmentRes{
			ID:                a.ID,
			ApplicationID:     a.ApplicationID,
			CandidateFullName: cName,
			NationalID:        nID,
			ExamRoomID:        a.ExamRoomID,
			RoomName:          rName,
			SeatNumber:        a.SeatNumber,
			CreatedAt:         a.CreatedAt.Format(time.RFC3339),
			UpdatedAt:         a.UpdatedAt.Format(time.RFC3339),
		})
	}
	if resData == nil {
		resData = []dto.ExamRoomAssignmentRes{}
	}
	return resData, nil
}
