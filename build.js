const esbuild = require('esbuild');

const watch = process.argv.includes('--watch');

const config = {
    entryPoints: [
        'assets/main.js',
        'assets/style.css',
    ],
    bundle: true,
    minify: !watch,
    sourcemap: watch,
    outdir: '/static',
    format: 'esm',
    splitting: !watch,
    loader: {
        '.js': 'js',
        '.css': 'css',
    },
    absWorkingDir: process.cwd()
};

if (watch) {
    esbuild.context(config).then(ctx => {
        ctx.watch();
        console.log('Watching for changes...');
    });
} else {
    esbuild.build(config).catch(() => process.exit(1));
}
