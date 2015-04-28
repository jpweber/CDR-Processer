/*"
* @Author: Jim Weber"
* @Date:   2015-01-28 10:09:26"
* @Last Modified by:   jpweber
* @Last Modified time: 2015-04-27 22:40:08
 */

package CDR

import (
	// "encoding/csv"
	"encoding/json"
	"fmt"
	"ko/CDRSubfields"
	"os"
	"strings"
	"sync"
)

// Cdrcollection holds too slices containing Stop records and attempt records
type CdrCollection struct {
	Stops    [][]string
	Attempts [][]string
	Starts   [][]string
}

//creates a complete record map
func CreateRecordMap(wg *sync.WaitGroup, records [][]string, recordType string) []map[string]string {
	recordMap := make([]map[string]string, len(records))
	var cdrData map[string]string
	for i, value := range records {
		switch recordType {
		case "attempts":
			cdrData = FillCDRMap(CdrAttemptKeys(), value)
		case "stops":
			cdrData = FillCDRMap(CdrStopKeys(), value)
		case "starts":
			cdrData = FillCDRMap(CdrStartKeys(), value) //normal
		}
		cdrData = BreakOutSubFields(cdrData)
		recordMap[i] = cdrData
	}
	wg.Done()
	// fmt.Println(recordType + "Done")
	return recordMap
}

// FillCDRMap takes an array of cdr keys and and array of cdr values
// and returns a map with the passed in key value pairs
func FillCDRMap(keys []string, values []string) map[string]string {

	cdrMap := make(map[string]string)

	// for i, value := range values {
	// 	fmt.Println(keys[i], value, "\n")
	// 	cdrMap[keys[i]] = value
	// }

	for i, key := range keys {
		if i < len(values) {
			cdrMap[key] = values[i]
		} else {
			// if we don't have all the values to fill the map
			// fill with blanks so we don't check beyond bounds of values array
			cdrMap[key] = ""
		}
	}

	return cdrMap
}

func KeysString(m map[string]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
		fmt.Println(k)
	}
	fmt.Println(len(keys))
	return strings.Join(keys, ", ")
}

func ValuesString(m map[string]string) string {
	values := make([]string, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	// fmt.Println(len(values))
	//      bullshit extra blank at start
	return "'','','" + strings.Join(values, "', '") + "' "
}

func BreakOutSubFields(cdrMap map[string]string) map[string]string {
	//this is ripe for an optimizing refactor. Working it out very non-DRY for simplicity of business logic
	//refactor this down later. Thinking on break out function with a pre-made map of fields that have subfields
	//to break out then just loop over them. Put a function in the subifields package to determine which subfield
	// function to use for the breakout/parsing.

	//0x0001A09A00720365
	if cdrMap["Accounting_ID"] != "" {
		accountingIdSfs := CDRSubfields.AccountingID(cdrMap["Accounting_ID"])
		for key, value := range accountingIdSfs {
			cdrMap[key] = value
		}
	} else {
		// fmt.Println("Accounting ID not found")
	}

	//CMHGSX3:03-CMHGSX3-NTAND-ISUP01
	if cdrMap["Route_Selected"] != "" {
		routeSelectedSfs := CDRSubfields.RouteSelected(cdrMap["Route_Selected"])
		cdrMap["RS_Gateway"] = routeSelectedSfs["RS_Gateway"]
		cdrMap["RS_TrunkGroup"] = routeSelectedSfs["RS_Trunkgroup"]
	} else {
		// fmt.Println("Accounting ID not found")
	}

	if cdrMap["Ingress_IP_Circuit_End_Point"] != "" {
		ingressIpSfs := CDRSubfields.IngressCirIPEndPoint(cdrMap["Ingress_IP_Circuit_End_Point"])
		cdrMap["IIPE_local"] = ingressIpSfs["ingressIPendpoint_local"]
		cdrMap["IIPE_remote"] = ingressIpSfs["ingressIPendpoint_remote"]
	} else {
		// fmt.Println("Accounting ID not found")
	}

	if cdrMap["Egress_IP_Circuit_End_Point"] != "" {
		egressIpSfs := CDRSubfields.EgressCirIPEndPoint(cdrMap["Egress_IP_Circuit_End_Point"])
		cdrMap["EIPE_local"] = egressIpSfs["egressIPendpoint_local"]
		cdrMap["EIPE_remote"] = egressIpSfs["egressIPendpoint_remote"]
	} else {
		// fmt.Println("Accounting ID not found")
	}

	if cdrMap["Ingress_Protocol_Variant_Spec_Data"] != "" {
		ingressProtVariSpecDataSfs := CDRSubfields.IngressProtocolVariantSpecData(cdrMap["Ingress_Protocol_Variant_Spec_Data"])
		for key, value := range ingressProtVariSpecDataSfs {
			cdrMap[key] = value
		}
	} else {
		// fmt.Println("Accounting ID not found")
	}

	if cdrMap["Egress_Protocol_Variant_Spec_Data"] != "" {
		egressProtVariSpecDataSfs := CDRSubfields.EgressProtocolVariantSpecData(cdrMap["Egress_Protocol_Variant_Spec_Data"])
		for key, value := range egressProtVariSpecDataSfs {
			cdrMap[key] = value
		}
	} else {
		// fmt.Println("Accounting ID not found")
	}

	if cdrMap["Ingress_Codec_Type"] != "" {
		ingressCodecTypeSfs := CDRSubfields.IngressCodecType(cdrMap["Ingress_Codec_Type"])
		for key, value := range ingressCodecTypeSfs {
			cdrMap[key] = value
		}
	} else {
		// fmt.Println("Accounting ID not found")
	}

	if cdrMap["Egress_Codec_Type"] != "" {
		egressCodecTypeSfs := CDRSubfields.EgressCodecType(cdrMap["Egress_Codec_Type"])
		for key, value := range egressCodecTypeSfs {
			cdrMap[key] = value
		}
	} else {
		// fmt.Println("Accounting ID not found")
	}

	if cdrMap["Call_Setup_Delay"] != "" {
		callSetupDelaySfs := CDRSubfields.CallSetupDelay(cdrMap["Call_Setup_Delay"])
		for key, value := range callSetupDelaySfs {
			cdrMap[key] = value
		}
	} else {
		// fmt.Println("Accounting ID not found")
	}

	if cdrMap["Ingress_DSP_Data"] != "" {
		ingressDspDataSfs := CDRSubfields.IngressDspData(cdrMap["Ingress_DSP_Data"])
		for key, value := range ingressDspDataSfs {
			cdrMap[key] = value
		}
	} else {
		// fmt.Println("Accounting ID not found")
	}

	if cdrMap["Egress_DSP_Data"] != "" {
		egressDspDataSfs := CDRSubfields.EgressDspData(cdrMap["Egress_DSP_Data"])
		for key, value := range egressDspDataSfs {
			cdrMap[key] = value
		}
	} else {
		// fmt.Println("Accounting ID not found")
	}

	return cdrMap
}

// CdrStopKeys generates and array of cdr keys to be used in a map
// if we need to capture new fields they would need to be added here
func CdrStopKeys() []string {

	keys := []string{
		"Record_Type",
		"Gateway_Name",
		"Accounting_ID",
		"Start_Time_in_System_Ticks",
		"Node_Time_Zone",
		"Start_Date",
		"Start_Time",
		"Policy_Server_Response_Setup_time",
		"Receipt_of_Alerting_ProcProg_Setup_Time",
		"Service_Established_Setup_Time",
		"Disconnect_Date",
		"Disconnect_Time",
		"Disconnect_to_Completion_of_Call_Time",
		"Call_Service_Duration",
		"Call_Disconnect_Reason",
		"Service_Delivered",
		"Call_Direction",
		"Service_Provider",
		"Transit_Network_Selection_Code",
		"Calling_Number",
		"Called_Number",
		"Extra_Called_Number_Address_Digits",
		"Number_of_Called_Num_Translations_Done_by_This_Node",
		"Called_Number_Before_Translation_1",
		"Translation_Type_1",
		"Called_Number_Before_Translation_2",
		"Translation_Type_2",
		"Billing_Number",
		"Route_Label",
		"Route_Attempt_Number",
		"Route_Selected",
		"Egress_Local_Signaling_IP_Addr",
		"Egress_Remote_Signaling_IP_Addr",
		"Ingress_Trunk_Group_Name",
		"Ingress_PSTN_Circuit_End_Point",
		"Ingress_IP_Circuit_End_Point",
		"Egress_PSTN_Circuit_End_Point",
		"Egress_IP_Circuit_End_Point",
		"Ingress_DSP_Audio_Bytes_Sent",
		"Ingress_DSP_Audio_Packets_Sent",
		"Ingress_DSP_Audio_Bytes_Received",
		"Ingress_DSP_Audio_Packets_Received",
		"OLIP",
		"JIP",
		"Carrier_Code",
		"Call_Group_ID",
		"Script_Log_Data",
		"Time_Exit_Msg_Receipt",
		"Time_Exit_Msg_Generation",
		"Calling_Party_Nature_of_Addr",
		"Called_Party_Nature_of_Addr",
		"Ingress_Protocol_Variant_Spec_Data",
		"Ingress_Signaling_Type",
		"Egress_Signaling_Type",
		"Ingress_Far_End_Switch_Type",
		"Egress_Far_End_Switch_Type",
		"Far_End_Ingress_TG_Carrier_Code",
		"Far_End_Egress_TG_Carrier_Code",
		"Calling_Party_Category",
		"Dialed_Number",
		"Carrier_Selection_Information",
		"Called_Number_Numbering_Plan",
		"Generic_Address_Parameter",
		"Disconnect_Initiator",
		"Ingress_Number_Packets_Lost",
		"Ingress_Interarrival_Packet_Jitter",
		"Ingress_Last_Measurement_for_Latency",
		"Egress_Trunk_Group_Name",
		"Egress_Protocol_Variant_Spec_Data",
		"Incoming_Calling_Number",
		"AMA_Call_Type",
		"Message_Billing_Index",
		"Originating_LATA",
		"Route_Index_Used",
		"Calling_Party_Number_Presentation_Restriction",
		"Incoming_ISUP_Charge_Number",
		"Incoming_ISUP_Charge_Number_NOA",
		"Dialed_Number_NOA",
		"Ingress_Codec_Type",
		"Egress_Codec_Type",
		"Ingress_RTP_Packetization_Time",
		"GSX_NBS_Call_ID",
		"Originator_Echo_Cancellation",
		"Terminator_Echo_Cancellation",
		"Charge_Flag",
		"AMA_Service_Logic_Identification",
		"AMA_BAF_Module",
		"AMA_Set_Hex_AB_Indication",
		"Service_Feature_ID",
		"FE_Parameter",
		"Satellite_Indicator",
		"PSX_Billing_Information",
		"Originating_TDM_Trunk_Group_Type",
		"Terminating_TDM_Trunk_Group_Type",
		"Ingress_Trunk_Member_Number",
		"Egress_Trunk_Group_ID",
		"Egress_Switch_ID",
		"Ingress_Local_ATM_Addr",
		"Ingress_Remote_ATM_Addr",
		"Egress_Local_ATM_Addr",
		"Egress_Remote_ATM_Addr",
		"Policy_Response_Call_Type",
		"Outgoing_Route_Identification",
		"Outgoing_Message_Identification",
		"Incoming_Route_Identification",
		"Calling_Name",
		"Calling_Name_Type",
		"Incoming_Calling_Party_Numbering_Plan",
		"Outgoing_Calling_Party_Numbering_Plan",
		"Calling_Party_Business_Group_ID",
		"Called_Party_Business_Group_ID",
		"CgP_Public_Presence_Directory_Number",
		"Time_Elapsed_Setup_Msg_RX_Last_Call_Rte_Att",
		"Billing_Number_NOA",
		"Incoming_Calling_Number_NOA",
		"Egress_Trunk_Member_Number",
		"Selected_Route_Type",
		"Telcordia_Long_Duration_Record_Type",
		"Time_Elapsed_Prevous_Record",
		"Cumulative_Route_Index",
		"Call_Disconnect_Reason_TX_Ingress",
		"Call_Disconnect_Reason_TX_Egress",
		"ISDN_PRI_Calling_Party_Sub_Addr",
		"Outgoing_Trunk_Group_Number_EXM",
		"Ingress_Local_Signaling_IP_Addr",
		"Ingress_Remote_Signaling_IP_Addr",
		"Record_Sequence_Number",
		"Transmission_Medium_Requirement",
		"Information_Transfer_Rate",
		"USI_User_Information_Layer_1",
		"Unrecognized_Raw_ISUP_Calling_Party_Category",
		"Egress_RLT_Feature_Spec_Data",
		"two_B_Chan_Transfer_Feature_Spec_Data",
		"Calling_Party_Business_Unit",
		"Called_Party_Business_Unit",
		"Redirection_Feature_Spec_Data",
		"Ingress_RLT_Feature_Spec_Data",
		"PSX_Index",
		"PSX_Congestion_Level",
		"PSX_Processing_Time",
		"Script_Name",
		"Ingress_External_Accounting_Data",
		"Egress_External_Accounting_Data",
		"Egress_RTP_Packetization_Time",
		"Egress_DSP_Audio_Bytes_Sent",
		"Egress_DSP_Audio_Packets_Sent",
		"Egress_DSP_Audio_Bytes_Received",
		"Egress_DSP_Audio_Packets_Received",
		"Egress_Packets_Lost",
		"Egress_Interarrival_Packet_Jitter",
		"Egress_Last_Measurement_for_Latency",
		"Ingress_Maximum_Packet_Outage",
		"Egress_Maximum_Packet_Outage",
		"Ingress_Packet_Playout_Buffer_Quality",
		"Egress_Packet_Playout_Buffer_Quality",
		"Call_Supervision_Type",
		"Ingress_SIP_Refer_Replaces_Feature_Spec_Data",
		"Egress_SIP_Refer_Replaces_Feature_Spec_Data",
		"Network_Transfer_Feature_Spec_Data",
		"Call_Condition",
		"Toll_Indicator",
		"Gen_Num_Number",
		"Gen_Num_Presentation_Restriction_Indicator",
		"Gen_Num_Numbering_Plan",
		"Gen_Num_NOA",
		"Gen_Num_Type",
		"Originating_Trunk_Type",
		"Terminating_Trunk_Type",
		"Remote_GSX_NBS_Billing_Indicator",
		"VPN_Calling_Private_Presence_Number",
		"VPN_Calling_Public_Presence_Number",
		"External_FCI",
		"Ingress_Policing_Discards",
		"Egress_Policing_Discards",
		"Announcement_ID",
		"Source_Information",
		"Network_ID",
		"Partition_ID",
		"NCOS",
		"Ingress_SRTP",
		"Egress_SRTP",
		"ISDN_Access_Indicator_From_FCI",
		"Call_Disconnect_Location",
		"Call_Disconnect_Location_Tx_Ingress",
		"Call_Disconnect_Location_Tx_Egress",
		"Network_Call_Ref_Call_Identity",
		"Network_Call_Ref_Signaling_PC",
		"Ingress_ISUP_MIME_Protocol_Variant_Spec_Data",
		"Egress_ISUP_MIME_Protocol_Variant_Spec_Data",
		"Modem_Tone_Type",
		"Modem_Tone_Signal_Level",
		"Video_Codec_Data",
		"Video_Codec_Statistics",
		"SVS_Customer",
		"SVS_Vendor",
		"Call_To_Test_PSX",
		"PSX_Overlap_Route_Requests",
		"Call_Setup_Delay",
		"Overload_Status",
		"Ingress_BICC_Info",
		"Egress_BICC_Info",
		"Ingress_DSP_Data",
		"Egress_DSP_Data",
		"Call_Recorded_Indicator",
		"Call_Recorded_RTP_Tx_IP_Address",
		"Call_Recorded_RTP_Tx_Port_Number",
		"Call_Recorded_RTP_Rv_IP_Address",
		"Call_Recorded_RTP_Rv_Port_Number",
		"MLPP_Precedence_Level",
		"MSRP_Service_Type_Field",
		"NPUKK_Special_Routing_Information",
		"NPUKK_Customer_or_Carrier_ID",
		"NPUKK_Service_Type_Identifier",
		"NPSSP_Special_Handling_Info",
		"NPSSP_Service_Type_Identifier",
		"Total_ITX_Change_Units",
		"Global_Charge_Reference",
		"IP_Call_Limit_at_Ingress_SIP_Peer",
		"IP_Call_Limit_at_Ingress_IPTG",
		"IP_BW_Limit_at_Ingress_IPTG",
		"IP_Call_Limit_at_Egress_SIP_Peer",
		"IP_Call_Limit_at_Egress_IPTG",
		"IP_BW_Limit_at_Egress_IPTG",
		"PSX_Name",
		"No_of_PSX_Tried",
		"Ingress_Inbound_R_Factor",
		"Egress_Inbound_R_Factor",
		"Ingress_Outbound_R_Factor",
		"Egress_Outbound_R_Factor",
		"Media_Stream_Data",
		"Media_Stream_Stats",
		"Transcode_Indicator",
		"HD_Codec_Rate",
		"Remote_Ingress_Audio_RTCP_Learned_Metrics",
		"Remote_Egress_Audio_RTCP_Learned_Metrics",
		"Final_Route_Label",
		"MTA_Information",
		"VBR_Common_Billing_Data",
		"VBR_Route_Billing_Data",
	}

	return keys

}

func CdrAttemptKeys() []string {
	keys := []string{
		"Record_Type",
		"Gateway_Name",
		"Accounting_ID",
		"Start_Time",
		"Node_Time_Zone",
		"Start_Date",
		"Start_Time",
		"Policy_Server_Response_Setup_time",
		"Receipt_of_Alerting_ProcProg_Setup_Time",
		"Disconnect_Time",
		"Disconnect_to_Completion_of_Call_Time",
		"Call_Disconnect_Reason",
		"Service_Delivered",
		"Call_Direction",
		"Service_Provider",
		"Transit_Network_Selection_Code",
		"Calling_Number",
		"Called_Number",
		"Extra_Called_Number_Address_Digits",
		"Num_of_CdNum_Trans_Done_This_Node",
		"Called_Number_Before_Translation_1",
		"Translation_Type_1",
		"Called_Number_Before_Translation_2",
		"Translation_Type_2",
		"Billing_Number",
		"Route_Label",
		"Route_Attempt_Number",
		"Route_Selected",
		"Egress_Local_Signaling_IP_Addr",
		"Egress_Remote_Signaling_IP_Addr",
		"Ingress_Trunk_Group_Name",
		"Ingress_PSTN_Circuit_End_Point",
		"Ingress_IP_Circuit_End_Point",
		"Egress_PSTN_Circuit_End_Point",
		"Egress_IP_Circuit_End_Point",
		"OLIP",
		"JIP",
		"Carrier_Code",
		"Call_Group_ID",
		"Script_Log_Data",
		"Time_Exit_Msg_Receipt",
		"Time_Exit_Msg_Generation",
		"Calling_Party_Nature_of_Addr",
		"Called_Party_Nature_of_Addr",
		"Ingress_Protocol_Variant_Spec_Data",
		"Ingress_Signaling_Type",
		"Egress_Signaling_Type",
		"Ingress_Far_End_Switch_Type",
		"Egress_Far_End_Switch_Type",
		"Far_End_Ingress_TG_Carrier_Code",
		"Far_End_Egress_TG_Carrier_Code",
		"Calling_Party_Category",
		"Dialed_Number",
		"Carrier_Selection_Information",
		"Called_Number_Numbering_Plan",
		"Generic_Address_Parameter",
		"Disconnect_Initiator",
		"Egress_Trunk_Group_Name",
		"Egress_Protocol_Variant_Spec_Data",
		"Incoming_Calling_Number",
		"AMA_Call_Type",
		"Message_Billing_Index",
		"Originating_LATA",
		"Route_Index_Used",
		"CgP_Number_Presentation_Restriction",
		"Incoming_ISUP_Charge_Number",
		"Incoming_ISUP_Charge_Number_NOA",
		"Dialed_Number_NOA",
		"Ingress_Codec_Type",
		"Egress_Codec_Type",
		"Ingress_RTP_Packetization_Time",
		"GSX_NBS_Call_ID",
		"Terminated_with_Script",
		"Originator_Echo_Cancellation",
		"Terminator_Echo_Cancellation",
		"Charge_Flag",
		"AMA_Service_Logic_Identification",
		"AMA_BAF_Module",
		"AMA_Set_Hex_AB_Indication",
		"Service_Feature_ID",
		"FE_Parameter",
		"Satellite_Indicator",
		"PSX_Billing_Information",
		"Originating_TDM_Trunk_Group_Type",
		"Terminating_TDM_Trunk_Group_Type",
		"Ingress_Trunk_Member_Number",
		"Egress_Trunk_Group_ID",
		"Egress_Switch_ID",
		"Ingress_Local_ATM_Addr",
		"Ingress_Remote_ATM_Addr",
		"Egress_Local_ATM_Addr",
		"Egress_Remote_ATM_Addr",
		"Policy_Response_Call_Type",
		"Outgoing_Route_Identification",
		"Outgoing_Message_Identification",
		"Incoming_Route_Identification",
		"Calling_Name",
		"Calling_Name_Type",
		"Incoming_CgP_Numbering_Plan",
		"Outgoing_CgP_Numbering_Plan",
		"Calling_Party_Business_Group_ID",
		"Called_Party_Business_Group_ID",
		"CgP_Public_Presence_Directory_Number",
		"Time_Elapsed_Setup_Msg_RX_Last_Call_Rte_Att",
		"Disconnect_Date",
		"Billing_Number_NOA",
		"Incoming_Calling_Number_NOA",
		"Egress_Trunk_Member_Number",
		"Selected_Route_Type",
		"Cumulative_Route_Index",
		"Call_Disconnect_Reason_TX_Ingress",
		"Call_Disconnect_Reason_TX_Egress",
		"ISDN_PRI_Calling_Party_Sub_Addr",
		"Outgoing_Trunk_Group_Number_EXM",
		"Ingress_Local_Signaling_IP_Addr",
		"Ingress_Remote_Signaling_IP_Addr",
		"Record_Sequence_Number",
		"Transmission_Medium_Requirement",
		"Information_Transfer_Rate",
		"USI_User_Information_Layer_1",
		"Unrecognized_Raw_ISUP_CgP_Category",
		"Egress_RLT_Feature_Spec_Data",
		"two_B_Chan_Transfer_Feature_Spec_Data",
		"Calling_Party_Business_Unit",
		"Called_Party_Business_Unit",
		"Redirection_Feature_Spec_Data",
		"Ingress_RLT_Feature_Spec_Data",
		"PSX_Index",
		"PSX_Congestion_Level",
		"PSX_Processing_Time",
		"Script_Name",
		"Ingress_External_Accounting_Data",
		"Egress_External_Accounting_Data",
		"Egress_RTP_Packetization_Time",
		"Call_Supervision_Type",
		"Ingress_SIP_Refer_Replaces_Feature_Spec_Data",
		"Egress_SIP_Refer_Replaces_Feature_Spec_Data",
		"Network_Transfer_Feature_Spec_Data",
		"Call_Condition",
		"Toll_Indicator",
		"Gen_Num_Number",
		"Gen_Num_Presentation_Restriction_Indicator",
		"Gen_Num_Numbering_Plan",
		"Gen_Num_NOA",
		"Gen_Num_Type",
		"Final_ATT_Indicator",
		"Originating_Trunk_Type",
		"Terminating_Trunk_Type",
		"Remote_GSX_NBS_Billing_Indicator",
		"Extra_Disconnect_Reason",
		"VPN_Calling_Private_Presence_Number",
		"VPN_Calling_Public_Presence_Number",
		"External_FCI",
		"Announcement_ID",
		"Source_Information",
		"Network_ID",
		"Partition_ID",
		"NCOS",
		"ISDN_Access_Indicator_From_FCI",
		"Call_Disconnect_Location",
		"Call_Disconnect_Location_Tx_Ingress",
		"Call_Disconnect_Location_Tx_Egress",
		"Network_Call_Ref_Call_Identity",
		"Network_Call_Ref_Signaling_PC",
		"Ingress_ISUP_MIME_Protocol_Variant_Spec_Data",
		"Egress_ISUP_MIME_Protocol_Variant_Spec_Data",
		"Video_Codec_Data",
		"SVS_Customer",
		"SVS_Vendor",
		"Call_To_Test_PSX",
		"PSX_Overlap_Route_Requests",
		"Call_Setup_Delay",
		"Overload_Status",
		"Ingress_BICC_Info",
		"Egress_BICC_Info",
		"MLPP_Precedence_Level",
		"MSRP_Service_Type_Field",
		"NPUKK_Special_Routing_Information",
		"NPUKK_Customer_or_Carrier_ID",
		"NPUKK_Service_Type_Identifier",
		"NPSSP_Special_Handling_Info",
		"NPSSP_Service_Type_Identifier",
		"Oni_Str",
		"Suppress_ONI",
		"Global_Charge_Reference",
		"3xx_Contact_Information",
		"PSX_Name",
		"No_of_PSX_tried",
		"Final_Route_Label",
		"MTA_Information",
		"VBR_Common_Billing_Data",
		"VBR_Route_Billing_Data",
	}

	return keys
}

func CdrStartKeys() []string {
	keys := []string{
		"Record_Type",
		"Gateway_Name",
		"Accounting_ID",
		"Start_Time",
		"Node_Time_Zone",
		"Start_Date",
		"Start_Time",
		"Receipt_of_Time_Elapsed_PSX_Response",
		"Receipt_of_Alerting_ProcProg_Setup_Time",
		"Receipt_of_Serv_Est_Setup_Time",
		"Service_Delivered",
		"Call_Direction",
		"Service_Provider",
		"Transit_Network_Selection_Code",
		"Calling_Number",
		"Called_Number",
		"Extra_Called_Number_Address_Digits",
		"Num_of_CdNum_Trans_Done_This_Node",
		"Called_Number_Before_Translation_1",
		"Translation_Type_1",
		"Called_Number_Before_Translation_2",
		"Translation_Type_2",
		"Billing_Number",
		"Route_Label",
		"Route_Attempt_Number",
		"Route_Selected",
		"Egress_Local_Signaling_IP_Addr",
		"Egress_Remote_Signaling_IP_Addr",
		"Ingress_Trunk_Group_Name",
		"Ingress_PSTN_Circuit_End_Point",
		"Ingress_IP_Circuit_End_Point",
		"Egress_PSTN_Circuit_End_Point",
		"Egress_IP_Circuit_End_Point",
		"OLIP",
		"JIP",
		"Carrier_Code",
		"Call_Group_ID",
		"Time_Exit_Msg_Receipt",
		"Time_Exit_Msg_Generation",
		"Calling_Party_Nature_of_Addr",
		"Called_Party_Nature_of_Addr",
		"Ingress_Protocol_Variant_Spec_Data",
		"Ingress_Signaling_Type",
		"Egress_Signaling_Type",
		"Ingress_Far_End_Switch_Type",
		"Egress_Far_End_Switch_Type",
		"Far_End_Ingress_TG_Carrier_Code",
		"Far_End_Egress_TG_Carrier_Code",
		"Calling_Party_Category",
		"Dialed_Number",
		"Carrier_Selection_Information",
		"Called_Number_Numbering_Plan",
		"Generic_Address_Parameter",
		"Egress_Trunk_Group_Name",
		"Egress_Protocol_Variant_Spec_Data",
		"Incoming_Calling_Number",
		"AMA_Call_Type",
		"Message_Billing_Index",
		"Originating_LATA",
		"Route_Index_Used",
		"CgP_Number_Presentation_Restriction",
		"Incoming_ISUP_Charge_Number",
		"Incoming_ISUP_Charge_Number_NOA",
		"Dialed_Number_NOA",
		"GSX_NBS_Call_ID",
		"Charge_Flag",
		"AMA_Service_Logic_Identification",
		"AMA_BAF_Module",
		"AMA_Set_Hex_AB_Indication",
		"Service_Feature_ID",
		"FE_Parameter",
		"Satellite_Indicator",
		"PSX_Billing_Information",
		"Originating_TDM_Trunk_Group_Type",
		"Terminating_TDM_Trunk_Group_Type",
		"Ingress_Trunk_Member_Number",
		"Egress_Trunk_Group_ID",
		"Egress_Switch_ID",
		"Ingress_Local_ATM_Addr",
		"Ingress_Remote_ATM_Addr",
		"Egress_Local_ATM_Addr",
		"Egress_Remote_ATM_Addr",
		"Policy_Response_Call_Type",
		"Outgoing_Route_Identification",
		"Outgoing_Message_Identification",
		"Incoming_Route_Identification",
		"Calling_Name",
		"Calling_Name_Type",
		"Incoming_CgP_Numbering_Plan",
		"Outgoing_CgP_Numbering_Plan",
		"Calling_Party_Business_Group_ID",
		"Called_Party_Business_Group_ID",
		"CgP_Public_Presence_Directory_Number",
		"Time_Elapsed_Setup_Msg_RX_Last_Call_Rte_Att",
		"Billing_Number_NOA",
		"Incoming_Calling_Number_NOA",
		"Egress_Trunk_Member_Number",
		"Selected_Route_Type",
		"Cumulative_Route_Index",
		"ISDN_PRI_Calling_Party_Sub_Addr",
		"Outgoing_Trunk_Group_Number_EXM",
		"Ingress_Local_Signaling_IP_Addr",
		"Ingress_Remote_Signaling_IP_Addr",
		"Record_Sequence_Number",
		"Transmission_Medium_Requirement",
		"Information_Transfer_Rate",
		"USI_User_Information_Layer_1",
		"Unrecognized_Raw_ISUP_CgP_Category",
		"Egress_RLT_Feature_Spec_Data",
		"two_B_Chan_Transfer_Feature_Spec_Data",
		"Calling_Party_Business_Unit",
		"Called_Party_Business_Unit",
		"Redirection_Feature_Spec_Data",
		"Ingress_RLT_Feature_Spec_Data",
		"PSX_Index",
		"PSX_Congestion_Level",
		"PSX_Processing_Time",
		"Script_Name",
		"Ingress_External_Accounting_Data",
		"Egress_External_Accounting_Data",
		"Call_Supervision_Type",
		"Ingress_SIP_Refer_Replaces_Feature_Spec_Data",
		"Egress_SIP_Refer_Replaces_Feature_Spec_Data",
		"Network_Transfer_Feature_Spec_Data",
		"Call_Condition",
		"Toll_Indicator",
		"Gen_Num_Number",
		"Gen_Num_Presentation_Restriction_Indicator",
		"Gen_Num_Numbering_Plan",
		"Gen_Num_NOA",
		"Gen_Num_Type",
		"Originating_Trunk_Type",
		"Terminating_Trunk_Type",
		"VPN_Calling_Private_Presence_Number",
		"VPN_Calling_Public_Presence_Number",
		"External_FCI",
		"Announcement_ID",
		"Source_Information",
		"Network_ID",
		"Partition_ID",
		"NCOS",
		"ISDN_Access_Indicator_From_FCI",
		"Network_Call_Ref_Call_Identity",
		"Network_Call_Ref_Signaling_PC",
		"Ingress_ISUP_MIME_Protocol_Variant_Spec_Data",
		"Egress_ISUP_MIME_Protocol_Variant_Spec_Data",
		"Video_Codec_Data",
		"SVS_Customer",
		"SVS_Vendor",
		"Remote_GSX_NBS_Billing_Indicator",
		"Call_To_Test_PSX",
		"PSX_Overlap_Route_Requests",
		"Call_Setup_Delay",
		"Overload_Status",
		"Ingress_BICC_Info",
		"Egress_BICC_Info",
		"MLPP_Precedence_Level",
		"MSRP_Service_Type_Field",
		"NPUKK_Special_Routing_Information",
		"NPUKK_Customer_or_Carrier_ID",
		"NPUKK_Service_Type_Identifier",
		"NPSSP_Special_Handling_Info",
		"NPSSP_Service_Type_Identifier",
		"Oni_Str",
		"Suppress_ONI",
		"Global_Charge_Reference",
		"IP_Call_Limit_at_Ingress_SIP_Peer ",
		"IP_Call_Limit_at_Ingress_IPTG ",
		"IP_BW_Limit_at_Ingress_IPTG ",
		"IP_Call_Limit_at_Egress_SIP_Peer ",
		"IP_Call_Limit_at_Egress_IPTG ",
		"IP_BW_Limit_at_Egress_IPTG ",
		"3xx_Contact_Information",
		"PSX_Name",
		"No_of_PSX_tried",
		"Final_Route_Label",
		"MTA_Information",
	}

	return keys
}

// JsonCdr converts a CDR records in a map to a json string
func JsonCdr(cdrRecord map[string]string) string {
	jsondata, err := json.Marshal(cdrRecord) // convert to JSON

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(string(jsondata))
	return string(jsondata)
}

// SplitTypes takes raw parsed csv data and splits the stop records
// and attempt records in to seperate collections
func SplitTypes(values [][]string) *CdrCollection {
	c := new(CdrCollection)

	for _, value := range values {
		if value[0] == "ATTEMPT" {
			// fmt.Println(value)
			c.Attempts = append(c.Attempts, value)
		}
		if value[0] == "STOP" {
			c.Stops = append(c.Stops, value)
		}
		if value[0] == "START" {
			c.Starts = append(c.Starts, value)
		}
	}

	return c
}
