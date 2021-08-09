// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <0.8.0;

/** 
 * 一个抽象的用来管理worker和requester节点的平台，在进行众包任务时
 * worker和requester首先需要在platform上面进行注册来获取相应的token.
 */
contract Platform {

    address _platform;
    mapping (address => uint256) _reputation;
    /**
     * 平台进行管理
     */
    constructor() public {
        _platform = msg.sender;
    }

    /**
     * 每个节点都需要向节点注册来获取相应的token
     */
    function register(address node) public {
        require(node != address(0x0), "cannot tranfer to zero address");
        _reputation[node] += 1;
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
        _reputation[_arr] += 1;
    }
}