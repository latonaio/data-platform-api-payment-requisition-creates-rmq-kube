package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-payment-requisition-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-payment-requisition-creates-rmq-kube/config"
	"sync"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient,

) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:  context.Background(),
		conf: conf,
		rmq:  rmq,
	}
}

func (c *DPFMAPICaller) AsyncPaymentRequisitionCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,

	log *logger.Logger,
	// msg rabbitmq.RabbitmqMessage,
) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	sqlUpdateFin := make(chan error)

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			go c.Header(&wg, &mtx, sqlUpdateFin, log, &errs, input)
		case "HeaderPDF":
			go c.HeaderPDF(&wg, &mtx, sqlUpdateFin, log, &errs, input)
		case "Item":
			go c.Item(&wg, &mtx, sqlUpdateFin, log, &errs, input)
		default:
			wg.Done()
		}
	}

	// 後処理
	ticker := time.NewTicker(10 * time.Second)
	select {
	case e := <-sqlUpdateFin:
		if e != nil {
			mtx.Lock()
			errs = append(errs, e)
			return errs
		}
	case <-ticker.C:
		mtx.Lock()
		errs = append(errs, xerrors.New("time out"))
		return errs
	}

	return nil
}

func (c *DPFMAPICaller) Header(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs *[]error, sdc *dpfm_api_input_reader.SDC) {
	var err error = nil
	defer wg.Done()
	defer func() {
		errFin <- err
	}()
	sessionID := sdc.RuntimeSessionID
	ctx := context.Background()

	// data_platform_payment_requisition_header_dataの更新
	headerData := sdc.PaymentRequisition
	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "PaymentRequisitionHeader", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return
	}
	res.Success()
	if !checkResult(res) {
		// err = xerrors.New("Header Data cannot insert")
		sdc.SQLUpdateResult = getBoolPtr(false)
		sdc.SQLUpdateError = "Header Data cannot insert"
		return
	}

	sdc.SQLUpdateResult = getBoolPtr(true)
	return
}

func (c *DPFMAPICaller) HeaderPDF(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs *[]error, sdc *dpfm_api_input_reader.SDC) {
	var err error = nil
	defer wg.Done()
	defer func() {
		errFin <- err
	}()
	sessionID := sdc.RuntimeSessionID
	ctx := context.Background()

	// data_platform_payment_requisition_item_dataの更新
	headerPDFData := sdc.PaymentRequisition.HeaderPDF
	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerPDFData, "function": "PaymentRequisitionHeaderPDF", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return
	}
	res.Success()
	if !checkResult(res) {
		// err = xerrors.New("Header PDF Data cannot insert")
		sdc.SQLUpdateResult = getBoolPtr(false)
		sdc.SQLUpdateError = "Header PDF Data cannot insert"
		return
	}

	sdc.SQLUpdateResult = getBoolPtr(true)
	return
}

func (c *DPFMAPICaller) Item(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs *[]error, sdc *dpfm_api_input_reader.SDC) {
	var err error = nil
	defer wg.Done()
	defer func() {
		errFin <- err
	}()
	sessionID := sdc.RuntimeSessionID
	ctx := context.Background()

	// data_platform_payment_requisition_item_dataの更新
	itemData := sdc.PaymentRequisition.Item
	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "PaymentRequisitionItem", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return
	}
	res.Success()
	if !checkResult(res) {
		// err = xerrors.New("Item Data cannot insert")
		sdc.SQLUpdateResult = getBoolPtr(false)
		sdc.SQLUpdateError = "Item Data cannot insert"
		return
	}

	sdc.SQLUpdateResult = getBoolPtr(true)
	return
}

func checkResult(msg rabbitmq.RabbitmqMessage) bool {
	data := msg.Data()
	_, ok := data["result"]
	if !ok {
		return false
	}
	result, ok := data["result"].(string)
	if !ok {
		return false
	}
	return result == "success"

}

func getBoolPtr(b bool) *bool {
	return &b
}
