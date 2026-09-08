class Retriever:
    def __init__(self, embedder, store):
        self.embedder = embedder
        self.store = store

    def retrieve(self, query, k=5):
        q_emb = self.embedder.encode(query)[0]
        return self.store.search(q_emb, k=k)
