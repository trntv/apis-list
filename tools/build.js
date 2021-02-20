const fs = require("fs");
const Handlebars = require("handlebars")
const slugify = require("slugify");
const {read} = require("./list");
const {categories, sortCategories} = require("./categories") 

module.exports = (source, destination) => {
    const apis = read(source)
    
    Handlebars.registerHelper('short', function (s) {
        if (!s || s.length <= 50) {
            return s
        }

        s = s.replace(/\n/g, ". ")
        s = s.split(/[.!?]/g)[0]

        return s.trim()
    })

    Handlebars.registerHelper('slug', function (s) {
        s = s.replace(/[^a-zA-Z0-9 ]/g, "")
        s = s.replace(/ /g, "-")
        
        return s.toLowerCase()
    })
    
    Handlebars.registerHelper('category_icon', (s) => {
        return categories[s] || "📃"
    })
    
    const categoriesNames = []
    const graveyard = []
    const byCategory = {}

    const apiTemplate = Handlebars.compile(fs.readFileSync(__dirname + "/../api.handlebars").toString());
    apis.forEach((api) => {
        api.slug = api.slug || slugify(api.name).toLowerCase()
        
        if (api.is_active === false) {
            graveyard.push(api)
            return
        }

        api.categories.sort(sortCategories)
        api.categories.forEach((c) => {
            byCategory[c] = byCategory[c] || []
            byCategory[c].push(api)
            if (!categoriesNames.includes(c)) {
                categoriesNames.push(c)
            }
        })

        const byPlatform = {}
        
        if (api.libraries) {
            api.libraries.forEach((l) => {
                byPlatform[l.platform] = byPlatform[l.platform] || []
                byPlatform[l.platform].push(l)
            })
        }
        
        let apiData = apiTemplate({
            api: api,
            editorLink: 'https://github.com/apis-list/apis-list/edit/main/apis-list.yaml',
            libraries: byPlatform
        })

        const dir = `${destination}/apis/${api.slug}`
        
        if (!fs.existsSync(dir)){
            fs.mkdirSync(dir);
        }
        
        fs.writeFileSync(`${dir}/${api.slug}.md`, apiData)
    })

    const readmeTemplate = Handlebars.compile(fs.readFileSync(__dirname + "/../README.handlebars").toString());
    const readme = readmeTemplate({
        categoriesNames: categoriesNames.sort(sortCategories),
        apis:            byCategory,
        Graveyard:       graveyard,
    })
    
    fs.writeFileSync(destination + "/README.md", readme)
}
