# Render the implementation note

The [canonical manuscript](../manuscript/README.md) combines a technical
description with an archived basic-demo plot. The build stages copy that plot
from its retained source so repeated renders use the same evidence.

Install uv, Python 3.14, Pandoc with a compatible `pandoc-crossref`, and XeLaTeX
with its packages. Use a committed Git checkout named `goference` and an unused
location for the Template engine. From the Goference repository root:

```bash
GOFERENCE_ROOT="$(pwd -P)"
TEMPLATE_ROOT="$GOFERENCE_ROOT/../goference-render-template"
git clone https://github.com/docxology/template.git "$TEMPLATE_ROOT"
git -C "$TEMPLATE_ROOT" switch --detach 87704feae0a4490ef7dad9c931d268e8c30273c2

mkdir -p output/figures
cp active-goference/output/png_simple/basic_belief_evolution_2025-09-22_08-39-19.png \
  output/figures/basic_belief_evolution.png
cp docs/manuscript/figure_registry.json output/figures/figure_registry.json

cd "$TEMPLATE_ROOT"
uv sync --locked --no-dev --python 3.14
mkdir -p projects/working
ln -s "$GOFERENCE_ROOT" projects/working/goference
uv run --no-sync python scripts/pipeline/stage_03_render.py --project working/goference --skip-manuscript-hydration
uv run --no-sync python scripts/pipeline/stage_04_validate.py --project working/goference
```

The lifecycle alias and both Git indexes supply the renderer's provenance inputs.
Read `output/pdf/goference_combined.pdf` and `output/web/index.html` under
the Goference checkout. Validation and provenance receipts are in
`output/reports/`. Run the [Go checks](setup.md) separately.
