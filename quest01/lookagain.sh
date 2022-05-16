#! bin/bash
find . -iname "*.sh" -exec basename {} .sh ';'| sort -r | tr -d './' 