from edu247model.mcp.protocol import MCPMessage
from edu247model.mcp.tools import list_tools, get_tool

class MCPServer:
    def handle(self, req):
        cmd = req.get("cmd")

        if cmd == "list_tools":
            return {"tools": list_tools()}

        if cmd == "call_tool":
            name = req.get("name")
            args = req.get("args", {})
            tool = get_tool(name)
            if not tool:
                return {"error": f"Tool '{name}' not found"}
            return {"result": tool(args)}

        return {"error": "unknown_cmd"}

def main():
    server = MCPServer()

    while True:
        line = input()
        if not line.strip():
            continue

        req = MCPMessage.decode(line)
        res = server.handle(req)
        print(MCPMessage.encode(res))

if __name__ == "__main__":
    main()
