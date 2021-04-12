const webpack = require('webpack');
process.env.VUE_APP_VERSION = require('./package.json').version;

const config = require('./config');

module.exports = {
    pwa: {
        iconPaths: {
            favicon16: 'img/icon/favicon-16x16.png',
            favicon32: 'img/icon/favicon-32x32.png',
            appleTouchIcon: 'img/icon/apple-touch-icon.png',
        }
    },
    configureWebpack: {
        plugins: [
            new webpack.DefinePlugin({
                process: {
                    env: {
                        VUE_APP_API_HOST_URL: `'${config.apiUrl || "http://localhost:5001"}'`,
                        VUE_APP_ENV_NAME: `'${config.environnementName || "Local"}'`,
                        VUE_APP_VERSION: `'${process.env.VUE_APP_VERSION}'`,
                        VUE_APP_GRAPHQL_HTTP: `'${config.graphQlUrl || "http://localhost:8080"}'`,
                        VUE_APP_GRAPHQL_WS: `'${config.wsGraphqlUrl || "ws://localhost:8080/query"}'`,
                    },
                    devServer: {
                        port: 8081
                    }
                }
            })
        ]
    },
    transpileDependencies: ['vuetify'],
    lintOnSave: false,
};