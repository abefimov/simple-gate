var Gatekeeper = artifacts.require("./Gatekeeper.sol");
var SNM = artifacts.require("./SNM.sol");

module.exports = function(deployer, network) {
    if (network === "private"){
        deployer.deploy(Gatekeeper, SNM.address, {gasPrice: 0});
    }else if (network === "rinkeby") {
        deployer.deploy(Gatekeeper, "0x06bda3cf79946e8b32a0bb6a3daa174b577c55b5");
    }else{
        deployer.deploy(Gatekeeper, SNM.address);
    }
};
