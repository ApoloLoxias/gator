#!/usr/bin/env bash

cfg="$HOME/.gatorconfig.json"

db_url=$(
  python3 <<EOF

import json, sys
with open("$cfg") as f:
        print(json.load(f)["db_url"])

EOF
)

goose -dir sql/schema postgres "$db_url" "$@"
