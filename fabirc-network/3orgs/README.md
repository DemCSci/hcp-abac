### 本文工作
以无排序组织的方式启动 Hyperledger Fabric 网络，其中包含三个组织——  org1(org1) 、 web(org2) 、 org2(org3) ， council 组织为网络提供 TLS-CA 服务，并且运行维护着三个 orderer 服务；其余每个组织都运行维护着一个 peer 节点。

|               项                | 运行端口 |                           说明                           |
| :-----------------------------: | :------: | :------------------------------------------------------: |
|     `council.lei.net`      |   7050   |  council 组织的 CA 服务， 为联盟链网络提供 TLS-CA 服务   |
| `orderer1.council.lei.net` |   7051   |               council 组织的 orderer1 服务               |
| `orderer1.council.lei.net` |   7052   |        council 组织的 orderer1 服务的 admin 服务         |
| `orderer2.council.lei.net` |   7054   |               council 组织的 orderer2 服务               |
| `orderer2.council.lei.net` |   7055   |        council 组织的 orderer2 服务的 admin 服务         |
| `orderer3.council.lei.net` |   7057   |               council 组织的 orderer3 服务               |
| `orderer3.council.lei.net` |   7058   |        council 组织的 orderer3 服务的 admin 服务         |
|       `org1.lei.net`       |   7250   | org1(org1) 组织的 CA 服务， 包含成员： peer1 、 admin1 、user1 |
|    `peer1.org1.lei.net`    |   7251   |                org1 组织的 peer1 成员节点                |
|       `web.lei.net`        |   7350   | web(org2) 组织的 CA 服务， 包含成员： peer1 、 admin1 、user1  |
|    `peer1.web.lei.net`     |   7351   |                web 组织的 peer1 成员节点                 |
|       `org2.lei.net`       |   7450   | org2(org3)组织的 CA 服务， 包含成员： peer1 、 admin1 、user1 |
|    `peer1.org2.lei.net`    |   7451   |                org2 组织的 peer1 成员节点                |
