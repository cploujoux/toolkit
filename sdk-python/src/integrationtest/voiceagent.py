from beamlit.agents import agent
from beamlit.common import init

settings = init()

from typing import AsyncIterator
from starlette.websockets import WebSocket


async def websocket_stream(websocket: WebSocket) -> AsyncIterator[str]:
    while True:
        data = await websocket.receive_text()
        yield data

@agent(
    agent={
        "metadata": {
            "name": "voice-agent",
            "environment": "production",
        },
        "spec": {
            "description": "A chat agent using Beamlit to handle your tasks.",
            "model": "gpt-4o-mini-realtime-preview",
        },
    },
    remote_functions=["brave-search"],
)
async def main(
    websocket: WebSocket, agent, functions,
):
    agent.bind_tools(functions)
    await agent.aconnect(websocket_stream, websocket.send_text)