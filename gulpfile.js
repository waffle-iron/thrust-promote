var gulp = require('gulp');
var less = require('gulp-less');
var webpack = require('webpack');

var config = {
    context: __dirname + "/public/js",
    entry: "./app.js",
    output: {
        path: __dirname + "/public/js",
        filename: "bundle.js"
    },
    module: {
        loaders: [
            { test: /\.js?$/, loaders: ['babel'], exclude: /node_modules/ },
            { test: /\.js$/, exclude: /node_modules/, loader: 'babel-loader'},
        ]
    },
};

gulp.task('scripts', function(done) {
  webpack(config).run(function(err, stats) {
    if(err) {
      console.log('Error', err);
    }
    else {
      console.log(stats.toString());
    }
    done();
  });
});

gulp.task('styles', function() {
  return gulp.src('public/less/app.less')
    .pipe(less())
    .pipe(gulp.dest('public/css'));
});

gulp.task('default', ['scripts', 'styles']);
