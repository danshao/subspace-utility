package main

import (
	"fmt"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/vpn"
	"time"
)

var account = vpn.Account{
	Username: "1_1495004926659038074",
	Password: "SFi5yU4lSh",
	Email: "carterlin@ecoworkinc.com",
	Description: "Default Admin VPN Profile",
	LoginCount: 1,
	ExpireTime: time.Unix(0, 0 * int64(time.Millisecond)),
	LastLoginTime: time.Unix(0, 1495665787182 * int64(time.Millisecond)),
	CreatedTime: time.Unix(0, 1494972526771 * int64(time.Millisecond)),
	UpdatedTime: time.Unix(0, 1494972526908 * int64(time.Millisecond)),
}

var hub = vpn.Hub{
	Name: "subspace",
	Accounts: []vpn.Account{account},
}

var server = vpn.Softether{
	Hub: hub,
	AdminPassword: "subspace",
	PreSharedKey: "subspace",
}

func ExampleGenerateSoftetherAccountConfig() {
	result, _ := vpn.GenerateSoftetherAccountConfig(account)
	fmt.Println(result)
	// Output:
	// 					declare 1_1495004926659038074
	// 					{
	// 						byte AuthNtLmSecureHash tmDCpoiYTNpWw/lQ4N3soA==
	// 						byte AuthPassword 8Tbpu20Wtd+56zbOyFmoInDE60k=
	// 						uint AuthType 1
	// 						uint64 CreatedTime 1494972526771
	// 						uint64 ExpireTime 0
	// 						uint64 LastLoginTime 1495665787182
	// 						string Note Default$20Admin$20VPN$20Profile
	// 						uint NumLogin 1
	// 						string RealName carterlin@ecoworkinc.com
	// 						uint64 UpdatedTime 1494972526908
	//
	// 						declare Traffic
	// 						{
	// 							declare RecvTraffic
	// 							{
	// 								uint64 BroadcastBytes 0
	// 								uint64 BroadcastCount 0
	// 								uint64 UnicastBytes 0
	// 								uint64 UnicastCount 0
	// 							}
	// 							declare SendTraffic
	// 							{
	// 								uint64 BroadcastBytes 0
	// 								uint64 BroadcastCount 0
	// 								uint64 UnicastBytes 0
	// 								uint64 UnicastCount 0
	// 							}
	// 						}
	// 					}
	//
}

func ExampleGenerateSoftetherHubConfig() {

	result, _ := vpn.GenerateSoftetherHubConfig(hub)
	fmt.Println(result)
	// Output:
	// 		declare subspace
	// 		{
	// 			uint64 CreatedTime 0
	// 			byte HashedPassword /HF6NyMfkCemAEYz1Ykjl5uJTK0=
	// 			uint64 LastCommTime 0
	// 			uint64 LastLoginTime 0
	// 			uint NumLogin 0
	// 			bool Online true
	// 			bool RadiusConvertAllMsChapv2AuthRequestToEap false
	// 			string RadiusRealm $
	// 			uint RadiusRetryInterval 0
	// 			uint RadiusServerPort 1812
	// 			string RadiusSuffixFilter $
	// 			bool RadiusUsePeapInsteadOfEap false
	// 			byte SecurePassword AAAAAAAAAAAAAAAAAAAAAAAAAAA=
	// 			uint Type 0
	//
	// 			declare AccessList
	// 			{
	// 			}
	// 			declare AdminOption
	// 			{
	// 				uint allow_hub_admin_change_option 0
	// 				uint deny_bridge 0
	// 				uint deny_change_user_password 0
	// 				uint deny_empty_password 0
	// 				uint deny_hub_admin_change_ext_option 0
	// 				uint deny_qos 0
	// 				uint deny_routing 0
	// 				uint max_accesslists 0
	// 				uint max_bitrates_download 0
	// 				uint max_bitrates_upload 0
	// 				uint max_groups 0
	// 				uint max_multilogins_per_user 0
	// 				uint max_sessions 0
	// 				uint max_sessions_bridge 0
	// 				uint max_sessions_client 0
	// 				uint max_sessions_client_bridge_apply 0
	// 				uint max_users 0
	// 				uint no_access_list_include_file 0
	// 				uint no_cascade 0
	// 				uint no_change_access_control_list 0
	// 				uint no_change_access_list 0
	// 				uint no_change_admin_password 0
	// 				uint no_change_cert_list 0
	// 				uint no_change_crl_list 0
	// 				uint no_change_groups 0
	// 				uint no_change_log_config 0
	// 				uint no_change_log_switch_type 0
	// 				uint no_change_msg 0
	// 				uint no_change_users 0
	// 				uint no_delay_jitter_packet_loss 0
	// 				uint no_delete_iptable 0
	// 				uint no_delete_mactable 0
	// 				uint no_disconnect_session 0
	// 				uint no_enum_session 0
	// 				uint no_offline 0
	// 				uint no_online 0
	// 				uint no_query_session 0
	// 				uint no_read_log_file 0
	// 				uint no_securenat 0
	// 				uint no_securenat_enabledhcp 0
	// 				uint no_securenat_enablenat 0
	// 			}
	// 			declare CascadeList
	// 			{
	// 			}
	// 			declare LogSetting
	// 			{
	// 				uint PacketLogSwitchType 4
	// 				uint PACKET_LOG_ARP 0
	// 				uint PACKET_LOG_DHCP 1
	// 				uint PACKET_LOG_ETHERNET 0
	// 				uint PACKET_LOG_ICMP 0
	// 				uint PACKET_LOG_IP 0
	// 				uint PACKET_LOG_TCP 0
	// 				uint PACKET_LOG_TCP_CONN 1
	// 				uint PACKET_LOG_UDP 0
	// 				bool SavePacketLog true
	// 				bool SaveSecurityLog true
	// 				uint SecurityLogSwitchType 4
	// 			}
	// 			declare Message
	// 			{
	// 			}
	// 			declare Option
	// 			{
	// 				uint AccessListIncludeFileCacheLifetime 30
	// 				uint AdjustTcpMssValue 0
	// 				bool ApplyIPv4AccessListOnArpPacket false
	// 				bool AssignVLanIdByRadiusAttribute false
	// 				bool BroadcastLimiterStrictMode false
	// 				uint BroadcastStormDetectionThreshold 0
	// 				uint ClientMinimumRequiredBuild 0
	// 				bool DenyAllRadiusLoginWithNoVlanAssign false
	// 				uint DetectDormantSessionInterval 0
	// 				bool DisableAdjustTcpMss false
	// 				bool DisableCheckMacOnLocalBridge false
	// 				bool DisableCorrectIpOffloadChecksum false
	// 				bool DisableHttpParsing false
	// 				bool DisableIPParsing false
	// 				bool DisableIpRawModeSecureNAT false
	// 				bool DisableKernelModeSecureNAT false
	// 				bool DisableUdpAcceleration false
	// 				bool DisableUdpFilterForLocalBridgeNic false
	// 				bool DisableUserModeSecureNAT false
	// 				bool DoNotSaveHeavySecurityLogs false
	// 				bool DropArpInPrivacyFilterMode true
	// 				bool DropBroadcastsInPrivacyFilterMode true
	// 				bool FilterBPDU false
	// 				bool FilterIPv4 false
	// 				bool FilterIPv6 false
	// 				bool FilterNonIP false
	// 				bool FilterOSPF false
	// 				bool FilterPPPoE false
	// 				uint FloodingSendQueueBufferQuota 33554432
	// 				bool ManageOnlyLocalUnicastIPv6 true
	// 				bool ManageOnlyPrivateIP true
	// 				uint MaxLoggedPacketsPerMinute 0
	// 				uint MaxSession 0
	// 				bool NoArpPolling false
	// 				bool NoDhcpPacketLogOutsideHub true
	// 				bool NoEnum false
	// 				bool NoIpTable false
	// 				bool NoIPv4PacketLog false
	// 				bool NoIPv6AddrPolling false
	// 				bool NoIPv6DefaultRouterInRAWhenIPv6 true
	// 				bool NoIPv6PacketLog false
	// 				bool NoLookBPDUBridgeId false
	// 				bool NoMacAddressLog true
	// 				bool NoManageVlanId false
	// 				bool NoPhysicalIPOnPacketLog false
	// 				bool NoSpinLockForPacketDelay false
	// 				bool RemoveDefGwOnDhcpForLocalhost true
	// 				uint RequiredClientId 0
	// 				uint SecureNAT_MaxDnsSessionsPerIp 0
	// 				uint SecureNAT_MaxIcmpSessionsPerIp 0
	// 				uint SecureNAT_MaxTcpSessionsPerIp 0
	// 				uint SecureNAT_MaxTcpSynSentPerIp 0
	// 				uint SecureNAT_MaxUdpSessionsPerIp 0
	// 				bool SecureNAT_RandomizeAssignIp false
	// 				bool SuppressClientUpdateNotification false
	// 				bool UseHubNameAsDhcpUserClassOption false
	// 				bool UseHubNameAsRadiusNasId false
	// 				string VlanTypeId 0x8100
	// 				bool YieldAfterStorePacket false
	// 			}
	// 			declare SecureNAT
	// 			{
	// 				bool Disabled false
	// 				bool SaveLog true
	//
	// 				declare VirtualDhcpServer
	// 				{
	// 					string DhcpDnsServerAddress 192.168.30.1
	// 					string DhcpDnsServerAddress2 0.0.0.0
	// 					string DhcpDomainName $
	// 					bool DhcpEnabled true
	// 					uint DhcpExpireTimeSpan 7200
	// 					string DhcpGatewayAddress 192.168.30.1
	// 					string DhcpLeaseIPEnd 192.168.30.200
	// 					string DhcpLeaseIPStart 192.168.30.10
	// 					string DhcpPushRoutes $
	// 					string DhcpSubnetMask 255.255.255.0
	// 				}
	// 				declare VirtualHost
	// 				{
	// 					string VirtualHostIp 192.168.30.1
	// 					string VirtualHostIpSubnetMask 255.255.255.0
	// 					string VirtualHostMacAddress 00-AC-7F-26-60-3F
	// 				}
	// 				declare VirtualRouter
	// 				{
	// 					bool NatEnabled true
	// 					uint NatMtu 1500
	// 					uint NatTcpTimeout 1800
	// 					uint NatUdpTimeout 60
	// 				}
	// 			}
	// 			declare SecurityAccountDatabase
	// 			{
	// 				declare CertList
	// 				{
	// 				}
	// 				declare CrlList
	// 				{
	// 				}
	// 				declare GroupList
	// 				{
	// 				}
	// 				declare IPAccessControlList
	// 				{
	// 				}
	// 				declare UserList
	// 				{
	// 					declare 1_1495004926659038074
	// 					{
	// 						byte AuthNtLmSecureHash tmDCpoiYTNpWw/lQ4N3soA==
	// 						byte AuthPassword 8Tbpu20Wtd+56zbOyFmoInDE60k=
	// 						uint AuthType 1
	// 						uint64 CreatedTime 1494972526771
	// 						uint64 ExpireTime 0
	// 						uint64 LastLoginTime 1495665787182
	// 						string Note Default$20Admin$20VPN$20Profile
	// 						uint NumLogin 1
	// 						string RealName carterlin@ecoworkinc.com
	// 						uint64 UpdatedTime 1494972526908
	//
	// 						declare Traffic
	// 						{
	// 							declare RecvTraffic
	// 							{
	// 								uint64 BroadcastBytes 0
	// 								uint64 BroadcastCount 0
	// 								uint64 UnicastBytes 0
	// 								uint64 UnicastCount 0
	// 							}
	// 							declare SendTraffic
	// 							{
	// 								uint64 BroadcastBytes 0
	// 								uint64 BroadcastCount 0
	// 								uint64 UnicastBytes 0
	// 								uint64 UnicastCount 0
	// 							}
	// 						}
	// 					}
	// 				}
	// 			}
	// 			declare Traffic
	// 			{
	// 				declare RecvTraffic
	// 				{
	// 					uint64 BroadcastBytes 0
	// 					uint64 BroadcastCount 0
	// 					uint64 UnicastBytes 0
	// 					uint64 UnicastCount 0
	// 				}
	// 				declare SendTraffic
	// 				{
	// 					uint64 BroadcastBytes 0
	// 					uint64 BroadcastCount 0
	// 					uint64 UnicastBytes 0
	// 					uint64 UnicastCount 0
	// 				}
	// 			}
	// 		}
}

func ExampleGenerateSoftetherServerConfig() {
	result, _ := vpn.GenerateSoftetherConfig(server)
	fmt.Println(result)
	// Output:
	// # Software Configuration File
	// # ---------------------------
	// #
	// # You may edit this file when the VPN Server / Client / Bridge program is not running.
	// #
	// # In prior to edit this file manually by your text editor,
	// # shutdown the VPN Server / Client / Bridge background service.
	// # Otherwise, all changes will be lost.
	// #
	// declare root
	// {
	// 	uint ConfigRevision 1
	// 	bool IPsecMessageDisplayed true
	// 	string Region US
	// 	bool VgsMessageDisplayed false
	//
	// 	declare DDnsClient
	// 	{
	// 		bool Disabled false
	// 		byte Key XW+hOCwdQgezlyygTgJfW8nDNMo=
	// 		string LocalHostname ip-10-0-0-169
	// 		string ProxyHostName $
	// 		uint ProxyPort 0
	// 		uint ProxyType 0
	// 		string ProxyUsername $
	// 	}
	// 	declare IPsec
	// 	{
	// 		bool EtherIP_IPsec false
	// 		string IPsec_Secret subspace
	// 		string L2TP_DefaultHub subspace
	// 		bool L2TP_IPsec true
	// 		bool L2TP_Raw false
	//
	// 		declare EtherIP_IDSettingsList
	// 		{
	// 		}
	// 	}
	// 	declare ListenerList
	// 	{
	// 		declare Listener0
	// 		{
	// 			bool DisableDos false
	// 			bool Enabled true
	// 			uint Port 992
	// 		}
	// 		declare Listener1
	// 		{
	// 			bool DisableDos false
	// 			bool Enabled true
	// 			uint Port 1194
	// 		}
	// 		declare Listener2
	// 		{
	// 			bool DisableDos false
	// 			bool Enabled true
	// 			uint Port 5555
	// 		}
	// 	}
	// 	declare LocalBridgeList
	// 	{
	// 		bool DoNotDisableOffloading false
	// 	}
	// 	declare ServerConfiguration
	// 	{
	// 		bool AcceptOnlyTls true
	// 		uint64 AutoDeleteCheckDiskFreeSpaceMin 104857600
	// 		uint AutoDeleteCheckIntervalSecs 300
	// 		uint AutoSaveConfigSpan 300
	// 		bool BackupConfigOnlyWhenModified true
	// 		string CipherName RC4-MD5
	// 		uint CurrentBuild 9634
	// 		bool DisableCoreDumpOnUnix false
	// 		bool DisableDeadLockCheck false
	// 		bool DisableDosProction false
	// 		bool DisableGetHostNameWhenAcceptTcp false
	// 		bool DisableIntelAesAcceleration false
	// 		bool DisableIPv6Listener false
	// 		bool DisableNatTraversal false
	// 		bool DisableOpenVPNServer false
	// 		bool DisableSessionReconnect false
	// 		bool DisableSSTPServer false
	// 		bool DontBackupConfig false
	// 		bool EnableVpnAzure false
	// 		bool EnableVpnOverDns false
	// 		bool EnableVpnOverIcmp false
	// 		byte HashedPassword /HF6NyMfkCemAEYz1Ykjl5uJTK0=
	// 		string KeepConnectHost keepalive.softether.org
	// 		uint KeepConnectInterval 50
	// 		uint KeepConnectPort 80
	// 		uint KeepConnectProtocol 1
	// 		uint64 LoggerMaxLogSize 1073741823
	// 		uint MaxConcurrentDnsClientThreads 512
	// 		uint MaxConnectionsPerIP 256
	// 		uint MaxUnestablishedConnections 1000
	// 		bool NoHighPriorityProcess false
	// 		bool NoLinuxArpFilter false
	// 		bool NoSendSignature false
	// 		string OpenVPNDefaultClientOption dev-type$20tun,link-mtu$201500,tun-mtu$201500,cipher$20AES-128-CBC,auth$20SHA1,keysize$20128,key-method$202,tls-client
	// 		string OpenVPN_UdpPortList 1194
	// 		bool SaveDebugLog false
	// 		byte ServerCert MIIDvjCCAqagAwIBAgIBADANBgkqhkiG9w0BAQsFADBeMRkwFwYDVQQDDBBpcC0xNzItMzEtMzctMTQ1MRkwFwYDVQQKDBBpcC0xNzItMzEtMzctMTQ1MRkwFwYDVQQLDBBpcC0xNzItMzEtMzctMTQ1MQswCQYDVQQGEwJVUzAeFw0xNzA1MTUxOTA5MDFaFw0zNzEyMzExOTA5MDFaMF4xGTAXBgNVBAMMEGlwLTE3Mi0zMS0zNy0xNDUxGTAXBgNVBAoMEGlwLTE3Mi0zMS0zNy0xNDUxGTAXBgNVBAsMEGlwLTE3Mi0zMS0zNy0xNDUxCzAJBgNVBAYTAlVTMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsBcC8qsjE1bKoF6GWFwDHFmXgBbn2rd536i9VIH/3nze6MoE5GRI08nASsVQnYilJAxG9I6Gu6QsPmKGZ++mfsHF4kxZere6I6pwioIRm3DisN7lhOtuAo3T3sHx4EH40afrI+CasZQ+IGbJyWTzF6jE95JsRgoiLGtc2nhoeUH5JC88cejBsFlQHURLcSqEmEmQGx2ZBe4kUczftYXzzWMCEUb2gZI4MUEutQKuJbPjpsrMPCvvssIEQvH8KU+qHK9Zu1alqBM9ZCsAjEpOdhUd0FlMwclOnkGoNaLVYbasZHWHm+ACV9IlMbbb2UeutjkZlp2viwr0YgpVhX2W2QIDAQABo4GGMIGDMA8GA1UdEwEB/wQFMAMBAf8wCwYDVR0PBAQDAgH2MGMGA1UdJQRcMFoGCCsGAQUFBwMBBggrBgEFBQcDAgYIKwYBBQUHAwMGCCsGAQUFBwMEBggrBgEFBQcDBQYIKwYBBQUHAwYGCCsGAQUFBwMHBggrBgEFBQcDCAYIKwYBBQUHAwkwDQYJKoZIhvcNAQELBQADggEBAIDPwSNQMCkqoDpmVW5sQOyqNNrHvZDC4mhnviI459gSOtwjel9bKC79lHwSLBtcoinnlNflkPV/UN1tbEa7EU8T8+cZvq0sBrnGjgTRF91MfL9ti5AHJqylgIVORdZKgCqBc2qvhU7OMFyfPKfcwlKQTd0NvYvOsPZMKlGwo6+7G2EJNf/D0Uq7d/7Ja+W8+ZiYCB++kF4XPyWtvyzxvDWVjL8nMm1qKR+mxxr2cUbe4Y/RGjWHHvv3hmtFqgHNOlEmk3QXBglqSDX9hWsTXdyzMQ6oZwCTA6I2Jazy62dYTt/JuevJw8TRBAwXFkvEW/c8UHTr1CO986aZEXWlNGU=
	// 		byte ServerKey MIIEpAIBAAKCAQEAsBcC8qsjE1bKoF6GWFwDHFmXgBbn2rd536i9VIH/3nze6MoE5GRI08nASsVQnYilJAxG9I6Gu6QsPmKGZ++mfsHF4kxZere6I6pwioIRm3DisN7lhOtuAo3T3sHx4EH40afrI+CasZQ+IGbJyWTzF6jE95JsRgoiLGtc2nhoeUH5JC88cejBsFlQHURLcSqEmEmQGx2ZBe4kUczftYXzzWMCEUb2gZI4MUEutQKuJbPjpsrMPCvvssIEQvH8KU+qHK9Zu1alqBM9ZCsAjEpOdhUd0FlMwclOnkGoNaLVYbasZHWHm+ACV9IlMbbb2UeutjkZlp2viwr0YgpVhX2W2QIDAQABAoIBAALvz0/GxOX8q5JzISB2IGMnp1wuhZ1jM/gj80xmgqLuSDVvsRcyCPjnQ3q7wbVu4vov7jsD6YFVoptDv2QaXhd9GdTFWzStvb/8hrX4p1yFTHKQtRk7Z/FI9kCf0W3hC8pLV6OYilvt0f9uV20xbtB1MxvB6GZeouVEwR6vpWZoSSfZgxtqoZNhylDhtGLth2ucrC4sGLy49VYZc7MhP0GNqbkMnGEd2aVY8zagSIVyMcgp+0J6qCnCyZ0xsK9bVp/jdm3Blx4qI/gp8+6HL7VP1V9AnUhW91cVWjJ+w9Y27Xw/bgO2rgGIHtA5Cvbt6KeLplKqAxIS+31xqDXl5RECgYEA6KVAq7GsEIk+ajvy/9zJLOR/Y9RU+DLx8ki72335KbHQF9Lk1cQqhTruOS0B8E9YgFdtx6P2+nbDPtCXiwywrizzGK8As9rh6EMMCB/bMDLgKx8K5GFQYmG7ouX5h6ZFEufBh+yAFFrmpyU7Cd7CRbpItjHrr+wDGB3Ket/mSu0CgYEAwcRWkXhkvrv7WW0kB0Yr3ACXZXZD2RQrQmrI7KiwiM7E+REZ43BLLh1MgP12wLYZ/N92cs2gyCqstQ5FY/wPLcWORepeo91LEKqgAIIwSuGlNlyFysJRLkYyzsYvMs1MqdddgtA8oOF8oGX9EIVxIXD6BaH7Xu17MpfUqcyuQh0CgYAEqwQwLcF5F3fkI83Nan2y4+bWdvU5gbQvG1yBm/Z66rF1OOsUivobcEJFbWzDiKQK9zYhsl2LSwTB5ueTh66n4AkErsZB1ZuA0I+WPiVkwzCSTO4oBPfa8YEVVK38Fc7/AR11/73WKrjEE9Aqc4jMY1+AIGJNRlwNiNN+Qa3aRQKBgQC8dbg1GMHz+WYInmySXp++BR07ZAGtmz5QBQiTlabOCT3vecCSQ9/7ZCfKtbvQx80S9E2Cb5lX1pnQJ6c6Dii9Pg6Y1dFi05N1DF8+32EVE87axraidMj8lu5Hyt6RLLjJ4FxlJSAy4d2TLF8suH1GJ1omLFwPAi/3D1GaLSe3yQKBgQCPlY+k07duoeuYx9Irl3JTyRTa6woWkHUyngHoDVftiWzxYBvMJ2NB1QHCUobp0TcoKxDntztbmcTNSWdHRNHSYik4htAKX4dvCbA/jmvKXbzAkzfk60tkeY2Paqp6WvSwoOjXIrFOvJohDGbZK98JxB6Fff6Epx95L8GtPrFo0g==
	// 		uint ServerLogSwitchType 4
	// 		uint ServerType 0
	// 		bool Tls_Disable1_0 false
	// 		bool Tls_Disable1_1 false
	// 		bool Tls_Disable1_2 false
	// 		bool UseKeepConnect true
	// 		bool UseWebTimePage false
	// 		bool UseWebUI false
	//
	// 		declare GlobalParams
	// 		{
	// 			uint FIFO_BUDGET 10240000
	// 			uint HUB_ARP_SEND_INTERVAL 5000
	// 			uint IP_TABLE_EXPIRE_TIME 60000
	// 			uint IP_TABLE_EXPIRE_TIME_DHCP 300000
	// 			uint MAC_TABLE_EXPIRE_TIME 600000
	// 			uint MAX_BUFFERING_PACKET_SIZE 2560000
	// 			uint MAX_HUB_LINKS 1024
	// 			uint MAX_IP_TABLES 65536
	// 			uint MAX_MAC_TABLES 65536
	// 			uint MAX_SEND_SOCKET_QUEUE_NUM 128
	// 			uint MAX_SEND_SOCKET_QUEUE_SIZE 2560000
	// 			uint MAX_STORED_QUEUE_NUM 1024
	// 			uint MEM_FIFO_REALLOC_MEM_SIZE 655360
	// 			uint MIN_SEND_SOCKET_QUEUE_SIZE 320000
	// 			uint QUEUE_BUDGET 2048
	// 			uint SELECT_TIME 256
	// 			uint SELECT_TIME_FOR_NAT 30
	// 			uint STORM_CHECK_SPAN 500
	// 			uint STORM_DISCARD_VALUE_END 1024
	// 			uint STORM_DISCARD_VALUE_START 3
	// 		}
	// 		declare ServerTraffic
	// 		{
	// 			declare RecvTraffic
	// 			{
	// 				uint64 BroadcastBytes 0
	// 				uint64 BroadcastCount 0
	// 				uint64 UnicastBytes 0
	// 				uint64 UnicastCount 0
	// 			}
	// 			declare SendTraffic
	// 			{
	// 				uint64 BroadcastBytes 0
	// 				uint64 BroadcastCount 0
	// 				uint64 UnicastBytes 0
	// 				uint64 UnicastCount 0
	// 			}
	// 		}
	// 		declare SyslogSettings
	// 		{
	// 			string HostName $
	// 			uint Port 0
	// 			uint SaveType 0
	// 		}
	// 	}
	// 	declare VirtualHUB
	// 	{
	// 		declare DEFAULT
	// 		{
	// 			uint64 CreatedTime 1494842941898
	// 			byte HashedPassword +WzqGYrR3VYXrAhKPZLGEHcIwO8=
	// 			uint64 LastCommTime 1494842941893
	// 			uint64 LastLoginTime 1494842941893
	// 			uint NumLogin 0
	// 			bool Online true
	// 			bool RadiusConvertAllMsChapv2AuthRequestToEap false
	// 			string RadiusRealm $
	// 			uint RadiusRetryInterval 0
	// 			uint RadiusServerPort 1812
	// 			string RadiusSuffixFilter $
	// 			bool RadiusUsePeapInsteadOfEap false
	// 			byte SecurePassword bpw3X/O5E8a6G6ccnl4uXmDtkwI=
	// 			uint Type 0
	//
	// 			declare AccessList
	// 			{
	// 			}
	// 			declare AdminOption
	// 			{
	// 				uint allow_hub_admin_change_option 0
	// 				uint deny_bridge 0
	// 				uint deny_change_user_password 0
	// 				uint deny_empty_password 0
	// 				uint deny_hub_admin_change_ext_option 0
	// 				uint deny_qos 0
	// 				uint deny_routing 0
	// 				uint max_accesslists 0
	// 				uint max_bitrates_download 0
	// 				uint max_bitrates_upload 0
	// 				uint max_groups 0
	// 				uint max_multilogins_per_user 0
	// 				uint max_sessions 0
	// 				uint max_sessions_bridge 0
	// 				uint max_sessions_client 0
	// 				uint max_sessions_client_bridge_apply 0
	// 				uint max_users 0
	// 				uint no_access_list_include_file 0
	// 				uint no_cascade 0
	// 				uint no_change_access_control_list 0
	// 				uint no_change_access_list 0
	// 				uint no_change_admin_password 0
	// 				uint no_change_cert_list 0
	// 				uint no_change_crl_list 0
	// 				uint no_change_groups 0
	// 				uint no_change_log_config 0
	// 				uint no_change_log_switch_type 0
	// 				uint no_change_msg 0
	// 				uint no_change_users 0
	// 				uint no_delay_jitter_packet_loss 0
	// 				uint no_delete_iptable 0
	// 				uint no_delete_mactable 0
	// 				uint no_disconnect_session 0
	// 				uint no_enum_session 0
	// 				uint no_offline 0
	// 				uint no_online 0
	// 				uint no_query_session 0
	// 				uint no_read_log_file 0
	// 				uint no_securenat 0
	// 				uint no_securenat_enabledhcp 0
	// 				uint no_securenat_enablenat 0
	// 			}
	// 			declare CascadeList
	// 			{
	// 			}
	// 			declare LogSetting
	// 			{
	// 				uint PacketLogSwitchType 4
	// 				uint PACKET_LOG_ARP 0
	// 				uint PACKET_LOG_DHCP 1
	// 				uint PACKET_LOG_ETHERNET 0
	// 				uint PACKET_LOG_ICMP 0
	// 				uint PACKET_LOG_IP 0
	// 				uint PACKET_LOG_TCP 0
	// 				uint PACKET_LOG_TCP_CONN 1
	// 				uint PACKET_LOG_UDP 0
	// 				bool SavePacketLog true
	// 				bool SaveSecurityLog true
	// 				uint SecurityLogSwitchType 4
	// 			}
	// 			declare Message
	// 			{
	// 			}
	// 			declare Option
	// 			{
	// 				uint AccessListIncludeFileCacheLifetime 30
	// 				uint AdjustTcpMssValue 0
	// 				bool ApplyIPv4AccessListOnArpPacket false
	// 				bool AssignVLanIdByRadiusAttribute false
	// 				bool BroadcastLimiterStrictMode false
	// 				uint BroadcastStormDetectionThreshold 0
	// 				uint ClientMinimumRequiredBuild 0
	// 				bool DenyAllRadiusLoginWithNoVlanAssign false
	// 				uint DetectDormantSessionInterval 0
	// 				bool DisableAdjustTcpMss false
	// 				bool DisableCheckMacOnLocalBridge false
	// 				bool DisableCorrectIpOffloadChecksum false
	// 				bool DisableHttpParsing false
	// 				bool DisableIPParsing false
	// 				bool DisableIpRawModeSecureNAT false
	// 				bool DisableKernelModeSecureNAT false
	// 				bool DisableUdpAcceleration false
	// 				bool DisableUdpFilterForLocalBridgeNic false
	// 				bool DisableUserModeSecureNAT false
	// 				bool DoNotSaveHeavySecurityLogs false
	// 				bool DropArpInPrivacyFilterMode true
	// 				bool DropBroadcastsInPrivacyFilterMode true
	// 				bool FilterBPDU false
	// 				bool FilterIPv4 false
	// 				bool FilterIPv6 false
	// 				bool FilterNonIP false
	// 				bool FilterOSPF false
	// 				bool FilterPPPoE false
	// 				uint FloodingSendQueueBufferQuota 33554432
	// 				bool ManageOnlyLocalUnicastIPv6 true
	// 				bool ManageOnlyPrivateIP true
	// 				uint MaxLoggedPacketsPerMinute 0
	// 				uint MaxSession 0
	// 				bool NoArpPolling false
	// 				bool NoDhcpPacketLogOutsideHub true
	// 				bool NoEnum false
	// 				bool NoIpTable false
	// 				bool NoIPv4PacketLog false
	// 				bool NoIPv6AddrPolling false
	// 				bool NoIPv6DefaultRouterInRAWhenIPv6 true
	// 				bool NoIPv6PacketLog false
	// 				bool NoLookBPDUBridgeId false
	// 				bool NoMacAddressLog true
	// 				bool NoManageVlanId false
	// 				bool NoPhysicalIPOnPacketLog false
	// 				bool NoSpinLockForPacketDelay false
	// 				bool RemoveDefGwOnDhcpForLocalhost true
	// 				uint RequiredClientId 0
	// 				uint SecureNAT_MaxDnsSessionsPerIp 0
	// 				uint SecureNAT_MaxIcmpSessionsPerIp 0
	// 				uint SecureNAT_MaxTcpSessionsPerIp 0
	// 				uint SecureNAT_MaxTcpSynSentPerIp 0
	// 				uint SecureNAT_MaxUdpSessionsPerIp 0
	// 				bool SecureNAT_RandomizeAssignIp false
	// 				bool SuppressClientUpdateNotification false
	// 				bool UseHubNameAsDhcpUserClassOption false
	// 				bool UseHubNameAsRadiusNasId false
	// 				string VlanTypeId 0x8100
	// 				bool YieldAfterStorePacket false
	// 			}
	// 			declare SecureNAT
	// 			{
	// 				bool Disabled true
	// 				bool SaveLog true
	//
	// 				declare VirtualDhcpServer
	// 				{
	// 					string DhcpDnsServerAddress 192.168.30.1
	// 					string DhcpDnsServerAddress2 0.0.0.0
	// 					string DhcpDomainName $
	// 					bool DhcpEnabled true
	// 					uint DhcpExpireTimeSpan 7200
	// 					string DhcpGatewayAddress 192.168.30.1
	// 					string DhcpLeaseIPEnd 192.168.30.200
	// 					string DhcpLeaseIPStart 192.168.30.10
	// 					string DhcpPushRoutes $
	// 					string DhcpSubnetMask 255.255.255.0
	// 				}
	// 				declare VirtualHost
	// 				{
	// 					string VirtualHostIp 192.168.30.1
	// 					string VirtualHostIpSubnetMask 255.255.255.0
	// 					string VirtualHostMacAddress 00-AC-5A-21-3E-20
	// 				}
	// 				declare VirtualRouter
	// 				{
	// 					bool NatEnabled true
	// 					uint NatMtu 1500
	// 					uint NatTcpTimeout 1800
	// 					uint NatUdpTimeout 60
	// 				}
	// 			}
	// 			declare SecurityAccountDatabase
	// 			{
	// 				declare CertList
	// 				{
	// 				}
	// 				declare CrlList
	// 				{
	// 				}
	// 				declare GroupList
	// 				{
	// 				}
	// 				declare IPAccessControlList
	// 				{
	// 				}
	// 				declare UserList
	// 				{
	// 				}
	// 			}
	// 			declare Traffic
	// 			{
	// 				declare RecvTraffic
	// 				{
	// 					uint64 BroadcastBytes 0
	// 					uint64 BroadcastCount 0
	// 					uint64 UnicastBytes 0
	// 					uint64 UnicastCount 0
	// 				}
	// 				declare SendTraffic
	// 				{
	// 					uint64 BroadcastBytes 0
	// 					uint64 BroadcastCount 0
	// 					uint64 UnicastBytes 0
	// 					uint64 UnicastCount 0
	// 				}
	// 			}
	// 		}
	// 		declare subspace
	// 		{
	// 			uint64 CreatedTime 0
	// 			byte HashedPassword /HF6NyMfkCemAEYz1Ykjl5uJTK0=
	// 			uint64 LastCommTime 0
	// 			uint64 LastLoginTime 0
	// 			uint NumLogin 0
	// 			bool Online true
	// 			bool RadiusConvertAllMsChapv2AuthRequestToEap false
	// 			string RadiusRealm $
	// 			uint RadiusRetryInterval 0
	// 			uint RadiusServerPort 1812
	// 			string RadiusSuffixFilter $
	// 			bool RadiusUsePeapInsteadOfEap false
	// 			byte SecurePassword AAAAAAAAAAAAAAAAAAAAAAAAAAA=
	// 			uint Type 0
	//
	// 			declare AccessList
	// 			{
	// 			}
	// 			declare AdminOption
	// 			{
	// 				uint allow_hub_admin_change_option 0
	// 				uint deny_bridge 0
	// 				uint deny_change_user_password 0
	// 				uint deny_empty_password 0
	// 				uint deny_hub_admin_change_ext_option 0
	// 				uint deny_qos 0
	// 				uint deny_routing 0
	// 				uint max_accesslists 0
	// 				uint max_bitrates_download 0
	// 				uint max_bitrates_upload 0
	// 				uint max_groups 0
	// 				uint max_multilogins_per_user 0
	// 				uint max_sessions 0
	// 				uint max_sessions_bridge 0
	// 				uint max_sessions_client 0
	// 				uint max_sessions_client_bridge_apply 0
	// 				uint max_users 0
	// 				uint no_access_list_include_file 0
	// 				uint no_cascade 0
	// 				uint no_change_access_control_list 0
	// 				uint no_change_access_list 0
	// 				uint no_change_admin_password 0
	// 				uint no_change_cert_list 0
	// 				uint no_change_crl_list 0
	// 				uint no_change_groups 0
	// 				uint no_change_log_config 0
	// 				uint no_change_log_switch_type 0
	// 				uint no_change_msg 0
	// 				uint no_change_users 0
	// 				uint no_delay_jitter_packet_loss 0
	// 				uint no_delete_iptable 0
	// 				uint no_delete_mactable 0
	// 				uint no_disconnect_session 0
	// 				uint no_enum_session 0
	// 				uint no_offline 0
	// 				uint no_online 0
	// 				uint no_query_session 0
	// 				uint no_read_log_file 0
	// 				uint no_securenat 0
	// 				uint no_securenat_enabledhcp 0
	// 				uint no_securenat_enablenat 0
	// 			}
	// 			declare CascadeList
	// 			{
	// 			}
	// 			declare LogSetting
	// 			{
	// 				uint PacketLogSwitchType 4
	// 				uint PACKET_LOG_ARP 0
	// 				uint PACKET_LOG_DHCP 1
	// 				uint PACKET_LOG_ETHERNET 0
	// 				uint PACKET_LOG_ICMP 0
	// 				uint PACKET_LOG_IP 0
	// 				uint PACKET_LOG_TCP 0
	// 				uint PACKET_LOG_TCP_CONN 1
	// 				uint PACKET_LOG_UDP 0
	// 				bool SavePacketLog true
	// 				bool SaveSecurityLog true
	// 				uint SecurityLogSwitchType 4
	// 			}
	// 			declare Message
	// 			{
	// 			}
	// 			declare Option
	// 			{
	// 				uint AccessListIncludeFileCacheLifetime 30
	// 				uint AdjustTcpMssValue 0
	// 				bool ApplyIPv4AccessListOnArpPacket false
	// 				bool AssignVLanIdByRadiusAttribute false
	// 				bool BroadcastLimiterStrictMode false
	// 				uint BroadcastStormDetectionThreshold 0
	// 				uint ClientMinimumRequiredBuild 0
	// 				bool DenyAllRadiusLoginWithNoVlanAssign false
	// 				uint DetectDormantSessionInterval 0
	// 				bool DisableAdjustTcpMss false
	// 				bool DisableCheckMacOnLocalBridge false
	// 				bool DisableCorrectIpOffloadChecksum false
	// 				bool DisableHttpParsing false
	// 				bool DisableIPParsing false
	// 				bool DisableIpRawModeSecureNAT false
	// 				bool DisableKernelModeSecureNAT false
	// 				bool DisableUdpAcceleration false
	// 				bool DisableUdpFilterForLocalBridgeNic false
	// 				bool DisableUserModeSecureNAT false
	// 				bool DoNotSaveHeavySecurityLogs false
	// 				bool DropArpInPrivacyFilterMode true
	// 				bool DropBroadcastsInPrivacyFilterMode true
	// 				bool FilterBPDU false
	// 				bool FilterIPv4 false
	// 				bool FilterIPv6 false
	// 				bool FilterNonIP false
	// 				bool FilterOSPF false
	// 				bool FilterPPPoE false
	// 				uint FloodingSendQueueBufferQuota 33554432
	// 				bool ManageOnlyLocalUnicastIPv6 true
	// 				bool ManageOnlyPrivateIP true
	// 				uint MaxLoggedPacketsPerMinute 0
	// 				uint MaxSession 0
	// 				bool NoArpPolling false
	// 				bool NoDhcpPacketLogOutsideHub true
	// 				bool NoEnum false
	// 				bool NoIpTable false
	// 				bool NoIPv4PacketLog false
	// 				bool NoIPv6AddrPolling false
	// 				bool NoIPv6DefaultRouterInRAWhenIPv6 true
	// 				bool NoIPv6PacketLog false
	// 				bool NoLookBPDUBridgeId false
	// 				bool NoMacAddressLog true
	// 				bool NoManageVlanId false
	// 				bool NoPhysicalIPOnPacketLog false
	// 				bool NoSpinLockForPacketDelay false
	// 				bool RemoveDefGwOnDhcpForLocalhost true
	// 				uint RequiredClientId 0
	// 				uint SecureNAT_MaxDnsSessionsPerIp 0
	// 				uint SecureNAT_MaxIcmpSessionsPerIp 0
	// 				uint SecureNAT_MaxTcpSessionsPerIp 0
	// 				uint SecureNAT_MaxTcpSynSentPerIp 0
	// 				uint SecureNAT_MaxUdpSessionsPerIp 0
	// 				bool SecureNAT_RandomizeAssignIp false
	// 				bool SuppressClientUpdateNotification false
	// 				bool UseHubNameAsDhcpUserClassOption false
	// 				bool UseHubNameAsRadiusNasId false
	// 				string VlanTypeId 0x8100
	// 				bool YieldAfterStorePacket false
	// 			}
	// 			declare SecureNAT
	// 			{
	// 				bool Disabled false
	// 				bool SaveLog true
	//
	// 				declare VirtualDhcpServer
	// 				{
	// 					string DhcpDnsServerAddress 192.168.30.1
	// 					string DhcpDnsServerAddress2 0.0.0.0
	// 					string DhcpDomainName $
	// 					bool DhcpEnabled true
	// 					uint DhcpExpireTimeSpan 7200
	// 					string DhcpGatewayAddress 192.168.30.1
	// 					string DhcpLeaseIPEnd 192.168.30.200
	// 					string DhcpLeaseIPStart 192.168.30.10
	// 					string DhcpPushRoutes $
	// 					string DhcpSubnetMask 255.255.255.0
	// 				}
	// 				declare VirtualHost
	// 				{
	// 					string VirtualHostIp 192.168.30.1
	// 					string VirtualHostIpSubnetMask 255.255.255.0
	// 					string VirtualHostMacAddress 00-AC-7F-26-60-3F
	// 				}
	// 				declare VirtualRouter
	// 				{
	// 					bool NatEnabled true
	// 					uint NatMtu 1500
	// 					uint NatTcpTimeout 1800
	// 					uint NatUdpTimeout 60
	// 				}
	// 			}
	// 			declare SecurityAccountDatabase
	// 			{
	// 				declare CertList
	// 				{
	// 				}
	// 				declare CrlList
	// 				{
	// 				}
	// 				declare GroupList
	// 				{
	// 				}
	// 				declare IPAccessControlList
	// 				{
	// 				}
	// 				declare UserList
	// 				{
	// 					declare 1_1495004926659038074
	// 					{
	// 						byte AuthNtLmSecureHash tmDCpoiYTNpWw/lQ4N3soA==
	// 						byte AuthPassword 8Tbpu20Wtd+56zbOyFmoInDE60k=
	// 						uint AuthType 1
	// 						uint64 CreatedTime 1494972526771
	// 						uint64 ExpireTime 0
	// 						uint64 LastLoginTime 1495665787182
	// 						string Note Default$20Admin$20VPN$20Profile
	// 						uint NumLogin 1
	// 						string RealName carterlin@ecoworkinc.com
	// 						uint64 UpdatedTime 1494972526908
	//
	// 						declare Traffic
	// 						{
	// 							declare RecvTraffic
	// 							{
	// 								uint64 BroadcastBytes 0
	// 								uint64 BroadcastCount 0
	// 								uint64 UnicastBytes 0
	// 								uint64 UnicastCount 0
	// 							}
	// 							declare SendTraffic
	// 							{
	// 								uint64 BroadcastBytes 0
	// 								uint64 BroadcastCount 0
	// 								uint64 UnicastBytes 0
	// 								uint64 UnicastCount 0
	// 							}
	// 						}
	// 					}
	// 				}
	// 			}
	// 			declare Traffic
	// 			{
	// 				declare RecvTraffic
	// 				{
	// 					uint64 BroadcastBytes 0
	// 					uint64 BroadcastCount 0
	// 					uint64 UnicastBytes 0
	// 					uint64 UnicastCount 0
	// 				}
	// 				declare SendTraffic
	// 				{
	// 					uint64 BroadcastBytes 0
	// 					uint64 BroadcastCount 0
	// 					uint64 UnicastBytes 0
	// 					uint64 UnicastCount 0
	// 				}
	// 			}
	// 		}
	// 	}
	// 	declare VirtualLayer3SwitchList
	// 	{
	// 	}
	// }
	//
}