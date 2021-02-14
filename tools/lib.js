const {read, write} = require("./list");
const axios = require('axios');
const url = require('url');
const github_token = process.env["GITHUB_TOKEN"]

module.exports = async (source) => {
    const apis = read(source)
    
    for (const i in apis) {
        console.log(`processing ${i} of ${apis.length}`)
        
        let api = apis[i]
        
        if (!api.libraries) {
            continue
        }

        
        for (const k in api.libraries) {
            let lib = api.libraries[k]
            
            if (lib.source_code_url) {
                lib = await enrich(lib, lib.source_code_url)
            }

            if (!lib.author && lib.stargazers_count === undefined && lib.documentation_url) {
                lib = await enrich(lib, lib.documentation_url)
            }

            api.libraries[k] = lib
        }

        apis[i] = api
    }
    
    write(source, apis)
}

async function enrich(lib, libUrl) {
    let u = new url.URL(libUrl)
    if (u.host !== "github.com" && u.host !== "www.github.com") {
        return lib
    }

    let owner = resolveOwner(u)
    let repo = resolveRepo(u)
    
    try {
        let info = await getInfo(owner, repo)
        if (info.owner.type === "User") {
            lib.author = info.owner.login
        }

        lib.stargazers_count = info.stargazers_count
    } catch (e) {
        console.error(e.message)
    }
    
    return lib
}

function resolveOwner(u) {
    return u.pathname.split("/")[1]
}

function resolveRepo(u) {
    return u.pathname.split("/")[2]
}

async function getInfo(owner, repo) {
    let resp = await axios.get(`https://api.github.com/repos/${owner}/${repo}`, {
        headers: {Authorization: `token ${github_token}`}
    })
    
    return resp.data
}
