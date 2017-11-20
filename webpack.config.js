var WebpackDevServer = require("webpack-dev-server");
var webpack = require("webpack");
var path = require("path");

var compiler = webpack({
    entry: {
	index: './client/index/index.js',
	404: './client/404/404.js'
    },
    output: {
	path: path.resolve(__dirname, 'dist'),
	publicPath: '/dist/',
	filename: '[name].bundle.js'
    },
    module: {
	loaders: [{
	    test: /\.vue$/,
	    loader: 'vue-loader'
	}, {
	    test: /\.js/,
	    loader: 'babel-loader'
	}]
    },
    resolve: {
	modules: [
	    path.resolve('./client'),
	    path.resolve('./node_modules')
	],
	alias: {
	    vue: 'vue/dist/vue.js'
	}
    }
});

var server = new WebpackDevServer(compiler, {
    contentBase: "/home/denlillemand/go/src/github.com/denlillemand/blog/dist",
    hot: true,
    historyApiFallback: false,
    compress: true,
    proxy: {
	"**": "http://localhost:3000"
    },
    clientLogLevel: "info",
    quiet: false,
    noInfo: false,
    lazy: true,
    //filename:"bundle.js",
    watchOptions: {
	aggregateTimeout: 300,
	poll: 1000
    },
    publicPath: "/dist/",
    headers: { "X-Custom-Header":"yes"},
    stats: { colors: true },
});
server.listen(3001, "localhost", function() {});
