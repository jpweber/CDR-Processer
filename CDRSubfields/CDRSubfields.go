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
		if i < len(parts) {
			ingressProtocolVariantSpecDataMap[subfield] = parts[i]
		} else {
			ingressProtocolVariantSpecDataMap[subfield] = ""
		}
	}

	return ingressProtocolVariantSpecDataMap
}

//SIP,b10e953ce9b911e298a600151737326e@204.124.15.58,<sip:7875315936@204.124.15.58:5065;user=phone>;tag=177540785-3792812521-352364184-1848784663,<sip:19106333559@76.10.220.14;user=phone>;tag=gK00f27408,0,,,,sip:19106333559@76.10.220.14;user=phone,,,,sip:7875315936@204.124.15.58:5065;user=phone,sip:19106333559@76.10.220.14:5060,,,,1,BYE,16,0,0,,0,0,,,,,,,,1,0,0,0,,
func EngressProtocolVariantSpecData(egressProtocolVariantSpecData string) map[string]string {
	parts := strings.Split(egressProtocolVariantSpecData, ",")
	fmt.Println(len(parts))
	egressProtocolVariantSpecDataMap := make(map[string]string)
	subfields := []string{
		"EPVSD_Protocol_Variant",
		"EPVSD_Call_ID",
		"EPVSD_From",
		"EPVSD_To,",
		"EPVSD_Blank_Field",
		"EPVSD_SIP-T_Version",
		"EPVSD_SIP_URI_PAI_Display Name",
		"EPVSD_P-K_CallFwdLast_User_Param",
		"EPVSD_SIP Req URI User/Host",
		"EPVSD_SIP URI PAI User/Host",
		"EPVSD_Proxy_Auth_Username",
		"EPVSD_Tel_URI_PAI_Display_Name",
		"EPVSD_Invite_Contact_Header",
		"EPVSD_200_OK_Invite_Contact_Header",
		"EPVSD_P-K_CallFwdOrig_Redir_Reason",
		"EPVSD_Tel_URI_PAI_User_Name",
		"EPVSD_P-Sig_Info_Contractor_Num",
		"EPVSD_ACK_Rx'd_for_200_OK",
		"EPVSD_Status_Msg_for_Call_Release",
		"EPVSD_Reason_Header_Value_Q850",
		"EPVSD_NAPT_Status_Signaling",
		"EPVSD_NAPT_Status_Media",
		"EPVSD_NAPT_Orig_Peer_SDP_Addr",
		"EPVSD_UUI_Sending_Count",
		"EPVSD_UUI_Receiving_Count",
		"EPVSD_Service_Info",
		"EPVSD_ICID",
		"EPVSD_Gen'd_Host",
		"EPVSD_Orig_IOI",
		"EPVSD_Term_IOI",
		"EPVSD_Special_Routing_Table_Num",
		"EPVSD_IP_Address_For_FQDN_Call",
		"EPVSD_SIP_Transport_Protocol",
		"EPVSD_Direct_Media",
		"EPVSD_Inbound_SMM_Indicator",
		"EPVSD_Outbound_SMM_Indicator",
		"EPVSD_Originating_Charge_Area",
		"EPVSD_Terminating_Charge_Area",
		"EPVSD_Feature_Tag_Contact",
		"EPVSD_Feature_Tag_Accept-Contact",
		"EPVSD_P-Charging-Function-Address",
		"EPVSD_P-Called-Party-Id",
		"EPVSD_P-Visited-Network-Id",
		"EPVSD_Direct_Media_with_NAPT_Call",
		"EPVSD_Inbound_SMM_Profile_Name",
		"EPVSD_Outbound_SMM_Profile_Name",
	}

	for i, subfield := range subfields {
		fmt.Println(subfield)
		fmt.Println(i)
		if i < len(parts) {
			egressProtocolVariantSpecDataMap[subfield] = parts[i]
		} else {
			egressProtocolVariantSpecDataMap[subfield] = ""
		}
	}

	return egressProtocolVariantSpecDataMap
}

func IngressCodecType(ingressCodecType string) map[string]string {
	parts := strings.Split(ingressCodecType, ":")
	ingressCodecTypeMap := make(map[string]string)
	ingressCodecTypeMap["ICT_Network_Type"] = parts[0]
	ingressCodecTypeMap["ICT_Codec_Type"] = parts[1]
	ingressCodecTypeMap["ICT_Audio_Encoding_Type"] = parts[2]

	return ingressCodecTypeMap
}

func EgressCodecType(egressCodecType string) map[string]string {
	parts := strings.Split(egressCodecType, ":")
	egressCodecTypeMap := make(map[string]string)
	egressCodecTypeMap["ECT_Network_Type"] = parts[0]
	egressCodecTypeMap["ECT_Codec_Type"] = parts[1]
	egressCodecTypeMap["ECT_Audio_Encoding_Type"] = parts[2]

	return egressCodecTypeMap
}
