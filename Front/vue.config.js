const webpack = require('webpack');
process.env.VUE_APP_VERSION = require('./package.json').version;

const config = require('./config');

module.exports = {
    configureWebpack: {
        plugins: [
            new webpack.DefinePlugin({
                process: {
                    env: {
                        VUE_APP_API_HOST_URL: `'${config.apiUrl || "http://localhost:5001"}'`,
                        VUE_APP_ENV_NAME: `'${config.environnementName || "Local"}'`,
                        VUE_APP_VERSION: `'${process.env.VUE_APP_VERSION}'`,
                        VUE_APP_GRAPHQL_HTTP:  `'${config.graphqlUrl || "http://localhost:8080"}'`,
                        VUE_APP_GRAPHQL_WS: `'${config.wsGraphqlUrl || "ws://localhost:8080/query"}'`,
                    },
                    devServer: {
                        port: 8081
                    }
                }
            })
        ]
    },
    transpileDependencies: ['vuetify']
};