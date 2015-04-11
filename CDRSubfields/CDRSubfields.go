package CDRSubfields

import (
	"fmt"
	"strconv"
	"strings"
)

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

func RouteSelected(routeSelected string) map[string]string {

	parts := strings.Split(routeSelected, ":")
	routeSelectedMap := make(map[string]string)
	routeSelectedMap["RS_Gateway"] = parts[0]
	routeSelectedMap["RS_Trunkgroup"] = parts[1]

	return routeSelectedMap
}
