const {read} = require("./list");
const axios = require('axios').default;

const client = axios.create({
    timeout: 10000,
    headers: {
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36 OPR/74.0.3911.107'
    }
});


module.exports = async (source) => {
    try {
        await checkApiLinks(source)
        await checkLibrariesLinks(source)
        console.log("All links were checked")
    } catch (error) {
        console.error(error)
    }
}

async function checkApiLinks(source) {
    const apis = read(source)
    
    for (const api of apis) {
        if (api.is_active === false || !api.links) {
            continue
        }

        for (const link of api.links) {
            if (!link.url) {
                console.error(`empty link (${link.name}) in API (${api.name})`)
                continue
            }
            
            try {
                await client.get(link.url);
            } catch (error) {
                if (error.response && error.response.status === 429) {
                    continue
                }
                console.error(`failed to fetch link (${link.url}) in API (${api.name}) due to "${error.message}"`);
            }
        }
    }
    
    console.log("APIs links were checked")
}

async function checkLibrariesLinks(source) {
    const apis = read(source)

    for (const api of apis) {
        if (api.is_active === false || !api.libraries) {
            continue
        }
        
        for (const lib of api.libraries) {
            if (lib.documentation_url === "" && lib.source_code_url === "") {
                console.error(`empty links (${lib.name}) in API (${api.name})`)
            }

            if (lib.documentation_url) {
                try {
                    await client.get(lib.documentation_url)
                } catch (error) {
                    if (error.response && error.response.status === 429) {
                        continue
                    }
                    console.error(`failed to fetch documentation_url (${lib.documentation_url}) in API (${api.name}) due to "${error.message}"`);
                }
            }

            if (lib.source_code_url) {
                try {
                    await client.get(lib.source_code_url)
                } catch (error) {
                    if (error.response && error.response.status === 429) {
                        continue
                    }
                    console.error(`failed to fetch source_code_url (${lib.source_code_url}) in API (${api.name}) due to "${error.message}"`);
                }
            }
        }
    }

    console.log("Libraries' links were checked")
}
