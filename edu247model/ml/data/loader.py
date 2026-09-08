import pandas as pd
import os

def load_africa_education_text_chunks():
    base = "edu247model/ml/data/africa_education/files"
    files = [f for f in os.listdir(base) if f.endswith(".csv")]

    chunks = []
    for f in files:
        df = pd.read_csv(os.path.join(base, f))
        for _, row in df.iterrows():
            text = " ".join([f"{col}: {row[col]}" for col in df.columns])
            chunks.append(text)

    return chunks
