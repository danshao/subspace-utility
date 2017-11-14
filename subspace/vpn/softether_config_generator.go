package vpn

import (
	"bytes"
	"fmt"
	"text/template"
)

func GenerateSoftetherConfig(softether Softether) (string, error) {
	hubConfig, err := GenerateSoftetherHubConfig(softether.Hub)
	if nil != err {
		return "", err
	}
	serverTemplate := fmt.Sprintf(TEMPLATE_SERVER, hubConfig)
	return formatTemplate("template.server", serverTemplate, softether)
}

func GenerateSoftetherHubConfig(hub Hub) (string, error) {
	var usersConfigBuffer bytes.Buffer
	for _, acc := range hub.Accounts {
		if userTemplate, err := GenerateSoftetherAccountConfig(acc); nil == err {
			usersConfigBuffer.WriteString(userTemplate)
		} else {
			return "", err
		}
	}
	usersConfig := usersConfigBuffer.String()

	var accessRulesConfigBuffer bytes.Buffer
	for _, accessRule := range hub.AccessRules {
		if accessRuleTemplate, err := GenerateSoftetherAccessRuleConfig(accessRule); nil == err {
			accessRulesConfigBuffer.WriteString(accessRuleTemplate)
		} else {
			return "", err
		}
	}
	accessRulesConfig := accessRulesConfigBuffer.String()

	hubTemplate := fmt.Sprintf(TEMPLATE_HUB, accessRulesConfig, usersConfig)
	if hubConfig, err := formatTemplate("template.hub", hubTemplate, hub); nil != err {
		return "", err
	} else {
		return hubConfig, nil
	}
}

func GenerateSoftetherAccountConfig(account Account) (string, error) {
	if userTemplate, err := formatTemplate("template.user", TEMPLATE_USER, account); nil == err {
		return userTemplate, nil
	} else {
		return "", err
	}
}

func GenerateSoftetherAccessRuleConfig(accessRule AccessRule) (string, error) {
	if accessRuleTemplate, err := formatTemplate("template.accessRule", TEMPLATE_ACCESS_RULE, accessRule); nil == err {
		return accessRuleTemplate, nil
	} else {
		return "", err
	}
}

func formatTemplate(tplName string, tplString string, data interface{}) (string, error) {
	tmpl, err := template.New("template.name").Parse(tplString)
	if nil != err {
		return "", err
	}
	var doc bytes.Buffer
	if err := tmpl.Execute(&doc, data); nil != err {
		return "", err
	} else {
		return doc.String(), nil
	}
}

const TEMPLATE_SERVER = "# Software Configuration File\n" +
	"# ---------------------------\n" +
	"#\n" +
	"# You may edit this file when the VPN Server / Client / Bridge program is not running.\n" +
	"#\n" +
	"# In prior to edit this file manually by your text editor,\n" +
	"# shutdown the VPN Server / Client / Bridge background service.\n" +
	"# Otherwise, all changes will be lost.\n" +
	"#\n" +
	"declare root\n" +
	"{\n" +
	"	uint ConfigRevision 1\n" +
	"	bool IPsecMessageDisplayed true\n" +
	"	string Region US\n" +
	"	bool VgsMessageDisplayed false\n" +
	"\n" +
	"	declare DDnsClient\n" +
	"	{\n" +
	"		bool Disabled false\n" +
	"		byte Key XW+hOCwdQgezlyygTgJfW8nDNMo=\n" +
	"		string LocalHostname ip-10-0-0-169\n" +
	"		string ProxyHostName $\n" +
	"		uint ProxyPort 0\n" +
	"		uint ProxyType 0\n" +
	"		string ProxyUsername $\n" +
	"	}\n" +
	"	declare IPsec\n" +
	"	{\n" +
	"		bool EtherIP_IPsec false\n" +
	"		string IPsec_Secret {{.PreSharedKey}}\n" +
	"		string L2TP_DefaultHub {{.GetDefaultHub}}\n" +
	"		bool L2TP_IPsec true\n" +
	"		bool L2TP_Raw false\n" +
	"\n" +
	"		declare EtherIP_IDSettingsList\n" +
	"		{\n" +
	"		}\n" +
	"	}\n" +
	"	declare ListenerList\n" +
	"	{\n" +
	"		declare Listener0\n" +
	"		{\n" +
	"			bool DisableDos true\n" +
	"			bool Enabled true\n" +
	"			uint Port {{.GetDefaultAdministrationPort}}\n" +
	"		}\n" +
	"		declare Listener1\n" +
	"		{\n" +
	"			bool DisableDos false\n" +
	"			bool Enabled true\n" +
	"			uint Port 1194\n" +
	"		}\n" +
	"		declare Listener2\n" +
	"		{\n" +
	"			bool DisableDos false\n" +
	"			bool Enabled true\n" +
	"			uint Port 5555\n" +
	"		}\n" +
	"	}\n" +
	"	declare LocalBridgeList\n" +
	"	{\n" +
	"		bool DoNotDisableOffloading false\n" +
	"	}\n" +
	"	declare ServerConfiguration\n" +
	"	{\n" +
	"		bool AcceptOnlyTls true\n" +
	"		uint64 AutoDeleteCheckDiskFreeSpaceMin 104857600\n" +
	"		uint AutoDeleteCheckIntervalSecs 300\n" +
	"		uint AutoSaveConfigSpan 300\n" +
	"		bool BackupConfigOnlyWhenModified true\n" +
	"		string CipherName RC4-MD5\n" +
	"		uint CurrentBuild 9634\n" +
	"		bool DisableCoreDumpOnUnix false\n" +
	"		bool DisableDeadLockCheck false\n" +
	"		bool DisableDosProction false\n" +
	"		bool DisableGetHostNameWhenAcceptTcp false\n" +
	"		bool DisableIntelAesAcceleration false\n" +
	"		bool DisableIPv6Listener false\n" +
	"		bool DisableNatTraversal false\n" +
	"		bool DisableOpenVPNServer false\n" +
	"		bool DisableSessionReconnect false\n" +
	"		bool DisableSSTPServer false\n" +
	"		bool DontBackupConfig false\n" +
	"		bool EnableVpnAzure false\n" +
	"		bool EnableVpnOverDns false\n" +
	"		bool EnableVpnOverIcmp false\n" +
	"		byte HashedPassword {{.GetAdminPasswordHash}}\n" + ///HF6NyMfkCemAEYz1Ykjl5uJTK0=
	"		string KeepConnectHost keepalive.softether.org\n" +
	"		uint KeepConnectInterval 50\n" +
	"		uint KeepConnectPort 80\n" +
	"		uint KeepConnectProtocol 1\n" +
	"		uint64 LoggerMaxLogSize 1073741823\n" +
	"		uint MaxConcurrentDnsClientThreads 512\n" +
	"		uint MaxConnectionsPerIP 256\n" +
	"		uint MaxUnestablishedConnections 1000\n" +
	"		bool NoHighPriorityProcess false\n" +
	"		bool NoLinuxArpFilter false\n" +
	"		bool NoSendSignature false\n" +
	"		string OpenVPNDefaultClientOption dev-type$20tun,link-mtu$201500,tun-mtu$201500,cipher$20AES-128-CBC,auth$20SHA1,keysize$20128,key-method$202,tls-client\n" +
	"		string OpenVPN_UdpPortList 1194\n" +
	"		bool SaveDebugLog false\n" +
	"		byte ServerCert MIIDvjCCAqagAwIBAgIBADANBgkqhkiG9w0BAQsFADBeMRkwFwYDVQQDDBBpcC0xNzItMzEtMzctMTQ1MRkwFwYDVQQKDBBpcC0xNzItMzEtMzctMTQ1MRkwFwYDVQQLDBBpcC0xNzItMzEtMzctMTQ1MQswCQYDVQQGEwJVUzAeFw0xNzA1MTUxOTA5MDFaFw0zNzEyMzExOTA5MDFaMF4xGTAXBgNVBAMMEGlwLTE3Mi0zMS0zNy0xNDUxGTAXBgNVBAoMEGlwLTE3Mi0zMS0zNy0xNDUxGTAXBgNVBAsMEGlwLTE3Mi0zMS0zNy0xNDUxCzAJBgNVBAYTAlVTMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsBcC8qsjE1bKoF6GWFwDHFmXgBbn2rd536i9VIH/3nze6MoE5GRI08nASsVQnYilJAxG9I6Gu6QsPmKGZ++mfsHF4kxZere6I6pwioIRm3DisN7lhOtuAo3T3sHx4EH40afrI+CasZQ+IGbJyWTzF6jE95JsRgoiLGtc2nhoeUH5JC88cejBsFlQHURLcSqEmEmQGx2ZBe4kUczftYXzzWMCEUb2gZI4MUEutQKuJbPjpsrMPCvvssIEQvH8KU+qHK9Zu1alqBM9ZCsAjEpOdhUd0FlMwclOnkGoNaLVYbasZHWHm+ACV9IlMbbb2UeutjkZlp2viwr0YgpVhX2W2QIDAQABo4GGMIGDMA8GA1UdEwEB/wQFMAMBAf8wCwYDVR0PBAQDAgH2MGMGA1UdJQRcMFoGCCsGAQUFBwMBBggrBgEFBQcDAgYIKwYBBQUHAwMGCCsGAQUFBwMEBggrBgEFBQcDBQYIKwYBBQUHAwYGCCsGAQUFBwMHBggrBgEFBQcDCAYIKwYBBQUHAwkwDQYJKoZIhvcNAQELBQADggEBAIDPwSNQMCkqoDpmVW5sQOyqNNrHvZDC4mhnviI459gSOtwjel9bKC79lHwSLBtcoinnlNflkPV/UN1tbEa7EU8T8+cZvq0sBrnGjgTRF91MfL9ti5AHJqylgIVORdZKgCqBc2qvhU7OMFyfPKfcwlKQTd0NvYvOsPZMKlGwo6+7G2EJNf/D0Uq7d/7Ja+W8+ZiYCB++kF4XPyWtvyzxvDWVjL8nMm1qKR+mxxr2cUbe4Y/RGjWHHvv3hmtFqgHNOlEmk3QXBglqSDX9hWsTXdyzMQ6oZwCTA6I2Jazy62dYTt/JuevJw8TRBAwXFkvEW/c8UHTr1CO986aZEXWlNGU=\n" +
	"		byte ServerKey MIIEpAIBAAKCAQEAsBcC8qsjE1bKoF6GWFwDHFmXgBbn2rd536i9VIH/3nze6MoE5GRI08nASsVQnYilJAxG9I6Gu6QsPmKGZ++mfsHF4kxZere6I6pwioIRm3DisN7lhOtuAo3T3sHx4EH40afrI+CasZQ+IGbJyWTzF6jE95JsRgoiLGtc2nhoeUH5JC88cejBsFlQHURLcSqEmEmQGx2ZBe4kUczftYXzzWMCEUb2gZI4MUEutQKuJbPjpsrMPCvvssIEQvH8KU+qHK9Zu1alqBM9ZCsAjEpOdhUd0FlMwclOnkGoNaLVYbasZHWHm+ACV9IlMbbb2UeutjkZlp2viwr0YgpVhX2W2QIDAQABAoIBAALvz0/GxOX8q5JzISB2IGMnp1wuhZ1jM/gj80xmgqLuSDVvsRcyCPjnQ3q7wbVu4vov7jsD6YFVoptDv2QaXhd9GdTFWzStvb/8hrX4p1yFTHKQtRk7Z/FI9kCf0W3hC8pLV6OYilvt0f9uV20xbtB1MxvB6GZeouVEwR6vpWZoSSfZgxtqoZNhylDhtGLth2ucrC4sGLy49VYZc7MhP0GNqbkMnGEd2aVY8zagSIVyMcgp+0J6qCnCyZ0xsK9bVp/jdm3Blx4qI/gp8+6HL7VP1V9AnUhW91cVWjJ+w9Y27Xw/bgO2rgGIHtA5Cvbt6KeLplKqAxIS+31xqDXl5RECgYEA6KVAq7GsEIk+ajvy/9zJLOR/Y9RU+DLx8ki72335KbHQF9Lk1cQqhTruOS0B8E9YgFdtx6P2+nbDPtCXiwywrizzGK8As9rh6EMMCB/bMDLgKx8K5GFQYmG7ouX5h6ZFEufBh+yAFFrmpyU7Cd7CRbpItjHrr+wDGB3Ket/mSu0CgYEAwcRWkXhkvrv7WW0kB0Yr3ACXZXZD2RQrQmrI7KiwiM7E+REZ43BLLh1MgP12wLYZ/N92cs2gyCqstQ5FY/wPLcWORepeo91LEKqgAIIwSuGlNlyFysJRLkYyzsYvMs1MqdddgtA8oOF8oGX9EIVxIXD6BaH7Xu17MpfUqcyuQh0CgYAEqwQwLcF5F3fkI83Nan2y4+bWdvU5gbQvG1yBm/Z66rF1OOsUivobcEJFbWzDiKQK9zYhsl2LSwTB5ueTh66n4AkErsZB1ZuA0I+WPiVkwzCSTO4oBPfa8YEVVK38Fc7/AR11/73WKrjEE9Aqc4jMY1+AIGJNRlwNiNN+Qa3aRQKBgQC8dbg1GMHz+WYInmySXp++BR07ZAGtmz5QBQiTlabOCT3vecCSQ9/7ZCfKtbvQx80S9E2Cb5lX1pnQJ6c6Dii9Pg6Y1dFi05N1DF8+32EVE87axraidMj8lu5Hyt6RLLjJ4FxlJSAy4d2TLF8suH1GJ1omLFwPAi/3D1GaLSe3yQKBgQCPlY+k07duoeuYx9Irl3JTyRTa6woWkHUyngHoDVftiWzxYBvMJ2NB1QHCUobp0TcoKxDntztbmcTNSWdHRNHSYik4htAKX4dvCbA/jmvKXbzAkzfk60tkeY2Paqp6WvSwoOjXIrFOvJohDGbZK98JxB6Fff6Epx95L8GtPrFo0g==\n" +
	"		uint ServerLogSwitchType 4\n" +
	"		uint ServerType 0\n" +
	"		bool Tls_Disable1_0 false\n" +
	"		bool Tls_Disable1_1 false\n" +
	"		bool Tls_Disable1_2 false\n" +
	"		bool UseKeepConnect true\n" +
	"		bool UseWebTimePage false\n" +
	"		bool UseWebUI false\n" +
	"\n" +
	"		declare GlobalParams\n" +
	"		{\n" +
	"			uint FIFO_BUDGET 10240000\n" +
	"			uint HUB_ARP_SEND_INTERVAL 5000\n" +
	"			uint IP_TABLE_EXPIRE_TIME 60000\n" +
	"			uint IP_TABLE_EXPIRE_TIME_DHCP 300000\n" +
	"			uint MAC_TABLE_EXPIRE_TIME 600000\n" +
	"			uint MAX_BUFFERING_PACKET_SIZE 2560000\n" +
	"			uint MAX_HUB_LINKS 1024\n" +
	"			uint MAX_IP_TABLES 65536\n" +
	"			uint MAX_MAC_TABLES 65536\n" +
	"			uint MAX_SEND_SOCKET_QUEUE_NUM 128\n" +
	"			uint MAX_SEND_SOCKET_QUEUE_SIZE 2560000\n" +
	"			uint MAX_STORED_QUEUE_NUM 1024\n" +
	"			uint MEM_FIFO_REALLOC_MEM_SIZE 655360\n" +
	"			uint MIN_SEND_SOCKET_QUEUE_SIZE 320000\n" +
	"			uint QUEUE_BUDGET 2048\n" +
	"			uint SELECT_TIME 256\n" +
	"			uint SELECT_TIME_FOR_NAT 30\n" +
	"			uint STORM_CHECK_SPAN 500\n" +
	"			uint STORM_DISCARD_VALUE_END 1024\n" +
	"			uint STORM_DISCARD_VALUE_START 3\n" +
	"		}\n" +
	"		declare ServerTraffic\n" +
	"		{\n" +
	"			declare RecvTraffic\n" +
	"			{\n" +
	"				uint64 BroadcastBytes 0\n" +
	"				uint64 BroadcastCount 0\n" +
	"				uint64 UnicastBytes 0\n" +
	"				uint64 UnicastCount 0\n" +
	"			}\n" +
	"			declare SendTraffic\n" +
	"			{\n" +
	"				uint64 BroadcastBytes 0\n" +
	"				uint64 BroadcastCount 0\n" +
	"				uint64 UnicastBytes 0\n" +
	"				uint64 UnicastCount 0\n" +
	"			}\n" +
	"		}\n" +
	"		declare SyslogSettings\n" +
	"		{\n" +
	"			string HostName $\n" +
	"			uint Port 0\n" +
	"			uint SaveType 0\n" +
	"		}\n" +
	"	}\n" +
	"	declare VirtualHUB\n" +
	"	{\n" +
	"		declare DEFAULT\n" +
	"		{\n" +
	"			uint64 CreatedTime 1494842941898\n" +
	"			byte HashedPassword +WzqGYrR3VYXrAhKPZLGEHcIwO8=\n" +
	"			uint64 LastCommTime 1494842941893\n" +
	"			uint64 LastLoginTime 1494842941893\n" +
	"			uint NumLogin 0\n" +
	"			bool Online true\n" +
	"			bool RadiusConvertAllMsChapv2AuthRequestToEap false\n" +
	"			string RadiusRealm $\n" +
	"			uint RadiusRetryInterval 0\n" +
	"			uint RadiusServerPort 1812\n" +
	"			string RadiusSuffixFilter $\n" +
	"			bool RadiusUsePeapInsteadOfEap false\n" +
	"			byte SecurePassword bpw3X/O5E8a6G6ccnl4uXmDtkwI=\n" +
	"			uint Type 0\n" +
	"\n" +
	"			declare AccessList\n" +
	"			{\n" +
	"			}\n" +
	"			declare AdminOption\n" +
	"			{\n" +
	"				uint allow_hub_admin_change_option 0\n" +
	"				uint deny_bridge 0\n" +
	"				uint deny_change_user_password 0\n" +
	"				uint deny_empty_password 0\n" +
	"				uint deny_hub_admin_change_ext_option 0\n" +
	"				uint deny_qos 0\n" +
	"				uint deny_routing 0\n" +
	"				uint max_accesslists 0\n" +
	"				uint max_bitrates_download 0\n" +
	"				uint max_bitrates_upload 0\n" +
	"				uint max_groups 0\n" +
	"				uint max_multilogins_per_user 0\n" +
	"				uint max_sessions 0\n" +
	"				uint max_sessions_bridge 0\n" +
	"				uint max_sessions_client 0\n" +
	"				uint max_sessions_client_bridge_apply 0\n" +
	"				uint max_users 0\n" +
	"				uint no_access_list_include_file 0\n" +
	"				uint no_cascade 0\n" +
	"				uint no_change_access_control_list 0\n" +
	"				uint no_change_access_list 0\n" +
	"				uint no_change_admin_password 0\n" +
	"				uint no_change_cert_list 0\n" +
	"				uint no_change_crl_list 0\n" +
	"				uint no_change_groups 0\n" +
	"				uint no_change_log_config 0\n" +
	"				uint no_change_log_switch_type 0\n" +
	"				uint no_change_msg 0\n" +
	"				uint no_change_users 0\n" +
	"				uint no_delay_jitter_packet_loss 0\n" +
	"				uint no_delete_iptable 0\n" +
	"				uint no_delete_mactable 0\n" +
	"				uint no_disconnect_session 0\n" +
	"				uint no_enum_session 0\n" +
	"				uint no_offline 0\n" +
	"				uint no_online 0\n" +
	"				uint no_query_session 0\n" +
	"				uint no_read_log_file 0\n" +
	"				uint no_securenat 0\n" +
	"				uint no_securenat_enabledhcp 0\n" +
	"				uint no_securenat_enablenat 0\n" +
	"			}\n" +
	"			declare CascadeList\n" +
	"			{\n" +
	"			}\n" +
	"			declare LogSetting\n" +
	"			{\n" +
	"				uint PacketLogSwitchType 4\n" +
	"				uint PACKET_LOG_ARP 0\n" +
	"				uint PACKET_LOG_DHCP 1\n" +
	"				uint PACKET_LOG_ETHERNET 0\n" +
	"				uint PACKET_LOG_ICMP 0\n" +
	"				uint PACKET_LOG_IP 0\n" +
	"				uint PACKET_LOG_TCP 0\n" +
	"				uint PACKET_LOG_TCP_CONN 1\n" +
	"				uint PACKET_LOG_UDP 0\n" +
	"				bool SavePacketLog true\n" +
	"				bool SaveSecurityLog true\n" +
	"				uint SecurityLogSwitchType 4\n" +
	"			}\n" +
	"			declare Message\n" +
	"			{\n" +
	"			}\n" +
	"			declare Option\n" +
	"			{\n" +
	"				uint AccessListIncludeFileCacheLifetime 30\n" +
	"				uint AdjustTcpMssValue 0\n" +
	"				bool ApplyIPv4AccessListOnArpPacket false\n" +
	"				bool AssignVLanIdByRadiusAttribute false\n" +
	"				bool BroadcastLimiterStrictMode false\n" +
	"				uint BroadcastStormDetectionThreshold 0\n" +
	"				uint ClientMinimumRequiredBuild 0\n" +
	"				bool DenyAllRadiusLoginWithNoVlanAssign false\n" +
	"				uint DetectDormantSessionInterval 0\n" +
	"				bool DisableAdjustTcpMss false\n" +
	"				bool DisableCheckMacOnLocalBridge false\n" +
	"				bool DisableCorrectIpOffloadChecksum false\n" +
	"				bool DisableHttpParsing false\n" +
	"				bool DisableIPParsing false\n" +
	"				bool DisableIpRawModeSecureNAT false\n" +
	"				bool DisableKernelModeSecureNAT false\n" +
	"				bool DisableUdpAcceleration false\n" +
	"				bool DisableUdpFilterForLocalBridgeNic false\n" +
	"				bool DisableUserModeSecureNAT false\n" +
	"				bool DoNotSaveHeavySecurityLogs false\n" +
	"				bool DropArpInPrivacyFilterMode true\n" +
	"				bool DropBroadcastsInPrivacyFilterMode true\n" +
	"				bool FilterBPDU false\n" +
	"				bool FilterIPv4 false\n" +
	"				bool FilterIPv6 false\n" +
	"				bool FilterNonIP false\n" +
	"				bool FilterOSPF false\n" +
	"				bool FilterPPPoE false\n" +
	"				uint FloodingSendQueueBufferQuota 33554432\n" +
	"				bool ManageOnlyLocalUnicastIPv6 true\n" +
	"				bool ManageOnlyPrivateIP true\n" +
	"				uint MaxLoggedPacketsPerMinute 0\n" +
	"				uint MaxSession 0\n" +
	"				bool NoArpPolling false\n" +
	"				bool NoDhcpPacketLogOutsideHub true\n" +
	"				bool NoEnum false\n" +
	"				bool NoIpTable false\n" +
	"				bool NoIPv4PacketLog false\n" +
	"				bool NoIPv6AddrPolling false\n" +
	"				bool NoIPv6DefaultRouterInRAWhenIPv6 true\n" +
	"				bool NoIPv6PacketLog false\n" +
	"				bool NoLookBPDUBridgeId false\n" +
	"				bool NoMacAddressLog true\n" +
	"				bool NoManageVlanId false\n" +
	"				bool NoPhysicalIPOnPacketLog false\n" +
	"				bool NoSpinLockForPacketDelay false\n" +
	"				bool RemoveDefGwOnDhcpForLocalhost true\n" +
	"				uint RequiredClientId 0\n" +
	"				uint SecureNAT_MaxDnsSessionsPerIp 0\n" +
	"				uint SecureNAT_MaxIcmpSessionsPerIp 0\n" +
	"				uint SecureNAT_MaxTcpSessionsPerIp 0\n" +
	"				uint SecureNAT_MaxTcpSynSentPerIp 0\n" +
	"				uint SecureNAT_MaxUdpSessionsPerIp 0\n" +
	"				bool SecureNAT_RandomizeAssignIp false\n" +
	"				bool SuppressClientUpdateNotification false\n" +
	"				bool UseHubNameAsDhcpUserClassOption false\n" +
	"				bool UseHubNameAsRadiusNasId false\n" +
	"				string VlanTypeId 0x8100\n" +
	"				bool YieldAfterStorePacket false\n" +
	"			}\n" +
	"			declare SecureNAT\n" +
	"			{\n" +
	"				bool Disabled true\n" +
	"				bool SaveLog true\n" +
	"\n" +
	"				declare VirtualDhcpServer\n" +
	"				{\n" +
	"					string DhcpDnsServerAddress 192.168.30.1\n" +
	"					string DhcpDnsServerAddress2 0.0.0.0\n" +
	"					string DhcpDomainName $\n" +
	"					bool DhcpEnabled true\n" +
	"					uint DhcpExpireTimeSpan 7200\n" +
	"					string DhcpGatewayAddress 192.168.30.1\n" +
	"					string DhcpLeaseIPEnd 192.168.30.200\n" +
	"					string DhcpLeaseIPStart 192.168.30.10\n" +
	"					string DhcpPushRoutes $\n" +
	"					string DhcpSubnetMask 255.255.255.0\n" +
	"				}\n" +
	"				declare VirtualHost\n" +
	"				{\n" +
	"					string VirtualHostIp 192.168.30.1\n" +
	"					string VirtualHostIpSubnetMask 255.255.255.0\n" +
	"					string VirtualHostMacAddress 00-AC-5A-21-3E-20\n" +
	"				}\n" +
	"				declare VirtualRouter\n" +
	"				{\n" +
	"					bool NatEnabled true\n" +
	"					uint NatMtu 1500\n" +
	"					uint NatTcpTimeout 1800\n" +
	"					uint NatUdpTimeout 60\n" +
	"				}\n" +
	"			}\n" +
	"			declare SecurityAccountDatabase\n" +
	"			{\n" +
	"				declare CertList\n" +
	"				{\n" +
	"				}\n" +
	"				declare CrlList\n" +
	"				{\n" +
	"				}\n" +
	"				declare GroupList\n" +
	"				{\n" +
	"				}\n" +
	"				declare IPAccessControlList\n" +
	"				{\n" +
	"				}\n" +
	"				declare UserList\n" +
	"				{\n" +
	"				}\n" +
	"			}\n" +
	"			declare Traffic\n" +
	"			{\n" +
	"				declare RecvTraffic\n" +
	"				{\n" +
	"					uint64 BroadcastBytes 0\n" +
	"					uint64 BroadcastCount 0\n" +
	"					uint64 UnicastBytes 0\n" +
	"					uint64 UnicastCount 0\n" +
	"				}\n" +
	"				declare SendTraffic\n" +
	"				{\n" +
	"					uint64 BroadcastBytes 0\n" +
	"					uint64 BroadcastCount 0\n" +
	"					uint64 UnicastBytes 0\n" +
	"					uint64 UnicastCount 0\n" +
	"				}\n" +
	"			}\n" +
	"		}\n" +
	"%s" + // Hub subspace here
	"	}\n" +
	"	declare VirtualLayer3SwitchList\n" +
	"	{\n" +
	"	}\n" +
	"}\n"

const TEMPLATE_HUB = "		declare {{.Name}}\n" +
	"		{\n" +
	"			uint64 CreatedTime 0\n" +
	"			byte HashedPassword /HF6NyMfkCemAEYz1Ykjl5uJTK0=\n" +
	"			uint64 LastCommTime 0\n" +
	"			uint64 LastLoginTime 0\n" +
	"			uint NumLogin 0\n" +
	"			bool Online true\n" +
	"			bool RadiusConvertAllMsChapv2AuthRequestToEap false\n" +
	"			string RadiusRealm $\n" +
	"			uint RadiusRetryInterval 0\n" +
	"			uint RadiusServerPort 1812\n" +
	"			string RadiusSuffixFilter $\n" +
	"			bool RadiusUsePeapInsteadOfEap false\n" +
	"			byte SecurePassword AAAAAAAAAAAAAAAAAAAAAAAAAAA=\n" +
	"			uint Type 0\n" +
	"\n" +
	"			declare AccessList\n" +
	"			{\n" +
	"%s" + // AccessRules here
	"			}\n" +
	"			declare AdminOption\n" +
	"			{\n" +
	"				uint allow_hub_admin_change_option 0\n" +
	"				uint deny_bridge 0\n" +
	"				uint deny_change_user_password 0\n" +
	"				uint deny_empty_password 0\n" +
	"				uint deny_hub_admin_change_ext_option 0\n" +
	"				uint deny_qos 0\n" +
	"				uint deny_routing 0\n" +
	"				uint max_accesslists 0\n" +
	"				uint max_bitrates_download 0\n" +
	"				uint max_bitrates_upload 0\n" +
	"				uint max_groups 0\n" +
	"				uint max_multilogins_per_user 0\n" +
	"				uint max_sessions 0\n" +
	"				uint max_sessions_bridge 0\n" +
	"				uint max_sessions_client 0\n" +
	"				uint max_sessions_client_bridge_apply 0\n" +
	"				uint max_users 0\n" +
	"				uint no_access_list_include_file 0\n" +
	"				uint no_cascade 0\n" +
	"				uint no_change_access_control_list 0\n" +
	"				uint no_change_access_list 0\n" +
	"				uint no_change_admin_password 0\n" +
	"				uint no_change_cert_list 0\n" +
	"				uint no_change_crl_list 0\n" +
	"				uint no_change_groups 0\n" +
	"				uint no_change_log_config 0\n" +
	"				uint no_change_log_switch_type 0\n" +
	"				uint no_change_msg 0\n" +
	"				uint no_change_users 0\n" +
	"				uint no_delay_jitter_packet_loss 0\n" +
	"				uint no_delete_iptable 0\n" +
	"				uint no_delete_mactable 0\n" +
	"				uint no_disconnect_session 0\n" +
	"				uint no_enum_session 0\n" +
	"				uint no_offline 0\n" +
	"				uint no_online 0\n" +
	"				uint no_query_session 0\n" +
	"				uint no_read_log_file 0\n" +
	"				uint no_securenat 0\n" +
	"				uint no_securenat_enabledhcp 0\n" +
	"				uint no_securenat_enablenat 0\n" +
	"			}\n" +
	"			declare CascadeList\n" +
	"			{\n" +
	"			}\n" +
	"			declare LogSetting\n" +
	"			{\n" +
	"				uint PacketLogSwitchType 4\n" +
	"				uint PACKET_LOG_ARP 0\n" +
	"				uint PACKET_LOG_DHCP 1\n" +
	"				uint PACKET_LOG_ETHERNET 0\n" +
	"				uint PACKET_LOG_ICMP 0\n" +
	"				uint PACKET_LOG_IP 0\n" +
	"				uint PACKET_LOG_TCP 0\n" +
	"				uint PACKET_LOG_TCP_CONN 1\n" +
	"				uint PACKET_LOG_UDP 0\n" +
	"				bool SavePacketLog true\n" +
	"				bool SaveSecurityLog true\n" +
	"				uint SecurityLogSwitchType 4\n" +
	"			}\n" +
	"			declare Message\n" +
	"			{\n" +
	"			}\n" +
	"			declare Option\n" +
	"			{\n" +
	"				uint AccessListIncludeFileCacheLifetime 30\n" +
	"				uint AdjustTcpMssValue 0\n" +
	"				bool ApplyIPv4AccessListOnArpPacket false\n" +
	"				bool AssignVLanIdByRadiusAttribute false\n" +
	"				bool BroadcastLimiterStrictMode false\n" +
	"				uint BroadcastStormDetectionThreshold 0\n" +
	"				uint ClientMinimumRequiredBuild 0\n" +
	"				bool DenyAllRadiusLoginWithNoVlanAssign false\n" +
	"				uint DetectDormantSessionInterval 0\n" +
	"				bool DisableAdjustTcpMss false\n" +
	"				bool DisableCheckMacOnLocalBridge false\n" +
	"				bool DisableCorrectIpOffloadChecksum false\n" +
	"				bool DisableHttpParsing false\n" +
	"				bool DisableIPParsing false\n" +
	"				bool DisableIpRawModeSecureNAT false\n" +
	"				bool DisableKernelModeSecureNAT false\n" +
	"				bool DisableUdpAcceleration false\n" +
	"				bool DisableUdpFilterForLocalBridgeNic false\n" +
	"				bool DisableUserModeSecureNAT false\n" +
	"				bool DoNotSaveHeavySecurityLogs false\n" +
	"				bool DropArpInPrivacyFilterMode true\n" +
	"				bool DropBroadcastsInPrivacyFilterMode true\n" +
	"				bool FilterBPDU false\n" +
	"				bool FilterIPv4 false\n" +
	"				bool FilterIPv6 false\n" +
	"				bool FilterNonIP false\n" +
	"				bool FilterOSPF false\n" +
	"				bool FilterPPPoE false\n" +
	"				uint FloodingSendQueueBufferQuota 33554432\n" +
	"				bool ManageOnlyLocalUnicastIPv6 true\n" +
	"				bool ManageOnlyPrivateIP true\n" +
	"				uint MaxLoggedPacketsPerMinute 0\n" +
	"				uint MaxSession 0\n" +
	"				bool NoArpPolling false\n" +
	"				bool NoDhcpPacketLogOutsideHub true\n" +
	"				bool NoEnum false\n" +
	"				bool NoIpTable false\n" +
	"				bool NoIPv4PacketLog false\n" +
	"				bool NoIPv6AddrPolling false\n" +
	"				bool NoIPv6DefaultRouterInRAWhenIPv6 true\n" +
	"				bool NoIPv6PacketLog false\n" +
	"				bool NoLookBPDUBridgeId false\n" +
	"				bool NoMacAddressLog true\n" +
	"				bool NoManageVlanId false\n" +
	"				bool NoPhysicalIPOnPacketLog false\n" +
	"				bool NoSpinLockForPacketDelay false\n" +
	"				bool RemoveDefGwOnDhcpForLocalhost true\n" +
	"				uint RequiredClientId 0\n" +
	"				uint SecureNAT_MaxDnsSessionsPerIp 0\n" +
	"				uint SecureNAT_MaxIcmpSessionsPerIp 0\n" +
	"				uint SecureNAT_MaxTcpSessionsPerIp 0\n" +
	"				uint SecureNAT_MaxTcpSynSentPerIp 0\n" +
	"				uint SecureNAT_MaxUdpSessionsPerIp 0\n" +
	"				bool SecureNAT_RandomizeAssignIp false\n" +
	"				bool SuppressClientUpdateNotification false\n" +
	"				bool UseHubNameAsDhcpUserClassOption false\n" +
	"				bool UseHubNameAsRadiusNasId false\n" +
	"				string VlanTypeId 0x8100\n" +
	"				bool YieldAfterStorePacket false\n" +
	"			}\n" +
	"			declare SecureNAT\n" +
	"			{\n" +
	"				bool Disabled false\n" +
	"				bool SaveLog true\n" +
	"\n" +
	"				declare VirtualDhcpServer\n" +
	"				{\n" +
	"					string DhcpDnsServerAddress 192.168.30.1\n" +
	"					string DhcpDnsServerAddress2 0.0.0.0\n" +
	"					string DhcpDomainName $\n" +
	"					bool DhcpEnabled true\n" +
	"					uint DhcpExpireTimeSpan 7200\n" +
	"					string DhcpGatewayAddress 192.168.30.1\n" +
	"					string DhcpLeaseIPEnd 192.168.30.200\n" +
	"					string DhcpLeaseIPStart 192.168.30.10\n" +
	"					string DhcpPushRoutes $\n" +
	"					string DhcpSubnetMask 255.255.255.0\n" +
	"				}\n" +
	"				declare VirtualHost\n" +
	"				{\n" +
	"					string VirtualHostIp 192.168.30.1\n" +
	"					string VirtualHostIpSubnetMask 255.255.255.0\n" +
	"					string VirtualHostMacAddress 00-AC-7F-26-60-3F\n" +
	"				}\n" +
	"				declare VirtualRouter\n" +
	"				{\n" +
	"					bool NatEnabled true\n" +
	"					uint NatMtu 1500\n" +
	"					uint NatTcpTimeout 1800\n" +
	"					uint NatUdpTimeout 60\n" +
	"				}\n" +
	"			}\n" +
	"			declare SecurityAccountDatabase\n" +
	"			{\n" +
	"				declare CertList\n" +
	"				{\n" +
	"				}\n" +
	"				declare CrlList\n" +
	"				{\n" +
	"				}\n" +
	"				declare GroupList\n" +
	"				{\n" +
	"				}\n" +
	"				declare IPAccessControlList\n" +
	"				{\n" +
	"				}\n" +
	"				declare UserList\n" +
	"				{\n" +
	"%s" + // Users here
	"				}\n" +
	"			}\n" +
	"			declare Traffic\n" +
	"			{\n" +
	"				declare RecvTraffic\n" +
	"				{\n" +
	"					uint64 BroadcastBytes 0\n" +
	"					uint64 BroadcastCount 0\n" +
	"					uint64 UnicastBytes 0\n" +
	"					uint64 UnicastCount 0\n" +
	"				}\n" +
	"				declare SendTraffic\n" +
	"				{\n" +
	"					uint64 BroadcastBytes 0\n" +
	"					uint64 BroadcastCount 0\n" +
	"					uint64 UnicastBytes 0\n" +
	"					uint64 UnicastCount 0\n" +
	"				}\n" +
	"			}\n" +
	"		}\n"

const TEMPLATE_USER = "					declare {{.Username}}\n" +
	"					{\n" +
	"						byte AuthNtLmSecureHash {{.GetNtLmSecureHash}}\n" +
	"						byte AuthPassword {{.GetPasswordHash}}\n" +
	"						uint AuthType 1\n" +
	"						uint64 CreatedTime {{.GetCreatedTimeInMilliseconds}}\n" +
	"						uint64 ExpireTime {{.GetExpireTime}}\n" +
	"						uint64 LastLoginTime {{.GetLastLoginTime}}\n" +
	"						string Note {{.GetNote}}\n" +
	"						uint NumLogin {{.LoginCount}}\n" +
	"						string RealName {{.GetRealName}}\n" +
	"						uint64 UpdatedTime {{.GetUpdatedTime}}\n" +
	"\n" +
	"						declare Traffic\n" +
	"						{\n" +
	"							declare RecvTraffic\n" +
	"							{\n" +
	"								uint64 BroadcastBytes 0\n" +
	"								uint64 BroadcastCount 0\n" +
	"								uint64 UnicastBytes 0\n" +
	"								uint64 UnicastCount 0\n" +
	"							}\n" +
	"							declare SendTraffic\n" +
	"							{\n" +
	"								uint64 BroadcastBytes 0\n" +
	"								uint64 BroadcastCount 0\n" +
	"								uint64 UnicastBytes 0\n" +
	"								uint64 UnicastCount 0\n" +
	"							}\n" +
	"						}\n" +
	"					}\n"

const TEMPLATE_ACCESS_RULE = "				declare {{.Index}}\n" +
	"				{\n" +
	"					bool Active true\n" +
	"					bool CheckDstMac false\n" +
	"					bool CheckSrcMac false\n" +
	"					bool CheckTcpState false\n" +
	"					uint Delay 0\n" +
	"					string DestIpAddress {{.DestIpAddress}}\n" +
	"					uint DestPortEnd 0\n" +
	"					uint DestPortStart 0\n" +
	"					string DestSubnetMask {{.DestSubnetMask}}\n" +
	"					string DestUsername $\n" +
	"					bool Discard {{.Discard}}\n" +
	"					bool Established false\n" +
	"					bool IsIPv6 false\n" +
	"					uint Jitter 0\n" +
	"					uint Loss 0\n" +
	"					string Note {{.Note}}\n" +
	"					uint Priority {{.Priority}}\n" +
	"					uint Protocol 0\n" +
	"					string RedirectUrl $\n" +
	"					string SrcIpAddress 0.0.0.0\n" +
	"					uint SrcPortEnd 0\n" +
	"					uint SrcPortStart 0\n" +
	"					string SrcSubnetMask 0.0.0.0\n" +
	"					string SrcUsername {{.SrcUsername}}\n" +
	"				}\n"
