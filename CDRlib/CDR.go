/*"
* @Author: Jim Weber"
* @Date:   2015-01-28 10:09:26"
* @Last Modified by:   jpweber
* @Last Modified time: 2015-01-29 11:39:28
 */

package CDR

import (
	// "encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

// Deprecated
// type StopCDR struct {
// 	Record_Type                                         string
// 	Gateway_Name                                        string
// 	Accounting_ID                                       string
// 	Start_Time_in_System_Ticks                          int64
// 	Node_Time_Zone                                      string
// 	Start_Date                                          string
// 	Start_Time                                          string
// 	Policy_Server_Response_Setup_time                   int64
// 	Receipt_of_Alerting_ProcProg_Setup_Time             int64
// 	Service_Established_Setup_Time                      int64
// 	Disconnect_Date                                     string
// 	Disconnect_Time                                     string
// 	Disconnect_to_Completion_of_Call_Time               int64
// 	Call_Service_Duration                               int64
// 	Call_Disconnect_Reason                              int
// 	Service_Delivered                                   string
// 	Call_Direction                                      string
// 	Service_Provider                                    string
// 	Transit_Network_Selection_Code                      string
// 	Calling_Number                                      string
// 	Called_Number                                       string
// 	Extra_Called_Number_Address_Digits                  string
// 	Number_of_Called_Num_Translations_Done_by_This_Node int
// 	Called_Number_Before_Translation_1                  string
// 	Translation_Type_1                                  int
// 	Called_Number_Before_Translation_2                  string
// 	Translation_Type_2                                  int
// 	Billing_Number                                      string
// 	Route_Label                                         string
// 	Route_Attempt_Number                                int
// 	Route_Selected                                      string
// 	Egress_Local_Signaling_IP_Addr                      string
// 	Egress_Remote_Signaling_IP_Addr                     string
// 	Ingress_Trunk_Group_Name                            string
// 	Ingress_PSTN_Circuit_End_Point                      string
// 	Ingress_IP_Circuit_End_Point                        string
// 	Egress_PSTN_Circuit_End_Point                       string
// 	Egress_IP_Circuit_End_Point                         string
// 	Ingress_DSP_Audio_Bytes_Sent                        int64
// 	Ingress_DSP_Audio_Packets_Sent                      int64
// 	Ingress_DSP_Audio_Bytes_Received                    int64
// 	Ingress_DSP_Audio_Packets_Received                  int64
// 	OLIP                                                int
// 	JIP                                                 string
// 	Carrier_Code                                        string
// 	Call_Group_ID                                       string
// 	Script_Log_Data                                     string
// 	Time_Exit_Msg_Receipt                               int64
// 	Time_Exit_Msg_Generation                            int64
// 	Calling_Party_Nature_of_Addr                        int
// 	Called_Party_Nature_of_Addr                         int64
// 	Ingress_Protocol_Variant_Spec_Data                  string
// 	Ingress_Signaling_Type                              int
// 	Egress_Signaling_Type                               int
// 	Ingress_Far_End_Switch_Type                         int
// 	Egress_Far_End_Switch_Type                          int
// 	Far_End_Ingress_TG_Carrier_Code                     string
// 	Far_End_Egress_TG_Carrier_Code                      string
// 	Calling_Party_Category                              string
// 	Dialed_Number                                       string
// 	Carrier_Selection_Information                       int
// 	Called_Number_Numbering_Plan                        int
// 	Generic_Address_Parameter                           string
// 	Disconnect_Initiator                                int
// 	Ingress_Number_Packets_Lost                         int64
// 	Ingress_Interarrival_Packet_Jitter                  int
// 	Ingress_Last_Measurement_for_Latency                int
// 	Egress_Trunk_Group_Name                             string
// 	Egress_Protocol_Variant_Spec_Data                   string
// 	Incoming_Calling_Number                             string
// 	AMA_Call_Type                                       string
// 	Message_Billing_Index                               string
// 	Originating_LATA                                    string
// 	Route_Index_Used                                    int
// 	Calling_Party_Number_Presentation_Restriction       int
// 	Incoming_ISUP_Charge_Number                         string
// 	Incoming_ISUP_Charge_Number_NOA                     int
// 	Dialed_Number_NOA                                   int
// 	Ingress_Codec_Type                                  string
// 	Egress_Codec_Type                                   string
// 	Ingress_RTP_Packetization_Time                      int
// 	GSX_NBS_Call_ID                                     string
// 	Originator_Echo_Cancellation                        bool
// 	Terminator_Echo_Cancellation                        bool
// 	Charge_Flag                                         int
// 	AMA_Service_Logic_Identification                    string
// 	AMA_BAF_Module                                      string
// 	AMA_Set_Hex_AB_Indication                           bool
// 	Service_Feature_ID                                  string
// 	FE_Parameter                                        string
// 	Satellite_Indicator                                 bool
// 	PSX_Billing_Information                             string
// 	Originating_TDM_Trunk_Group_Type                    int
// 	Terminating_TDM_Trunk_Group_Type                    int
// 	Ingress_Trunk_Member_Number                         int
// 	Egress_Trunk_Group_ID                               string
// 	Egress_Switch_ID                                    string
// 	Ingress_Local_ATM_Addr                              string
// 	Ingress_Remote_ATM_Addr                             string
// 	Egress_Local_ATM_Addr                               string
// 	Egress_Remote_ATM_Addr                              string
// 	Policy_Response_Call_Type                           int
// 	Outgoing_Route_Identification                       int
// 	Outgoing_Message_Identification                     int
// 	Incoming_Route_Identification                       int
// 	Calling_Name                                        string
// 	Calling_Name_Type                                   int
// 	Incoming_Calling_Party_Numbering_Plan               int
// 	Outgoing_Calling_Party_Numbering_Plan               int
// 	Calling_Party_Business_Group_ID                     int64
// 	Called_Party_Business_Group_ID                      int64
// 	Calling_Party_Public_Presence_Directory_Number      string
// 	Time_Last_Call_Routing_Attempt                      int64
// 	Billing_Number_NOA                                  int
// 	Incoming_Calling_Number_NOA                         int
// 	Egress_Trunk_Member_Number                          int
// 	Selected_Route_Type                                 int
// 	Telcordia_Long_Duration_Record_Type                 int
// 	Time_Elapsed_Prevous_Record                         int64
// 	Cumulative_Route_Index                              int
// 	Call_Disconnect_Reason_TX_Ingress                   int
// 	Call_Disconnect_Reason_TX_Egress                    int
// 	ISDN_PRI_Calling_Party_Sub_Addr                     string
// 	Outgoing_Trunk_Group_Number_EXM                     string
// 	Ingress_Local_Signaling_IP_Addr                     string
// 	Ingress_Remote_Signaling_IP_Addr                    string
// 	Record_Sequence_Number                              int
// 	Transmission_Medium_Requirement                     int
// 	Information_Transfer_Rate                           int
// 	USI_User_Information_Layer_1                        int
// 	Unrecognized_Raw_ISUP_Calling_Party_Category        int
// 	Egress_RLT_Feature_Spec_Data                        string
// 	two_B_Chan_Transfer_Feature_Spec_Data               string
// 	Calling_Party_Business_Unit                         int64
// 	Called_Party_Business_Unit                          int64
// 	Redirection_Feature_Spec_Data                       string
// 	Ingress_RLT_Feature_Spec_Data                       string
// 	PSX_Index                                           int
// 	PSX_Congestion_Level                                int
// 	PSX_Processing_Time                                 int
// 	Script_Name                                         string
// 	Ingress_External_Accounting_Data                    string
// 	Egress_External_Accounting_Data                     string
// 	Egress_RTP_Packetization_Time                       int
// 	Egress_DSP_Audio_Bytes_Sent                         int64
// 	Egress_DSP_Audio_Packets_Sent                       int64
// 	Egress_DSP_Audio_Bytes_Received                     int64
// 	Egress_DSP_Audio_Packets_Received                   int64
// 	Egress_Packets_Lost                                 int64
// 	Egress_Interarrival_Packet_Jitter                   int
// 	Egress_Last_Measurement_for_Latency                 int64
// 	Ingress_Maximum_Packet_Outage                       int64
// 	Egress_Maximum_Packet_Outage                        int64
// 	Ingress_Packet_Playout_Buffer_Quality               string
// 	Egress_Packet_Playout_Buffer_Quality                string
// 	Call_Supervision_Type                               int
// 	Ingress_SIP_Refer_Replaces_Feature_Spec_Data        string
// 	Egress_SIP_Refer_Replaces_Feature_Spec_Data         string
// 	Network_Transfer_Feature_Spec_Data                  string
// 	Call_Condition                                      int
// 	Toll_Indicator                                      int
// 	Gen_Num_Number                                      string
// 	Gen_Num_Presentation_Restriction_Indicator          int
// 	Gen_Num_Numbering_Plan                              int
// 	Gen_Num_NOA                                         int
// 	Gen_Num_Type                                        int
// 	Originating_Trunk_Type                              int
// 	Terminating_Trunk_Type                              int
// 	Remote_GSX_NBS_Billing_Indicator                    bool
// 	VPN_Calling_Private_Presence_Number                 string
// 	VPN_Calling_Public_Presence_Number                  string
// 	External_FCI                                        string
// 	Ingress_Policing_Discards                           int64
// 	Egress_Policing_Discards                            int64
// 	Announcement_ID                                     int
// 	Source_Information                                  int
// 	Network_ID                                          int
// 	Partition_ID                                        int
// 	NCOS                                                int
// 	Ingress_SRTP                                        string
// 	Egress_SRTP                                         string
// 	ISDN_Access_Indicator_From_FCI                      int
// 	Call_Disconnect_Location                            int
// 	Call_Disconnect_Location_Tx_Ingress                 int
// 	Call_Disconnect_Location_Tx_Egress                  int
// 	Network_Call_Ref_Call_Identity                      int64
// 	Network_Call_Ref_Signaling_PC                       int
// 	Ingress_ISUP_MIME_Protocol_Variant_Spec_Data        string
// 	Egress_ISUP_MIME_Protocol_Variant_Spec_Data         string
// 	Modem_Tone_Type                                     int
// 	Modem_Tone_Signal_Level                             int
// 	Video_Codec_Data                                    string
// 	Video_Codec_Statistics                              string
// 	SVS_Customer                                        int
// 	SVS_Vendor                                          int
// 	Call_To_Test_PSX                                    bool
// 	PSX_Overlap_Route_Requests                          int
// 	Call_Setup_Delay                                    string
// 	Overload_Status                                     int
// 	Ingress_BICC_Info                                   string
// 	Egress_BICC_Info                                    string
// 	Ingress_DSP_Data                                    int
// 	Egress_DSP_Data                                     int
// 	Call_Recorded_Indicator                             string
// 	Call_Recorded_RTP_Tx_IP_Address                     string
// 	Call_Recorded_RTP_Tx_Port_Number                    int
// 	Call_Recorded_RTP_Rv_IP_Address                     string
// 	Call_Recorded_RTP_Rv_Port_Number                    int
// 	MLPP_Precedence_Level                               int
// 	MSRP_Service_Type_Field                             int
// 	NPUKK_Special_Routing_Information                   string
// 	NPUKK_Customer_or_Carrier_ID                        int
// 	NPUKK_Service_Type_Identifier                       int
// 	NPSSP_Special_Handling_Info                         int
// 	NPSSP_Service_Type_Identifier                       string
// 	Total_ITX_Change_Units                              int
// 	Global_Charge_Reference                             string
// 	IP_Call_Limit_at_Ingress_SIP_Peer                   string
// 	IP_Call_Limit_at_Ingress_IPTG                       bool
// 	IP_BW_Limit_at_Ingress_IPTG                         bool
// 	IP_Call_Limit_at_Egress_SIP_Peer                    bool
// 	IP_Call_Limit_at_Egress_IPTG                        bool
// 	IP_BW_Limit_at_Egress_IPTG                          bool
// 	PSX_Name                                            string
// 	Number_of_PSX_Tried                                 int
// 	Ingress_Inbound_R_Factor                            int
// 	Egress_Inbound_R_Factor                             int
// 	Ingress_Outbound_R_Factor                           int
// 	Egress_Outbound_R_Factor                            int
// 	Media_Stream_Data                                   string
// 	Media_Stream_Stats                                  string
// 	Transcode_Indicator                                 bool
// 	HD_Codec_Rate                                       int
// 	Remote_Ingress_Audio_RTCP_Learned_Metrics           string
// 	Remote_Egress_Audio_RTCP_Learned_Metrics            string
// 	Final_Route_Label                                   string
// 	MTA_Information                                     string
// 	VBR_Common_Billing_Data                             string
// 	VBR_Route_Billing_Data                              string
// 	//bluetone augments
// 	DT      int
// 	dsi     string
// 	rawfile string
// }

func FillCDRMap(keys []string, values []string) map[string]string {

	cdrMap := make(map[string]string)

	for i, value := range values {
		cdrMap[keys[i]] = value
	}

	return cdrMap
}

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
		"Calling_Party_Public_Presence_Directory_Number",
		"Time_Last_Call_Routing_Attempt",
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
		"Number_of_PSX_Tried",
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
		"DT",
		"dsi",
		"rawfile",
	}

	return keys

}

func JsonCdr(cdrRecord map[string]string) {
	jsondata, err := json.Marshal(cdrRecord) // convert to JSON

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsondata))
}
