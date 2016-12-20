# vpn-status-checker
Simple text parser that will see who's logged into openvpn server.  I wrote this
as an exercise and also to run in Consul.  This will show who is logged into
VPN via the Consul web UI.

## Overview
Openvpn server provides a status log that will show who's currently logged in.

```
OpenVPN CLIENT LIST
Updated,Mon Dec 19 18:12:57 2016
Common Name,Real Address,Bytes Received,Bytes Sent,Connected Since
user1,10.10.10.10:36243,306736,1641137,Mon Dec 19 16:31:26 2016
user2,10.10.10.11:47810,94365,132645,Mon Dec 19 18:07:34 2016
ROUTING TABLE
Virtual Address,Common Name,Real Address,Last Ref
10.8.0.3,user1,10.10.10.10:36243,Mon Dec 19 18:12:19 2016
10.8.0.4,user2,10.10.10.11:47810,Mon Dec 19 18:12:56 2016
GLOBAL STATS
Max bcast/mcast queue length,0
END
```

This application will parse the text in between the 3rd line and the line that
has 'ROUTING TABLE'.

## Build
```
git clone https://github.com/nocode99/vpn-status-checker.git
cd vpn-status-checker
go build vpn-status-checker
```

## Run
vpn-status-checker /etc/opennvpn/openvpn-status.log
