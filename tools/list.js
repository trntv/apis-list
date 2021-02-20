const fs = require("fs");
const YAML = require("yaml");

function read(source) {
    const data = fs.readFileSync(source, {encoding: 'utf8'})

    return YAML.parse(data)
}

function write(source, apis) {
    for (let i = 0; i < apis.length; i++) {
        apis[i] = compact(apis[i])
    }
    fs.writeFileSync(source, YAML.stringify(apis))
}

function add(source, api) {
    let apis = read(source)
    apis.push(api)
    write(source, apis)
}

function compact(obj) {
    for (const [key, value] of Object.entries(obj)) {
        if (value === "" || value === undefined) {
            delete obj[key]
        }
        
        if (typeof value === 'object') {
            obj[key] = compact(value)
        }
    }
    
    return obj
}

exports.read = read
exports.write = write
exports.add = add
