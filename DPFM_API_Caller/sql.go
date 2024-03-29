package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-payment-requisition-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-payment-requisition-creates-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-payment-requisition-creates-rmq-kube/sub_func_complementer"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var headerCreates *dpfm_api_output_formatter.HeaderCreates
	var headerDoc *[]dpfm_api_output_formatter.HeaderDoc
	var item *[]dpfm_api_output_formatter.Item
	for _, fn := range accepter {
		switch fn {
		case "Header":
			c.headerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
			headerCreates = dpfm_api_output_formatter.ConvertToHeaderCreates(subfuncSDC)
			// headerDoc = dpfm_api_output_formatter.ConvertToHeaderDoc(subfuncSDC)
		case "Item":
			c.itemCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
			item = dpfm_api_output_formatter.ConvertToItem(subfuncSDC)
		default:

		}
	}

	data := &dpfm_api_output_formatter.CreatesMessage{
		HeaderCreates: headerCreates,
		HeaderDoc:     headerDoc,
		Item:          item,
	}

	return data
}

// func (c *DPFMAPICaller) updateSqlProcess(
// 	ctx context.Context,
// 	mtx *sync.Mutex,
// 	input *dpfm_api_input_reader.SDC,
// 	output *dpfm_api_output_formatter.SDC,
// 	accepter []string,
// 	errs *[]error,
// 	log *logger.Logger,
// ) interface{} {
// 	var headerUpdates *dpfm_api_output_formatter.HeaderUpdates
// 	for _, fn := range accepter {
// 		switch fn {
// 		case "Header":
// 			res := c.headerUpdateSql(mtx, input, output, errs, log)
// 			headerUpdates = dpfm_api_output_formatter.ConvertToHeaderUpdates(*res)
// 		// case "Item":
// 		// go c.itemUpdateSql(wg, mtx, subFuncFin, log, errs, input, output)
// 		default:

// 		}
// 	}

// 	data := &dpfm_api_output_formatter.UpdatesMessage{
// 		HeaderUpdates: headerUpdates,
// 	}

// 	return data
// }

func (c *DPFMAPICaller) headerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_orders_header_dataの更新
	headerData := subfuncSDC.Message.Header
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "PaymentRequisitionHeader", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return
	}
	res.Success()
	if !checkResult(res) {
		// err = xerrors.New("Header Data cannot insert")
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	return
}

func (c *DPFMAPICaller) itemCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_orders_item_dataの更新
	for _, itemData := range *subfuncSDC.Message.Item {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "PaymentRequisitionItem", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return
		}
		res.Success()
		if !checkResult(res) {
			// err = xerrors.New("Item Data cannot insert")
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Data cannot insert"
			return
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}
	return
}

// func (c *DPFMAPICaller) headerUpdateSql(
// 	mtx *sync.Mutex,
// 	input *dpfm_api_input_reader.SDC,
// 	output *dpfm_api_output_formatter.SDC,
// 	errs *[]error,
// 	log *logger.Logger,
// ) *dpfm_api_input_reader.HeaderUpdates {
// 	b, _ := json.Marshal(input.Header)
// 	var req dpfm_api_input_reader.HeaderUpdates
// 	err := json.Unmarshal(b, &req)
// 	if err != nil {
// 		err = xerrors.Errorf("unmarshal error: %w", err)
// 		return nil
// 	}
// 	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": req, "function": "OrdersHeader", "runtime_session_id": 123})
// 	if err != nil {
// 		err = xerrors.Errorf("rmq error: %w", err)
// 		return nil
// 	}
// 	res.Success()
// 	if !checkResult(res) {
// 		// err = xerrors.New("Header Data cannot insert")
// 		output.SQLUpdateResult = getBoolPtr(false)
// 		output.SQLUpdateError = "Header Data cannot insert"
// 		return nil
// 	}
// 	if output.SQLUpdateResult == nil {
// 		output.SQLUpdateResult = getBoolPtr(true)
// 	}

// 	return &req
// }
