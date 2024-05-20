const { parallel, src, dest } = require('gulp');
const purgeCSS = require('gulp-purgecss')
const cleanCSS = require('gulp-clean-css');
const concat = require('gulp-concat');

function adminCSS() {
    return src('node_modules/ng-admin/build/ng-admin.min.css', { sourcemaps: true }).pipe(dest("static/css", { sourcemaps: '.' }))
}

function adminJS() {
    return src('node_modules/ng-admin/build/ng-admin.min.js', { sourcemaps: true }).pipe(dest("static/js", { sourcemaps: '.' }))
}

function mainCSS() {
    return src("src/css/*.css").pipe(src('node_modules/bootstrap/dist/css/bootstrap.min.css')).pipe(purgeCSS({
        content: [
            'internal/web/templates/*.gohtml',
            'internal/web/templates/**/*.gohtml'
        ],
        safelist: {
            deep: [/^modal/]
        }
    })).pipe(src('node_modules/select2/dist/css/select2.min.css'))
       .pipe(src('node_modules/@ttskch/select2-bootstrap4-theme/dist/select2-bootstrap4.min.css'))
       .pipe(cleanCSS({compatibility: 'ie8'})).pipe(concat('main.min.css')).pipe(dest("static/css"))
}

function mainJS() {
    return src("src/js/*.js").pipe(dest("static/js"))
}

function bootstrapJS() {
    return src('node_modules/bootstrap/dist/js/bootstrap.min.js', {sourcemaps: true}).pipe(dest('static/js', {sourcemaps: '.'}))
}

function bootstrapIcons() {
    return src('node_modules/bootstrap-icons/**/*').pipe(dest('static/bootstrap-icons'))
}

function selectJS() {
    return src('node_modules/select2/dist/js/select2.min.js', {sourcemaps: true}).pipe(dest('static/js', {sourcemaps: '.'}))
}

function jqueryJS() {
    return src('node_modules/jquery/dist/jquery.slim.min.js', {sourcemaps: true}).pipe(dest('static/js', {sourcemaps: '.'}))
}

function katex() {
    return src('node_modules/katex/dist/**/*').pipe(dest('static/katex'))
}

function mainFavicon() {
    return src("src/favicon.ico").pipe(src("src/*.png")).pipe(src("src/site.webmanifest")).pipe(dest("static"))
}

const admin = parallel(adminCSS, adminJS)
const main = parallel(mainCSS, mainJS, mainFavicon, bootstrapJS, selectJS, jqueryJS, katex, bootstrapIcons)

exports.default = parallel(admin, main)