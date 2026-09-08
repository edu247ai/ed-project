from .base import Skill
from edu247model.rag.pipeline import RAGPipeline

class RAGSkill(Skill):
    name = "rag"

    def __init__(self, supervisor):
        super().__init__(supervisor)
        self.rag = RAGPipeline()
        self.rag.build()

    def run(self, query):
        return self.rag.query(query)
