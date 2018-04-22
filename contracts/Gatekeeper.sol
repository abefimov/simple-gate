pragma solidity ^0.4.21;


import "zeppelin-solidity/contracts/token/ERC20/StandardToken.sol";
import "./SNMLToken.sol";
import "zeppelin-solidity/contracts/ownership/Ownable.sol";


contract Gatekeeper is Ownable {

    using SafeMath for uint256;

    SNMLToken token;

    function Gatekeeper(address _token) public {
        token = SNMLToken(_token);
        owner = msg.sender;
    }

    uint256 public transactionAmount = 0;
    mapping(bytes32 => bool) public paid;

    event PayInTrx(address indexed from, uint256 indexed txNumber, uint256 indexed value);
    event PayOut(address indexed from, uint256 indexed txNumber, uint256 indexed value);
    event Suicide(uint block);

    function PayIn(uint256 _value) public {
        require(token.transferFrom(msg.sender, this, _value));
        transactionAmount = transactionAmount + 1;
        PayInTrx(this, transactionAmount, _value);
    }

    function Payout(address _to, uint256 _value, uint256 _txNumber) public onlyOwner{
        bytes32 txHash = keccak256(_to, _txNumber, _value);
        require(!paid[txHash]);
        require(token.transfer(_to, _value));
        paid[txHash] = true;
        PayOut(_to, _txNumber, _value);
    }

    function kill() public onlyOwner{
        uint balance = token.balanceOf(this);
        require(token.transfer(owner, balance));
        Suicide(block.timestamp);
        selfdestruct(owner);
    }

}
