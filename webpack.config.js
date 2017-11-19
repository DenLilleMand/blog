var path = require('path');

module.exports = {
    entry: './client/main.js',
    output: {
	path: path.resolve(__dirname, 'dist'),
	publicPath: '/dist/',
	filename: 'bundle.js'
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
	/*alias: {
	  'vue$': 'node_modules/vue/dist/vue.js'
	},
	extensions: ['.js', '.vue' , '.json']*/
    }
}
