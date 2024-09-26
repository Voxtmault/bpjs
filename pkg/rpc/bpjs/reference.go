package bpjs

import (
	"context"
	"fmt"
	"time"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
	"github.com/voxtmault/bpjs-rs-module/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/voxtmault/bpjs-service-proto/go"
)

type BPJSReferenceRPCService struct {
	pb.UnimplementedReferenceServiceServer
	Service interfaces.Reference
}

func InitReferenceService() *BPJSReferenceRPCService {
	// Init Services
	s := BPJSReferenceRPCService{
		Service: &services.ReferenceService{
			HttpHandler: &services.RequestHandlerService{
				Security: &services.BPJSSecurityService{},
			},
		},
	}
	return &s
}

func (s *BPJSReferenceRPCService) DiagnosisReference(ctx context.Context, in *pb.DiagnosisReferenceRequest) (*pb.ReferenceResponse, error) {

	obj := models.DiagnosisReferenceParams{
		Code: in.GetDiagnosisCode(),
	}

	if err := utils.GetValidator().Struct(obj); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("invalid request: %w", err).Error())
	}

	data, err := s.Service.DiagnoseReference(ctx, obj.Code)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}

func (s *BPJSReferenceRPCService) DoctorReference(ctx context.Context, in *pb.DoctorReferenceRequest) (*pb.ReferenceResponse, error) {

	obj := models.DoctorReferenceParams{
		ServiceType:    in.GetServiceType(),
		ServiceDate:    in.GetServiceDate(),
		SpecialistCode: in.GetSpecialistCode(),
	}

	// Checks if the Service Date is not included in the request, if it isn't then default to now / today
	if obj.ServiceDate == "" {
		obj.ServiceDate = time.Now().Format(time.DateOnly)
	}

	if err := utils.GetValidator().Struct(obj); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("invalid request: %w", err).Error())
	}

	data, err := s.Service.DoctorReference(ctx, obj.ServiceType, obj.ServiceDate, obj.SpecialistCode)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}

func (s *BPJSReferenceRPCService) PoliclinicsReference(ctx context.Context, in *pb.PoliclinicsReferenceRequest) (*pb.ReferenceResponse, error) {
	obj := models.PoliReferenceParams{
		PoliCode: in.GetPoliCode(),
		PoliName: in.GetPoliName(),
	}

	if err := utils.GetValidator().Struct(obj); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("invalid request: %w", err).Error())
	}

	var params string
	if obj.PoliCode != "" {
		params = obj.PoliCode
	} else if obj.PoliName != "" {
		params = obj.PoliName
	} else {
		return nil, status.Errorf(codes.Internal, "Invalid Poli Reference Params. Somehow bypassed the validation step")
	}

	data, err := s.Service.PoliclinicsReference(ctx, params)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}

func (s *BPJSReferenceRPCService) HealthFacilityReference(ctx context.Context, in *pb.HealthFacilityReferenceRequest) (*pb.ReferenceResponse, error) {
	obj := models.HealthFacilityReferenceParams{
		HealthFacilityName: in.GetHcName(),
		HealthFacilityType: in.GetHcType(),
	}

	if err := utils.GetValidator().Struct(obj); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("invalid request: %w", err).Error())
	}

	data, err := s.Service.HealthFacilityReference(ctx, obj.HealthFacilityName, obj.HealthFacilityType)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}

func (s *BPJSReferenceRPCService) ProcedureReference(ctx context.Context, in *pb.ProcedureReferenceRequest) (*pb.ReferenceResponse, error) {
	obj := models.ProcedureReferenceParams{
		ProcedureCode: in.GetProcedureCode(),
	}

	if err := utils.GetValidator().Struct(obj); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("invalid request: %w", err).Error())
	}

	data, err := s.Service.ProcedureReference(ctx, obj.ProcedureCode)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}

func (s *BPJSReferenceRPCService) NursingClassReference(ctx context.Context, in *pb.ReferenceRequest) (*pb.ReferenceResponse, error) {
	data, err := s.Service.NursingClassReference(ctx)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}

func (s *BPJSReferenceRPCService) SpecialistReference(ctx context.Context, in *pb.ReferenceRequest) (*pb.ReferenceResponse, error) {
	data, err := s.Service.SpecialistReference(ctx)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}

func (s *BPJSReferenceRPCService) DischargeMethodReference(ctx context.Context, in *pb.ReferenceRequest) (*pb.ReferenceResponse, error) {
	data, err := s.Service.DischargeMethodReference(ctx)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}

func (s *BPJSReferenceRPCService) PostDischargeReference(ctx context.Context, in *pb.ReferenceRequest) (*pb.ReferenceResponse, error) {
	data, err := s.Service.PostDischargeReference(ctx)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}

func (s *BPJSReferenceRPCService) ProvinceReference(ctx context.Context, in *pb.ReferenceRequest) (*pb.ReferenceResponse, error) {
	data, err := s.Service.ProvinceReference(ctx)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}

func (s *BPJSReferenceRPCService) RegencyReference(ctx context.Context, in *pb.RegencyReferenceRequest) (*pb.ReferenceResponse, error) {
	obj := models.RegencyReferenceParams{
		ProvinceCode: in.GetProvinceCode(),
	}

	if err := utils.GetValidator().Struct(obj); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("invalid request: %w", err).Error())
	}

	data, err := s.Service.RegencyReference(ctx, obj.ProvinceCode)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}

func (s *BPJSReferenceRPCService) DistrictReference(ctx context.Context, in *pb.DistrictReferenceRequest) (*pb.ReferenceResponse, error) {
	obj := models.DistrictReferenceParams{
		RegencyCode: in.GetRegencyCode(),
	}

	if err := utils.GetValidator().Struct(obj); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("invalid request: %w", err).Error())
	}

	data, err := s.Service.DistrictReference(ctx, obj.RegencyCode)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}

func (s *BPJSReferenceRPCService) AttendingPhysicianReference(ctx context.Context, in *pb.AttendingPhysicianReferenceRequest) (*pb.ReferenceResponse, error) {
	obj := models.AttendingPhysicianReferenceParams{
		DoctorName: in.GetDoctorCode(),
	}

	if err := utils.GetValidator().Struct(obj); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("invalid request: %w", err).Error())
	}

	data, err := s.Service.AttendingPhysicianReference(ctx, obj.DoctorName)
	if err != nil {
		// log.Println(eris.Cause(err))
		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	arrObj := []*pb.Reference{}
	for _, item := range data {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return &pb.ReferenceResponse{
		StatusCode: int32(codes.OK),
		Message:    "Success",
		Reference:  arrObj,
	}, nil
}
