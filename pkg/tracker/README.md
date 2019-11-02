## Tracker

### 機能

位置情報の追跡，保存

### 実装

Tracking, SignalingにおけるすべてのClient-Server間の通信は
一つのWebSocket上で行う.
そのため処理の制御のためにWebSocket上でやり取りするデータのフォーマットを予め定義する.

#### データ構造

*Redis*

| Key | Value |
| :-: | :-: |
| ID | Pod IP:Port |

*MySQL*

| Column | Type | Description |
| :-: | :-: | :-: |
| id | int64 | インクリメント主キー. フレームワークでは特に使用しない |
| peer_id | char(128) | ピアの一意なID |
| h3_hash | char(128) | longitude, latitudeから算出したh3 hash. シャーディングに使用 |
| h3_resolution | int(10) | h3ハッシュの解像度(0~15) |
| latitude | decimal(8,6) | 緯度 |
| longitude | decimal(9,6)| 経度 |


#### コネクション確立
1. WebSocketでClientを待ち受ける
2. Clientはgreetメッセージを送信する
3. Tracker-DBに (Peer-ID, STUN or Relay,IP:Port, Longitude, Latitude) の情報を書き込む
4. ServerはClientのIDを鍵にredisに自身のLocal-IPとポートを書き込む
5. ServerはClientのIDを鍵に自身のメモリ上にコネクションのソケットをマッピングする

#### トラッキング
1. ClientはTrackingメッセージを送信する
2. Serverはデータに含まれる位置情報をClientのIDとともにTracker-DBに書き込む

#### Peer情報取得
1. ClientはLookUpメッセージを送信する
2. Serverはデータに含まれる起点となる経度，緯度から半径内に存在する
Peerの一覧をTracker-DBから取得する
3. Clientにその一覧を返信する

#### シグナリング
1. RoomOwnerとなるClientは通信したいPeerのIDのリストをServerに送信する
2. ServerはIDを用いてコネクションを管理するServerのリストをRedisを介して取得し，それぞれのサーバへ対象となるPeerが通信可能かをProtocolBufferを用いて確認する.
3. 他Serverは各PeerへPingメッセージを送信し，返信を確認した後にServerへ返信する.

ここからPeer同士の通信ステータスによって処理が変わる

- 1:1でPeer同士がSTUNでグローバルIP, Portを持っていた場合
    1. そのままお互いのIP, Portを用いてPure P2Pに通信する
- 1:1でどちらかのPeerがSTUNを超えられなかった場合
    1. RoomOwnerのPeerはこれ以降Relayコンポーネントと通信を行い
        Hybrid P2P通信を行う
- 1:nの場合
    1. すべてのピアがSTUNで通信可能の場合でもRelayコンポーネントを介して通信を行う
