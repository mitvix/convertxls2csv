# convertxls2csv
Simple program in Go to convert XLS files to CSV using excelize

# Usage:

```
git clone https://github.com/mitvix/convertxls2csv
cd convertxls2csv

go mod init convertxls2csv
go get github.com/xuri/excelize/v2
go build -o convertxls2csv main.go

./convertxls2csv <file.xlsx>
```

