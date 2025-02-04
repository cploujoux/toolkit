from beamlit.agents import agent
from logging import getLogger
from beamlit.common import init
from starlette.websockets import WebSocket

settings = init()
logger = getLogger(__name__)


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
    try:
        agent.bind_tools(functions)
        await agent.aconnect(websocket)
    except Exception as e:
        logger.error(f"Error connecting to agent: {str(e)}")
        await websocket.send_text(str(e))
        await websocket.close()
