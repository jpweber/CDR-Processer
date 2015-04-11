package CDRSubfields

import (
	"fmt"
	"strconv"
	"strings"
)

//0x0001A09A00720365
func AccountingID(accountingid string) map[string]string {
	shelf, _ := strconv.ParseInt(accountingid[2:6], 0, 64)
	bootCount, _ := strconv.ParseInt("0x"+accountingid[6:10], 0, 64)
	callId := accountingid[10:16]

	accountingidMap := make(map[string]string)
	accountingidMap["shelf"] = fmt.Sprintf("%d", shelf)
	accountingidMap["bootCount"] = fmt.Sprintf("%d", bootCount)
	accountingidMap["callId"] = callId

	//!debug
	//fmt.Println(accountingidMap)
	return accountingidMap
}

//CMHGSX3:03-CMHGSX3-NTAND-ISUP01
func RouteSelected(routeSelected string) map[string]string {

	parts := strings.Split(routeSelected, ":")
	routeSelectedMap := make(map[string]string)
	routeSelectedMap["RS_Gateway"] = parts[0]
	routeSelectedMap["RS_Trunkgroup"] = parts[1]

	return routeSelectedMap
}

////76.10.220.16:12808/204.124.15.102:62864
func IngressCirIPEndPoint(ingressCirIPEndPoint string) map[string]string {
	parts := strings.Split(ingressCirIPEndPoint, "/")
	ingressCirIPEndPointMap := make(map[string]string)
	ingressCirIPEndPointMap["ingressIPendpoint_local"] = parts[0]
	ingressCirIPEndPointMap["ingressIPendpoint_remote"] = parts[1]

	return ingressCirIPEndPointMap
}

func EgressCirIPEndPoint(egressCirIPEndPoint string) map[string]string {
	fmt.Println(egressCirIPEndPoint)
	parts := strings.Split(egressCirIPEndPoint, "/")
	egressCirIPEndPointMap := make(map[string]string)
	egressCirIPEndPointMap["egressIPendpoint_local"] = parts[0]
	egressCirIPEndPointMap["egressIPendpoint_remote"] = parts[1]

	return egressCirIPEndPointMap
}
