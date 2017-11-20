var webpack = require('webpack');
var WebpackDevServer = require('webpack-dev-server');
var config = require('./webpack.config.js');
var host = 'localhost';
var port = 3001;

new WebpackDevServer(webpack(config), {
    publicPath: config.output.publicPath,
    hot: true,
    historyApiFallback: true,
    proxy: {
        "**": "http://localhost:3000",
    }
}).listen(port, host, (err, result) => {
    if(err) {
        console.log("Webpack dev server err:", err);
    }
    console.log('Webpack dev server listening on port:' + port + ' . Host:' + host);
});
