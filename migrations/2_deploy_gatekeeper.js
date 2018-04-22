var Gatekeeper = artifacts.require("./Gatekeeper.sol");
var SNMLToken = artifacts.require("./SNMLToken.sol");

module.exports = function(deployer, network) {
    if (network === "private"){
        deployer.deploy(Gatekeeper, SNMLToken.address, {gasPrice: 0});
    }else if (network === "rinkeby") {
        deployer.deploy(Gatekeeper, "0x06bda3cf79946e8b32a0bb6a3daa174b577c55b5");
    }else{
        deployer.deploy(Gatekeeper, SNMLToken.address);
    }
};
