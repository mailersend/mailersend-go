package mailersend_test

import (
	"testing"

	"github.com/mailersend/mailersend-go"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateDmarcMonitorListOptions(t *testing.T) {
	options := mailersend.ListDmarcMonitorOptions{
		Page:  1,
		Limit: 25,
	}

	assert.Equal(t, 1, options.Page)
	assert.Equal(t, 25, options.Limit)
}

func TestCanCreateDmarcMonitorOptions(t *testing.T) {
	options := mailersend.CreateDmarcMonitorOptions{
		DomainID: "domain-id",
	}

	assert.Equal(t, "domain-id", options.DomainID)
}

func TestCanCreateUpdateDmarcMonitorOptions(t *testing.T) {
	options := mailersend.UpdateDmarcMonitorOptions{
		MonitorID:         "monitor-id",
		WantedDmarcRecord: "v=DMARC1; p=reject;",
	}

	assert.Equal(t, "monitor-id", options.MonitorID)
	assert.Equal(t, "v=DMARC1; p=reject;", options.WantedDmarcRecord)
}

func TestCanCreateDmarcReportOptions(t *testing.T) {
	options := mailersend.ListDmarcReportOptions{
		MonitorID: "monitor-id",
		Page:      1,
		Limit:     25,
	}

	assert.Equal(t, "monitor-id", options.MonitorID)
	assert.Equal(t, 1, options.Page)
	assert.Equal(t, 25, options.Limit)
}

func TestCanCreateDmarcReportSourcesOptions(t *testing.T) {
	options := mailersend.ListDmarcReportSourcesOptions{
		MonitorID: "monitor-id",
		Page:      1,
		Limit:     25,
	}

	assert.Equal(t, "monitor-id", options.MonitorID)
	assert.Equal(t, 1, options.Page)
	assert.Equal(t, 25, options.Limit)
}
