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
                        VUE_APP_VERSION: `'${process.env.VUE_APP_VERSION}'`
                    }
                }
            })
        ]
    },
    transpileDependencies: ['vuetify']
};