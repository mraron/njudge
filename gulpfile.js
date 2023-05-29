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
    return src("static_src/css/*.css").pipe(src('node_modules/bootstrap/dist/css/bootstrap.min.css')).pipe(purgeCSS({
        content: ['templates/*.gohtml', 'templates/**/*.gohtml'],
        safelist: {
            deep: [/^modal/]
        }
    })).pipe(cleanCSS({compatibility: 'ie8'})).pipe(concat('main.min.css')).pipe(dest("static/css"))
}

function mainJS() {
    return src("static_src/js/*.js").pipe(dest("static/js"))
}

function mainFavicon() {
    return src("static_src/favicon.ico").pipe(src("static_src/*.png")).pipe(src("static_src/site.webmanifest")).pipe(dest("static"))
}

const admin = parallel(adminCSS, adminJS)
const main = parallel(mainCSS, mainJS, mainFavicon)

exports.default = parallel(admin, main)