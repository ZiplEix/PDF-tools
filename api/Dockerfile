# ---- Stage 1: build Go app ---------------------------------------------------
FROM golang:1.24-bookworm AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
# Build statique pour un binaire portable
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o /app/main .

# ---- Stage 2: runtime with PDF/Image tools ----------------------------------
FROM debian:bookworm-slim

# Évite les prompts APT et réduit le bruit
ENV DEBIAN_FRONTEND=noninteractive

# IMPORTANT: regroupe les installs dans une même couche (meilleur cache)
# Si tu ajoutes/supprimes un paquet, modifie la ligne TOOL_PACKAGES ci-dessous.
ARG TOOL_PACKAGES="\
  qpdf \
  ghostscript \
  poppler-utils \
  imagemagick \
  img2pdf \
  exiftool \
  tesseract-ocr \
  tesseract-ocr-eng \
  tesseract-ocr-fra \
  tesseract-ocr-ita \
  ocrmypdf \
  pngquant \
  unpaper \
  fonts-dejavu \
  ca-certificates \
  curl \
"

# MàJ + install des outils PDF/Image
RUN apt-get update \
 && apt-get install -y --no-install-recommends $TOOL_PACKAGES \
 && rm -rf /var/lib/apt/lists/*

# ImageMagick désactive souvent la lecture PDF par défaut pour raisons de sécu.
# On autorise la lecture/écriture de PDF/PS/EPS (si le fichier de policy existe).
# NB: pour convertir PDF -> images, préfère pdftocairo (Poppler) quand c'est possible.
RUN if [ -f /etc/ImageMagick-6/policy.xml ]; then \
      sed -i 's#<policy domain="coder" rights="none" pattern="PDF" />#<policy domain="coder" rights="read|write" pattern="PDF" />#' /etc/ImageMagick-6/policy.xml && \
      sed -i 's#<policy domain="coder" rights="none" pattern="PS" />#<policy domain="coder" rights="read|write" pattern="PS" />#' /etc/ImageMagick-6/policy.xml && \
      sed -i 's#<policy domain="coder" rights="none" pattern="EPS" />#<policy domain="coder" rights="read|write" pattern="EPS" />#' /etc/ImageMagick-6/policy.xml || true ; \
    fi

# (Optionnel) Réglage Tesseract si tu ajoutes d’autres langues
ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/4.00/tessdata

# Création d’un user non-root (bonnes pratiques)
RUN useradd -m -u 10001 appuser
USER appuser

WORKDIR /app
COPY --from=builder /app/main /app/main

EXPOSE 8080
CMD ["/app/main"]
