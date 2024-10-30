package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type Monitoring interface {
	GetMonitoringDataKunjungan(ctx context.Context, tanggalPelayanan, jenisPelayanan string) ([]*models.MonitoringDataKunjungan, error)
	GetMonitoringDataKlaim(ctx context.Context, tanggalPelayanan, jenisPelayanan, statusKlaim string) ([]*models.DataKlaim, error)
	GetMonitoringHistoryPelayananPeserta(ctx context.Context, noKartu, tanggalMulai, tanggalAkhir string) ([]*models.MonitoringDataKunjungan, error)
	GetMonitoringKlaimJaminanJasaraharja(ctx context.Context, jenisPelayanan, tanggalMulai, tanggalAkhir string) ([]*models.DataJasaraharja, error)
}
