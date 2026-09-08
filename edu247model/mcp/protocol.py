import json

class MCPMessage:
    @staticmethod
    def encode(obj):
        return json.dumps(obj)

    @staticmethod
    def decode(line):
        return json.loads(line)
