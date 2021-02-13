# How to add my API?
1. Add information about your in [`apis-list.yaml`](https://github.com/apis-list/apis-list/blob/main/apis-list.yaml) file
2. Run `npm run build`
3. Make PR with changes

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
No, they are generated from APIs database in `apis-list.yaml` file

So, if you want to make changes in `README.md` texts, make changes in [`README.handlebars`](https://github.com/apis-list/apis-list/blob/main/README.handlebars)
If you want to make changes in API's template change [`api.handlebars`](https://github.com/apis-list/apis-list/blob/main/api.handlebars) file
If you want to modify API's information, make changes in corresponding specification in [`apis-list.yaml`](https://github.com/apis-list/apis-list/blob/main/apis-list.yaml) file

# APIs to add
Search for issues with labels ["help wanted" and "new api"](https://github.com/apis-list/apis-list/issues?q=is%3Aissue+is%3Aopen+label%3A"new+api"+label%3A"new+api"). These list one of more APIs that should be added.

# Does API definition has schema?
Yes, it has - [schema.json](https://github.com/apis-list/apis-list/blob/main/schema.json)
