pragma solidity ^0.4.18;


import "zeppelin-solidity/contracts/token/ERC20/StandardToken.sol";
import "zeppelin-solidity/contracts/ownership/Ownable.sol";


contract SNMLToken is StandardToken, Ownable {

    using SafeMath for uint256;

    string public name = "SONM sidechain test token";

    string public symbol = "SNML";

    uint public decimals = 18;

    event GiveAway(address indexed whom, uint amount);

    function SNMLToken() public{
        owner = msg.sender;
        totalSupply_ = 444 * 1e6 * 1e18;
        balances[msg.sender] = totalSupply_;
    }

}
