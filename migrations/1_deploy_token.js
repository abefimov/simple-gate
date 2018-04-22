var SNMLToken = artifacts.require("./SNMLToken.sol");

module.exports = function (deployer, network) {
    if (network === "private") {
        deployer.deploy(SNMLToken, {gasPrice: 0})
    } else if (network === "rinkeby") { } else {
        deployer.deploy(SNMLToken)
    }
};
