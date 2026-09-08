from sentence_transformers import SentenceTransformer

model = SentenceTransformer("all-MiniLM-L6-v2")

def embed_text(text: str):
    vec = model.encode(text)
    return vec.tolist()  # SQLAlchemy expects Python list
