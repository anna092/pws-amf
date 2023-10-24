package producer

import (
	"encoding/json"
	"log"

	amf_context "github.com/free5gc/amf/internal/context"
	ngap_message "github.com/free5gc/amf/internal/ngap/message"
	"github.com/free5gc/openapi/models"
)

func NonUeN2MessageTransferProcedure(amfSelf *amf_context.AMFContext, message models.NonUeN2MessageTransferRequest) {
	var keyValueN2Information map[string]string
	err := json.Unmarshal(message.BinaryDataN2Information, &keyValueN2Information)
	if err != nil {
		log.Fatal(err)
	}
	// var globalRanNodeId models.GlobalRanNodeId
	// // var plmnId models.PlmnId
	// // plmnId.Mcc = "2018"
	// // plmnId.Mnc = "93"
	// var gnbId models.GNbId
	// gnbId.BitLength = 32
	// gnbId.GNBValue = "00000001"
	// globalRanNodeId.GNbId = &gnbId
	// // globalRanNodeId.PlmnId = &plmnId
	// ran, ok := amfSelf.AmfRanFindByRanID(globalRanNodeId)
	// if !ok {
	// 	log.Fatal("UE Not found")
	// }
	amf_context.AMF_Self().AmfRanPool.Range(func(key, value interface{}) bool {
		amfRan := value.(*amf_context.AmfRan)
		ngap_message.SendWriteReplaceWarningRequest(amfRan, keyValueN2Information)
		return true
	})

}
