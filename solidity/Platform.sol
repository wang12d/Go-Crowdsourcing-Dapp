// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <0.8.0;
import "./Task.sol";

/** 
 * 一个抽象的用来管理worker和requester节点的平台，在进行众包任务时
 * worker和requester首先需要在platform上面进行注册来获取相应的token.
 */
contract Platform {

    address _platform;
    mapping (address => uint256) _reputation;
    uint _deposition;
    /**
     * 平台进行管理
     */
    constructor() public {
        _platform = msg.sender;
    }

    receive() external payable {

    }

    // 根据Solidity 0.6.2的描述进行修改 https://docs.soliditylang.org/en/v0.6.0/contracts.html#receive-ether-function. 
    // 只有requester上传押金到该智能合约
    fallback() external payable {
        _deposition += msg.value;
    }
    /**
     * 每个节点都需要向节点注册来获取相应的token
     */
    function register(address node) public {
        require(node != address(0x0), "cannot tranfer to zero address");
        _reputation[node] += 1;
    }

    /**
    * 进行任务奖励
    * 按照我们设计的协议，任务应该是由Workers直接提交到Requesters的，
    * 但是目前在这里作为测试用例，在智能合约上面实现该操作
    */
    function Rewarding(address payable worker, uint maskedRewards, uint unmaskedRewards, Task t) public {
        // 只有Requester才能调用此信息
        address payable requester = t.requester();
        require(msg.sender == requester, "Only requester can call this to workers who participants its task and has not been rewarded.");
        require(maskedRewards <= _deposition, "Not enough rewards to award.");
        t.addingFinishWorkers();
        if (maskedRewards >= 0) {
            worker.transfer(maskedRewards);
            // 如果该worker任务完成很好，那么则授予一个token给他进行奖励
            // 于是它下次也可以参与任务
            increaseReputation(worker);
            t.rewarding(unmaskedRewards);
            _deposition -= maskedRewards;
        }
        else {
            decreaseRepuation(worker);
            // 该worker并没有诚实的参与任务，需要销毁一个token
            // p.burnFrom(worker, 1);
        }
        if (t.finishedWorkers() == t.workerRequired() && t.totalRewards() >= 0) {
            requester.transfer(t.totalRewards());
        }
    }

    function platformAddress() public view returns (address) {
        return _platform;
    }

    function reputation(address _arr) public view returns (uint256) {
        return _reputation[_arr];
    }

    function increaseReputation(address _arr) public {
        _reputation[_arr] += 1;
    }

    function decreaseRepuation(address _arr) public {
        if (_reputation[_arr] > 0) {
            _reputation[_arr] -= 1;
        }
    }
}