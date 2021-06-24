// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20Burnable.sol";

/** 
 * 一个抽象的用来管理worker和requester节点的平台，在进行众包任务时
 * worker和requester首先需要在platform上面进行注册来获取相应的token.
 */
contract Platform is ERC20Burnable {

    address _platform;
    /**
     * 平台进行管理
     */
    constructor() public ERC20("Ticket", "TT") {
        _setupDecimals(0);
        _platform = msg.sender;
        _mint(_platform, 10000);
    }

    /**
     * 每个节点都需要向节点注册来获取相应的token
     */
    function register(address node) public {
        require(node != address(0x0), "cannot tranfer to zero address");
        transfer(node, 1);
    }

    function platformAddress() public view returns (address) {
        return _platform;
    }
}