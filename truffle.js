module.exports = {
    networks: {
        development: {
            host: 'localhost',
            port: 8545,
            network_id: '*'
        },
        private: {
            host: 'localhost',
            port: 9045,
            network_id: '444'
        },
        rinkeby: {
            host: 'localhost',
            port: 8545,
            network_id: '*'
        },
    },
    solc: {
        optimizer: {
            enabled: true,
            runs: 200
        }
    }
};
