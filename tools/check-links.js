const YAML = require('yaml')
const fs = require('fs')
const axios = require('axios');

const file = fs.readFileSync(process.argv[2], 'utf8')
const apis = YAML.parse(file)
apis.forEach(api => {
    checkLink(api.uri).then(valid => {
        if (valid === false) {
            throw new Error()
        }
    }).catch(err => {
        console.log(`invalid link: ${api.name} - ${api.uri} - ${err.message}`)
    })

    if (api.libraries in Array) {
        ali.libraries.forEach(lib => {
            checkLink(lib.source_code_uri).then(valid => {
                if (valid === false) {
                    throw new Error()
                }
            }).catch(err => {
                console.log(`invalid library link: ${api.name} - ${lib.source_code_uri} - ${err.message}`)
            })

            checkLink(lib.homepage_uri).then(valid => {
                if (valid === false) {
                    throw new Error()
                }
            }).catch(err => {
                console.log(`invalid library link: ${api.name} - ${lib.source_code_uri} - ${err.message}`)
            })
        })
    }

})

function checkLink(uri) {
    return new Promise((resolve, reject) => {
        axios.get(uri, {timeout: 60 * 1000})
            .then(response => {
                if (response.status === 200) {
                    resolve(true)
                } else {
                    reject(new Error(`invalid status code ${response.status}`))
                }
            })
            .catch(reject)
    })
}
