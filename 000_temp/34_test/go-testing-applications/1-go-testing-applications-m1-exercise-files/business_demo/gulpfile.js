var gulp = require('gulp');
var less = require('gulp-less');
var path = require('path');
var shell = require('gulp-shell');

var lessPath = './less/**/*.less';
var goPath = 'src/poms/**/*.go';

gulp.task('less', function () {
  return gulp.src('less/app.less')
    .pipe(less({
      paths: [ path.join(__dirname, 'less') ]
    }))
    .pipe(gulp.dest('./res'));
});

gulp.task('compilepkg', function() {
  return gulp.src(goPath, {read: false})
    .pipe(shell(['go install <%= stripPath(file.path) %>'],
      {
          templateData: {
            stripPath: function(filePath) {
              var subPath = filePath.substring(process.cwd().length + 5);
              var pkg = subPath.substring(0, subPath.lastIndexOf(path.sep));
              return pkg;
            }
          }
      })
    );
});

gulp.task('watch', function() {
  gulp.watch(lessPath, ['less']);
  gulp.watch(goPath, ['compilepkg']);
});
