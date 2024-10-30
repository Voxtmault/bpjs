package controllers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
)

type ReferenceController struct {
	service  interfaces.Reference
	validate echo.Validator
}

func NewReferenceController(service interfaces.Reference, validate echo.Validator) *ReferenceController {
	return &ReferenceController{
		service:  service,
		validate: validate,
	}
}

func (s ReferenceController) GetPolyclinic(c echo.Context) error {
	var res Response

	generalName := c.Param("general_name")
	if generalName == "" || generalName == ":general_name" {
		res.Message = "kode / nama poli tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.PolyclinicsReference(c.Request().Context(), generalName)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetPolyclinic", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetHealthFacility(c echo.Context) error {
	var res Response

	healthFacilityCode := c.Param("hf_code")
	if healthFacilityCode == "" || healthFacilityCode == ":hf_code" {
		res.Message = "nama / kode faskes tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	healthFacilityType := c.Param("hf_type")
	if healthFacilityType == "" || healthFacilityType == ":hf_type" {
		res.Message = "tipe faskes tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.HealthFacilityReference(c.Request().Context(), healthFacilityCode, healthFacilityType)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetHealthFacility", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetNursingClass(c echo.Context) error {
	var res Response

	var err error
	res.Data, err = s.service.NursingClassReference(c.Request().Context())
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetNursingClass", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetTreatmentRoom(c echo.Context) error {
	var res Response

	var err error
	res.Data, err = s.service.TreatmentRoomReference(c.Request().Context())
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetNursingClass", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetDischargeMethod(c echo.Context) error {
	var res Response

	var err error
	res.Data, err = s.service.DischargeMethodReference(c.Request().Context())
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetDischargeMethod", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetPostDischargeCondition(c echo.Context) error {
	var res Response

	var err error
	res.Data, err = s.service.PostDischargeReference(c.Request().Context())
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetDischargeMethod", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetDoctorSpecialists(c echo.Context) error {
	var res Response

	var err error
	res.Data, err = s.service.SpecialistReference(c.Request().Context())
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetDoctorSpecialists", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetDoctors(c echo.Context) error {
	var res Response

	generalName := c.Param("general_name")
	if generalName == "" || generalName == ":general_name" {
		res.Message = "nama dokter / dpjp tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.DoctorReference(c.Request().Context(), generalName)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetDoctors", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetAttendingPhysicians(c echo.Context) error {
	var res Response

	serviceType := c.Param("service_type")
	if serviceType == "" || serviceType == ":service_type" {
		res.Message = "jenis layanan tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	serviceDate := c.Param("service_date")
	if serviceDate == "" || serviceDate == ":service_date" {
		res.Message = "tanggal pelayanan / SEP tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	specialistCode := c.Param("specialist_code")
	if specialistCode == "" || specialistCode == ":specialist_code" {
		res.Message = "kode spesialis / sub-spesialis tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.AttendingPhysicianReference(c.Request().Context(), serviceType, serviceDate, specialistCode)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetAttendingPhysicians", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetDiagnosis(c echo.Context) error {
	var res Response

	icdXCode := c.Param("icd_x_code")
	if icdXCode == "" || icdXCode == ":icd_x_code" {
		res.Message = "kode ICD-X tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.DiagnoseReference(c.Request().Context(), icdXCode)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetDiagnosis", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetProcedure(c echo.Context) error {
	var res Response

	icd9Code := c.Param("icd_9_code")
	if icd9Code == "" || icd9Code == ":icd_9_code" {
		res.Message = "kode ICD-9 tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.ProcedureReference(c.Request().Context(), icd9Code)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetProcedure", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetProvince(c echo.Context) error {
	var res Response

	var err error
	res.Data, err = s.service.ProvinceReference(c.Request().Context())
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetProvince", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetRegency(c echo.Context) error {
	var res Response

	provinceCode := c.Param("province_code")
	if provinceCode == "" || provinceCode == ":province_code" {
		res.Message = "kode provinsi tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.RegencyReference(c.Request().Context(), provinceCode)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetRegency", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetDistrict(c echo.Context) error {
	var res Response

	regencyCode := c.Param("regency_code")
	if regencyCode == "" || regencyCode == ":regency_code" {
		res.Message = "kode kabupaten tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.DistrictReference(c.Request().Context(), regencyCode)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetDistrict", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetPRBDiagnosis(c echo.Context) error {
	var res Response

	var err error
	res.Data, err = s.service.DiagnosePRBReference(c.Request().Context())
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetPRBDiagnosis", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferenceController) GetPRBMedicine(c echo.Context) error {
	var res Response

	generalName := c.Param("general_name")
	if generalName == "" || generalName == ":general_name" {
		res.Message = "nama obat generik tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.MedicinePRBReference(c.Request().Context(), generalName)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("reference_controller -> GetPRBMedicine", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
