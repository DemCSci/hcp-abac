### 本文工作
以无排序组织的方式启动 Hyperledger Fabric 网络，其中包含五个组织—— council 、 soft(org1) 、 web(org2) 、 hard(org3) 、org4、org5， council 组织为网络提供 TLS-CA 服务，并且运行维护着5个 orderer 服务；其余每个组织都运行维护着一个 peer 节点。

|               项                | 运行端口 |                           说明                           |
| :-----------------------------: | :------: | :------------------------------------------------------: |
|     `council.lei.net`      |   7050   |  council 组织的 CA 服务， 为联盟链网络提供 TLS-CA 服务   |
| `orderer1.council.lei.net` |   7051   |               council 组织的 orderer1 服务               |
| `orderer1.council.lei.net` |   7052   |        council 组织的 orderer1 服务的 admin 服务         |
| `orderer2.council.lei.net` |   7054   |               council 组织的 orderer2 服务               |
| `orderer2.council.lei.net` |   7055   |        council 组织的 orderer2 服务的 admin 服务         |
| `orderer3.council.lei.net` |   7057   |               council 组织的 orderer3 服务               |
| `orderer3.council.lei.net` |   7058   |        council 组织的 orderer3 服务的 admin 服务         |
| `orderer4.council.lei.net` |   7060   |               council 组织的 orderer4 服务               |
| `orderer4.council.lei.net` |   7061   |        council 组织的 orderer4 服务的 admin 服务         |
| `orderer5.council.lei.net` |   7063   |               council 组织的 orderer5 服务               |
| `orderer5.council.lei.net` |   7064   |        council 组织的 orderer5 服务的 admin 服务         |
|       `soft.lei.net`       |   7250   | soft 组织的 CA 服务， 包含成员： peer1 、 admin1 、user1 |
|    `peer1.soft.lei.net`    |   7251   |                soft 组织的 peer1 成员节点                |
|       `web.lei.net`        |   7350   | web 组织的 CA 服务， 包含成员： peer1 、 admin1 、user1  |
|    `peer1.web.lei.net`     |   7351   |                web 组织的 peer1 成员节点                 |
|       `hard.lei.net`       |   7450   | hard 组织的 CA 服务， 包含成员： peer1 、 admin1 、user1 |
|    `peer1.hard.lei.net`    |   7451   |                hard 组织的 peer1 成员节点                |
|       `org4.lei.net`       |   7550   | org4 组织的 CA 服务， 包含成员： peer1 、 admin1 、user1 |
|    `peer1.org4.lei.net`    |   7551   |                org4 组织的 peer1 成员节点                |
|       `org5.lei.net`       |   7650   | org5 组织的 CA 服务， 包含成员： peer1 、 admin1 、user1 |
|    `peer1.org5.lei.net`    |   7651   |                org5 组织的 peer1 成员节点                |
