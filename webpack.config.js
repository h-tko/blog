var path = require('path');

module.exports = {
    entry: {
        "bundle": path.join(__dirname, 'assets/js/app.js'),
    },
    output: {
        path: path.join(__dirname, 'assets/dist'),
        filename: "[name].js"
    },
    cache: true,
    module: {
        loaders: [{
            test: /\.css$/, loader: 'style-loader!css-loader'
        }, {
            test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/, loader: 'url-loader?limit=10000&mimetype=application/font-woff'
        }, {
            test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/, loader: 'file-loader',
        }, {
            test: /\.js$/,
            exclude: /node_modules/,
            loader: 'babel',
            query: {
                presets: ['es2015']
            }
        }, {
            test: /\.js$/,
            exclude: /node_modules/,
            loader: 'eslint-loader'
        }]
    },
    eslint: {
        configFile: './.eslintrc'
    },
    resolve: {
        extensions: ['', '.js']
    },
    devtool: 'source-map'
};
