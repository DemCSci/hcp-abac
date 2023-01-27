'use strict';

const { WorkloadModuleBase } = require('@hyperledger/caliper-core');

class MyWorkload extends WorkloadModuleBase {

    constructor() {
        super();
    }
    
    async initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext) {
        await super.initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext);
        this.id = 0;
    }
    
    async submitTransaction() { 
        const myArgs = {
            contractId: this.roundArguments.contractId,
            contractFunction: 'CreateAsset',
            invokerIdentity: 'User1',
            //id string, color string, size int, owner string, appraisedValu    e int
            contractArguments: [`worker_${this.workerIndex}_${this.id}`, "blue", 14, "user1", 100],
            readOnly: false
        };

        await this.sutAdapter.sendRequests(myArgs);
        this.id++;
    }
    
    async cleanupWorkloadModule() {
        for (let i=0; i<this.id; i++) {
            const assetID = `worker_${this.workerIndex}_${i}`;
            console.log(`Worker ${this.workerIndex}/${this.totalWorkers}: Deleting asset ${assetID}`);
            const request = {
                contractId: this.roundArguments.contractId,
                contractFunction: 'DeleteAsset',
                invokerIdentity: 'User1',
                contractArguments: [assetID],
                readOnly: false
            };
            await this.sutAdapter.sendRequests(request);
        }
    }
}

function createWorkloadModule() {
    return new MyWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
