'use strict';

const { WorkloadModuleBase } = require('@hyperledger/caliper-core');

class MyWorkload extends WorkloadModuleBase {
    constructor() {
        super();
    }
    
    async initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext) {
        await super.initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext);

    }
    
    async submitTransaction() {
        let id = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
            var r = Math.random() * 16 | 0,
                v = c == 'x' ? r : (r & 0x3 | 0x8);
            return v.toString(16);
        });
        const request = {
            "id": id,
            "requesterId": "user:x509::CN=User1@org1.example.com,OU=client,L=San Francisco,ST=California,C=US::CN=ca.org1.example.com,O=org1.example.com,L=San Francisco,ST=California,C=US",
            "resourceId": "resource:ea913278-e617-4270-ac37-ae48377c64f4"
        }
        let requestJson = JSON.stringify(request)
        const myArgs = {
            contractId: this.roundArguments.contractId,
            version: "2.0",
            contractFunction: 'DecideWithRecord',
            invokerIdentity: 'User1',
            contractArguments: [requestJson],
            readOnly: false
        };
        
        await this.sutAdapter.sendRequests(myArgs);
    }
    
    async cleanupWorkloadModule() {
    }
}

function createWorkloadModule() {
    return new MyWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
