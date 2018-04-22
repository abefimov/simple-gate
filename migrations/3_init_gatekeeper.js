var Gatekeeper = artifacts.require("./Gatekeeper.sol");
var SNMLToken = artifacts.require("./SNMLToken.sol");

module.exports = function(deployer, network) {
    if (network === "private"){
        SNMLToken.deployed()
            .then(function (res) {
                res.transfer(Gatekeeper.address, 444*1e6*1e18, {gasPrice: 0});
            })
    }else if (network === "rinkeby"){

    }else{
        SNMLToken.deployed()
            .then(function (res) {
                res.transfer(Gatekeeper.address, 444*1e6*1e18);
            })
    }


};
