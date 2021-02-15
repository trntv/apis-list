const {read} = require("./list");
const fs = require("fs");
const path = require("path");

module.exports = (source) => {
    const apis = read(source)
    const slugs = new Map()
    for (let api of apis) {
        slugs.set(api.slug, true)
    }
    
    let apisDir = path.dirname(source) + '/apis'
    let dirs = fs.readdirSync(apisDir)
    for (let slug of dirs) {
        if (slug.startsWith('.')) {
            continue
        }

        if (slugs.has(slug)) {
            continue
        }

        let apiDir = apisDir + "/" + slug
        let dirty = false
        let files = fs.readdirSync(apiDir)
        for (let file of files) {
            if (!file.endsWith(".md")) {
                dirty = true
                break
            }
        }
        
        if (!dirty) {
            fs.rmdirSync(apiDir, { recursive: true });
            console.log("removed directory: " + apiDir)
        } else {
            console.log("orphan couldn't be deleted: " + apiDir)
        }
    }
}
