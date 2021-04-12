const env = process.env.CURRENT_ENV;
const defaultConfig = require('./config.local');
const prodConfig = require('./config.prod');
const vercelConfig = require('./config.vercel');

let currentConfig;

switch (env) {
    case "production":
        currentConfig = prodConfig;
        break;
    case "vercel":
        currentConfig = vercelConfig;
        break;
    case "local":
        currentConfig = defaultConfig;
    default:
        currentConfig = vercelConfig;
}

console.log(currentConfig);
module.exports = currentConfig;