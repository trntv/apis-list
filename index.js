const { Command } = require('commander');
const program = new Command();

program.version('0.1.0');

const build = require('./tools/build');
const { checkApiLinks, checkLibrariesLinks } = require('./tools/check-links');

const defaultSource = __dirname + '/apis-list.yaml';
const defaultDestination = __dirname;

program
    .command('build [source] [destination]')
    .description('build apis files from database')
    .action((source, destination) => build(source || defaultSource, destination || defaultDestination));

program
    .command('check-links [source]')
    .action((source) => {
        checkApiLinks(source || defaultSource)
        checkLibrariesLinks(source || defaultSource)
    });

program
    .command('check-orphans [source]')
    .action((source) => require('./tools/check-orphans')(source || defaultSource));

program
    .command('lib [source]')
    .action(async source => await require('./tools/lib')(source || defaultSource));

program
    .command('add [source]')
    .action(async source => await require('./tools/add')(source || defaultSource));

program.parseAsync(process.argv)
    .then(() => {
        console.log("Goobye!");
    })
