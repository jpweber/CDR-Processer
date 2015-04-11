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
	parts := strings.Split(egressCirIPEndPoint, "/")
	egressCirIPEndPointMap := make(map[string]string)
	egressCirIPEndPointMap["egressIPendpoint_local"] = parts[0]
	egressCirIPEndPointMap["egressIPendpoint_remote"] = parts[1]

	return egressCirIPEndPointMap
}

//SIP,b10e953ce9b911e298a600151737326e@204.124.15.58,<sip:7875315936@204.124.15.58:5065;user=phone>;tag=177540785-3792812521-352364184-1848784663,<sip:19106333559@76.10.220.14;user=phone>;tag=gK00f27408,0,,,,sip:19106333559@76.10.220.14;user=phone,,,,sip:7875315936@204.124.15.58:5065;user=phone,sip:19106333559@76.10.220.14:5060,,,,1,BYE,16,0,0,,0,0,,,,,,,,1,0,0,0,,
func IngressProtocolVariantSpecData(ingressProtocolVariantSpecData string) map[string]string {
	parts := strings.Split(ingressProtocolVariantSpecData, ",")
	fmt.Println(len(parts))
	ingressProtocolVariantSpecDataMap := make(map[string]string)
	subfields := []string{
		"IPVSD_Protocol_Variant",
		"IPVSD_Call_ID",
		"IPVSD_From",
		"IPVSD_To,",
		"IPVSD_Blank_Field",
		"IPVSD_SIP-T_Version",
		"IPVSD_SIP_URI_PAI_Display Name",
		"IPVSD_P-K_CallFwdLast_User_Param",
		"IPVSD_SIP Req URI User/Host",
		"IPVSD_SIP URI PAI User/Host",
		"IPVSD_Proxy_Auth_Username",
		"IPVSD_Tel_URI_PAI_Display_Name",
		"IPVSD_Invite_Contact_Header",
		"IPVSD_200_OK_Invite_Contact_Header",
		"IPVSD_P-K_CallFwdOrig_Redir_Reason",
		"IPVSD_Tel_URI_PAI_User_Name",
		"IPVSD_P-Sig_Info_Contractor_Num",
		"IPVSD_ACK_Rx'd_for_200_OK",
		"IPVSD_Status_Msg_for_Call_Release",
		"IPVSD_Reason_Header_Value_Q850",
		"IPVSD_NAPT_Status_Signaling",
		"IPVSD_NAPT_Status_Media",
		"IPVSD_NAPT_Orig_Peer_SDP_Addr",
		"IPVSD_UUI_Sending_Count",
		"IPVSD_UUI_Receiving_Count",
		"IPVSD_Service_Info",
		"IPVSD_ICID",
		"IPVSD_Gen'd_Host",
		"IPVSD_Orig_IOI",
		"IPVSD_Term_IOI",
		"IPVSD_Special_Routing_Table_Num",
		"IPVSD_IP_Address_For_FQDN_Call",
		"IPVSD_SIP_Transport_Protocol",
		"IPVSD_Direct_Media",
		"IPVSD_Inbound_SMM_Indicator",
		"IPVSD_Outbound_SMM_Indicator",
		"IPVSD_Originating_Charge_Area",
		"IPVSD_Terminating_Charge_Area",
		"IPVSD_Feature_Tag_Contact",
		"IPVSD_Feature_Tag_Accept-Contact",
		"IPVSD_P-Charging-Function-Address",
		"IPVSD_P-Called-Party-Id",
		"IPVSD_P-Visited-Network-Id",
		"IPVSD_Direct_Media_with_NAPT_Call",
		"IPVSD_Inbound_SMM_Profile_Name",
		"IPVSD_Outbound_SMM_Profile_Name",
	}

	for i, subfield := range subfields {
		fmt.Println(subfield)
		fmt.Println(i)
		if i < len(parts) {
			ingressProtocolVariantSpecDataMap[subfield] = parts[i]
		} else {
			ingressProtocolVariantSpecDataMap[subfield] = ""
		}
	}

	return ingressProtocolVariantSpecDataMap
}
