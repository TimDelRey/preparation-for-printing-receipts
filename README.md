# Receipter

A small Go CLI tool for generating receipt sheets from an Excel workbook.

The program:
- reads the source file `work_doc_4.xlsx`
- processes calculation data from the workbook
- generates a new `Receipts.xlsx` file with separate receipt sheets

## Build

```bash
go build -o ./Receipter ./cmd/worker
```

## Run

Place `work_doc_4.xlsx` next to the binary, then run:

```bash
./Receipter
```

## Output

After execution, the following file will be created:

```text
Receipts.xlsx
```

## Project Structure

- `cmd/worker` — application entry point
- `pkg/service` — processing flow and orchestration
- `pkg/domain` — core models and receipt-building logic
- `pkg/domain/sample` — Excel templates, styles, and static data
