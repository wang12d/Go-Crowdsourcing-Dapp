const Platform = artifacts.require("Platform");
const Task = artifacts.require("Task");

module.exports = async function (deployer, network, accounts) {
    await deployer.deploy(Task);
    await deployer.deploy(Platform);
};