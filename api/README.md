# PDF Tools API

A lightweight REST API for manipulating PDF files.
Built with [Go](https://golang.org) + [Echo](https://echo.labstack.com/), and powered by [qpdf](https://qpdf.sourceforge.io/).

## Features
- **Merge** multiple PDFs into one
- **Split** a PDF into single-page PDFs
- **Extract** specific page ranges
- **Reorder / Duplicate / Delete** pages
- **Rotate** selected pages

## Running the API

Build and run via Docker:

```bash
docker compose up --build
```

By default the API listens on port 8080.

## API Endpoints

All routes accept multipart/form-data with one or more PDF files.

---

`POST /pdf/merge`

Merge multiple PDF files into one.
 - Form field: files[] – at least 2 PDF files
 - Response: merged.pdf

Example:
```bash
curl -X POST http://localhost:8080/pdf/merge \
  -F "files=@doc1.pdf" \
  -F "files=@doc2.pdf" \
  -o merged.pdf
```

---

`POST /pdf/split`

Split a PDF into single-page PDFs.
Returns a ZIP archive containing one PDF per page.
 - Form field: file
 - Response: split.zip

Example:
```bash
curl -X POST http://localhost:8080/pdf/split \
  -F "file=@input.pdf" \
  -o split.zip
```

---

`POST /pdf/extract`

Extract specific page ranges into a new PDF.

Example:
```bash
curl -X POST "http://localhost:8080/pdf/extract?ranges=1-3,5" \
  -F "file=@input.pdf" \
  -o extracted.pdf
```

---

`POST /pdf/reorder`

Reorder, duplicate, or remove pages.
 - Query param: order (e.g. 3,1,1,4-7)
 - Form field: file
 - Response: reordered.pdf

Example:
```bash
curl -X POST "http://localhost:8080/pdf/reorder?order=3,1,1,4-7" \
  -F "file=@input.pdf" \
  -o reordered.pdf
```

---

`POST /pdf/rotate`

Rotate selected pages.
 - Query param:
   - angle – 90, 180, 270, +90, -90
   - pages – page ranges (e.g. 1,3-5) or all
 - Form field: file
 - Response: rotated.pdf

Example:
```bash
curl -X POST "http://localhost:8080/pdf/rotate?angle=90&pages=1-2" \
  -F "file=@input.pdf" \
  -o rotated.pdf
```

---

Notes

 - All operations are performed using temporary directories; no files are persisted on the server.
 - Large inputs are subject to timeout (default: 2 minutes).
 - Maximum number of files for merging: 50.
