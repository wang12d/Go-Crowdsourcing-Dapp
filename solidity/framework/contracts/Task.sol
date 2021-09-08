// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <0.8.0;

// import "@openzeppelin/contracts/token/ERC20/ERC20Burnable.sol";
// import "./Platform.sol";

// @title 移动众包-智能合约
// 主要负责处理Workers和Requesters之间的关系，它需要
// 进行：
//      1. 负责审核Workers参与任务的请求
//      2. 当Requesters创建任务时，保存Requesters所需要的押金
contract Task {
    // 对于一个任务，剩下的Workers数量
    uint _workerRequired;
    // 计算每个Worker应该获得的奖励
    uint _totalRewards;
    address[] _workerAddresses;
    string _description;
    address _requester;
    uint _deposition;
    uint _remainingWorkers;
    uint _finishedWorkers;
    bytes[] _data;
    // uint256 _reputation;
     
    // 负责在特定条件下用来进行交互的Event，进行交易的触发
    // 假若Requester完成了任务并进行了评估，那么它就会触发一个交易来
    // 奖励Workers，同时Workers也可以通过监听该信息来判断其
    // 是否收到了相应的奖励
    event TaskPublished(uint rewards, address indexed _requester, string _description);
    event TaskRegistered(address indexed _worker, address indexed _requester, uint _remainingWorkers);
    event DataSubmitted(address indexed _from, bytes _data);

    /** 为了保证该Smart Contract能够接受来自其他用户的押金
    * 需要实现Fallback Function. 
    * Requester需要将押金从传递到智能合约的地址
    * 后期为了进行权限控制，需要判断消息发送者是否真的为Requesters
    * Workers需要将押金从传递到智能合约的地址 
    * 后期为了进行权限控制，需要判断消息发送者是否真的为Workers
    */
    constructor(uint workerRequired, uint totalRewards, string memory description) public  {
        _workerRequired = workerRequired;
        _totalRewards = totalRewards;
        _description = description;
        _requester = msg.sender;
        _remainingWorkers = workerRequired;
        _deposition = _totalRewards;
        // _reputation = reputation;
    }

    receive() external payable {

    }

    // 根据Solidity 0.6.2的描述进行修改 https://docs.soliditylang.org/en/v0.6.0/contracts.html#receive-ether-function.
    // 只有requester上传押金到该智能合约
    fallback() external payable {
        require(_requester == msg.sender, "Only requester can upload transaction to this contract");
        _deposition += msg.value;
    }
   
    /** 将该众包任务发布到链上，所有成员都可以选择是否参加该任务
    *  该过程会触发一个event，链上节点看到该event可以自由选择是
    *  否参加该任务.     
    */
    function publish() public {
        // 对任务的信息进行记录，包括保存任务的押金，记录任务需要的人数
        require (_requester == msg.sender, "Not enough amount to pay the collaterals");
        require (_totalRewards <= _deposition, "The deposition must larger than rewards to workers to pubulish task.");
        // 在requester使用自己的伪名进行任务发布以后，requester能够获得token的分发
        // 并且在workers完成任务的时候才能够获得token作为奖励
        _totalRewards = _deposition;
        emit TaskPublished(_totalRewards, _requester, _description);
    }
    // 给定一个地址，查看其所需要的Workers数量是否已满
    function remainingWorkers() public view returns (uint rem) {
        rem = _remainingWorkers;
    }
    /** Workers在链上看到该任务就可以参与该任务
    * 但是workers在注册时会被基于一个不记名的token，该token时一个
    * 参与的凭证
    */
    function register() public {
        require(_remainingWorkers > 0, "The task do not need workers anymore");
        // require(p.reputation(msg.sender) >= _reputation, "Not enough reputation to participant the task");
        address worker = msg.sender;
        // require (balanceOf(worker) > 0, "Only worker with token can participant the task");
        // 此时任务算是被workers接受了
        _remainingWorkers--;
        // 保存所有用户地址
        _workerAddresses.push(worker);
        emit TaskRegistered(worker, _requester, _remainingWorkers);
    }

    /** Workers进行任务上传
    * 它应该上传的是一个加密的数据包，或者当数据直接在区块链上面
    * 进行传递时不划算了，那么则应该上传相应对应数据块的索引，这样
    * 相应Requester就能够找到并对其进行访问
    */
    function SubmitData(bytes memory data) public {
        // 只有注册了的worker才可以上传数据
        uint currentWorkers = _workerRequired - _remainingWorkers;
        bool inWorkers = false;
        for (uint i = 0; i < currentWorkers; i++) {
            if (_workerAddresses[i] == msg.sender) {
                inWorkers = true; break;
            }
        }
        _data.push(data);
        require(inWorkers, "Only registered worker can submit data");
        emit DataSubmitted(msg.sender, data); // Requester by subscription to track the process of the current task
    }
    /**
    * 进行任务奖励
    * 按照我们设计的协议，任务应该是由Workers直接提交到Requesters的，
    * 但是目前在这里作为测试用例，在智能合约上面实现该操作
    */
    // function Rewarding(address payable worker, uint rewards, Platform p) public {
    //     // 只有Requester才能调用此信息
    //     require(msg.sender == _requester, "Only requester can call this to workers who participants its task and has not been rewarded.");
    //     require(_totalRewards >= rewards, "Not enough rewards to award.");
    //     _finishedWorkers += 1;
    //     if (rewards >= 0) {
    //         worker.transfer(rewards);
    //         // 如果该worker任务完成很好，那么则授予一个token给他进行奖励
    //         // 于是它下次也可以参与任务
    //         p.increaseReputation(worker);
    //         _totalRewards -= rewards;
    //     }
    //     else {
    //         p.decreaseRepuation(worker);
    //         // 该worker并没有诚实的参与任务，需要销毁一个token
    //         // p.burnFrom(worker, 1);
    //     }
    //     if (_finishedWorkers == _workerRequired && _totalRewards >= 0) {
    //         msg.sender.transfer(_totalRewards);
    //     }
    // }

    function addingFinishWorkers() public {
        _finishedWorkers += 1;
    }

    function rewarding(uint rewards) public {
        _totalRewards -= rewards;
    }

    function requester() public view returns (address payable r) {
        r = payable(_requester);
    }

    function finishedWorkers() public view returns (uint) {
        return _finishedWorkers;
    }

    function workerRequired() public view returns (uint) {
        return _workerRequired;
    }

    function totalRewards() public view returns (uint) {
        return _totalRewards;
    }

    // 辅助函数，用来进行从uint到string的转化，这一步主要是帮助Contracts进行任务的转化
    function uintToString(uint v) internal pure returns (string memory str) {
        uint maxlength = 100;
        bytes memory reversed = new bytes(maxlength);
        uint i = 0;
        while (v != 0) {
            uint remainder = v % 10;
            v = v / 10;
            reversed[i++] = byte(uint8(48 + remainder));
        }
        bytes memory s = new bytes(i);
        for (uint j = 0; j < i; j++) {
            s[j] = reversed[i - 1 - j];
        }
        str = string(s);
    }

    function addString(string memory a, string memory b) internal pure returns (string memory) {
        return string(abi.encodePacked(a, b));
    }
    // 将字符串转化为固定长度的Bytes32，这样客户端可以很方便的将Event重新转化为字符串再进行描述
    function stringToBytes32(string memory source) internal pure returns (bytes32 result) {
        bytes memory tempEmptyStringTest = bytes(source);
        if (tempEmptyStringTest.length == 0) {
            return 0x0;
        }
        assembly {
            result := mload(add(source, 32))
        }
    }
}
