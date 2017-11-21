var path = require('path');

module.exports = {
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
	    'vue$': 'vue/dist/vue.esm.js'
	}
    },
    devServer: {
	hot: true,
	historyApiFallback: false,
	compress: true,
	quiet: false,
	port: 3001,
	stats: { colors: true },
	proxy: {
	    "**":"http://localhost:3000"
	}

    }
};
