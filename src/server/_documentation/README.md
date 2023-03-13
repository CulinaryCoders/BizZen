Using gomarkdoc package to create markdown files for this project's documentation.

Example usage (generate markdown documentation for all files in models directory):

```
gomarkdoc .\models\ > .\_documentation\models-package.md
```

Additional formatting tweaks applied using 'Replace Rules' extension in VS code to make markdown files more readable.

Added the following to `settings.json` file in VS Code after installing the 'Replace Rules' extension.

```
"replacerules.rules": {

        "gomarkdoc - Reformat title heading":{
            "find": "^# (.*)$",
            "replace":"# '*$1*' Package",
            "languages": ["markdown"]
        },
        "gomarkdoc - Reformat index heading":{
            "find": "^## Index$",
            "replace":"## Index\n----",
            "languages": ["markdown"]
        },
        "gomarkdoc - Reformat type headings":{
            "find": "^## type (.*)$",
            "replace":"<br />\n\n----\n\n## **type $1**\n\n<br />\n\n",
            "languages": ["markdown"]
        },
        "gomarkdoc - Reformat func headings":{
            "find": "^### func (.*)$",
            "replace":"<br />\n\n----\n\n### **func $1**",
            "languages": ["markdown"]
        },
        "gomarkdoc - Unescape markdown italics":{
            "find": "\\\\\\*(.*)\\\\\\*",
            "replace":"<br />\n\n***$1***",
            "languages": ["markdown"]
        },
        "gomarkdoc - Pad whitespace for description subheadings":{
            "find": "^\\*\\*\\*Description\\*\\*\\*\\n((.*\\n)*?)^\\*\\*\\*",
            "replace":"<pre>\n***Description***\n$1\n</pre>\n***",
            "languages": ["markdown"]
        }
    },
    "replacerules.rulesets": {
        "gomarkdoc - Reformat documentation markdown":{
            "rules": ["gomarkdoc - Reformat title heading",
                      "gomarkdoc - Reformat index heading", 
                      "gomarkdoc - Reformat type headings", 
                      "gomarkdoc - Reformat func headings",
                      "gomarkdoc - Unescape markdown italics"]
        }
    }
```

To reformat the particular markdown file, press `CTRL+SHIFT+P` to open up the command pallette in VS Code. Then select the `Replace Rules: Run Ruleset` option and select the `gomarkdoc - Reformat documentation markdown` ruleset (defined in `settings.json` above).