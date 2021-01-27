# How can I help?
- add new API
- suggest changes to API's information
- mark API is inactive  
- add link to API specification if missing or invalid
- add/edit library
- categorization improvements (new categories, more accurate names for categories, suggestions to eliminate a category)
- anything else, really...

# Who can add new APIs?
Anyone can add or suggest changes to an API

# What kind of APIs can be added?
Publicly available APIs of any kind (paid or free)

# Should I make changes in *.md files directly?
No, they are generated from APIs specifications located in directories inside `apis` and named after API name

So, if you want to make changes in `README.md` texts, make changes in `templates/readme.gomd`. 
If you want to modify api's data, make changes in corresponding specification `apis/[API NAME]/[API NAME].yaml`

# APIs to add
Search for issues with labels ["help wanted" and "new api"](https://github.com/apis-list/apis-list/issues?q=is%3Aissue+is%3Aopen+label%3A"new+api"+label%3A"new+api"). These list one of more APIs that should be added.

# Does API definition has schema?
Yes, it has - [schema.json](https://github.com/apis-list/apis-list/blob/main/schema.json)
