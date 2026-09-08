from .embedder import Embedder
from .store import VectorStore
from .loader import load_africa_education_text
from .retriever import Retriever

class RAGPipeline:
    def __init__(self):
        self.embedder = Embedder()
        self.store = VectorStore()
        self.retriever = Retriever(self.embedder, self.store)

    def build(self):
        texts = load_africa_education_text()
        embeddings = self.embedder.encode(texts)
        self.store.add(embeddings, texts)

    def query(self, question, k=5):
        results = self.retriever.retrieve(question, k=k)
        return {
            "question": question,
            "context": results,
            "answer": self.generate_answer(question, results)
        }

    def generate_answer(self, question, context):
        # Simple answer generator — you can replace with LLM later
        return f"Based on the dataset, relevant information includes: {context[:2]}"
