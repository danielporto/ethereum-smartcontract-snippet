#!/usr/bin/env bash
for fin in $(ls solidity | grep ".sol"); do
  echo "Processing ${fin}..."
  # convert to lowercase
  fout=$(echo "$fin" | tr '[:upper:]' '[:lower:]')
  # replace extension
  fout="${fout%????}.go"
  echo "Generating ${fout}..."
  abigen --sol solidity/$fin --pkg contracts --out contracts/$fout
done;
echo "Done."
