const env = process.env.CURRENT_ENV;
const defaultConfig = require('./config.local');
const prodConfig = require('./config.prod');
const vercelConfig = require('./config.vercel');

let currentConfig;

console.log(env);

switch (env) {
    case "production":
        currentConfig = prodConfig;
        break;
    case "vercel":
        currentConfig = vercelConfig;
        break;
    case "local":
        currentConfig = defaultConfig;
        break;
    default:
        currentConfig = vercelConfig;
}

console.log(currentConfig);
module.exports = currentConfig;