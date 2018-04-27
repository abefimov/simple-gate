const Gatekeeper = artifacts.require('./Gatekeeper.sol');
const SNM = artifacts.require("./SNM.sol");

const assertRevert = require('./helpers/assertRevert.js');


contract('Gatekeeper', async function (accounts) {
    let gatekeeper;
    let token;
    let owner = accounts[0];
    let sender = accounts[1];
    let creeper = accounts[3];
    let test_value = 100;

    before(async function () {
        gatekeeper = await Gatekeeper.deployed();
        token = await SNM.deployed();

        await token.approve(gatekeeper.address, 1000, {from: sender});
        let allowance = await token.allowance(sender, gatekeeper.address);

        let balance = await token.balanceOf(gatekeeper.address);
    });

    it('test Payout', async function () {
        let start_balance = await token.balanceOf(gatekeeper.address);

        await gatekeeper.Payout(sender, 100, 1, {from: owner});

        let end_balance = await token.balanceOf(gatekeeper.address);

        assert.equal(start_balance.toNumber() - 100, end_balance.toNumber());

        try{
            await gatekeeper.Payout(sender, 100, 1, {from: owner});
        }catch (e) {
            assertRevert(e);
        }
        try{
            await gatekeeper.Payout(sender, 100, 1, {from: creeper});
        }catch (e) {
            assertRevert(e);
        }
    });

    it('test PayIn', async function () {
        let start_balance = await token.balanceOf(sender);

        let allowance = await token.allowance(sender, gatekeeper.address);
        assert.equal(allowance.toNumber() >= test_value, true);

        await gatekeeper.PayIn(test_value, {from: sender});

        let txAmount = await gatekeeper.transactionAmount.call();
        assert.equal(1, txAmount.toNumber());
    });

    it('test kill', async function () {
        try {
            await gatekeeper.kill({from: creeper});
        }catch (e) {
            assertRevert(e)
        }

        let start_owner_balance = await token.balanceOf(owner);
        let start_gk_balance = await token.balanceOf(gatekeeper.address);
        let diff = start_gk_balance.toNumber() - start_owner_balance.toNumber();

        await gatekeeper.kill({from: owner});

        let end_owner_balance = await token.balanceOf(owner);
        let end_gk_balance = await token.balanceOf(gatekeeper.address);

        assert.equal(end_gk_balance.toNumber(), 0);
        assert.equal(end_owner_balance.toNumber() - end_gk_balance.toNumber(), diff);
    });


});
