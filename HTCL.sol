pragma solidity ^0.8.7;

contract sendether {
    address payable public owner;
    uint256 public contractbalance; //contract value
    uint256 public amount;
    uint256 public balanceto;
    address payable recipient;
    uint256 public startTime;
    uint256 public lockTime = 10 seconds;
    string public secret; //abbbc
    bytes32 public hash =
        0x529bd57c54687cfd2bf23415e7b13d342f623c24a21387d669b7c6537020007d;

    constructor(address payable _recipient) payable {
        owner = payable(msg.sender);
        contractbalance = address(this).balance;
        recipient = _recipient;
        amount = msg.value;
        startTime = block.timestamp;
    }

    // function sendviacall()public payable{//returns(bool sent,byte memory data) {
    //     recipient.transfer(amount);
    //     balanceto=recipient.balance;
    //     contractbalance = address(this).balance;

    // }

    // withdraw money if the secret is correct

    function withdraw(string memory _secret) public payable {
        require(keccak256(abi.encodePacked(_secret)) == hash, "wrong secret");
        secret = _secret;
        recipient.transfer(amount);
    }

    // refund once Hash time is expired
    function refund() public payable {
        require(block.timestamp > startTime + lockTime, "too early");
        owner.transfer(amount);
    }
}
