# gf - go find ( not girlfriend )
## gf is a tool for searching for files and directories with nice features like regex.

## Syntax:
```$ gf [options ...] path```
## Options:
| Flag | Usage |
| - | - |
| `-d` | Set the max depth of file checking. Default is 3. |
| `-n` | Search for files with the given phrase in name. |
| `-t` | Specify if should search for files or directories. |
| `-e` | Search for files with a given extension. |
| `-s` | Sort files by name, size, or modification time. |
| `-x` | Search name with regular expressions. |
### Take a look ( they are nicely bold and blue and nice and cool ):
```
julian@archlinux ~/projects $ gf -x "[a-c]" -t f -s size -d 3 .
project/
   ├-  script.py
   ├-┬ .py_env/
     └-  config.py
   ├-┬ logs/
     └-  app.log
   ├-┬ assets/
     ├-  script.js
     └-  style.css
   ├-┬ docs/
     └-  tutorial.md
   ├-┬ src/
     └-  main.py
   └-┬ tests/
     └-  test_main.py
```
