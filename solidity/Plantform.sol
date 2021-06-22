// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <0.8.0;

/**
 * 移动众包任务合约，该合约在Requester发布众包任务时创建，目的是对Requester和Worker进行管理
 */
contract Plantform {
    string description;
    uint requiredWorkers;
    uint remainingWorkers;
    uint rewards;
    uint workerDeposite;
    uint requesterDeposite;
    address[] workerAddress;
    address requesterAddress;

    constructor(uint workers, uint taskReward, string memory taskDescription) public {
        requiredWorkers = workers;
        rewards = taskReward;
        description = taskDescription;
        workerDeposite = taskReward;
        requesterDeposite = workers*taskReward;
        requesterAddress = msg.sender;
    }

    function participant() public {
        remainingWorkers--;
        workerAddress.push(msg.sender);
    }

    function rewarding() public {
        
    }

    function evaluation() private {
        
    }

}