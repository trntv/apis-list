const prompt = require('prompt');
const colors = require("colors/safe");
const schema = require('../schema.json')
const list = require("./list");

let promptSchema = {
    properties: {
        url: {
            description: "API's homepage",
            type: "string"
        },
        ...convertSchema(schema),
        sdk: {
            properties: {
                source_code_url: {
                    description: "SDK's source code URL (VCS etc)",
                },
                documentation_url: {
                    description: "SDK's documentation URL",
                }
            },
        }
    }
}

module.exports = (source) => {
    prompt.message = colors.green("New API");
    prompt.delimiter = colors.red(" > ");
    
    return new Promise((resolve) => {
        prompt.start();

        prompt.get(promptSchema, function (err, result) {
            if (err) {
                return resolve()
            }
            
            let api = {
                name: result.name,
                description: result.description,
                categories: result.categories,
                logo: result.logo,
                type: result.type,
                is_free: result.is_free,
                links: [
                    {
                        name: "API's homepage",
                        url: result.url
                    }
                ],
                is_active: true
            }
            
            if (result.specification.url) {
                api.specification = result.specification
            }
            
            if (result.sdk.source_code_url || result.sdk.documentation_url) {
                api.libraries = [
                    {
                        name: "Library",
                        source_code_url: result.sdk.source_code_url || undefined,
                        documentation_url: result.sdk.documentation_url || undefined,
                    }
                ]
            }
            list.add(source, api);
            resolve()
        });
    })

}

function convertSchema(obj, prefix) {
    const skipped = [
        "slug",
        "is_active",
        "libraries",
        "links",
        "discussion_url"
    ]
    
    let properties = {}
    
    for (const [key, value] of Object.entries(obj.properties)) {
        if (skipped.includes(key)) {
            continue
        }
        
        let name = prefix !== undefined ?  prefix + '.' + key : key
        
        properties[name] = {
            description: value.description + (value.type === "array" ? " (Ctrl + C to finish)" : "") + (value.type === "boolean" ? " (true or false)" : ""),
            type: value.type,
            required: schema.required.includes(key),
            minItems: value.type === "array" && schema.required.includes(key) ? 1 : 0,
            properties: value.type === "object" ? convertSchema(value) : null,
        }
    }

    return properties
} 
