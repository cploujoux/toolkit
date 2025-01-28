import asyncio

from beamlit.agents import agent, get_chat_model
from beamlit.common import init
from beamlit.functions import function
from langgraph.graph.graph import CompiledGraph

settings = init()
chat = get_chat_model("gpt-4o-mini")

@function()
async def get_weather(location: str):
    """Get the weather for a given location"""
    print("Using getWeather");
    return "The weather in " + location + " is sunny"

@agent(
    override_model=chat,
    override_functions=[get_weather],
)
async def main(request, agent: CompiledGraph):
    config = {"configurable": {"thread_id": "thread_id"}}
    agent_body = {"messages": [("user", "What's the weather in san francisco ?")]}
    response = None
    async for chunk in agent.astream(agent_body, config=config):
        response = chunk
    print(response)


asyncio.run(main({}))
