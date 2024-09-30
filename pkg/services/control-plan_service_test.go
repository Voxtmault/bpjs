package services

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

func TestGetViaSEP(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ControlPlanService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.GetViaSEP(context.Background(), "030107010217Y001465")
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error getting control plan via SEP: %v", err)
	}

	log.Println("Data: ", data)
}

func TestGetViaControlLetterNumber(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ControlPlanService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.GetViaControlLetterNumber(context.Background(), "0301R0010120K000003")
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error getting control plan via control letter number: %v", err)
	}

	log.Println("Data: ", data)
}

func TestGetControlPlanFromCardNumber(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ControlPlanService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.GetControlPlanFromCardNumber(context.Background(), &models.ControlPlansFromCardNumberParams{
		Month:      "01",
		Year:       "2021",
		CardNumber: "0002017051402",
		Filter:     "1",
	})
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error getting control plan via control letter number: %v", err)
	}

	log.Println("Data: ", data)
}

func TestGetControlPlans(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ControlPlanService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.GetControlPlans(context.Background(), &models.ControlPlanParams{
		StartDate: time.Now().Format(time.DateOnly),
		EndDate:   time.Now().AddDate(0, 0, 7).Format(time.DateOnly),
		Filter:    "1",
	})
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error getting control plan via control letter number: %v", err)
	}

	log.Println("Data: ", data)
}

func TestGetClinicControlPlans(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ControlPlanService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.GetClinicControlPlans(context.Background(), &models.ClinicControlParams{
		ControlType:        "2",
		Identifier:         "0301R0010323V000039",
		ControlPlannedDate: time.Now().Format(time.DateOnly),
	})
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error getting clinic assigned control plans: %v", err)
	}

	log.Println("Data: ", data)
}

func TestGetDoctorPracticeSchedule(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ControlPlanService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.GetDoctorPracticeSchedule(context.Background(), &models.DoctorScheduleParams{
		ControlType:        "2",
		PoliCode:           "INT",
		ControlPlannedDate: time.Now().Format(time.DateOnly),
	})
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error getting clinic assigned control plans: %v", err)
	}

	log.Println("Data: ", data)
}

func TestCreateControlPlan(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ControlPlanService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.CreateControlPlan(context.Background(), &models.ControlPlanCreate{
		SEPNumber:   "0301R0111018V000006",
		DoctorCode:  "12345",
		ClinicCode:  "INT",
		ControlDate: time.Now().AddDate(0, 0, 7).Format(time.DateOnly),
		User:        "Test Create Rencana Kontrol",
	})
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating control plan: %v", err)
	}

	log.Println("Data: ", data)
}

func TestUpdateControlPlan(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ControlPlanService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.UpdateControlPlan(context.Background(), &models.UpdateControlPlans{
		ControlLetterNumber: "0301R0110321K000002",
		SEPNumber:           "0301R0111018V000006",
		DoctorCode:          "12345",
		TargetPoliCode:      "INT",
		ControlPlannedDate:  time.Now().AddDate(0, 0, 14).Format(time.DateOnly),
		User:                "Test Create Rencana Kontrol",
	})
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating control plan: %v", err)
	}

	log.Println("Data: ", data)
}

func TestDeleteControlPlan(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ControlPlanService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	err := service.DeleteControlPlan(context.Background(), "", "Test Delete Surat Kontrol")
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating control plan: %v", err)
	}
}

func TestCreateInpatientCareOrder(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ControlPlanService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.CreateInpatientCareOrder(context.Background(), &models.ControlPlanCreate{
		CardNumber:  "0002017051402",
		DoctorCode:  "12345",
		ClinicCode:  "INT",
		ControlDate: time.Now().Format(time.DateOnly),
		User:        "Test Create SPRI",
	})
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating control plan: %v", err)
	}

	log.Println("Data: ", data)
}

func TestUpdateInpatientCareORder(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ControlPlanService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.UpdateInpatientCareOrder(context.Background(), &models.UpdateControlPlans{
		InpatientCareOrderNumber: "0301R0110421K000116",
		DoctorCode:               "12345",
		TargetPoliCode:           "INT",
		ControlPlannedDate:       time.Now().AddDate(0, 0, 14).Format(time.DateOnly),
		User:                     "Test Update SPRI",
	})
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating control plan: %v", err)
	}

	log.Println("Data: ", data)
}
