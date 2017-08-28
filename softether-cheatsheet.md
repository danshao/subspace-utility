# 安裝步驟

注意！此段與 Subspace 安裝的方式不同！

在 Ubuntu 16.04 安裝 SoftEther

## 更新環境

```sh
sudo apt-get update -y
sudo apt-get upgrade -y
sudo apt-get install libreadline-dev build-essential libssl-dev libncurses5-dev -y
```

## 下載安裝 SoftEther

### 下載解壓縮

```sh
curl -O http://www.softether-download.com/files/softether/v4.22-9634-beta-2016.11.27-tree/Source_Code/softether-src-v4.22-9634-beta.tar.gz
tar zxvf softether-src-v4.22-9634-beta.tar.gz
```

### 安裝

```sh
cd v4.22-9634
./configure
make
sudo make install
```

### 建立開機啟動(Ubuntu 16.04 LTS)

建立 `/etc/systemd/system/vpnserver.service`，內容為：

```service
[Unit]
Description=SoftEther VPN Server Daemon
After=network.target

[Service]
Type=forking
ExecStart=/usr/vpnserver/vpnserver start
ExecStop=/usr/vpnserver/ stop

[Install]
WantedBy=multi-user.target
```

如果路徑不同請自行修改。

執行 `sudo systemctl enable vpnserver`

### 建立開機啟動(Ubuntu 14.04 LTS)

在 `/etc/init.d/` 下建立 `vpnserver` 這個檔案。內容如下：

```sh
#!/bin/sh
# chkconfig: 2345 99 01
# description: SoftEther VPN Server

# 看你安裝的位置去調整此路徑
DAEMON=/usr/vpnserver/vpnserver
LOCK=/var/lock/subsys/vpnserver

test -x $DAEMON || exit 0

case "$1" in
  start)
    $DAEMON start
    touch $LOCK
    ;;
  stop)
    $DAEMON stop
    rm $LOCK
    ;;
  restart)
    $DAEMON stop
    sleep 3
    $DAEMON start
    ;;
  *)
    echo "Usage: $0 {start|stop|restart}"
    exit 1
    esac

exit 0
```

### 將 vpnserver 加進開機自動啟動

```sh
sudo update-rc.d vpnserver defaults
```

### 將 vpnserver 從開機自動啟動拿掉

```sh
sudo update-rc.d vpnserver remove
```

### 確認 vpnserver 是否正在跑

```sh
ps aux | grep vpnserver
```



## 環境設定

### 建立 Hub

```sh
sudo vpncmd localhost /server /cmd HubCreate MY_HUB /password:
```

### 建立使用者

建立使用者和設定該使用者的密碼是分開的兩個指令，無法在一個指令做完。

```sh
sudo vpncmd localhost /server /hub:MY_HUB /cmd UserCreate USER1 /GROUP:none /REALNAME:none /NOTE:none
sudo vpncmd localhost /server /hub:MY_HUB /cmd UserPasswordSet USER1 /PASSWORD:password
```

### 設定 Admin 密碼

這個留到最後做。

```sh
sudo vpncmd /server localhost /cmd ServerPasswordSet $(curl http://169.254.169.254/latest/meta-data/instance-id)
```

TODO

- 先建立 Hub 和 User 再設定上密碼或許可以避掉麻煩。
- 用 cloud-init 跑這件事情。


# Softether 操作

參考連結：

- [6.2 General Usage of vpncmd](https://www.softether.org/4-docs/1-manual/6._Command_Line_Management_Utility_Manual/6.2_General_Usage_of_vpncmd)
- [6.4 VPN Server / VPN Bridge Management Command Reference (For Virtual Hub)](https://www.softether.org/4-docs/1-manual/6._Command_Line_Management_Utility_Manual/6.4_VPN_Server_//_VPN_Bridge_Management_Command_Reference_%28For_Virtual_Hub%29)

## 備份還原

進入 SoftEther 後執行操作。進入方式：`vpncmd localhost /server /hub:subspace /password:subspace`

- 備份：記得先 Flush 會將設定檔匯出到 vpnserver 資料夾下，不能給路徑只能給檔名。

  ```
  Flush
  ConfigGet softether.config
  ```

   或是

  ```
  vpncmd localhost /server /password:my_password /cmd Flush
  vpncmd localhost /server /password:my_password /cmd ConfigGet
  ```

- 還原：將設定檔從 vpnserver 資料夾下匯入，不能給路徑只能給檔名。

  ```
  ConfigSet softether.config
  ```

# Softether 內部機制

## Server 密碼 Hash 方式

密碼 Hash 方式為 SHA0 後得到的 binary 再做 base64 encode。

以 instance ID `i-022b7a1921784a4c8` 做為密碼舉例如下。

```sh
php -r "echo base64_encode(openssl_digest('i-022b7a1921784a4c8', 'sha', true));"
```

結果：`uoHUBG6kJHZNNjPw9tsKjHRYbyc=`

```sh
echo -n "i-022b7a1921784a4c8" | openssl sha -binary | openssl base64
```

結果：`uoHUBG6kJHZNNjPw9tsKjHRYbyc=`

*這種方式可作為忘記 Server 密碼，想要進入 EC2 強制從 config 直接設定密碼時使用。或是以後有可以改 SoftEther 密碼時，做為重設密碼為 instance id 使用。*

## User 密碼 Hash 方式

將原始密碼與轉大寫的帳號串在一起後，用 SHA0 做 Hash 再做 base64 encode

```
php -r "echo base64_encode(openssl_digest('Password' . strtoupper('Username'), 'sha', true));"
```

```
echo -n "PasswordUSERNAME" | openssl sha -binary | openssl base64
```



# Known Issue

- Transfer Bytes, Transfer Packets 等等資訊在數字超過 20 位（21 位數）時會顯示錯亂。接近 1ZB 理論上目前根本用不到，所以應該可忽略。

# 疑難雜症排解

1. 遇到 VPN 有通過驗證，但是在 Client 都是嘗試連線一段時間然後連線失敗。關鍵 Log 如下：

   ```
   Acquiring an IP address from the DHCP server failed. To accept a PPP session, you need to have a DHCP server. Make sure that a DHCP server is working normally in the Ethernet segment which the Virtual Hub belongs to. If you do not have a DHCP server, you can use the Virtual DHCP function of the SecureNAT on the Virtual Hub instead.
   ```

   解法：進入 Hub 設定 > Virtual NAT and Virtual DHCP Server (SecureNAT) > Enable SecureNAT。

2. 遇到 Mac 上 Softether Server Manager 無法開啟，重新安裝也無法解決問題。
   執行 `ps aux | grep vpn` 找到正在執行的 `vpnsmgr.exe` 的 pid。
   執行 `kill -9 <PID>` 填入 pid 砍掉它。

3. 對 user 變更其 group，不會對該 user 已存在的 session 產生效果。

   - Network ACL 如果是對 group 設定，再把使用者加進該 group 時，要先將該 user 的 session 全部斷掉才會完全生效，否則只會對該 user 後續的連線有影響。

4. 修改完的值不會馬上被寫進去 vpn_server.config 內，必須要等待一段時間或是執行 vpncmd 的 flush。

   ​

