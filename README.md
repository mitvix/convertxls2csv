# convertxls2csv
Simple program in Go to convert XLS files to CSV using excelize

# Usage:

```
git clone https://github.com/mitvix/convertxls2csv
cd convertxls2csv

go mod init convertxls2csv
go get github.com/xuri/excelize/v2
go build -o convertxls2csv main.go

# <optional>
sudo cp convertxls2csv /usr/local/bin/

# To convert:
./convertxls2csv <file.xlsx>
```

# Result
<img width="698" height="109" alt="image" src="https://github.com/user-attachments/assets/f721d678-be9d-4555-897c-a52d0ca48733" />
