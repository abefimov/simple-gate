var Gatekeeper = artifacts.require("./Gatekeeper.sol");
var SNM = artifacts.require("./SNM.sol");

module.exports = function(deployer, network) {
    if (network === "private"){
        SNM.deployed()
            .then(function (res) {
                res.transfer(Gatekeeper.address, 444*1e6*1e18, {gasPrice: 0});
            })
    }else if (network === "rinkeby"){

    }else{
        SNM.deployed()
            .then(function (res) {
                res.transfer(Gatekeeper.address, 444*1e6*1e18);
            })
    }


};
