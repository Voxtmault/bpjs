package sub_routes

import "github.com/labstack/echo/v4"

func Reference(group *echo.Group) {

	reference := group.Group("/reference")

	reference.GET("/polyclinic/:general_name", nil)
	reference.GET("/health_facility/:hf_code/:hf_type", nil)
	reference.GET("/nursing_class", nil)
	reference.GET("/treatment_room", nil)
	reference.GET("/discharge_method", nil)
	reference.GET("/post_discharge_condition", nil)

	// Doctor Related
	reference.GET("/doctor/specialist", nil)
	reference.GET("/doctor/:general_name", nil)
	reference.GET("/doctor/service/:service_type/:service_date/:specialist_code", nil)

	// Reconciliation Program / Program Rujuk Balik
	reference.GET("/prb/diagnosis", nil)
	reference.GET("/prb/medicine/:general_name", nil)

	// International Classification
	reference.GET("/diagnosis/:icd_x_code", nil)
	reference.GET("/procedure/:icd_9_code", nil)

	// Geographic Reference
	reference.GET("/province", nil)
	reference.GET("/provinsi", nil)
	reference.GET("/regency/:province_code", nil)
	reference.GET("/kabupaten/:province_code", nil)
	reference.GET("/district/:regency_code", nil)
	reference.GET("/kecamatan/:regency_code", nil)
}
