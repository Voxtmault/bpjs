package sub_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/controllers"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
)

func Reference(e *echo.Echo, group *echo.Group) {

	referenceController := controllers.NewReferenceController(
		services.NewReferenceService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	reference := group.Group("/reference")

	// Uncategorized
	reference.GET("/polyclinic/:general_name", referenceController.GetPolyclinic)
	reference.GET("/health_facility/:hf_code/:hf_type", referenceController.GetHealthFacility)
	reference.GET("/nursing_class", referenceController.GetNursingClass)
	reference.GET("/treatment_room", referenceController.GetTreatmentRoom)
	reference.GET("/discharge_method", referenceController.GetDischargeMethod)
	reference.GET("/discharge_condition", referenceController.GetDischargeCondition)

	// Doctor Related
	reference.GET("/doctor/specialist", referenceController.GetDoctorSpecialists)
	reference.GET("/doctor/:general_name", referenceController.GetDoctors)
	reference.GET("/doctor/service/:service_type/:service_date/:specialist_code", referenceController.GetAttendingPhysicians)

	// Reconciliation Program / Program Rujuk Balik
	reference.GET("/prb/diagnosis", referenceController.GetPRBDiagnosis)
	reference.GET("/prb/medicine/:general_name", referenceController.GetPRBMedicine)

	// International Classification
	reference.GET("/diagnosis/:icd_x_code", referenceController.GetDiagnosis)
	reference.GET("/procedure/:icd_9_code", referenceController.GetProcedure)

	// Geographic Reference
	reference.GET("/province", referenceController.GetProvince)
	reference.GET("/provinsi", referenceController.GetProvince)
	reference.GET("/regency/:province_code", referenceController.GetRegency)
	reference.GET("/kabupaten/:province_code", referenceController.GetRegency)
	reference.GET("/district/:regency_code", referenceController.GetDistrict)
	reference.GET("/kecamatan/:regency_code", referenceController.GetDistrict)
}
