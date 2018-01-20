var path = require('path');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
    entry: {
	index: './client/index/index.js',
	404: './client/404/404.js',
	math: './client/math/math.js'
    },
    output: {
	path: path.resolve(__dirname, 'dist'),
	publicPath: '/static/',
	filename: '[name].bundle.js'
    },
    module: {
	loaders: [{
	    test: /\.vue$/,
	    loader: 'vue-loader'
	}, {
	    test: /\.js$/,
	    loader: 'babel-loader'
	}, {
	    test: /.css$/,
	    loader: ExtractTextPlugin.extract('style-loader', 'css-loader')
	}, {
	    test: /.(ttf|otf|eot|svg|woff(2)?)(\?[a-z0-9]+)?$/,
	    use: [{
	      loader: 'file-loader',
	      options: {
		name: '[name].[ext]',
		outputPath: 'fonts',
		publicPath: 'fonts'
	      }
	    }]
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
    }, 
    plugins: [
	new ExtractTextPlugin('[name].css', {
	    allChunks: true
	})
    ]
};
