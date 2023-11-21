#!/usr/bin/env bash

# https://github.com/olekukonko/tablewriter

wget -qO table.go https://raw.githubusercontent.com/olekukonko/tablewriter/master/table.go
wget -qO table_unicode.go https://raw.githubusercontent.com/olekukonko/tablewriter/master/table_unicode.go
wget -qO table_color.go https://raw.githubusercontent.com/olekukonko/tablewriter/master/table_with_color.go
wget -qO util.go https://raw.githubusercontent.com/olekukonko/tablewriter/master/util.go
wget -qO wrap.go https://raw.githubusercontent.com/olekukonko/tablewriter/master/wrap.go

sed -i 's/package tablewriter/package tablew/g' ./*.go

sed -i '1,8d' table.go
sed -i '1,7d' util.go
sed -i '1,7d' wrap.go
