package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-customer-product-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-customer-product-reads-rmq-kube/DPFM_API_Output_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var customerProduct *[]dpfm_api_output_formatter.CustomerProduct
	for _, fn := range accepter {
		switch fn {
		case "CustomerProduct":
			func() {
				customerProduct = c.CustomerProduct(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		CustomerProduct: customerProduct,
	}

	return data
}
