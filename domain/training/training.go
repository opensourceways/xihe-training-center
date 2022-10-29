package training

import (
	"github.com/opensourceways/xihe-training-center/domain"
)

type Training interface {
	Create(*domain.UserTraining) (domain.JobInfo, error)
	Delete(string) error
	Terminate(string) error
	GetLogDownloadURL(string) (string, error)
	GetDetail(string) (domain.JobDetail, error)

	// GetLogFilePath return the obs path of log
	GetLogFilePath(logDir string) (string, error)

	// GenOutput generates the zip file of output dir and
	// return the obs path of that file.
	GenOutput(outputDir string) (string, error)

	// GenAim generates the zip file of aim dir
	// and return the obs path of that file.
	GenAim(aimDir string) (string, error)
}
