package service

import (
	"errors"
	"math"
	"time"

	"go_be_enrollment/internal/modules/examinerassignment/dto"
	"go_be_enrollment/internal/modules/examinerassignment/entity"
	"go_be_enrollment/internal/modules/examinerassignment/repository"
	examiner_repo "go_be_enrollment/internal/modules/examiner/repository"
	room_repo "go_be_enrollment/internal/modules/examroom/repository"
)

type ExaminerAssignmentService interface {
	GetList(filter *dto.ExaminerAssignmentFilter) (*dto.PaginatedExaminerAssignmentRes, error)
	GetDetail(id uint) (*dto.ExaminerAssignmentRes, error)
	Create(req *dto.ExaminerAssignmentCreateReq) (*dto.ExaminerAssignmentRes, error)
	Update(id uint, req *dto.ExaminerAssignmentUpdateReq) (*dto.ExaminerAssignmentRes, error)
	Delete(id uint) error
	GetListByRoomID(roomID uint) ([]dto.ExaminerAssignmentRes, error)
}

type examinerAssignmentService struct {
	repo         repository.ExaminerAssignmentRepository
	examinerRepo examiner_repo.ExaminerRepository
	roomRepo     room_repo.ExamRoomRepository
}

func NewExaminerAssignmentService(repo repository.ExaminerAssignmentRepository, examinerRepo examiner_repo.ExaminerRepository, roomRepo room_repo.ExamRoomRepository) ExaminerAssignmentService {
	return &examinerAssignmentService{repo: repo, examinerRepo: examinerRepo, roomRepo: roomRepo}
}

func (s *examinerAssignmentService) GetList(filter *dto.ExaminerAssignmentFilter) (*dto.PaginatedExaminerAssignmentRes, error) {
	list, total, err := s.repo.GetList(filter)
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

	var resData []dto.ExaminerAssignmentRes
	for _, a := range list {
		eName := ""
		rName := ""
		if a.Examiner != nil {
			eName = a.Examiner.FullName
		}
		if a.ExamRoom != nil {
			rName = a.ExamRoom.RoomName
		}
		resData = append(resData, dto.ExaminerAssignmentRes{
			ID:               a.ID,
			ExaminerID:       a.ExaminerID,
			ExaminerFullName: eName,
			ExamRoomID:       a.ExamRoomID,
			RoomName:         rName,
			Role:             a.Role,
			CreatedAt:        a.CreatedAt.Format(time.RFC3339),
			UpdatedAt:        a.UpdatedAt.Format(time.RFC3339),
		})
	}
	
	if resData == nil {
		resData = []dto.ExaminerAssignmentRes{}
	}

	return &dto.PaginatedExaminerAssignmentRes{
		Data:       resData,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}, nil
}

func (s *examinerAssignmentService) GetDetail(id uint) (*dto.ExaminerAssignmentRes, error) {
	a, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("dữ liệu phân công không tồn tại")
	}

	eName := ""
	rName := ""
	if a.Examiner != nil {
		eName = a.Examiner.FullName
	}
	if a.ExamRoom != nil {
		rName = a.ExamRoom.RoomName
	}

	return &dto.ExaminerAssignmentRes{
		ID:               a.ID,
		ExaminerID:       a.ExaminerID,
		ExaminerFullName: eName,
		ExamRoomID:       a.ExamRoomID,
		RoomName:         rName,
		Role:             a.Role,
		CreatedAt:        a.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        a.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *examinerAssignmentService) Create(req *dto.ExaminerAssignmentCreateReq) (*dto.ExaminerAssignmentRes, error) {
	_, err := s.examinerRepo.FindByID(req.ExaminerID)
	if err != nil {
		return nil, errors.New("cán bộ coi thi không tồn tại")
	}

	_, err = s.roomRepo.FindByID(req.ExamRoomID)
	if err != nil {
		return nil, errors.New("phòng thi không tồn tại")
	}

	existing, _ := s.repo.FindByUniqueKey(req.ExaminerID, req.ExamRoomID, req.Role)
	if existing != nil {
		return nil, errors.New("cán bộ coi thi đã được phân công với vai trò này ở phòng thi hiện tại")
	}

	assignment := &entity.ExaminerAssignment{
		ExaminerID: req.ExaminerID,
		ExamRoomID: req.ExamRoomID,
		Role:       req.Role,
	}

	if err := s.repo.Create(assignment); err != nil {
		return nil, err
	}

	return s.GetDetail(assignment.ID)
}

func (s *examinerAssignmentService) Update(id uint, req *dto.ExaminerAssignmentUpdateReq) (*dto.ExaminerAssignmentRes, error) {
	assignment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("dữ liệu phân công không tồn tại")
	}

	_, err = s.examinerRepo.FindByID(req.ExaminerID)
	if err != nil {
		return nil, errors.New("cán bộ coi thi không tồn tại")
	}

	_, err = s.roomRepo.FindByID(req.ExamRoomID)
	if err != nil {
		return nil, errors.New("phòng thi không tồn tại")
	}

	if assignment.ExaminerID != req.ExaminerID || assignment.ExamRoomID != req.ExamRoomID || assignment.Role != req.Role {
		existing, _ := s.repo.FindByUniqueKey(req.ExaminerID, req.ExamRoomID, req.Role)
		if existing != nil && existing.ID != assignment.ID {
			return nil, errors.New("cán bộ coi thi đã được phân công với vai trò này ở phòng thi hiện tại")
		}
	}

	assignment.ExaminerID = req.ExaminerID
	assignment.ExamRoomID = req.ExamRoomID
	assignment.Role = req.Role

	if err := s.repo.Update(assignment); err != nil {
		return nil, err
	}

	return s.GetDetail(assignment.ID)
}

func (s *examinerAssignmentService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("dữ liệu phân công không tồn tại")
	}
	return s.repo.Delete(id)
}

func (s *examinerAssignmentService) GetListByRoomID(roomID uint) ([]dto.ExaminerAssignmentRes, error) {
	_, err := s.roomRepo.FindByID(roomID)
	if err != nil {
		return nil, errors.New("phòng thi không tồn tại")
	}

	list, err := s.repo.GetListByRoomID(roomID)
	if err != nil {
		return nil, err
	}

	var resData []dto.ExaminerAssignmentRes
	for _, a := range list {
		eName := ""
		rName := ""
		if a.Examiner != nil {
			eName = a.Examiner.FullName
		}
		if a.ExamRoom != nil {
			rName = a.ExamRoom.RoomName
		}
		resData = append(resData, dto.ExaminerAssignmentRes{
			ID:               a.ID,
			ExaminerID:       a.ExaminerID,
			ExaminerFullName: eName,
			ExamRoomID:       a.ExamRoomID,
			RoomName:         rName,
			Role:             a.Role,
			CreatedAt:        a.CreatedAt.Format(time.RFC3339),
			UpdatedAt:        a.UpdatedAt.Format(time.RFC3339),
		})
	}
	if resData == nil {
		resData = []dto.ExaminerAssignmentRes{}
	}
	return resData, nil
}
