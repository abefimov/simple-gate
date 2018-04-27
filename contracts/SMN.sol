pragma solidity ^0.4.21;


import "zeppelin-solidity/contracts/token/ERC20/StandardToken.sol";
import "zeppelin-solidity/contracts/ownership/Ownable.sol";


contract SNM is StandardToken, Ownable {

    using SafeMath for uint256;

    string public name = "SONM token";

    string public symbol = "SNM";

    uint public decimals = 18;

    address market;

    uint totalSupply_ = 444 * 1e6 * 1e18;

    function SNM() public{
        owner = msg.sender;

        balances[msg.sender] = totalSupply_;
    }

    function SetMarketAddress(address _newMarket) onlyOwner public {
        market = _newMarket;
    }

    //override for market
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
        if(_to == market){
            require(market != address(0));
            require(msg.sender == market);
            require(_value <= balances[_from]);

            balances[_from] = balances[_from].sub(_value);
            balances[_to] = balances[_to].add(_value);
            Transfer(_from, _to, _value);

            return true;
        } else {
            require(_to != address(0));
            require(_value <= balances[_from]);
            require(_value <= allowed[_from][msg.sender]);

            balances[_from] = balances[_from].sub(_value);
            balances[_to] = balances[_to].add(_value);
            allowed[_from][msg.sender] = allowed[_from][msg.sender].sub(_value);
            Transfer(_from, _to, _value);

            return true;
        }
    }
}
