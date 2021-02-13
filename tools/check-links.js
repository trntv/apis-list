const {read} = require("./list");
const axios = require('axios');

exports.checkApiLinks = (source) => {
    const apis = read(source)
    
    apis.forEach((api) => {
        if (!api.is_active && !api.links) {
            return
        }
        
        api.links.forEach((link) => {
            if (link.url === "") {
                console.error(`empty link (${link.name}) in API (${api.name})`)
                return
            }
            
            axios.get(link.url)
                .catch(function (error) {
                    console.error(`failed to fetch link (${link.url}) in API (${api.name}) due to "${error.message}"`);
                })
        })
    })
}

exports.checkLibrariesLinks = (source) => {
    const apis = read(source)

    apis.forEach((api) => {
        if (!api.is_active || !api.libraries) {
            return
        }

        api.libraries.forEach((lib) => {
            if (lib.documentation_url === "" && lib.source_code_url === "") {
                console.error(`empty links (${lib.name}) in API (${api.name})`)
            }

            if (lib.documentation_url) {
                axios.get(lib.documentation_url)
                    .catch(function (error) {
                        if (!error.response || (error.response.status && error.response.status !== 429)) {
                            console.error(`failed to fetch documentation_url (${lib.documentation_url}) in API (${api.name}) due to "${error.message}"`);
                        }

                    })
            }

            if (lib.source_code_url) {
                axios.get(lib.source_code_url)
                    .catch(function (error) {
                        if (!error.response || (error.response.status && error.response.status !== 429)) {
                            console.error(`failed to fetch source_code_url (${lib.source_code_url}) in API (${api.name}) due to "${error.message}"`);
                        }
                    })
            }

        })
    })
}
